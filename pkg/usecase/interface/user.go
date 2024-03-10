package _interface

import "github.com/volkankocali/hotel-store-case-go/pkg/schema"

type UserUseCase interface {
	UserSignUp(user schema.UserSchema) (schema.TokenUsers, error)
	UserLogin(user schema.UserLoginSchema) (schema.TokenUsers, error)
	UserRefreshToken(refreshToken string) (schema.TokenUsers, error)
}
