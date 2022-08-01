package app

import (
	"capi/domain"
	"capi/logger"
	"capi/service"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func Start() {

	// *wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// * create ServeMux
	// mux := http.NewServeMux()
	mux := mux.NewRouter()
	// ServeMux is used to define routes. If not implemented then it will use the default
	authR := mux.PathPrefix("/auth").Subrouter()

	authR.Use(loggingMiddleware)

	// * defining routes
	// mux.HandleFunc("/customers", addCustomer).Methods(http.MethodPost)

	// mux.HandleFunc("/customers/{customer_id}", getCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods(http.MethodGet)
	// mux.HandleFunc("/customers/{customer_id:[0-9]+}", updateCustomer).Methods(http.MethodPut)
	// mux.HandleFunc("/customers/{customer_id:[0-9]+}", deleteCustomer).Methods(http.MethodDelete)

	mux.Use(authMiddleware)
	// * starting the server
	http.ListenAndServe(":8080", mux)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := time.Now()
		next.ServeHTTP(w, r)
		logger.Info(fmt.Sprintf("%v %v %v", r.Method, r.URL, time.Since(timer)))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// next.ServeHTTP(w, r)
		token := r.Header.Get("Authorization")
		logger.Info(token)

		// check if token has bearer
		if !strings.Contains(token, "Bearer") {
			logger.Error("invalid token, must contain 'Bearer'")
			return
		}
		// split token -> ambil tokennya buang "Bearer" nya
		getToken := ""
		tokenArray := strings.Split(token, " ")
		if len(tokenArray) == 2 {
			getToken = tokenArray[1]
		}
		fmt.Println(getToken)

		// parsing token, err := jwt.Parse(
		signedToken, err := jwt.Parse(getToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("rahasia"), nil
		})
		if err != nil {
			logger.Error("Failed to parse token:" + err.Error())
		}

		// check token validation
		if signedToken.Valid {
			fmt.Println("token is valid")
			writeResponse(w, http.StatusOK, signedToken)
		}

	})
}
