package models

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint64
	UserName string
	Email    string
	Password string
}

func UserFactory() *User {
	return &User{
		UserName: fmt.Sprintf(faker.FirstName()),
		Email:    fmt.Sprintf(faker.Email()),
		Password: fmt.Sprintf(faker.Word()),
	}
}
