package models

import "github.com/volkankocali/hotel-store-case-go/pkg/models/traits"

type Users struct {
	ID           uint          `json:"id" gorm:"unique;not null"`
	Name         string        `json:"name"`
	Email        string        `json:"email" validate:"email"`
	Password     string        `json:"password" validate:"min=8,max=20"`
	Phone        string        `json:"phone"`
	Reservations []Reservation `json:"reservations" gorm:"foreignKey:UserID"`
	traits.Timestampable
}
