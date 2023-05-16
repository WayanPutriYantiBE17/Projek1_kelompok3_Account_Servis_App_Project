package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type user struct{
	ID uint
	nama string
	no_tlp string
	pasword string
	tgl_lahir string
	read_at  sql.NullTime
	created_at  sql.NullTime
	updated_at  sql.NullTime
	deleted_at  sql.NullTime
}

type top_Up struct{
	ID uint
	user_id int
	Jumlah_Topup float64
	read_at time.Time
	created_at time.Time `sql:"DEFAULT:current_timestamp"`
	updated_at time.Time 
	deleted_at time.Time 
}

type Saldo struct{
	user_id int
	Jumlah_Saldo float64
	read_at time.Time
	created_at time.Time
	updated_at time.Time 
	deleted_at time.Time 
}
// type Time struct{
// 	read_at time
// 	created_at time
// 	updated_at time.Time 
// 	deleted_at time.Time 
// }

type Transfer struct{
	ID uint
	user_id int
	nama_penerima string
	no_tlp string
	Jumlah_Saldo float64
	read_at time.Time
	created_at time.Time `sql:"DEFAULT:current_timestamp"`
	updated_at time.Time 
	deleted_at time.Time 
}


func main(){
	var connectionString = os.Getenv("Db_connectionString")
	db,err :=sql.Open("mysql",connectionString)
	if err !=nil{
		log.Fatal("error open connection",err.Error())
	}else{
		fmt.Println("berhasil")
	}

	db.SetConnMaxLifetime(time.Minute*3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	pingErr := db.Ping()
	if pingErr !=nil{
		log.Fatal("ping:",pingErr)
	}
	fmt.Println("connected!")

	defer db.Close()
	var newUser = user{
		nama: "Sandi",
		no_tlp: "08567889230",
		pasword:"891237",
		tgl_lahir: "13 April 2013",
	}
	idNew, err := RegisterUser(db, newUser)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println("id porduct", idNew)
	}
}
func RegisterUser(db *sql.DB, newUser user) (int64, error) {
    result, err := db.Exec("INSERT INTO user (nama,No_telepon,password,tanggal_lahir) VALUES (?, ?,?,?)", newUser.nama, newUser.no_tlp,newUser.pasword,newUser.tgl_lahir)
    if err != nil {
        return 0, fmt.Errorf("AddAlbum: %v", err)
    }

    // Get the new album's generated ID for the client.
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("AddAlbum: %v", err)
    }
    // Return the new album's ID.
    return id, nil
}