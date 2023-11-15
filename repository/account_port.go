package repository

type Account struct {
	AccountID    int     `db:"account_id"`
	CustomerID   int     `db:"customer_id"`
	OpendingDate string  `db:"opening_date"`
	AccountType  string  `db:"account_type"`
	Amount       float64 `db:"amount"`
	Status       int     `db:"status"`
} 

type IAccountRepository interface {
	Create_Repository(Account) (*Account, error)
	GetAll_Repository(int) ([]Account, error)
}

