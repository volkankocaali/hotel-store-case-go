package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/volkankocali/hotel-store-case-go/cmd/routes"
	"github.com/volkankocali/hotel-store-case-go/pkg/config"
	"github.com/volkankocali/hotel-store-case-go/pkg/database"
	"gorm.io/gorm"
	"log"
)

type App struct {
	Config   *config.Config
	FiberApp *fiber.App
	DB       *gorm.DB
}

func main() {
	app := App{}

	app.initialize()
	app.FiberApp = app.initFiber(app.Config)

	// setup routes
	routes.SetupRoutes(app.FiberApp, app.Config, app.DB)

	// run fiber app
	app.run()
}

func (app *App) initialize() {
	cfg := config.LoadConfig()
	mysql, _ := database.NewMysqlDB(*cfg)

	app.Config = cfg
	app.DB = mysql
}

func (app *App) initFiber(cfg *config.Config) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		AppName:     cfg.AppName,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	return fiberApp
}

func (app *App) run() {
	log.Printf("Starting server on port %s", app.Config.AppPort)
	if err := app.FiberApp.Listen(":" + app.Config.AppPort); err != nil {
		log.Fatalf("Error while starting server: %s", err)
	}
}
