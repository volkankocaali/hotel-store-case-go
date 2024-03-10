package _interfacepackage

import (
	"github.com/golang-jwt/jwt"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
)

type Helper interface {
	GenerateTokenClients(admin schema.UserSchemaResponse) (string, error)
	GeneratePasswordHash(password string) (string, error)
	CompareHashAndPassword(a string, b string) error
	GenerateReferenceCode() string
	ParseToken(tokenString string) (jwt.MapClaims, error)
}
