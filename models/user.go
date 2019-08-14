package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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
