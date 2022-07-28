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

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	query := "select * from customers"

	rows, err := d.client.Query(query)
	if err != nil {
		log.Println("error querying customers", err.Error())
		return nil, err
	}
	var customers []Customer
	for rows.Next() {
		var c Customer
		rows.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
		if err != nil {
			log.Println("error scanning customer", err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDB) FindByID(customerID string) (*Customer, *errs.AppErr) {
	query := "select * from customers where customer_id = $1"

	// row := d.client.QueryRow(query, customerID)
	// err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)

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
