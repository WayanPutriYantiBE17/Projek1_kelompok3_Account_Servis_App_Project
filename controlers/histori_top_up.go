package controlers

import (
	"database/sql"
	"fmt"
	"projek1/account_management/structs"

	_ "github.com/go-sql-driver/mysql"
)


func HistoriTopUP(db *sql.DB, no_tlp string, password string) error {
	// Mengambil ID pengguna berdasarkan nomor telepon
//var newUser structs.Users
var newTopUp structs.Top_Ups
query := "SELECT t.id_user,t. FROM user u INNER JOIN top_Up t ON u.id = t.user_id WHERE u.No_telepon = ? and u.password=?"
err := db.QueryRow(query, no_tlp,password).Scan(&newTopUp.User_id,&newTopUp.Jumlah_Topup)
if err != nil {
	return fmt.Errorf("HistoriTopUP: %v", err)
}
return nil
}