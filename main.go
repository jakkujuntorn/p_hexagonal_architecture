package main

import (
	_ "database/sql"
	"fmt"

	// "log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	"hexagonal_architecture/errs"
	"hexagonal_architecture/handler"
	"hexagonal_architecture/logs"
	"hexagonal_architecture/repository"
	"hexagonal_architecture/service"


	
)

func main() {

	// db,err:=sqlx.Open("mysql","root:P@ssw0rd@tcp(127.0.0.1:3306)/banking?parseTime=true")  parseTime=true จะใช้กับ ชนิดข้อมูล time.Time
	initTimeZone()
	initConfig()
	db := initDB()

	// ******************* Customer ***************************
	//***********REpo*************************
	// ข้อมูลมากจาก DB  Repository ไปดึงมา *****
	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	_ = customerRepositoryDB
	// customerRepositoryDB.GetAll()
	// Mock Data
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryMock

	//***************** Service ********************
	//  ข้อมูลมากจาก Repository  service ไปดึงมา *****
	// ค่าที่ return ออกมามัน struct ไป implement กับ interface แล้ว ****
	customerService := service.NewCustomerService(customerRepositoryMock)
	// customerService5 := service.NewCustomerService(customerRepositoryMock)
	// customerService5.GetCustomer()

	//****************Handler**********************
	// ข้อมูลมากจาก service  handler ไปดึงมา *****
	customerHandler := handler.NewCustomerHandler(customerService)
	// customerHandler.HandlerGetCustomer()



	// **************************** Account **************************
	// **Repo ***********
	accountRepositoryDB := repository.NewAccountRepository(db)
	// *** Service ******
	accountService := service.NewAccountService(accountRepositoryDB)
	// **** Handler *****
	accountHandler := handler.NewAccountHandler(accountService)
	// accountHandler.Test()
	
	// ใช้ mux *************************************
	router := mux.NewRouter()

	//********************** Customers **************************
	router.HandleFunc("/customers", customerHandler.HandlerGetCustomers).Methods(http.MethodGet)
	// mux paramiters ใช้  {}  {customerID:[0-9]+} ใส regular expression เพ่อกำหนดรูปแบบ ตัวอักษร
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.HandlerGetCustomer).Methods(http.MethodGet)

	// ******************* Account ******************************
	// mux paramiters ใช้  {}   {customerID:[0-9]+}
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)
	// router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.TestMuxRouter).Methods(http.MethodPost)
	
	// Create Mock Data ***
	router.HandleFunc("/create", repository.CreateMockData).Methods(http.MethodPost)
	router.HandleFunc("/delete/{customerID:[0-9]+}", repository.DeleteMockData).Methods(http.MethodDelete)

	// Log ******
	// log.Printf("Start DB Port %v", viper.GetInt("app.port"))
	logs.Info("Start DB Port" + viper.GetString("app.port"))

	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)

}

func initTimeZone() {
	// set TimeZone
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initConfig() {
	// ใช้ viper  เพื่อ เก็บค่า config ต่างๆ
	//vdo 1  นาที ที่ 1.13.00
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Set Env ได้
	viper.AutomaticEnv()
	// set ค่าใหม่กับ config เช่น APP_PORT=5000  ใช้ _  *******
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// APP_PORT=5000 go run .
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// * ส่ง pointer sqlx.DB
func initDB() *sqlx.DB {

	// .Open("mysql","root:P@ssw0rd@tcp(127.0.0.1:3306)/banking?parseTime=true")

	// เอาค่ามาต่กันให้เหมือนการตั้งค่าปกติ
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	// fmt.Println(dsn)
// ข้อมูลรับค่าเป็น 2 ชุดเพราะ .Operมี "","" ต้แงใช้ข้อมูล 2 ชุด 
	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	// ไม่ต้อง return err เพราะ มถ้ามีมันจะติดตรงนี้ *****
	if err != nil {
		panic(errs.AppError{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Connect Server ...",
		})
	}

	// เพื่อ *******
	db.SetConnMaxLifetime(3 * time.Minute)
	//เปิดใช้กี่คนพร้อมกัน
	db.SetMaxOpenConns(10)
	//
	db.SetMaxIdleConns(10)

	return db
}
