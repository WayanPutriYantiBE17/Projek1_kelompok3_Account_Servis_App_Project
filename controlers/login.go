package controlers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetLoginUsers(db *sql.DB, No_telp string, password string) (bool, error) {
	// Query the database to check if the user exists
	query := "SELECT COUNT(*) FROM user WHERE No_telepon = ? AND password = ?"
	var count int
	err := db.QueryRow(query, No_telp, password).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("Login: %v", err)
	}

	if count == 1 {
		return true, nil
	}

	return false, nil
}
