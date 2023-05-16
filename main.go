package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"projek1/account_management/controlers"
	"projek1/account_management/structs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)



func main(){
	var connectionString = os.Getenv("Db_connectionString")
	db,err :=sql.Open("mysql",connectionString)
	if err !=nil{
		log.Fatal("error open connection",err.Error())
	}else{
		fmt.Println("berhasil")
	}

	db.SetConnMaxLifetime(time.Minute*3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	pingErr := db.Ping()
	if pingErr !=nil{
		log.Fatal("ping:",pingErr)
	}
	fmt.Println("connected!")

	defer db.Close()
	
	fmt.Println("Menu: \n1. Register")
	fmt.Println("masukkan menu")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	// case 1:
	// 	dataProducts := controllers.GetAllProducts(db)
	// 	// fmt.Println(dataProducts)
	// 	for _, value := range dataProducts {
	// 		fmt.Printf("id: %d, nama: %s, harga: %d \n", value.ID, value.NamaProduk, value.Harga)
	// 	}

	case 1:
		var newuser structs.Users
		fmt.Println("Masukkan Nama User:")
		fmt.Scanln(&newuser.Nama)
		fmt.Println("Masukkan No Telepon:")
		fmt.Scanln(&newuser.No_tlp)
		fmt.Println("Masukkan Password:")
		fmt.Scanln(&newuser.Pasword)
		fmt.Println("Masukkan Tanggal Lahir:")
		fmt.Scanln(&newuser.Tgl_lahir )

	idNew, err := controlers.RegisterUser(db,newuser)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println("id porduct", idNew)
	}
}
}