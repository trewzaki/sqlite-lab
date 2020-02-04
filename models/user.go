package models

type User struct {
	ID          uint32  `gorm:"primary_key;AUTO_INCREMENT" json:"user_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phone_number"`
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
