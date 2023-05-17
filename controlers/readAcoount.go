package controlers

// kami menggunakan file login untuk membaca account user yang login, jadi file ini tidak dipakai
//
// import (
// 	"database/sql"

// 	"projek1/account_management/structs"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func GetLoginUser(db *sql.DB,No_telefon, passwords string) ([]structs.Users,error) {

// 	query := "SELECT * FROM user WHERE No_telefon=? and passwords=?"
// 	rows, err := db.Query(query, No_telefon, passwords)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var userData []structs.Users
// 	for rows.Next() {
// 		var user structs.Users
// 		err := rows.Scan(&user.ID, &user.Nama, &user.No_tlp, &user.Pasword, &user.Tgl_lahir)
// 		if err != nil {
// 			return nil, err
// 		}
// 		userData = append(userData, user)
// 	}

// 	err = rows.Err()
// 	if err != nil{
// 		return nil, err
// 	}
// 	return userData,nil
// }