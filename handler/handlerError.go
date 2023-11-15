package handler

import (
	"fmt"
	"hexagonal_architecture/errs"
	"net/http"
)

//หน้าที่ ไฟล์นี้ helpe function

func handlerCustomerError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	//  case errs.AppError
	case errs.AppError:
		// e.Code ได้เพราะมาจาก type AppError
		w.WriteHeader(e.Code)
		fmt.Fprint(w, e)
		// error ปกติ 
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, e)
		
	}
}
