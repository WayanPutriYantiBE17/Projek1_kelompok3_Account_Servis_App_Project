package controlers

import (
	"database/sql"
	"fmt"
	"log"
	"projek1/account_management/structs"

	_ "github.com/go-sql-driver/mysql"
)

func HistoriTransfer(db *sql.DB, no_tlp string, password string) error {
	// Mengambil ID pengguna berdasarkan nomor telepon
	var newUser structs.Users
	var newTransfer structs.Transfers

	query := "SELECT u.nama,t.Jumlah_Transfer,t.nama_penerima,t.created_at FROM user u INNER JOIN Transfer t ON u.id = t.user_id WHERE u.No_telepon = ? and u.password=?"
	rows, err := db.Query(query, no_tlp, password)
	if err != nil {
		log.Fatal("error query select", err.Error())
	}

	var allHistory []interface{} // menampung semua data
	for rows.Next() {            // proses pembacaan per baris                                                                              // menampung data per baris nya
		errScan := rows.Scan(&newUser.Nama, &newTransfer.Jumlah_Transfer, &newTransfer.Nama_penerima, &newTransfer.Created_at) // mapping
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}

		allHistory = append(allHistory, newUser)
		allHistory = append(allHistory, newTransfer)

		fmt.Println()
		fmt.Println("Data Pengirim")
		fmt.Println("Nama Pengguna:", newUser.Nama)
		fmt.Println("Jumlah Transfer:", newTransfer.Jumlah_Transfer)
		fmt.Println("Tanggal Transfer:", newTransfer.Created_at)
	}
	return nil
}

func HistoriTransferPenerima(db *sql.DB, no_tlp string, password string) error {
	// Mengambil ID pengguna berdasarkan nomor telepon
	var newUser structs.Users
	var newTransfer structs.TransfersPenerima

	query := "SELECT u.nama,t.Jumlah_Transfer,t.nama_pengirim,t.created_at FROM user u INNER JOIN TransferPenerima t ON u.id = t.user_id WHERE u.No_telepon = ? and u.password=?"
	rows, err := db.Query(query, no_tlp, password)
	if err != nil {
		log.Fatal("error query select", err.Error())
	}

	var allHistory []interface{} // menampung semua data
	for rows.Next() {            // proses pembacaan per baris                                                                              // menampung data per baris nya
		errScan := rows.Scan(&newUser.Nama, &newTransfer.Jumlah_Transfer, &newTransfer.Nama_pengirim, &newTransfer.Created_at) // mapping
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}

		allHistory = append(allHistory, newUser)
		allHistory = append(allHistory, newTransfer)

		fmt.Println()
		fmt.Println("Data Penerima")
		fmt.Println("Nama Pengguna:", newUser.Nama)
		fmt.Println("Nama Pengirim:", newTransfer.Nama_pengirim)
		fmt.Println("Jumlah Transfer:", newTransfer.Jumlah_Transfer)
		fmt.Println("Tanggal Transfer:", newTransfer.Created_at)
	}
	return nil
}
