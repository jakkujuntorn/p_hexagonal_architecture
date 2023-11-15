package service

import (
	// "fmt"
	"database/sql"
	_ "errors"
	"hexagonal_architecture/errs"
	"hexagonal_architecture/logs"
	"hexagonal_architecture/repository"
	_ "log"

	// "net/http"
)

// adapter ทำหน้าที่ส่งข้อมูลจาก DB ไป handle
type customerService struct {

	// ไม่ได้อ้างถึง db แต่อ้างถึง interface ที่เป็น port ที่ต่อ db
	custRepo repository.ICustomerRepository
}

// การที่จะ return เป็น ICustomerService ต้อง confrom ตาม interface ก่อน
func NewCustomerService(custRepo repository.ICustomerRepository) ICustomerService {
	// return struct ด้านบน ****
	return customerService{
		custRepo: custRepo,
	}
}

// ใช้ interface ของ repo มา GetAll
func (s customerService) GetCustomers_Sevice() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll_Repository()

	if err != nil {
	
		// log err จะอยู่ในั้นนี้ **********
		// log.Println(err)

		// logs ที่ทำมาใหม่รับค่าเป็น string ต้องส่ง err.Error() เข้าไป *****
		// ไป set ที่ logs ใหม่ ให้ส่งค่า err หรือ "Error" ก็ได้ ถ้าไม่ปรับที่ log ต้องส่ง err.Error() อย่างเดียวเท่านั้น *****
		logs.Error(err)
		// logs.Error("Error")

	
		// return nil, errs.AppError{
		// 	Code: http.StatusNotFound,
		// 	Message:"Customer Nut Found",
		// }

		return nil, errs.NewInternalServerError()
		
		
	}
	
	// ปั้น obj ใหม่ ออกไปใช้งาน เพราะไม่อยากให้ font end รู้ว่ามีข้อมูลอะไรบ้าง ส่งไปที่จำเป็น ******
	custReponses := []CustomerResponse{}
	// เอา customers ที่ได้จาก DB มาวนใสค่าใหม่
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		// ใสค่าใหม่ลงไป *****
		custReponses = append(custReponses, custResponse)
	}
	return custReponses,nil
}

func (s customerService) GetCustomer_Sevice(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById_Repository(id)
	if err != nil {

		// ดัก Error  จาก Repo ก่อนส่งออกไป handler
		if err == sql.ErrNoRows{
			// error แบบเก่า
			// return nil, errors.New("Customer not Found (Error from service)")

			// error แบบใหม่
			return nil, errs.NewNotFoundError("Customer not Found (Error Func from service ***)")
		}

		// log err จะอยู่ในั้นนี้ **********
		// log.Println(err)
		logs.Error(err)
		return nil, errs.NewInternalServerError()
	}
	custReponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	// return pointer ออกไป
	return &custReponse, nil
}

