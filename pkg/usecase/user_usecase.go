package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/volkankocali/hotel-store-case-go/pkg/config"
	constant "github.com/volkankocali/hotel-store-case-go/pkg/constant/middleware"
	user_constant "github.com/volkankocali/hotel-store-case-go/pkg/constant/user"
	helper_interface "github.com/volkankocali/hotel-store-case-go/pkg/helper/interface"
	interfaces "github.com/volkankocali/hotel-store-case-go/pkg/repository/interface"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
	"strings"
)

type UserUseCase struct {
	userRepository interfaces.UserRepository
	config         config.Config
	helper         helper_interface.Helper
}

func NewUserUseCase(userRepository interfaces.UserRepository, cfg *config.Config, h helper_interface.Helper) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
		config:         *cfg,
		helper:         h,
	}
}

func (u *UserUseCase) UserSignUp(user schema.UserSchema) (schema.TokenUsers, error) {
	// first check user already exist
	userCheck := u.userRepository.CheckUserExist(user.Email)
	if userCheck {
		return schema.TokenUsers{}, errors.New(user_constant.AlreadyExist)
	}

	// check user password and confirm password
	if user.Password != user.PasswordConfirmation {
		return schema.TokenUsers{}, errors.New(user_constant.PasswordNotMatch)
	}

	// hashed password
	hashedPassword, err := u.helper.GeneratePasswordHash(user.Password)
	if err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.HashError)
	}

	user.Password = hashedPassword

	// create user and save database
	userCreate, err := u.userRepository.SignUp(user)
	if err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.SignUpError)
	}

	// create jwt token
	tokenClients, err := u.helper.GenerateTokenClients(userCreate)
	if err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.TokenError)
	}

	return schema.TokenUsers{
		Users: userCreate,
		Token: tokenClients,
	}, nil
}

func (u *UserUseCase) UserLogin(user schema.UserLoginSchema) (schema.TokenUsers, error) {
	// first check user already exist
	userCheck := u.userRepository.CheckUserExist(user.Email)
	if !userCheck {
		return schema.TokenUsers{}, errors.New(user_constant.NotFound)
	}

	// check user email
	userDetails, err := u.userRepository.FindUserByEmail(user.Email)
	if err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.NotFound)
	}

	// check find db password and request password compare
	err = u.helper.CompareHashAndPassword(userDetails.Password, user.Password)
	if err != nil {
		return schema.TokenUsers{}, errors.New(user_constant.PasswordNotCorrect)
	}

	// user schema
	userSchema := schema.UserSchemaResponse{
		Id:    userDetails.ID,
		Name:  userDetails.Name,
		Email: userDetails.Email,
		Phone: userDetails.Phone,
	}

	// generate token clients
	tokenString, err := u.helper.GenerateTokenClients(userSchema)

	return schema.TokenUsers{
		Token: tokenString,
		Users: userSchema,
	}, nil
}

func (u *UserUseCase) UserRefreshToken(refreshToken string) (schema.TokenUsers, error) {
	refreshToken = strings.TrimPrefix(refreshToken, "Bearer ")

	// Parse and validate the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("ecommercegoapplication"), nil
	})

	if err != nil || !token.Valid {
		error.Error(errors.New(constant.InvalidRefreshToken))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		error.Error(errors.New(constant.InvalidRefreshToken))
	}

	// Retrieve user details from the refresh token
	id, ok := claims["id"].(float64)
	if !ok || id == 0 {
		error.Error(errors.New(constant.ErrorInRetrievingId))
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		error.Error(errors.New(constant.ErrorInRetrievingEmail))
	}

	// Generate a new access token
	user := schema.UserSchemaResponse{
		Id:    uint(id),
		Email: email,
	}
	accessToken, err := u.helper.GenerateTokenClients(user)
	if err != nil {
		error.Error(errors.New(constant.ErrorInGeneratingToken))
	}

	return schema.TokenUsers{
		Token: accessToken,
		Users: schema.UserSchemaResponse{},
	}, nil
}
