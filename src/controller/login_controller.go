package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"src/src/auth"
	"src/src/db"
	"src/src/models"
	"src/src/responses"
	"src/src/util"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	err = validateUser(user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Id, user.Password)
	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(userID, password string) (string, error) {

	var err error

	var hashedPass string
	var role string
	if strings.Contains(userID, "sa") {
		admin, err := db.NewAdminOperation().Get(userID)
		if err != nil {
			return "", err
		}
		hashedPass = admin.Password
		role = models.AdminRole
	} else if strings.Contains(userID, "se") {
		emp, err := db.NewEmployeeOperation().Get(userID)
		if err != nil {
			return "", err
		}
		hashedPass = emp.Password
		role = models.EmployeeRole
	} else {
		return "", fmt.Errorf("User is neither an employee nor admin")
	}

	err = util.VerifyPassword(hashedPass, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(userID, role)
}

func validateUser(u models.User) error {

	if u.Password == "" {
		return errors.New("Required Password")
	}
	if u.Id == "" {
		return errors.New("Required ID")
	}
	// if u.Email == "" {
	// 	return errors.New("Required Email")
	// }
	// if err := checkmail.ValidateFormat(emp.Email); err != nil {
	// 	return errors.New("Invalid Email")
	// }
	return nil

}
