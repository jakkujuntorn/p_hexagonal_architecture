package repository

import (
	"github.com/jmoiron/sqlx"
)

// adaptors *****
type customerRepositoryDB struct {
	db *sqlx.DB
}

// Constructor
func NewCustomerRepositoryDB(dataBase *sqlx.DB) ICustomerRepository { // return ICustomerRepository ก็ได้เพราะเป็น  interface ของ customer
	return customerRepositoryDB{db: dataBase}
}

func (r customerRepositoryDB) GetAll_Repository() ([]Customer, error) {
	customers := []Customer{}
	// ไปสร้างฐานข้อมูลด้วย
	query := "select customer_id, name, date_of_birth, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		// repository ไม่ต้องจัดการ error แค่โยนออกไป
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById_Repository(id int) (*Customer, error) {
	customer := Customer{}
	query := "select customer_id, name, date_of_birth, zipcode, status from customers where customer_id =?"

	// db.Get เราต้องการแค่ 1 ต้องส่ง id เข้าไปด้วย
	err := r.db.Get(&customer, query, id)
	
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
