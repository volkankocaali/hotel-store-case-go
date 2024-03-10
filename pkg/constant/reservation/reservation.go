package reservation

const (
	CreatedStatus   = "created"
	ConfirmedStatus = "confirmed"
	CancelledStatus = "cancelled"
	CompletedStatus = "completed"
)

const (
	StatusCreated   = "10"
	StatusConfirmed = "20"
	StatusCancelled = "30"
	StatusCompleted = "40"
)

var Status = map[string]string{
	StatusCreated:   CreatedStatus,
	StatusConfirmed: ConfirmedStatus,
	StatusCancelled: CancelledStatus,
	StatusCompleted: CompletedStatus,
}

var AccommodationTypes = map[string]string{
	"hotel":       "Otel",
	"resort":      "Tatil Köyü",
	"apartment":   "Apartman",
	"villa":       "Villa",
	"guest_house": "Pansiyon",
}

const (
	CreateSuccessMessage = "Your reservation was successfully created"
	UpdateSuccessMessage = "Your reservation was successfully updated"

	CheckInDateMustBeLessThanCheckOutDate = "check in date must be less than check out date"
	AccommodationTypeNotFound             = "accommodation type not found"
	NotFound                              = "reservation not found"
	GetSuccessMessage                     = "Your reservation was successfully fetched"
	GetAllSuccessMessage                  = "Your reservations were successfully fetched"
)
