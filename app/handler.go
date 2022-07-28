package app

import (
	"capi/service"
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
)

// type Customer struct {
// 	ID      int    `json:"id" xml:"id"`
// 	Name    string `json:"name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	ZipCode string `json:"zip_code" xml:"zipcode"`
// }

// var customers []Customer = []Customer{
// 	{1, "User1", "Jakarta", "12345"},
// 	{2, "User2", "Surabaya", "67890"},
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello Celerates!")
// }

// func getAllCustomers(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprint(w, "get customer endpoint\n")
// 	if r.Header.Get("Content-Type") == "application/xml" {
// 		w.Header().Add("Content-Type", "application/xml")
// 		xml.NewEncoder(w).Encode(customers)
// 	} else {
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(customers)

// 	}
// }
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}
}

func (ch *CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {

	// * get route variable
	vars := mux.Vars(r)

	customerID := vars["customer_id"]

	customer, err := ch.service.GetCustomerByID(customerID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w, err.Error())
	}

	// * return customer data
	// w.Header().Add("Content-Type", "application/json")
	writeResponse(w, http.StatusOK, customer)
	// json.NewEncoder(w).Encode(customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	// json.NewEncoder(w).Encode(data)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	// *get route variable
// 	vars := mux.Vars(r)
// 	customerId := vars["customer_id"]

// 	// *convert string to int
// 	// id, _ := strconv.Atoi(customerId)
// 	id, err := strconv.Atoi(customerId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "Invalid customer id")
// 	}

// 	// *search for customer data
// 	var cust Customer

// 	for _, data := range customers {
// 		if data.ID == id {
// 			cust = data
// 		}
// 	}

// 	if cust.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "Invalid customer found")
// 		return
// 	}

// 	// *return customer data
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(cust)
// }

// func addCustomer(w http.ResponseWriter, r *http.Request) {
// 	// *decode request body
// 	var cust Customer
// 	json.NewDecoder(r.Body).Decode(&cust)

// 	// *generate new id
// 	nextID := getNextID()
// 	cust.ID = nextID

// 	// *save data to array
// 	customers = append(customers, cust)

// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintln(w, "customer successfully created")

// }

// func getNextID() int {
// 	lastIndex := len(customers) - 1
// 	lastCustomer := customers[lastIndex]
// 	// cust := customers[len(customers)-1]
// 	return lastCustomer.ID + 1
// }

// func updateCustomer(w http.ResponseWriter, r *http.Request) {
// 	// *get route variable
// 	vars := mux.Vars(r)
// 	customerId := vars["customer_id"]

// 	// *convert string to int
// 	id, err := strconv.Atoi(customerId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "Invalid customer id")
// 		return
// 	}

// 	// *search for customer data
// 	var cust Customer

// 	for customerIndex, data := range customers {
// 		if data.ID == id {
// 			cust = data

// 			// decode req body
// 			var newCust Customer
// 			json.NewDecoder(r.Body).Decode(&newCust)

// 			// do update
// 			customers[customerIndex].Name = newCust.Name
// 			customers[customerIndex].City = newCust.City
// 			customers[customerIndex].ZipCode = newCust.ZipCode

// 			w.WriteHeader(http.StatusOK)
// 			fmt.Fprintln(w, "customer data updated")
// 		}
// 	}

// 	if cust.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "Invalid customer found")
// 		return
// 	}

// }

// func deleteCustomer(w http.ResponseWriter, r *http.Request) {

// 	// *get route variable
// 	vars := mux.Vars(r)
// 	customerId := vars["customer_id"]

// 	// *convert string to int
// 	id, err := strconv.Atoi(customerId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "Invalid customer id")
// 		return
// 	}

// 	// *search for customer data
// 	// var cust Customer

// 	// *delete function
// 	for customerIndex, data := range customers {
// 		if data.ID == id {
// 			customers = append(customers[:customerIndex], customers[customerIndex+1:]...)
// 			// code above will take data other than the selected one and put it in the original array
// 			w.WriteHeader(http.StatusOK)
// 			fmt.Fprintln(w, "customer data deleted")
// 		}
// 	}

// if cust.ID == 0 {
// 	w.WriteHeader(http.StatusNotFound)
// 	fmt.Fprint(w, "Invalid customer found")
// 	return
// }
// }
