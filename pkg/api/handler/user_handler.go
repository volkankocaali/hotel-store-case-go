package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	constant "github.com/volkankocali/hotel-store-case-go/pkg/constant/middleware"
	request_constant "github.com/volkankocali/hotel-store-case-go/pkg/constant/request"
	user_constant "github.com/volkankocali/hotel-store-case-go/pkg/constant/user"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema/response"
	services "github.com/volkankocali/hotel-store-case-go/pkg/usecase/interface"
	"net/http"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (u *UserHandler) SignUp(ctx *fiber.Ctx) error {
	var input schema.UserSchema

	if err := ctx.BodyParser(&input); err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, request_constant.BadRequest, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	// validation
	validate := validator.New()

	validationError := validate.Struct(&input)

	if validationError != nil {
		res := response.ClientResponse(fiber.StatusUnprocessableEntity, user_constant.NotValidate, nil, validationError.(validator.ValidationErrors).Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	userCreate, err := u.userUseCase.UserSignUp(input)

	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, user_constant.NotSignIn, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	success := response.ClientResponse(fiber.StatusCreated, user_constant.SuccessSignUp, userCreate, nil, 0)
	return ctx.Status(success.StatusCode).JSON(success)
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	var user schema.UserLoginSchema

	if err := ctx.BodyParser(&user); err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, request_constant.BadRequest, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	// validation
	validate := validator.New()

	validationError := validate.Struct(&user)

	if validationError != nil {
		res := response.ClientResponse(fiber.StatusUnprocessableEntity, user_constant.NotValidate, nil, validationError.(validator.ValidationErrors).Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	login, err := u.userUseCase.UserLogin(user)
	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, user_constant.NotFound, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	return ctx.JSON(login)
}

func (u *UserHandler) RefreshToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Get("Authorization")
	if refreshToken == "" {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": constant.MissingRefreshToken})
	}

	userRefreshToken, err := u.userUseCase.UserRefreshToken(refreshToken)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"refresh_token": userRefreshToken.Token})
}
