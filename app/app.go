package app

import (
	"capi/domain"
	"capi/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// *wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// * create ServeMux
	// mux := http.NewServeMux()
	mux := mux.NewRouter()
	// ServeMux is used to define routes. If not implemented then it will use the default

	// * defining routes
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// mux.HandleFunc("/customers", addCustomer).Methods(http.MethodPost)

	// mux.HandleFunc("/customers/{customer_id}", getCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods(http.MethodGet)
	// mux.HandleFunc("/customers/{customer_id:[0-9]+}", updateCustomer).Methods(http.MethodPut)
	// mux.HandleFunc("/customers/{customer_id:[0-9]+}", deleteCustomer).Methods(http.MethodDelete)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}
