package service


type CustomerResponse struct {
	CustomerID  int    `json:"customer_id"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
}

//port service
type ICustomerService interface {
	GetCustomers_Sevice()([]CustomerResponse,error)
	GetCustomer_Sevice(int) (*CustomerResponse,error)
}
