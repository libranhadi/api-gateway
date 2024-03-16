package model

import "fmt"

type Employee struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name"`
}

func (u Employee) Validate() error {
	if len(u.Name) == 0 {
		return fmt.Errorf("name field is required")
	}

	if len(u.Name) < 3 || len(u.Name) > 100 {
		return fmt.Errorf("name length must be between 3 and 100 characters")
	}

	return nil
}
