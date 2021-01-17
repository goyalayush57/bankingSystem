package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"src/src/models"
	"src/src/responses"
	"src/src/service"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (server *Server) CreateEmployee(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	emp := models.Employee{}
	err = json.Unmarshal(body, &emp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	empCreated, err := service.NewEmployeeService().CreateEmployee(emp)

	if err != nil {

		//formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, empCreated.ID))
	responses.JSON(w, http.StatusCreated, empCreated)
}

func (server *Server) FindEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// eid, err := strconv.ParseUint(vars["id"], 10, 32)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	eid := vars["id"]
	emp, err := service.NewEmployeeService().Get(eid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, emp)

}

func (server *Server) GetEmployees(w http.ResponseWriter, r *http.Request) {

	emps, err := service.NewEmployeeService().GetAll()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, emps)
}

func (server *Server) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	emp := models.Employee{}
	err = json.Unmarshal(body, &emp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vars := mux.Vars(r)
	// eid, err := strconv.ParseUint(vars["id"], 10, 32)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	eid := vars["id"]
	emp.ID = eid
	emp.UpdatedAt = time.Now()
	emp, err = service.NewEmployeeService().Update(emp)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, emp)
}

func (server *Server) DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	// eid, err := strconv.ParseUint(vars["id"], 10, 32)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	eid := vars["id"]

	err := service.NewEmployeeService().Delete(eid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) LinkEmployeeKYC(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	kyc := models.KYCDetails{}
	err = json.Unmarshal(body, &kyc)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vars := mux.Vars(r)
	cid := vars["id"]

	c, err := service.NewEmployeeService().LinkKyc(kyc, cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	c.KycDetails = &kyc
	responses.JSON(w, http.StatusOK, c)
}

func (server *Server) UnLinkEmployeeKYC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["id"]
	kycID, err := strconv.ParseUint(vars["kyc_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	c, err := service.NewEmployeeService().UnLinkKyc(uint32(kycID), cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, c)
}

func (server *Server) LinkEmployeeAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	acc := models.Account{}
	err = json.Unmarshal(body, &acc)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vars := mux.Vars(r)
	cid := vars["id"]

	c, err := service.NewEmployeeService().LinkAccount(acc, cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	c.Account = &acc
	responses.JSON(w, http.StatusOK, c)
}
func (server *Server) UnLinkEmployeeAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["id"]
	accID, err := strconv.ParseUint(vars["acc_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	c, err := service.NewEmployeeService().UnLinkAccount(accID, cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, c)
}
