package controlers

import (
	"database/sql"
	"fmt"

	"projek1/account_management/structs"

	_ "github.com/go-sql-driver/mysql"
)

func GetProfileByID(db *sql.DB, No_telepon string) (*structs.Users, error) {
	// Query the database to retrieve user profile
	query := "SELECT id, nama, No_telepon FROM user WHERE No_telepon = ?"
	row := db.QueryRow(query, No_telepon)

	// Initialize a User struct to store the retrieved data
	user := &structs.Users{}
	err := row.Scan(&user.ID, &user.Nama, &user.No_tlp)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("GetProfileByID: User with No_telepon %d not found", No_telepon)
		}
		return nil, fmt.Errorf("GetProfileByID: %v", err)
	}

	return user, nil
}
