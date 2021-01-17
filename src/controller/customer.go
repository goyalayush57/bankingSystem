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

func (server *Server) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	emp := models.Customer{}
	err = json.Unmarshal(body, &emp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	empCreated, err := service.NewCustomerService().CreateCustomer(emp)

	if err != nil {

		//formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, empCreated.ID))
	responses.JSON(w, http.StatusCreated, empCreated)
}

func (server *Server) FindCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["id"]

	c, err := service.NewCustomerService().Get(cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, c)

}

func (server *Server) GetCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := service.NewCustomerService().GetAll()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, customers)
}

func (server *Server) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	emp := models.Customer{}
	err = json.Unmarshal(body, &emp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vars := mux.Vars(r)
	cid := vars["id"]

	emp.ID = cid
	emp.UpdatedAt = time.Now()
	emp, err = service.NewCustomerService().Update(emp)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, emp)
}

func (server *Server) DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cid := vars["id"]

	err := service.NewCustomerService().Delete(cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) LinkCustomerKYC(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	kyc := models.KYCDetails{}
	err = json.Unmarshal(body, &kyc)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vars := mux.Vars(r)
	cid := vars["id"]

	c, err := service.NewCustomerService().LinkKyc(&kyc, cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, c)
}

func (server *Server) UnLinkCustomerKYC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["id"]
	kycID, err := strconv.ParseUint(vars["kyc_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = service.NewCustomerService().UnLinkKyc(uint32(kycID), cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, "")
}

func (server *Server) LinkCustomerAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	acc := models.Account{}
	err = json.Unmarshal(body, &acc)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vars := mux.Vars(r)
	cid := vars["id"]

	c, err := service.NewCustomerService().LinkAccount(&acc, cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, c)
}
func (server *Server) UnLinkCustomerAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["id"]
	accID, err := strconv.ParseUint(vars["acc_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = service.NewCustomerService().UnLinkAccount(accID, cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, "")
}
