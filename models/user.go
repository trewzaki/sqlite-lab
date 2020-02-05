package models

import "time"

// User model
type User struct {
	ID          uint32     `gorm:"primary_key;AUTO_INCREMENT" json:"user_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	PhoneNumber *string    `json:"phone_number"`
	DeletedAt   *time.Time `gorm:"DEFAULT:NULL" json:"deleted_at"`
}

// Create model function TODO
func (user *User) Create() error {
	return DB.Create(user).Error
}

// Update model function TODO
func (user *User) Update() error {
	return nil
}

// Get model function TODO
func (user *User) Get() ([]User, error) {
	return nil, nil
}

// GetAll model function
func (user *User) GetAll() ([]User, error) {
	userModel := []User{}
	if err := DB.Find(&userModel).Error; err != nil {
		return nil, err
	}

	return userModel, nil
}

// Delete model function TODO
func (user *User) Delete() error {
	return nil
}

// UnDelete model function TODO
func (user *User) UnDelete() error {
	return nil
}
