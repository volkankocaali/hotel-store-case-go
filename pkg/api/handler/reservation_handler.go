package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	request_constant "github.com/volkankocali/hotel-store-case-go/pkg/constant/request"
	reservation_const "github.com/volkankocali/hotel-store-case-go/pkg/constant/reservation"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema/response"
	services "github.com/volkankocali/hotel-store-case-go/pkg/usecase/interface"
)

type ReservationHandler struct {
	reservationUseCase services.ReservationUseCase
}

func NewReservationHandler(usecase services.ReservationUseCase) *ReservationHandler {
	return &ReservationHandler{
		reservationUseCase: usecase,
	}
}

func parseRequestBody(ctx *fiber.Ctx, input interface{}) error {
	if err := ctx.BodyParser(input); err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, request_constant.BadRequest, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	validate := validator.New()
	if validationError := validate.Struct(input); validationError != nil {
		res := response.ClientResponse(fiber.StatusUnprocessableEntity, request_constant.UnprocessableEntity, nil, validationError.(validator.ValidationErrors).Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	return nil
}

func (r *ReservationHandler) CreateReservation(ctx *fiber.Ctx) error {
	var input schema.CreateReservationSchema

	if err := parseRequestBody(ctx, &input); err != nil {
		return err
	}

	createReservation, err := r.reservationUseCase.CreateReservation(ctx, input)
	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, request_constant.BadRequest, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	success := response.ClientResponse(fiber.StatusCreated, reservation_const.CreateSuccessMessage, createReservation, nil, 0)
	return ctx.Status(success.StatusCode).JSON(success)
}

func (r *ReservationHandler) UpdateReservation(ctx *fiber.Ctx) error {
	var input schema.CreateReservationSchema

	if err := parseRequestBody(ctx, &input); err != nil {
		return err
	}

	updateReservation, err := r.reservationUseCase.UpdateReservation(ctx, input)
	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, request_constant.BadRequest, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	success := response.ClientResponse(fiber.StatusOK, reservation_const.UpdateSuccessMessage, updateReservation, nil, 0)
	return ctx.Status(success.StatusCode).JSON(success)
}

func (r *ReservationHandler) GetReservation(ctx *fiber.Ctx) error {
	referenceCode := ctx.Params("referenceCode")

	reservation, err := r.reservationUseCase.GetReservation(referenceCode)
	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, request_constant.BadRequest, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	success := response.ClientResponse(fiber.StatusOK, reservation_const.GetSuccessMessage, reservation, nil, 0)
	return ctx.Status(success.StatusCode).JSON(success)
}

func (r *ReservationHandler) AllReservation(ctx *fiber.Ctx) error {
	reservations, err := r.reservationUseCase.AllReservation(ctx)
	if err != nil {
		res := response.ClientResponse(fiber.StatusBadRequest, request_constant.BadRequest, nil, err.Error(), 0)
		return ctx.Status(res.StatusCode).JSON(res)
	}

	success := response.ClientResponse(fiber.StatusOK, reservation_const.GetAllSuccessMessage, reservations, nil, 0)
	return ctx.Status(success.StatusCode).JSON(success)

}
