package repository

import (
	"github.com/volkankocali/hotel-store-case-go/pkg/models"
	interfaces "github.com/volkankocali/hotel-store-case-go/pkg/repository/interface"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (u *userDatabase) CheckUserExist(email string) bool {
	var user models.Users
	u.DB.Where("email = ?", email).First(&user)

	if user.ID != 0 {
		return true
	}

	return false
}

func (u *userDatabase) FindUserByEmail(email string) (models.Users, error) {
	var user models.Users
	u.DB.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return models.Users{}, nil
	}

	return user, nil
}

func (u *userDatabase) SignUp(user schema.UserSchema) (schema.UserSchemaResponse, error) {
	createdUser := models.Users{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}

	err := u.DB.Create(&createdUser).Error

	if err != nil {
		return schema.UserSchemaResponse{}, err
	}

	return schema.UserSchemaResponse{
		Id:    createdUser.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}, nil
}

func (u *userDatabase) Create(user models.Users) (models.Users, error) {
	err := u.DB.Create(&user).Error

	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}
