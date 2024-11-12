package services

import (
	"book_online_shop_visa/models"
	"book_online_shop_visa/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// userService interface
type UserService interface{
	RegisterUser (user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

// userService Struct
type userService struct{
	userRepo repository.UserRepository
}

// NewUserService function
func NewUserRepository(userRepo repository.UserRepository) userService{
	return userService{userRepo: userRepo}
}

// RegisterUser function
    func (s userService) RegisterUser (user *models.User) error {
	// validasi input (email tidak boleh kosong)
	if user.Email == "" || user.Name == "" || user.Password == "" {
        return errors.New("email, name, dan password harus diisi lengkap" )
    }
} 

    func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
        var user models.User
        query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1"
        row := r.db.QueryRow(query, email)
    
        if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
            if err == sql.ErrNoRows {
                return nil, fmt.Errorf("user not found")
            }
            return nil, err
        }
    
        return &user, nil
    }

    // cek apakah email sudah terdaftar
    eexistingUser, err := s.userRepo.GetUserByEmail(user.Email)
    if err != nil && err != sql.ErrNoRows {
        return fmt.Errorf("error saat mengecek email: %v", err)
    }
    if existingUser != nil {
        return errors.New("email sudah terdaftar")
    }
    
    // Hash Password 
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("gagal melakukan hashing password: %v", err)
    }
    user.Password = string(hashedPassword)
    
    // simpan user ke repository
    err = s.userRepo.Create(user)
    if err!= nil {
        return fmt.Errorf("gagal menyimpan user: %v", err)
    }
    return nil
}
