package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"projek1/account_management/controlers"
	"projek1/account_management/structs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var connectionString = os.Getenv("Db_connectionString")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error open connection", err.Error())
	} else {
		fmt.Println("berhasil")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("ping:", pingErr)
	}
	fmt.Println("connected!")

	defer db.Close()

	fmt.Println("Menu: \n1. Register \n2. Login \n3. Read Account \n4. Delete \n5. Update \n6. Read Profil User Lain \n7. Transfer")
	fmt.Println("masukkan menu")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {

	case 1:
		var newuser structs.Users
		fmt.Println("Masukkan Nama User:")
		fmt.Scanln(&newuser.Nama)
		fmt.Println("Masukkan No Telepon:")
		fmt.Scanln(&newuser.No_tlp)
		fmt.Println("Masukkan Password:")
		fmt.Scanln(&newuser.Pasword)
		fmt.Println("Masukkan Tanggal Lahir:")
		fmt.Scanln(&newuser.Tgl_lahir)

		idNew, err := controlers.RegisterUser(db, newuser)
		if err != nil {
			fmt.Println("error", err.Error())
		} else {
			fmt.Println("Register Sukses", idNew)
		}

	case 2:
		var No_telefon string
		var passwords string
		fmt.Println("Masukkan No Telepon: ")
		fmt.Scanln(&No_telefon)
		fmt.Println("Masukkan Password: ")
		fmt.Scanln(&passwords)

		// Panggil fungsi login
		success, err, UserData := controlers.GetLoginUsers(db, No_telefon, passwords)
		if err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			if success && len(UserData) > 0 {
				fmt.Println("Login berhasil!")
			} else {
				fmt.Println("Login gagal. No_telepon atau password salah.")
			}
		}

	case 3:
		var No_telefon string
		var passwords string
		fmt.Println("Masukkan No Telepon: ")
		fmt.Scanln(&No_telefon)
		fmt.Println("Masukkan Password: ")
		fmt.Scanln(&passwords)

		// Panggil fungsi login
		success, err, UserData := controlers.GetLoginUsers(db, No_telefon, passwords)
		if err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			if success && len(UserData) > 0 {
				fmt.Println("\n\nBerikut adalah data user Account\n")
				for _, user := range UserData {
					fmt.Printf("Nama: %s\nNo Telepon: %s\nPassword: %s\nTanggal Lahir: %s", user.Nama, user.No_tlp, user.Pasword, user.Tgl_lahir)
				}
			} else {
				fmt.Println("\nTidak ada data user ditemukan")
			}
		}
	case 4:
		fmt.Println("Masukkan ID Pengguna yang akan dihapus:")
		var userID int
		fmt.Scanln(&userID)

		// Panggil fungsi DeleteUserByID
		err := controlers.DeleteUser(db, userID)
		if err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			fmt.Println("Pengguna dengan ID", userID, "telah dihapus.")
		}

	case 5:
		var newnama structs.Users
		fmt.Println("Masukkan Nama User baru:")
		fmt.Scanln(&newnama.Nama)
		fmt.Println("Masukkan No Telepon baru:")
		fmt.Scanln(&newnama.No_tlp)
		fmt.Println("Masukkan Password baru:")
		fmt.Scanln(&newnama.Pasword)
		fmt.Println("Masukkan Tanggal Lahir baru:")
		fmt.Scanln(&newnama.Tgl_lahir)
		fmt.Println("Masukkan id:")
		fmt.Scanln(&newnama.ID)
		nameUser, err := controlers.GetUpdateNama(db, newnama)
		if err != nil {
			fmt.Println("error", err.Error())
		} else {
			fmt.Println("", nameUser)
		}

	case 6:
		fmt.Println("Masukkan Nomor Telepon Pengguna:")
		var No_telepon string
		fmt.Scanln(&No_telepon)

		// Panggil fungsi GetProfileByID
		user, err := controlers.GetProfileByID(db, No_telepon)
		if err != nil {
			fmt.Println("Error:", err.Error())
			fmt.Println("Data tidak ditemukan!!!")
		} else {
			fmt.Println("ID:", user.ID)
			fmt.Println("Nama:", user.Nama)
			fmt.Println("No Telepon:", user.No_tlp)
			// Tampilkan informasi profil lainnya
		}

	case 7:
		// Ambil input dari pengguna
		var userID, jumlahTransfer float64
		var namaPenerima, nomorTelepon string

		fmt.Print("Masukkan ID pengirim: ")
		fmt.Scanln(&userID)

		fmt.Print("Masukkan nama penerima: ")
		fmt.Scanln(&namaPenerima)

		fmt.Print("Masukkan nomor telepon penerima: ")
		fmt.Scanln(&nomorTelepon)

		fmt.Print("Masukkan jumlah transfer: ")
		fmt.Scanln(&jumlahTransfer)

		// Panggil fungsi transfer
		err = controlers.Transfer(db, int(userID), namaPenerima, nomorTelepon, jumlahTransfer)
		if err != nil {
			log.Println("Transfer failed:", err)
		} else {
			log.Println("Transfer successful")
		}

	}
}
