package service

import (
	"capi/domain"
	"capi/dto"
	"capi/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppErr)
	GetCustomerByID(string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(customerStatus string) ([]dto.CustomerResponse, *errs.AppErr) {
	customers, err := s.repository.FindAll(customerStatus)
	if err != nil {
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// *put values after converted to DTO
	var dtoCust []dto.CustomerResponse
	for _, customer := range customers {
		dtoCust = append(dtoCust, customer.ToDTO())
	}
	return dtoCust, nil
}

func (s DefaultCustomerService) GetCustomerByID(customerID string) (*dto.CustomerResponse, *errs.AppErr) {
	cust, err := s.repository.FindByID(customerID)
	if err != nil {
		return nil, err
	}

	response := cust.ToDTO()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
