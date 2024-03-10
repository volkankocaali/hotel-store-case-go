package repository

import (
	reservation_const "github.com/volkankocali/hotel-store-case-go/pkg/constant/reservation"
	"github.com/volkankocali/hotel-store-case-go/pkg/models"
	interfaces "github.com/volkankocali/hotel-store-case-go/pkg/repository/interface"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
	"gorm.io/gorm"
)

type reservationDatabase struct {
	DB *gorm.DB
}

func NewReservationRepository(DB *gorm.DB) interfaces.ReservationRepository {
	return &reservationDatabase{DB}
}

func (r *reservationDatabase) Create(reservation schema.CreateReservationSchema) (schema.CreateReservationSchema, error) {
	reservationModel, err := reservation.ToReservationModel()
	if err != nil {
		return schema.CreateReservationSchema{}, err
	}

	err = r.DB.Create(&reservationModel).Error

	if err != nil {
		return schema.CreateReservationSchema{}, err
	}

	return reservation, nil
}

func (r *reservationDatabase) Update(reservationId string, reservation schema.CreateReservationSchema) (schema.CreateReservationSchema, error) {
	reservationModel, err := reservation.ToReservationModel()
	if err != nil {
		return schema.CreateReservationSchema{}, err
	}

	err = r.DB.Model(&reservationModel).Where("id = ?", reservationId).Updates(&reservationModel).Error

	if err != nil {
		return schema.CreateReservationSchema{}, err
	}

	return reservation, nil
}

func (r *reservationDatabase) GetReservationByReferenceCode(referenceCode string) (schema.CreateReservationSchema, error) {
	var reservation models.Reservation
	err := r.DB.Where("reference_code = ?", referenceCode).First(&reservation).Error

	if err != nil {
		return schema.CreateReservationSchema{}, err
	}

	return ConvertToSchema(reservation), nil
}

func (r *reservationDatabase) AllReservation(userId uint) ([]schema.CreateReservationSchema, error) {
	var reservations []models.Reservation

	// user id get all reservation
	err := r.DB.Where("user_id = ?", userId).Find(&reservations).Error

	if err != nil {
		return []schema.CreateReservationSchema{}, err
	}

	var reservationSchemas []schema.CreateReservationSchema
	for _, reservation := range reservations {
		reservationSchemas = append(reservationSchemas, ConvertToSchema(reservation))
	}

	return reservationSchemas, nil
}

func ConvertToSchema(reservation models.Reservation) schema.CreateReservationSchema {
	return schema.CreateReservationSchema{
		UserID:            reservation.UserID,
		ReferenceCode:     reservation.ReferenceCode,
		Destination:       reservation.Destination,
		CheckInDate:       reservation.CheckInDate.Format("2006-01-02"),
		CheckOutDate:      reservation.CheckOutDate.Format("2006-01-02"),
		Accommodation:     reservation_const.AccommodationTypes[reservation.Accommodation],
		NumberOfGuests:    reservation.NumberOfGuests,
		ReservationStatus: reservation_const.Status[reservation.ReservationStatus],
	}
}
