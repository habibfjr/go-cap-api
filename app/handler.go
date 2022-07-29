package app

import (
	"capi/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// *contain the query string url and get "status" key
	customerStatus := r.URL.Query().Get("status")

	// *get data with desired status
	customers, err := ch.service.GetAllCustomer(customerStatus)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {

	// * get route variable
	vars := mux.Vars(r)

	customerID := vars["customer_id"]

	customer, err := ch.service.GetCustomerByID(customerID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	// * return customer data
	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
