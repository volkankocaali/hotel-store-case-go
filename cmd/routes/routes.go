package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/volkankocali/hotel-store-case-go/cmd/middleware"
	"github.com/volkankocali/hotel-store-case-go/pkg/api/handler"
	"github.com/volkankocali/hotel-store-case-go/pkg/config"
	"github.com/volkankocali/hotel-store-case-go/pkg/helper"
	"github.com/volkankocali/hotel-store-case-go/pkg/repository"
	"github.com/volkankocali/hotel-store-case-go/pkg/usecase"
	"gorm.io/gorm"
)

func SetupRoutes(
	app *fiber.App,
	cfg *config.Config,
	db *gorm.DB,
) {
	h := helper.NewHelper(cfg)

	// Repository
	userRepository := repository.NewUserRepository(db)
	reservationRepository := repository.NewReservationRepository(db)

	// Usecase
	userUseCase := usecase.NewUserUseCase(userRepository, cfg, h)
	reservationUseCase := usecase.NewReservationUseCase(reservationRepository, cfg, h)

	// Handler
	userHandler := handler.NewUserHandler(userUseCase)
	reservationHandler := handler.NewReservationHandler(reservationUseCase)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/sign-up", userHandler.SignUp)
	v1.Post("/login", userHandler.Login)

	// Jwt Middleware for Client User
	v1.Use(middleware.UserAuthMiddleware)

	v1.Post("/refresh-token", userHandler.RefreshToken)

	// Reservation services
	v1.Get("/reservations/:referenceCode", reservationHandler.GetReservation)
	v1.Get("/reservations", reservationHandler.AllReservation)
	v1.Post("/reservations", reservationHandler.CreateReservation)
	v1.Put("/reservations/:id", reservationHandler.UpdateReservation)

	v1.Get("/health-check", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status": "success",
		})
	})
}
