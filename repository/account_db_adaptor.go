package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

// ทำไม return  IAccountRepository หรือ accountRepositoryDB *****
// เพราะ มัน confrom ตาม interface IAccountRepository แล้ว ******
// แต่แนะนำ return ออกไปเป็น interface *********
func NewAccountRepository(db *sqlx.DB) IAccountRepository {
	return accountRepositoryDB{db: db}
}

// confrom interface ตาม IAccountRepository
func (r accountRepositoryDB) Create_Repository(acc Account) (*Account, error) {
	query := "insert into account(customer_id,opening_date,account_type,amount,status) values(?,?,?,?,?)"
	result, err := r.db.Exec(
		query,
		// เรียงลำดับข้อมูลก่อนส่งตาม query ****
		// การ insert ข้อมูล แต่ละ database ไม่เหมือนกัน *******
		acc.CustomerID,
		acc.OpendingDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)

	if err != nil {
		return nil, err
	}

	// ดึงค่า id ที่สร้างเสร็จ ล่าสุดออกมา
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// แปลง id จาก int64 เป็น int ปกติ ****
	acc.AccountID = int(id)

	// ที่ Return &acc ได้ เพราะ Func ให้ Return *Account  ****
	return &acc, nil
}

func (r accountRepositoryDB) GetAll_Repository(customerID int) ([]Account, error) {
	accounts := []Account{}
	query := "select account_id,customer_id,opening_date,account_type,amount,status from account where customer_id=?"
	err := r.db.Select(&accounts, query, customerID)

	if err != nil {
		// repository ไม่ต้องจัดการ error แค่โยนออกไป
		return nil, err
	}

	return accounts, nil
}
