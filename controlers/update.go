package controlers

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"projek1/account_management/structs"

	_ "github.com/go-sql-driver/mysql"
)

func GetUpdateNama(db *sql.DB, newNama structs.Users, No_Tlp string, PW string) (int64, error) {
	var userID int64
	query := "SELECT id FROM user WHERE No_telepon = ? AND password = ?"
	err := db.QueryRow(query, No_Tlp, PW).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Kasus ketika nomor telepon dan password tidak sesuai
			return 0, fmt.Errorf("Delete Gagal !!!: Nomor telepon atau password tidak sesuai")
		}
		// Kasus lainnya
		return 0, fmt.Errorf("Delete: %v", err)
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	updateColumn := []string{}
	nilai := []interface{}{}

	//untuk memeriksa apakah namanya terupdate
	if newNama.Nama != "" {
		updateColumn = append(updateColumn, "nama=?")
		nilai = append(nilai, newNama.Nama)
	}

	//untuk memeriksa apakah no telponnya terupdate
	if newNama.No_tlp != "" {
		updateColumn = append(updateColumn, "No_telepon=?")
		nilai = append(nilai, newNama.No_tlp)
	}

	//untuk memeriksa apakah pw terupdate
	if newNama.Pasword != "" {
		updateColumn = append(updateColumn, "password=?")
		nilai = append(nilai, newNama.Pasword)
	}

	//untuk memeriksa apakah tanggal lahir terupdate
	if newNama.Tgl_lahir != "" {
		updateColumn = append(updateColumn, "tanggal_lahir=?")
		nilai = append(nilai, newNama.Tgl_lahir)
	}

	updateColumnStr := ""
	if len(updateColumn) > 0 {
		updateColumnStr = "SET " + strings.Join(updateColumn, ", ")
	}
	updateColumnStr += ", updated_at=?"
	nilai = append(nilai, currentTime)

	// Menghapus koma dan spasi terakhir dari updateColumns
	// = updateColumn[:len(updateColumn)-2]

	// menambahkan Nomor Telepon User Lama
	nilai = append(nilai, No_Tlp)

	// Membangun pernyataan SQL dengan kolom yang akan diperbarui dan placeholder yang sesuai
	queryy := fmt.Sprintf("UPDATE user %s WHERE No_telepon = ?", updateColumnStr)

	// Menjalankan pernyataan SQL dengan nilai-nilai yang diberikan
	result, err := db.Exec(queryy, nilai...)
	if err != nil {
		return 0, fmt.Errorf("Error executing statement: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Failed to get affected rows: %v", err)
	}

	if rowsAffected != 1 {
		return 0, fmt.Errorf("Update failed, no rows affected")
	}

	return rowsAffected, nil
}
