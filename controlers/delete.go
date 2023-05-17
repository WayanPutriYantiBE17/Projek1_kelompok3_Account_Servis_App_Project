package controlers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteUser(db *sql.DB, userID int) error {
	// Prepare the delete statement
	stmt, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return fmt.Errorf("DeleteUser: %v", err)
	}
	defer stmt.Close()

	// Execute the delete statement
	_, err = stmt.Exec(userID)
	if err != nil {
		return fmt.Errorf("DeleteUser: %v", err)
	}

	return nil
}
