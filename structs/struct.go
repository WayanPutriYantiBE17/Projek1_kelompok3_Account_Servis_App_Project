package structs

import "database/sql"

type Users struct {
	ID         uint
	Nama       string
	No_tlp     string
	Pasword    string
	Tgl_lahir  string
	//Read_at    sql.NullTime
	Created_at sql.NullTime
	Updated_at sql.NullTime
	Deleted_at sql.NullTime
}

type Top_Ups struct {
	ID           uint
	User_id      int
	Jumlah_Topup float64
	Read_at      sql.NullTime
	Created_at   sql.NullTime
	Updated_at   sql.NullTime
	Deleted_at   sql.NullTime
}

type Saldos struct {
	User_id      int
	Jumlah_Saldo float64
	Read_at      sql.NullTime
	Created_at   sql.NullTime
	Updated_at   sql.NullTime
	Deleted_at   sql.NullTime
}

// type Time struct{
// 	read_at time
// 	created_at time
// 	updated_at time.Time
// 	deleted_at time.Time
// }

type Transfers struct {
	ID            uint
	User_id       int
	Nama_penerima string
	No_tlp        string
	Jumlah_Saldo  float64
	Read_at       sql.NullTime
	Created_at    sql.NullTime
	Updated_at    sql.NullTime
	Deleted_at    sql.NullTime
}
