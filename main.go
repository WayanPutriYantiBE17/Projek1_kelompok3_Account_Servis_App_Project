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

	fmt.Println("Menu: \n1. Register \n2. Login \n3. Read Account \n4. Delete \n5. Update \n6. Read Profil User Lain \n7. Top Up\n8. Transfer\n9. Histori Top Up \n10. Histori Transfer\n11. Logout")
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
				fmt.Printf("\n\nBerikut adalah data user Account\n")
				for _, user := range UserData {
					fmt.Printf("Nama: %s\nNo Telepon: %s\nPassword: %s\nTanggal Lahir: %s", user.Nama, user.No_tlp, user.Pasword, user.Tgl_lahir)
				}
			} else {
				fmt.Println("\nTidak ada data user ditemukan")
			}
		}
	case 4:
		fmt.Println("Masukkan Nomor telepon yang akan dihapus:")
		var No_telepon string
		fmt.Scanln(&No_telepon)

		fmt.Println("Masukkan Password Anda:")
		var password string
		fmt.Scanln(&password)

		// Panggil fungsi DeleteUserByID
		err := controlers.DeleteUser(db, No_telepon, password)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Pengguna Nomor Telepon", No_telepon, "telah dihapus.")
		}

	case 5:
		var newnama structs.Users
		var Nomor_telefon string
		var password string
		fmt.Println("Masukkan Nomor Teleon Anda:")
		fmt.Scanln(&Nomor_telefon)
		fmt.Println("Masukkan Password Anda:")
		fmt.Scanln(&password)
		fmt.Println("Masukkan Nama User baru:")
		fmt.Scanln(&newnama.Nama)
		fmt.Println("Masukkan No Telepon baru:")
		fmt.Scanln(&newnama.No_tlp)
		fmt.Println("Masukkan Password baru:")
		fmt.Scanln(&newnama.Pasword)
		fmt.Println("Masukkan Tanggal Lahir baru:")
		fmt.Scanln(&newnama.Tgl_lahir)

		_, err := controlers.GetUpdateNama(db, newnama, Nomor_telefon, password)
		if err != nil {
			fmt.Println("error", err.Error())
		} else {
			fmt.Println("Update Berhasil")
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
			fmt.Println()
			fmt.Println("ID:", user.ID)
			fmt.Println("Nama:", user.Nama)
			fmt.Println("No Telepon:", user.No_tlp)
			// Tampilkan informasi profil lainnya
		}

	case 7:
		fmt.Println("Masukkan Nomor telepon anda:")
		var No_telepon string
		fmt.Scanln(&No_telepon)
		fmt.Println("Masukkan Nomor password anda:")
		var Password string
		fmt.Scanln(&Password)
		fmt.Println("Masukkan Jumlah top up:")
		var TopUp int
		fmt.Scanln(&TopUp)

		// Panggil fungsi GetProfileByID
		_, err := controlers.TopUP(db, No_telepon, Password, TopUp)
		if err != nil {
			log.Println("Top Up failed:", err)
		} else {
			log.Println("Top Up successful")
		}

	case 8:
		// Ambil input dari pengguna
		fmt.Println("Masukkan Nomor telepon anda: ")
		var No_telepon string
		fmt.Scanln(&No_telepon)
		fmt.Println("Masukkan Nama anda: ")
		var Nama string
		fmt.Scanln(&Nama)
		fmt.Println("Masukkan Nomor password anda: ")
		var Password string
		fmt.Scanln(&Password)
		fmt.Println("Masukkan Nomor telepon penerima: ")
		var No_telfn_penerima string
		fmt.Scanln(&No_telfn_penerima)
		fmt.Println("Masukkan Nama Penerima: ")
		var Nama_penerima string
		fmt.Scanln(&Nama_penerima)
		fmt.Println("Masukkan Jumlah Transfer: ")
		var Transfers int
		fmt.Scanln(&Transfers)

		// Panggil fungsi GetProfileByID
		_, err := controlers.Transfer(db, No_telepon, Password, Nama_penerima, No_telfn_penerima, Transfers)
		if err != nil {
			log.Println("Transfer Gagal:", err)
		} else {
			log.Println("Transfer Berhasil :)")
		}

		_, errr := controlers.PenerimaTrasfer(db, No_telepon, Nama, No_telfn_penerima, Transfers)
		if errr != nil {
			log.Println("Transfer Gagal Diterima:", errr)
		} else {
			log.Println("Transfer Berhasil Diterima :)")
		}

	case 9:
		fmt.Println("Masukkan Nomor telepon anda:")
		var No_telepon string
		fmt.Scanln(&No_telepon)
		fmt.Println("Masukkan Nomor password anda:")
		var Password string
		fmt.Scanln(&Password)

		err := controlers.HistoriTopUP(db, No_telepon, Password)
		if err != nil {
			log.Println("Top Up failed:", err)
		} else {
			log.Println("data Histori")

		}

	case 10:
		fmt.Println("Masukkan Nomor telepon anda:")
		var No_telepon string
		fmt.Scanln(&No_telepon)
		fmt.Println("Masukkan password anda: ")
		var Password string
		fmt.Scanln(&Password)

		err := controlers.HistoriTransfer(db, No_telepon, Password)
		if err != nil {
			log.Println("Transfer Gagal:", err)
		} else {
			log.Println("data History")

		}

		errr := controlers.HistoriTransferPenerima(db, No_telepon, Password)
		if errr != nil {
			log.Println("Transfer Gagal:", errr)
		} else {
			log.Println("data History")

		}
	case 11:
		fmt.Println()
		fmt.Println("Terimakasih telah bertransaksi")
		fmt.Println("Semoga harimu menyenangkan :)")

	default:
		fmt.Println()
		fmt.Println("Input salah")
	}

}
