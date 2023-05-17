package controlers

import (
	"database/sql"
	"fmt"

	"time"

	_ "github.com/go-sql-driver/mysql"
)



func TopUP(db *sql.DB, no_tlp string, password string, jumlah_topup int) (int64, error) {
	// Mengambil ID pengguna berdasarkan nomor telepon
	var userID int64
	query := "SELECT s.user_id FROM user u INNER JOIN Saldo s ON u.id = s.user_id WHERE u.No_telepon = ? and u.password=?"
	err := db.QueryRow(query, no_tlp,password).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("TopUP: %v", err)
	}

	if err != nil {
		return 0, fmt.Errorf("TopUP: %v", err)
	}

	// Menyimpan histori top-up di tabel "topup"
	queri := "INSERT INTO top_Up (user_id, Jumlah_Topup, created_at) VALUES (?, ?, ?)"
	result, err := db.Exec(queri, userID, jumlah_topup, time.Now())
	if err != nil {
		return 0, fmt.Errorf("TopUP: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("TopUP: %v", err)
	}

	// Update saldo dalam tabel "Saldo"
	query = "UPDATE Saldo SET Jumlah_Saldo = Jumlah_Saldo + ? WHERE user_id = ?"
	_, err = db.Exec(query, jumlah_topup, userID)
	if err != nil {
		return 0, fmt.Errorf("TopUP: %v", err)
	}

	return id, nil
}