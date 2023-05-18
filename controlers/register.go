package controlers

import (
	"database/sql"
	"fmt"

	//"log"

	"projek1/account_management/structs"

	_ "github.com/go-sql-driver/mysql"
)

func RegisterUser(db *sql.DB, newUser structs.Users) (int64, error) {
    result, err := db.Exec("INSERT INTO user (nama,No_telepon,password,tanggal_lahir) VALUES (?,?,?,?)", newUser.Nama, newUser.No_tlp,newUser.Pasword,newUser.Tgl_lahir)
    if err != nil {
        return 0, fmt.Errorf("AddAlbum: %v", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("AddAlbum: %v", err)
    }


    return id,nil
    }


    // Get the new album's generated ID for the client.

    // Return the new album's ID.
  
