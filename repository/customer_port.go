package repository

type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateofBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      int    `db:"status"`
}

// *** port ***
type ICustomerRepository interface {
	GetAll_Repository() ([]Customer, error)
	// return 1 customer  ต้องใช้ pointer
	GetById_Repository(int)(*Customer, error)
}
