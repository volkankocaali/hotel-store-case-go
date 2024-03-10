package schema

import (
	"github.com/volkankocali/hotel-store-case-go/pkg/models"
	"time"
)

type CreateReservationSchema struct {
	UserID            uint   `json:"user_id" form:"user_id" validate:"omitempty"`
	ReferenceCode     string `json:"reference_code" form:"reference_code" validate:"omitempty"`
	Destination       string `json:"destination" form:"destination" validate:"required"`
	CheckInDate       string `json:"check_in_date" form:"check_in_date" validate:"required"`
	CheckOutDate      string `json:"check_out_date" form:"check_out_date" validate:"required"`
	Accommodation     string `json:"accommodation" form:"accommodation" validate:"required"`
	NumberOfGuests    int    `json:"number_of_guests" form:"number_of_guests" validate:"omitempty,min=1"`
	ReservationStatus string `json:"reservation_status" form:"reservation_status" validate:"omitempty"`
}

func (s *CreateReservationSchema) ToReservationModel() (models.Reservation, error) {
	checkInDate, err := time.Parse("2006-01-02", s.CheckInDate)

	if err != nil {
		return models.Reservation{}, err
	}

	checkOutDate, err := time.Parse("2006-01-02", s.CheckOutDate)
	if err != nil {
		return models.Reservation{}, err
	}

	return models.Reservation{
		UserID:            s.UserID,
		ReferenceCode:     s.ReferenceCode,
		Destination:       s.Destination,
		CheckInDate:       checkInDate,
		CheckOutDate:      checkOutDate,
		Accommodation:     s.Accommodation,
		NumberOfGuests:    s.NumberOfGuests,
		ReservationStatus: s.ReservationStatus,
	}, nil
}
