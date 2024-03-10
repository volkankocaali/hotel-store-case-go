package _interface

import (
	"github.com/gofiber/fiber/v2"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
)

type ReservationUseCase interface {
	CreateReservation(ctx *fiber.Ctx, reservation schema.CreateReservationSchema) (schema.CreateReservationSchema, error)
	UpdateReservation(ctx *fiber.Ctx, reservation schema.CreateReservationSchema) (schema.CreateReservationSchema, error)
	GetReservation(referenceCode string) (schema.CreateReservationSchema, error)
	AllReservation(ctx *fiber.Ctx) ([]schema.CreateReservationSchema, error)
}
