package models

import "time"

type User struct {
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	PhoneNumber *string    `json:"phone_number"`
	CreatedAt   time.Time  `gorm:"DEFAULT:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"DEFAULT:now()" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"DEFAULT:NULL" json:"deleted_at"`
}

func (user *User) Create() error {
	return DB.Create(user).Error
}

func (user *User) GetAll() ([]User, error) {
	userModel := []User{}
	if err := DB.Find(&userModel).Error; err != nil {
		return nil, err
	}

	return userModel, nil
}
