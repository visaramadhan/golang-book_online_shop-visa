package repository

import (
    "database/sql"
    "book_online_shop_visa/models"
)
//userRepository interface
type UserRepository interface {
    CreateUser(user *models.User) error
    GetUserByID(id int) (*models.User, error)
}
//userRepository struct
type userRepository struct {
    db *sql.DB
}

//NewUserRepository creates a new userRepository
func NewUserRepository(db *sql.DB) userRepository {
    return userRepository{db: db}
}

//CreateUser inserts a new user into the database
func (r *userRepository) CreateUser (user *models.User) error{
    query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
    err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
    if err!= nil {
        return err
    }
    return nil
}

//GetUserByID unruk mengambil data user berdasarkan ID
func (r *userRepository) GetUserByID(id int) (*models.User, error){
    query := `SELECT id, name, email, password FROM users WHERE id = $1`
    user := &models.User{}
    err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err!= nil {
        return nil, err
    }
    return user, nil
}