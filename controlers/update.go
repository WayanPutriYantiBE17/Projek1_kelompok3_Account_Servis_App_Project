package controlers

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"projek1/account_management/structs"

	_ "github.com/go-sql-driver/mysql"
)
func GetUpdateNama(db *sql.DB, newNama structs.Users) (int64, error) {

    currentTime := time.Now().Format("2006-01-02 15:04:05")

    updateColumn :=[]string{}
    nilai :=[]interface{}{}

    //untuk memeriksa apakah namanya terupdate
    if newNama.Nama != "" {
        updateColumn =append(updateColumn, "nama=?")
        nilai = append(nilai, newNama.Nama)
    }

    //untuk memeriksa apakah no telponnya terupdate
    if newNama.No_tlp != "" {
        updateColumn = append(updateColumn,"No_telepon=?")
        nilai = append(nilai, newNama.No_tlp)
    }

    //untuk memeriksa apakah pw terupdate
    if newNama.Pasword != "" {
        updateColumn = append(updateColumn,"password=?")
        nilai = append(nilai, newNama.Pasword)
    }

    //untuk memeriksa apakah tanggal lahir terupdate
    if newNama.Tgl_lahir != "" {
        updateColumn = append(updateColumn,"tanggal_lahir=?")
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
    
    // menambahkan id data
    nilai = append(nilai, newNama.ID)


	// Membangun pernyataan SQL dengan kolom yang akan diperbarui dan placeholder yang sesuai
	query := fmt.Sprintf("UPDATE user %s WHERE id = ?", updateColumnStr)

	// Menjalankan pernyataan SQL dengan nilai-nilai yang diberikan
	result, err := db.Exec(query,nilai...)
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