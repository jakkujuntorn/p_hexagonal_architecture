package service

type NewAccountRequest struct {
	// AccountID DB สร้างให้แล้ว
	// AccountID    int     `json:"account_id"`

	// CustomerID เอามาจาก url ได้ ให้ handler layer ดึงมาจาก url
	// CustomerID   int     `json:"customer_id"`

	// OpendingDate เอาจาก DB ชัวร์กว่า
	// OpendingDate string  `json:"opening_date"`

	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`

	// Status มีค่า default อยู่แล้ว
	// Status      int     `json:"status"`
}

type AccountResponse struct {
	AccountID    int     `json:"account_id"`
	OpendingDate string  `json:"opening_date"`
	AccountType  string  `json:"account_type"`
	Amount       float64 `json:"amount"`
	Status       int     `json:"status"`
}

type IAccountService interface {
	NewAccount_Sevice(int, NewAccountRequest) (*AccountResponse, error)
	GetAccounts_Sevice(int) ([]AccountResponse, error)
}
