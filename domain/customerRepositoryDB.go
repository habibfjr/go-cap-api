package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	// *connection string
	// connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	connStr := "postgres://postgres:Yhf171999@localhost/banking?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindAll(customerStatus string) ([]Customer, *errs.AppErr) {

	var customers []Customer

	if customerStatus == "" {
		query := "select * from customers"
		err := d.client.Select(&customers, query)
		if err != nil {
			logger.Error("error querying customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	} else {
		if customerStatus == "active" {
			customerStatus = "1"
		} else {
			customerStatus = "0"
		}
		query := "select * from customers where status = $1"
		err := d.client.Select(&customers, query, customerStatus)
		if err != nil {
			logger.Error("error querying customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return customers, nil
}

func (d CustomerRepositoryDB) FindByID(customerID string) (*Customer, *errs.AppErr) {
	query := "select * from customers where customer_id = $1"

	var c Customer

	err := d.client.Get(&c, query, customerID)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("customer data not found" + err.Error())
			return nil, errs.NewNotFoundError("customer data not found")
		} else {
			logger.Error("error scanning customer data " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil
}
