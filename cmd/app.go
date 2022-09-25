package main

import (
	"coditeach/database"
	"coditeach/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
}

var infoLog = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "*",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	routes.Setup(app)

	app.Use(logger.New(logger.ConfigDefault))

	app.Get("/dashboard", monitor.New())

	database.Connect()

	log.Fatal(app.Listen(":8080"))

	infoLog.Printf("Server running on port 8080")
}
