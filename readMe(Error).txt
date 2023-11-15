repository หน้าที่ คือดึงข้อมูลและเก็บข้อมูลจาก db เท่านั้น ถ้ามี error แค่ส่งออกไป
service
    -service ต้องคิดเรื่อง error  ต้องแยกให้ออกว่า เป็น err ประเภทไหน (สร้าง error แยกออกมา)
    -service ต้องทำ error ที่ส่งมาจาก repository ด้วยการปั้น error ใหม่ออกไปให้  bussiness
    -service  ต้องส่ง status ด้วย 404,500,200
    -ปั้นข้อมูลใหม่ออกไปด้วย struct ตัวใหม่เอาแต่ค่าที่จำเป็นเท่านั้น

 log จะมีตั้งแต่ service ขึ้นไป   (repository  ไม่ต้องมี)

 Error (service จะเป็นคนจัดการ **** ส่งไปตรงๆหรือ ปั้น  err ใหม่ ***** ส่วน  handler จะแค่ส่งผ่าน err ออกไป)
    -ร้ายแรงให้ log ออกมาดู (tactical error)
    -ไม่ร้ายแรงให้ส่งไป fontend ได้ (logic error)

********************************************
handler
    -handler มีหน้าที่ ถอด statusCode มาใช้ และ setHeader 
    -handler ทำหน้าที่ฝั่ง  adapter ไม่มี interface มีแต่  struct
    -handler ถ้าได้รับ Error จาก Service ให้ส่งต่อให้ Font End ได้เลย เพราะ service ทำาพร้อมใช้อยู่แล้ว
    -handler จะ set JSON,Header
    -้handler layer นี้จะต้องมีตัวทำ API เช่น fiber, mux, http

******************************************************************
struct
- ใน struct คือบริการต่างๆที่จะเรียกใช้ (มันคือ interface ของอีก layer)
    - หมายถึงบริการที่รับมาด้วย รับมาจาก อีก layer นึง / มัคือการอ้างถึงตัว  port (interface)
- ค่าที่ return ออกมาเป็น interface มัน implement โดย Struct แล้ว ******************


*********************************
- ทำไม service กับ handler ถึง return interface ออกมาเหมือนกัน *****
        - เพราะตั้งชื่อเหมือนกัน (แก้ชื่อต่างกันแล้ว) ****

******************************
การจัดการ Error  
    -error มี 2 ที่ ของ service กับ handler
        -service   จะเซตค่าไว้หมดล้ว แต่บางอันต้องส่ง string  เข้ามา
        -handler ต้องส่งค่าเข้ามาเพื่อเช็คื type ก่อน และ พารามิเตอร์
    -สร้าง Error ใหม่ แยกเป็น Error ต่างๆ
    -handler จะเป็นตัวเรียกใช้ โดยส่ง  status และ message
    -handler สร้าง  handler สำหรับ เรีบกใช้  error
    -สิ่งที่ต้องการ http.REsponse, ตัว  error
    -และใช้ switch แยก Type error จะทำให้ error ที่ส่งมา เป็น error ปกติ หรือ AppError ก็ได้
    - ถอด typr error  err.(errs.AppError) type err ที่ได้มาเป็น AppError หรือไม่ ***

************************************
     ( ตัว Func  คือ ตัวเสียบ ตัวผู้ / paramiter คือช่องเสียบตัวเมีย) **********
     - อยากจะเสียบช่องนี้ต้องทำตัวเสียบให้ตรงช่องด้วย โดยใช้ interface ****
    -repository.NewCustomerRepositoryDB(db)

**********************************************************
    -ตอบคำถามที่ Func New ของแต่ละที่ จะ return ของอีกที่ 
    - เช่น repo รับ DB return DB -  แต่ interface ที่ return คือ IAccountRepository
    - servie รับ Repo return Repo  -  แต่ interface ที่ return คือ IAccountService
    - handler รับ servie return  ตัวมันเอง เพราะเอาไปใช้งานแล้ว - แต่ interface ที่ return คือ struct ของมันเอง
-Return ของ Func New ทำหน้าที่ instrand struct ขึ้นมา สร้าง struct ให้จะไม่ให้เข้าถึงโดยตรง ***************
    -ทำไมถึง return ทั้ง struct  ของมัน และ interface ได้ ****
        -มัน return struct ที่ confrom ตาม interface ****
        -มัน ถึง return ทั้ง struct หรือ interface ของมันเอง ได้ โดยที่ไม่ error *************
        -ที่ต้อง return เป็น interface เพราะจะเช็คเรื่อง confrom  ด้วย  
            -return ตาม struct Func Reciver ไม่ต้อง comfrom ก็ได้ แต่จะถูกบังคับด้วย port  *****
            - แต่ layer  จะ return struct เพราะไม่ต้อง confrom ตามใคร