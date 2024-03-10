package _interface

import (
	"github.com/volkankocali/hotel-store-case-go/pkg/models"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
)

type UserRepository interface {
	CheckUserExist(email string) bool
	FindUserByEmail(email string) (models.Users, error)
	SignUp(user schema.UserSchema) (schema.UserSchemaResponse, error)
	Create(user models.Users) (models.Users, error)
}
