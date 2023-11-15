package handler

import (
	"encoding/json"
	"fmt"
	_ "hexagonal_architecture/errs"
	"hexagonal_architecture/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// adapter  ต้องใช้ service เท่านั้น ****
type customerHandler struct {
	// บริการที่ handler จะใช้
	custSrv service.ICustomerService
}

// ทำหน้าที่ส่งผ่าน custSrv service.CustomerService เพื่อจะไม่ให้ใช้โดยตรง  ****
func NewCustomerHandler(custSrv service.ICustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}
}

// return ตาม muk HandleFunc
func (h customerHandler) HandlerGetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSrv.GetCustomers_Sevice()
	if err != nil {
		// status 500
		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Fprintln(w, err)

		handlerCustomerError(w, err)
		return
	}
	// ถ้าไม่ใส header จะเป็น  Content-Type: text/plain; charset=utf-8
	// แปลง customers เป็น json ก่อนส่งออกไป *********
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func (h customerHandler) HandlerGetCustomer(w http.ResponseWriter, r *http.Request) {

	// รับค่าจาก url จะได้ค่าที่ต่อ path นั้นมา  ด้วยคำสั่ง mux.Vars(r)["customerID"] *******
	// strconv.Atoi แปลง string to int
	customerID, err := strconv.Atoi(mux.Vars(r)["customerID"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err)
		return
	}

	// ดึงค่าแบบนี้ก็ได้ แปลงที่ละขั้นตอน และส่ง number ไป
	// params := mux.Vars(r)
	// number,err:=strconv.Atoi(params["customerID"])

	customer, err := h.custSrv.GetCustomer_Sevice(customerID)
	if err != nil {

		// // เช็คว่า errs.AppError มีตัวตนไหม
		// appErr, ok := err.(errs.AppError)
		// if ok {
		// 	// ตรงนี้ service คิดมาให้แล้ว handler แค่ส่งตัวไป  fontend *****
		// 	w.WriteHeader(appErr.Code)
		// 	fmt.Fprintln(w, appErr.Message)
		// 	return
		// }

		// เขียนแบบย่อด้วย func แยกออกไปอีกไฟล์ เอาข้างบนไปอยู่อีกไฟล์
		handlerCustomerError(w, err)

		// ตรงนี้  handler ต้องคิด err เอง *****
		// status 500
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	// แปลง customers เป็น json ก่อนส่งออกไป *********
	json.NewEncoder(w).Encode(customer)
}
