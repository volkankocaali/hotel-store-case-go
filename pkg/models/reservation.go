package models

import "time"

type Reservation struct {
	ID                uint      `json:"id"`
	UserID            uint      `json:"user_id"`
	User              Users     `gorm:"foreignKey:UserID"`
	ReferenceCode     string    `json:"reference_code"`
	Destination       string    `json:"destination"`
	CheckInDate       time.Time `json:"check_in_date"`
	CheckOutDate      time.Time `json:"check_out_date"`
	Accommodation     string    `json:"accommodation"`
	NumberOfGuests    int       `json:"number_of_guests"`
	ReservationStatus string    `json:"reservation_status"`
}
