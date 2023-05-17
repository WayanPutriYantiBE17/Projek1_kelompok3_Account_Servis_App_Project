package controlers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Transfer(db *sql.DB, userID int, namaPenerima string, nomorTelepon string, jumlahTransfer float64) error {
	// Check if the sender's balance is sufficient
	senderBalance, err := getBalance(db, userID)
	if err != nil {
		return fmt.Errorf("Transfer: Failed to get sender's balance: %v", err)
	}

	if senderBalance < jumlahTransfer {
		return fmt.Errorf("Transfer: Insufficient balance")
	}

	// Start the database transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Transfer: Failed to begin transaction: %v", err)
	}

	// Deduct the transfer amount from the sender's balance
	if err := deductBalance(tx, userID, jumlahTransfer); err != nil {
		tx.Rollback()
		return fmt.Errorf("Transfer: Failed to deduct balance: %v", err)
	}

	// Add the transfer amount to the receiver's balance
	receiverID, err := getUserIDByPhoneNumber(tx, nomorTelepon)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Transfer: Failed to get receiver's ID: %v", err)
	}

	if err := addBalance(tx, receiverID, jumlahTransfer); err != nil {
		tx.Rollback()
		return fmt.Errorf("Transfer: Failed to add balance: %v", err)
	}

	// Insert the transfer record
	if err := insertTransferRecord(tx, userID, namaPenerima, nomorTelepon, jumlahTransfer); err != nil {
		tx.Rollback()
		return fmt.Errorf("Transfer: Failed to insert transfer record: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("Transfer: Failed to commit transaction: %v", err)
	}

	return nil
}

// Helper function to get the user's balance
func getBalance(db *sql.DB, userID int) (float64, error) {
	query := "SELECT Jumlah_Saldo FROM Saldo WHERE user_id = ?"
	var balance float64
	err := db.QueryRow(query, userID).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("getBalance: User with ID %d not found", userID)
		}
		return 0, err
	}
	return balance, nil
}

// Helper function to deduct the transfer amount from the sender's balance
func deductBalance(tx *sql.Tx, userID int, amount float64) error {
	query := "UPDATE Saldo SET Jumlah_Saldo = Jumlah_Saldo - ? WHERE user_id = ?"
	_, err := tx.Exec(query, amount, userID)
	if err != nil {
		return err
	}
	return nil
}

// Helper function to add the transfer amount to the receiver's balance
func addBalance(tx *sql.Tx, userID int, amount float64) error {
	query := "UPDATE Saldo SET Jumlah_Saldo = Jumlah_Saldo + ? WHERE user_id = ?"
	_, err := tx.Exec(query, amount, userID)
	if err != nil {
		return err
	}
	return nil
}

// Helper function to get the user ID based on the phone number
func getUserIDByPhoneNumber(tx *sql.Tx, phoneNumber string) (int, error) {
	query := "SELECT id FROM user WHERE Nomor_telefon = ?"
	var userID int
	err := tx.QueryRow(query, phoneNumber).Scan(&userID)
	if err != nil {
		return userID, err
	}
	return userID, nil
}

// Helper function
// to insert the transfer record
func insertTransferRecord(tx *sql.Tx, userID int, namaPenerima string, nomorTelepon string, jumlahTransfer float64) error {
	query := "INSERT INTO Transfer(user_id, nama_penerima, Nomor_telefon, Jumlah_Transfer) VALUES (?, ?, ?, ?)"
	_, err := tx.Exec(query, userID, namaPenerima, nomorTelepon, jumlahTransfer)
	if err != nil {
		return err
	}
	return nil
}
