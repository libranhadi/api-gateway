package model

import "fmt"

type User struct {
	Id       string `json:"id" bson:"_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) Validate() error {
	if len(u.Email) == 0 {
		return fmt.Errorf("email field is required")
	}

	if len(u.Password) == 0 {
		return fmt.Errorf("password field is required")
	}

	if len(u.Email) < 3 || len(u.Email) > 100 {
		return fmt.Errorf("email length must be between 3 and 100 characters")
	}

	if len(u.Password) < 6 {
		return fmt.Errorf("password must be bigger than 6 characters")
	}

	return nil
}
