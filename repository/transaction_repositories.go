package repository

import (
	"book_online_shop_visa/models"
	"database/sql"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindByID(id int) (*models.Transaction, error)
	FindByUserID(userID int) ([]*models.Transaction, error)
	Delete(id int) error
	Update(transaction *models.Transaction) error
	Close() error
	CountByUserID(userID int) (int, error)
	CountByStatus(status string) (int, error)
	SumByUserID(userID int) (float64, error)
	SumByStatus(status string) (float64, error)
	FindByUserIDAndStatus(userID int, status string) ([]models.Transaction, error)
	FindByStatusAndDate(status string, startDate, endDate string) ([]models.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) transactionRepository {
	return transactionRepository{db: db}
}

func (r transactionRepository) Create(transaction *models.Transaction) error {
	sqlStatement := `INSERT INTO transactions (user_id, final_amount, status) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(sqlStatement, transaction.UserID, transaction.FinalAmount, transaction.Status)
	return err
}

func (r transactionRepository) FindByID(id int) (*models.Transaction, error) {
	sqlStatement := `SELECT id, user_id, final_amount, status FROM transactions WHERE id=$1`
	var transaction models.Transaction
	err := r.db.QueryRow(sqlStatement, id).Scan(&transaction.ID, &transaction.UserID, &transaction.FinalAmount, &transaction.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r transactionRepository) FindByUserID(userID int) ([]*models.Transaction, error) {
	sqlStatement := `SELECT id, user_id, final_amount, status FROM transactions WHERE user_id=$1`
	rows, err := r.db.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.FinalAmount, &transaction.Status)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}

func (r transactionRepository) Delete(id int) error {
	sqlStatement := `DELETE FROM transactions WHERE id=$1`
	_, err := r.db.Exec(sqlStatement, id)
	return err
}

func (r transactionRepository) Update(transaction *models.Transaction) error {
	sqlStatement := `UPDATE transactions SET user_id=$1, final_amount=$2, status=$3 WHERE id=$4`
	_, err := r.db.Exec(sqlStatement, transaction.UserID, transaction.FinalAmount, transaction.Status, transaction.ID)
	return err
}

func (r transactionRepository) Close() error {
	return r.db.Close()
}
