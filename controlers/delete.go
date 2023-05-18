package controlers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteUser(db *sql.DB, No_telepon string, password string) error {
	// Prepare the delete statement
	var userID int64
	query := "SELECT id FROM user WHERE No_telepon = ? AND password = ?"
	err := db.QueryRow(query, No_telepon, password).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Kasus ketika nomor telepon dan password tidak sesuai
			return fmt.Errorf("Delete Gagal !!!: Nomor telepon atau password tidak sesuai")
		}
		// Kasus lainnya
		return fmt.Errorf("Delete: %v", err)
	}

	// Prepare the delete statement
	stmt, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return fmt.Errorf("Delete: %v", err)
	}
	defer stmt.Close()

	// Execute the delete statement
	_, err = stmt.Exec(userID)
	if err != nil {
		return fmt.Errorf("Delete: %v", err)
	}
	return nil
}
