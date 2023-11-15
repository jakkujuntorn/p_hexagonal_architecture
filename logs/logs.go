package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"hexagonal_architecture/errs"
)

var log *zap.Logger

func init() {
	// Log, _ = zap.NewProduction()

	config := zap.NewProductionConfig()

	// เปลี่ยนค่า config ค่า TimeKey = "timestamp"
	config.EncoderConfig.TimeKey = "timestamp"

	// เปลี่ยนค่า config ค่า EncodeTime = zapcore.ISO8601TimeEncoder แสดงผล วันที่ ปี-เดือน-วัน
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// StacktraceKey ข้อมูลการ error ถ้าไม่ใส "" จะมีข้อมูลการ error ละเอียดมาก
	config.EncoderConfig.StacktraceKey = ""

	var err error
	// ******* แก้ คนที่เรียกใช้ ให้เป็นบรรทัดที่เรียกจริงๆ *****
	// zap.AddCallerSkip(1) skip ไป 1 ขั้นตอน จะทำให้ตอนเรียก log จะไม่แสดงว่า logs.go เป็นคนเรียก จะแสดงคนที่เรียกจริงๆ ********
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
func Debug(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// interface{} แบบนี้  คือ OBJ
func Error(message interface{}, fields ...zap.Field) {
	// log.Info(message, fields...)

	// 1.เช็คว่า message เป็น error จริงรึป่าว
	// msg,ok:=message.(error)
	// if ok { // ถ้า ok ถือว่าเป็น type error
	// 	msg.Error()
	// }

	//2. switch message.(type) เช็ค type เลย
	// switch type
	// รับ err.Error() หรือ  "Error Na ja" *****
	// เช็ค err ที่เข้ามา เป็น string หรือ  err.Error()

	switch v := message.(type) { // ถ้า message.(string) จะได้ค่าอะไร ****
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	case errs.AppError: // เพิ่ม case AppError
		log.Error(v.Message, fields...)
	}

}

// เสริม 
	// // เช็คว่า errs.AppError มีตัวตนไหม
		// appErr, ok := err.(errs.AppError)
		// if ok {
		// 	// ตรงนี้ service คิดมาให้แล้ว handler แค่ส่งตัวไป  fontend *****
		// 	w.WriteHeader(appErr.Code)
		// 	fmt.Fprintln(w, appErr.Message)
		// 	return
		// } 
