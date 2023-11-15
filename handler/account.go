package handler

import (
	"encoding/json"

	"hexagonal_architecture/errs"
	_ "hexagonal_architecture/logs"
	"hexagonal_architecture/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accSrv service.IAccountService
}

// type IHandler interface{
// 	NewAccountHandler(w http.ResponseWriter, r *http.Request)
// 	GetAccountsHanlder(w http.ResponseWriter, r *http.Request)
// }

// ที่ไม่ return IAccountService เพราะ ไม่ต้องส่งให้ใครแล้ว *****
// แต่จะ return IAccountService ก็ได้
// เอา Func ของตัวเองไปใช้ได้เลย *******
func NewAccountHandler(accSrv service.IAccountService) *accountHandler {
	return &accountHandler{accSrv: accSrv}
}

func (h *accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	// เช็ค header ก่อน ว่าถุกต้องรึป่าว เช็คว่าส่ง json มาให้รึป่าว 
	if r.Header.Get("content-type") != "application/json" {
		// errs.NewValidationError เป็น AppError
		handlerCustomerError(w, errs.NewValidationError("Request body incorrect format application/json"))
		return
	}
	
	request := service.NewAccountRequest{}
	// fmt.Println(json.NewDecoder(r.Body))
	// ถอดค่า Body ด้วย NewDecoder ตาม format NewAccountRequest *******
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		handlerCustomerError(w, errs.NewValidationError("Request body cannot Decoder"))
		return
	}

	response, err := h.accSrv.NewAccount_Sevice(customerID, request)
	if err != nil {
		handlerCustomerError(w, err)
		return
	}

	// เช็คทุกอย่างแล้วผ่าน
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "appliccation/json")
	json.NewEncoder(w).Encode(response)

}

func (h *accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	response, err := h.accSrv.GetAccounts_Sevice(customerID)
	if err != nil {
		// ไม่ต้องคิด error แค่ส่งออกไปก็พอ
		handlerCustomerError(w, err)
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// handler Mux คือ Func ที่รับ (w http.ResponseWriter, r *http.Request) 
func (h *accountHandler)TestMuxRouter(w http.ResponseWriter, r *http.Request)  {
	
}


