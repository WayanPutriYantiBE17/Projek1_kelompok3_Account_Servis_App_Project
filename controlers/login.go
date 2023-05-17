package controlers

import (
	"database/sql"
	"fmt"
	"projek1/account_management/structs"

	_ "github.com/go-sql-driver/mysql"
)

func GetLoginUsers(db *sql.DB, No_telp string, password string) (bool, error, []structs.Users) {
	// Query the database to check if the user exists
	query := "SELECT COUNT(*) FROM user WHERE No_telepon = ? AND password = ?"
	var count int
	var data structs.Users
	var data1 []structs.Users
	err := db.QueryRow(query, No_telp, password).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("Login: %v", err), data1
	}

	if count == 1 {
		// User exists, retrieve user data
		userQuery := "SELECT Nama, No_telepon, password FROM user WHERE No_telepon = ?"
		err := db.QueryRow(userQuery, No_telp).Scan(&data.Nama, &data.No_tlp, &data.Pasword)
		if err != nil {
			return false, fmt.Errorf("Login: failed to retrieve user data: %v", err), data1
		}

		data1 = append(data1, data)

		return true, nil, data1
	}

	return false, nil, data1
}
