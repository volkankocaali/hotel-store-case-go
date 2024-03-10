package _interface

import (
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
)

type ReservationRepository interface {
	Create(reservation schema.CreateReservationSchema) (schema.CreateReservationSchema, error)
	Update(reservationId string, reservation schema.CreateReservationSchema) (schema.CreateReservationSchema, error)
	GetReservationByReferenceCode(referenceCode string) (schema.CreateReservationSchema, error)
	AllReservation(userId uint) ([]schema.CreateReservationSchema, error)
}
