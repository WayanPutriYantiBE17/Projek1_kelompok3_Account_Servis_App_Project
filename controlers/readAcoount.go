package controlers

import (
	"database/sql"
	//"fmt"

	"log"

	"projek1/account_management/structs"

	_ "github.com/go-sql-driver/mysql"
)

func GetUser(db *sql.DB) []structs.Users {
	rows, errSelect := db.Query("SELECT nama, No_telepon, password, tanggal_lahir from user")
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var allUser []structs.Users // menampung semua data
	for rows.Next() {                 // proses pembacaan per baris
		var user structs.Users                                                                                // menampung data per baris nya
		errScan := rows.Scan(&user.Nama, &user.No_tlp, &user.Pasword, &user.Tgl_lahir) // mapping
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}

		allUser = append(allUser, user)
	}

	// fmt.Println("all:", allProduct)
	return allUser
}