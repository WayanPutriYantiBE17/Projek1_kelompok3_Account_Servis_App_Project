package controlers

import (
	"database/sql"
	"fmt"
	"log"
	"projek1/account_management/structs"

	_ "github.com/go-sql-driver/mysql"
)


func HistoriTopUP(db *sql.DB, no_tlp string, password string) error {
	// Mengambil ID pengguna berdasarkan nomor telepon
	var newUser structs.Users
	var newTop_Ups structs.Top_Ups

	query := "SELECT u.nama,t.Jumlah_Topup,t.created_at FROM user u INNER JOIN top_Up t ON u.id = t.user_id WHERE u.No_telepon = ? and u.password=?"
	rows, err := db.Query(query, no_tlp, password)
	if err != nil {
		log.Fatal("error query select", err.Error())
	}

	var allHistory []interface{} // menampung semua data
	for rows.Next() {            // proses pembacaan per baris                                                                              // menampung data per baris nya
		errScan := rows.Scan(&newUser.Nama, &newTop_Ups.Jumlah_Topup, &newTop_Ups.Created_at) // mapping
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}

		allHistory = append(allHistory, newUser)
		allHistory = append(allHistory,newTop_Ups)

		fmt.Println()
		fmt.Println("Nama Pengguna:", newUser.Nama)
		fmt.Println("Jumlah TopUp:", newTop_Ups.Jumlah_Topup)
		fmt.Println("Tanggal TopUp:", newTop_Ups.Created_at)
	}
	return nil
}