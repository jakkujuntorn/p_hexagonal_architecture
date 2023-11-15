package errs

import "net/http"

// ทำเรียนแบบ Error() string ของจริง
// confrom ตาม error interface *****
// เพราะ จะให้ error ที่ทำมาใหม่ ทำงานคล้าย Eror ของจริง เลย
type AppError struct {
	Code    int
	Message string
	// ใส error ไปเลยแทนการ conform interface error
	error
}

//recivever function ถ้าไม่ทำ func แบบนี้จะ return error ที่ func ไม่ได้ *****
// confrom ตาม error interface *****
// ด้วยการทำ recivever Error() ****
// func (e AppError) Error() string {
// 	return e.Message
// }

// ทำ func NotFound รองรับไว้เลย ตอนใช้แค่เรียก func และส่ง message เข้ามา
// จะได้ไม่ต้องใสค่า  Code and Message ทุกครั้งที่เรียกใช้งาน แต่ต้อง สร้าง Error ให้ครอบคลุมให้หมด
// return จะ เป็น interface error ออกไป
func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "Server cannot connect (Error from Func errs.go)",
	}
}

func NewValidationError(message string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}

}
