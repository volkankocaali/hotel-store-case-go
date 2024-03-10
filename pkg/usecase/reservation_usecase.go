package usecase

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/volkankocali/hotel-store-case-go/pkg/config"
	reservation_const "github.com/volkankocali/hotel-store-case-go/pkg/constant/reservation"
	helper_interface "github.com/volkankocali/hotel-store-case-go/pkg/helper/interface"
	interfaces "github.com/volkankocali/hotel-store-case-go/pkg/repository/interface"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
)

type ReservationUseCase struct {
	reservationRepository interfaces.ReservationRepository
	config                config.Config
	helper                helper_interface.Helper
}

func NewReservationUseCase(
	reservationRepository interfaces.ReservationRepository,
	cfg *config.Config,
	h helper_interface.Helper,
) *ReservationUseCase {
	return &ReservationUseCase{
		reservationRepository: reservationRepository,
		config:                *cfg,
		helper:                h,
	}
}

func (r *ReservationUseCase) CreateReservation(ctx *fiber.Ctx, reservation schema.CreateReservationSchema) (schema.CreateReservationSchema, error) {
	// reference code generate
	referenceCode := r.helper.GenerateReferenceCode()

	reservation.ReferenceCode = referenceCode

	// check in and check out control
	if reservation.CheckInDate >= reservation.CheckOutDate {
		return schema.CreateReservationSchema{}, errors.New("check in date must be less than check out date")
	}

	// check accommodation type
	if _, ok := reservation_const.AccommodationTypes[reservation.Accommodation]; !ok {
		return schema.CreateReservationSchema{}, errors.New(reservation_const.AccommodationTypeNotFound)
	}

	// new reservation record set status created
	reservation.ReservationStatus = reservation_const.StatusCreated

	// find bearer token user id
	token := ctx.Get("Authorization")
	claims, err := r.helper.ParseToken(token)
	if err != nil {
		return schema.CreateReservationSchema{}, err
	}

	reservation.UserID = uint(claims["id"].(float64))

	// create reservation and save database
	reservationCreate, err := r.reservationRepository.Create(reservation)

	if err != nil {
		return schema.CreateReservationSchema{}, err
	}

	return reservationCreate, nil
}

func (r *ReservationUseCase) UpdateReservation(ctx *fiber.Ctx, reservation schema.CreateReservationSchema) (schema.CreateReservationSchema, error) {
	reservationId := ctx.Params("id")

	// check in and check out control
	if reservation.CheckInDate >= reservation.CheckOutDate {
		return schema.CreateReservationSchema{}, errors.New(reservation_const.CheckInDateMustBeLessThanCheckOutDate)
	}

	// check accommodation type
	if _, ok := reservation_const.AccommodationTypes[reservation.Accommodation]; !ok {
		return schema.CreateReservationSchema{}, errors.New(reservation_const.AccommodationTypeNotFound)
	}

	if reservation.ReservationStatus != "" {
		// check reservation status
		if _, ok := reservation_const.Status[reservation.ReservationStatus]; !ok {
			return schema.CreateReservationSchema{}, errors.New(reservation_const.NotFound)
		}
	}

	// find bearer token user id
	token := ctx.Get("Authorization")
	claims, err := r.helper.ParseToken(token)
	if err != nil {
		return schema.CreateReservationSchema{}, err
	}

	reservation.UserID = uint(claims["id"].(float64))
	// update reservation
	update, err := r.reservationRepository.Update(reservationId, reservation)

	if err != nil {
		return schema.CreateReservationSchema{}, err
	}

	return update, nil
}

func (r *ReservationUseCase) GetReservation(referenceCode string) (schema.CreateReservationSchema, error) {
	// get referenceId by reservation find and return
	reservation, err := r.reservationRepository.GetReservationByReferenceCode(referenceCode)

	if err != nil {
		return schema.CreateReservationSchema{}, err
	}

	return reservation, nil
}

func (r *ReservationUseCase) AllReservation(ctx *fiber.Ctx) ([]schema.CreateReservationSchema, error) {
	// find bearer token user id
	token := ctx.Get("Authorization")
	claims, err := r.helper.ParseToken(token)
	if err != nil {
		return []schema.CreateReservationSchema{}, err
	}

	userId := uint(claims["id"].(float64))
	// get all reservation
	reservations, err := r.reservationRepository.AllReservation(userId)

	if err != nil {
		return nil, err
	}

	return reservations, nil
}
