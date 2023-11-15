package repository

import (
	"encoding/json"
	"errors"
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerRepositoryMock struct {
	customers []Customer
}



var customers []Customer

func NewCustomerRepositoryMock() ICustomerRepository { // ใช้ ICustomerRepository เพราะ interface ของ customer
	// Mock Data
	// customers := []Customer{
	// 	{CustomerID: 1, Name: "Rattha", City: "New Bangkok", ZipCode: "21547", Status: 1},
	// 	{CustomerID: 2, Name: "YingYo", City: "Bankkape", ZipCode: "10255", Status: 0},
	// 	{CustomerID: 3, Name: "Big GG", City: "Chaingmai", ZipCode: "12345", Status: 0},
	// 	{CustomerID: 4, Name: "Rarin", City: "Panum", ZipCode: "24158", Status: 1},
	// 	{CustomerID: 5, Name: "Kwan", City: "Pareriew", ZipCode: "24000", Status: 1},
	// }
	// return customerRepositoryMock{customers}

	// add Mock data

	customers = append(customers, Customer{CustomerID: 2, Name: "YingYo", City: "Bankkape", ZipCode: "10255", Status: 0})
	customers = append(customers, Customer{CustomerID: 1, Name: "Rattha", City: "New Bangkok", ZipCode: "21547", Status: 1})
	customers = append(customers, Customer{CustomerID: 4, Name: "Rarin", City: "Panum", ZipCode: "24158", Status: 1})
	return customerRepositoryMock{customers}
}

func (r customerRepositoryMock) GetAll_Repository() ([]Customer, error) {
	// bangkok code
	// return r.customers, nil

	// born to dev *****
	return customers, nil
}

func (r customerRepositoryMock) GetById_Repository(id int) (*Customer, error) {
	// loop เพื่อหา id ที่ตรงกันกับที่ส่งเข้ามา
	// bangkok code
	// for _, customer := range r.customers {
	// 	if customer.CustomerID == id {
	// 		return &customer, nil
	// 	}
	// }

	// born to dev *****
	for _, customer := range customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("Customer not Found")
}

// Born to DEv **************
// create Mock data
func CreateMockData(w http.ResponseWriter, r *http.Request) {

	var customer Customer
	// ถอดค่าจากตัวแปร
	_ = json.NewDecoder(r.Body).Decode(&customer)

	customers = append(customers, customer)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// update Mock data
func UpdateMockData(w http.ResponseWriter, r *http.Request) {

}

//DElete
func DeleteMockData(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	params, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	status := false

	for index, item := range customers {
		if item.CustomerID == params {
			customers = append(customers[:index], customers[index+1:]...)
			status = true
		}
		// else {
		// 	json.NewEncoder(w).Encode("Can not Found ID")
		// 	break
		// }
	}

	if status {
		w.Header().Set("content-type", "application/json")
		// ทำหน้าที่ return ค่าออกไป
		json.NewEncoder(w).Encode(customers)

	} else {
		json.NewEncoder(w).Encode("Can not Found ID")

	}

	// w.Header().Set("content-type", "application/json")
	// // ทำหน้าที่ return ค่าออกไป
	// json.NewEncoder(w).Encode(customers)
}
