package controlers

import (
	"database/sql"
	"fmt"
)

func Transfer(db *sql.DB, senderAccountID int, receiverAccountID int, amount float64) error {
	// Check if sender's account has sufficient balance
	senderQuery := "SELECT Jumlah_transfer FROM account WHERE account_id = ?"
	var senderBalance float64
	err := db.QueryRow(senderQuery, senderAccountID).Scan(&senderBalance)
	if err != nil {
		return fmt.Errorf("Transfer: Failed to retrieve sender's balance: %v", err)
	}

	if senderBalance < amount {
		return fmt.Errorf("Transfer: Insufficient balance in sender's account")
	}

	// Perform the transfer
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Transfer: Failed to begin transaction: %v", err)
	}

	// Deduct the amount from sender's account
	updateSenderQuery := "UPDATE account SET saldo = saldo - ? WHERE account_id = ?"
	_, err = tx.Exec(updateSenderQuery, amount, senderAccountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Transfer: Failed to deduct amount from sender's account: %v", err)
	}

	// Add the amount to receiver's account
	updateReceiverQuery := "UPDATE account SET saldo = saldo + ? WHERE account_id = ?"
	_, err = tx.Exec(updateReceiverQuery, amount, receiverAccountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Transfer: Failed to add amount to receiver's account: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Transfer: Failed to commit transaction: %v", err)
	}

	return nil
}
