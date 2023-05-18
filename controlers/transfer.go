package controlers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Transfer(db *sql.DB, no_tlp string, password string, Nama_penerima string, No_telepon string, Jumlah_Transfer int) (int64, error) {
	// Mengambil ID USER berdasarkan nomor telepon
	var userID int64
	var Jumlah_Saldo int64
	query := "SELECT s.user_id, s.Jumlah_Saldo FROM user u INNER JOIN Saldo s ON u.id = s.user_id WHERE u.No_telepon = ? and u.password= ?"
	err := db.QueryRow(query, no_tlp, password).Scan(&userID, &Jumlah_Saldo)
	if err != nil {
		return 0, fmt.Errorf("Transfer: %v", err)
	}

	if Jumlah_Saldo < int64(Jumlah_Transfer) {
		return 0, fmt.Errorf("Transfer: %v", err)
	} else {
		// Menyimpan histori Transfer di tabel "Transfer"
		queri := "INSERT INTO Transfer (user_id, nama_penerima,Nomor_telefon, Jumlah_Transfer, created_at) VALUES (?, ?, ?, ?, ?)"
		result, err := db.Exec(queri, userID, Nama_penerima, No_telepon, Jumlah_Transfer, time.Now())
		if err != nil {
			return 0, fmt.Errorf("Transfer: %v", err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("Transfer: %v", err)
		}

		// Untuk Menyesuaikan Saldo penerima menjadi balance/ setara
		query = "UPDATE Saldo SET Jumlah_Saldo = Jumlah_Saldo - ? WHERE user_id = ?"
		_, err = db.Exec(query, Jumlah_Transfer, userID)
		if err != nil {
			return 0, fmt.Errorf("Transfer: %v", err)
		}
		return id, nil
	}
	return 0, nil
}

func PenerimaTrasfer(db *sql.DB, No_tlp string, Nama_pengirim string, No_telepon string, Jumlah_Transfer int) (int64, error) {
	// Mengambil ID USER berdasarkan nomor telepon
	// Mencari data penerima
	var penerimaID int64
	queryy := "SELECT s.user_id FROM user u INNER JOIN Saldo s ON u.id = s.user_id WHERE u.No_telepon = ?"
	errr := db.QueryRow(queryy, No_telepon).Scan(&penerimaID)
	if errr != nil {
		return 0, fmt.Errorf("Transfer: %v", errr)
	}

	// Mengecek saldo sesuai pengirim apakah memungkinakan untuk transfer
	var Jumlah_Saldo int64
	querii := "SELECT s.Jumlah_Saldo FROM user u INNER JOIN Saldo s ON u.id = s.user_id WHERE u.No_telepon = ?"
	err := db.QueryRow(querii, No_tlp).Scan(&Jumlah_Saldo)
	if err != nil {
		return 0, fmt.Errorf("Transfer: %v", err)
	}

	// Update saldo Untuk Penerima "Saldo"
	if Jumlah_Saldo < int64(Jumlah_Transfer) {
		return 0, fmt.Errorf("Transfer: %v", err)
	} else {
		// Menyimpan histori Transfer di tabel "Transfer"
		queri := "INSERT INTO TransferPenerima (user_id, nama_pengirim, Nomor_telefon, Jumlah_Transfer, created_at) VALUES (?, ?, ?, ?, ?)"
		result, err := db.Exec(queri, penerimaID, Nama_pengirim, No_telepon, Jumlah_Transfer, time.Now())
		if err != nil {
			return 0, fmt.Errorf("Transfer: %v", err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("Transfer: %v", err)
		}
		queryy = "UPDATE Saldo SET Jumlah_Saldo = Jumlah_Saldo + ? WHERE user_id = ?"
		_, err = db.Exec(queryy, Jumlah_Transfer, penerimaID)
		if err != nil {
			return 0, fmt.Errorf("Transfer: %v", err)
		}
		return id, nil
	}
	return 0, nil
}
