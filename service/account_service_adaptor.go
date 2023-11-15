package service

import (
	"hexagonal_architecture/errs"
	"hexagonal_architecture/logs"
	"hexagonal_architecture/repository"
	"strings"
	"time"
)

type accountService struct {
	// บริการที่จะใช้ อันนี้มาจาก repository.IAccountRepository *
	accRepo repository.IAccountRepository
}

// สั่งให้ return IAccountService(เป็น interface ของ sevice เอง)
// แต่ return  repository.IAccountRepository (เป็น interface ของ repository) ทำไม ****************
func NewAccountService(accRepo repository.IAccountRepository) IAccountService {
	// return repository.IAccountRepository
	return accountService{accRepo: accRepo}
}



func (s accountService) NewAccount_Sevice(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	// ******** validate input  ว่าถูกต้องไหม *******
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("Amount at least 5000")
	}
	// แปลงให้เป็นตัวอักษรเล็กก่อน แล้วเช็คค่า AccountType
	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("Account Type should be saving or checking")
	}

	// ********* ปั้น data ใหม่ ก่อนส่งไป repo *********
	account := repository.Account{
		CustomerID: customerID,
		// format วันที่่
		OpendingDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType:  request.AccountType,
		Amount:       request.Amount,
		Status:       1,
	}

	// ******* create DB ********
	newAcc, err := s.accRepo.Create_Repository(account)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewInternalServerError()
	}

	// ก่อน return ต้องปั้น AccountResponse
	// จะเขียน func มาช้วยในการแปลงข้อมูลก็ได้ เพราะมันแปลงหลายที่
	response := AccountResponse{
		AccountID:    newAcc.AccountID,
		OpendingDate: newAcc.OpendingDate,
		AccountType:  newAcc.AccountType,
		Amount:       newAcc.Amount,
		Status:       newAcc.Status,
	}

	return &response, nil
}

func (s accountService) GetAccounts_Sevice(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll_Repository(customerID)
	if err != nil {
		logs.Error(err)
		// ส่ง error ที่เราทำออกไป
		return nil, errs.NewInternalServerError()
	}

	//accounts ที่ได้ต้องแปลง เป็น AccountRespons
	responses := []AccountResponse{}

	// loop accounts ที่มาจาก DB ใสใน response
	// แสดงข้อมูลที่จำเป็นออกไป ไม่เอาทั้งหมด *****
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:    account.AccountID,
			OpendingDate: account.OpendingDate,
			AccountType:  account.AccountType,
			Amount:       account.Amount,
			Status:       account.Status,
		})
	}

	return responses, nil
}
