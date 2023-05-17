package controlers

import (
	"database/sql"
	"fmt"

	"projek1/account_management/structs"

	_ "github.com/go-sql-driver/mysql"
)

func GetProfileByID(db *sql.DB, userID int) (*structs.Users, error) {
	// Query the database to retrieve user profile
	query := "SELECT id, nama, No_telepon FROM user WHERE id = ?"
	row := db.QueryRow(query, userID)

	// Initialize a User struct to store the retrieved data
	user := &structs.Users{}
	err := row.Scan(&user.ID, &user.Nama, &user.No_tlp)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("GetProfileByID: User with ID %d not found", userID)
		}
		return nil, fmt.Errorf("GetProfileByID: %v", err)
	}

	return user, nil
}
