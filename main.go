package main

import (
	_ "embed"
	"log"
	"os"
	"strconv"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/efectn/rest-countries/docs"
	"github.com/efectn/rest-countries/helpers"
	"github.com/efectn/rest-countries/routes"
	_ "github.com/efectn/rest-countries/statik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	fs "github.com/rakyll/statik/fs"
)

// @title Rest Countries Go
// @version 1.1
// @description Get information about countries via a RESTful API made by Go.
// @description Source Code: https://github.com/efectn/rest-countries
// @license.name GPL-3
// @license.url https://www.gnu.org/licenses/gpl-3.0.en.html
// @host localhost:3000
// @BasePath /api/v1
func main() {
	// Load Enviroment Variables
	err := godotenv.Load()
	helpers.ThrowError(err)

	prefork, err := strconv.ParseBool(os.Getenv("PREFORK"))
	helpers.ThrowError(err)
	rateLimit, err := strconv.ParseBool(os.Getenv("RATE_LIMIT"))
	helpers.ThrowError(err)
	connection, err := strconv.Atoi(os.Getenv("MAX_CONNECTION_COUNT"))
	helpers.ThrowError(err)
	expiration, err := strconv.Atoi(os.Getenv("EXPIRATION_SECOND"))
	helpers.ThrowError(err)

	// Webserver Configuration
	app := fiber.New(fiber.Config{
		Prefork:      prefork,
		ServerHeader: os.Getenv("SERVER_HEADER"),
		AppName:      os.Getenv("APP_NAME"),
	})

	// Statik FS Configuration
	statikFS, err := fs.New()
	helpers.ThrowError(err)

	// Register Middlewares (compress, filesystem, limiter)
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use("/data", filesystem.New(filesystem.Config{
		Root: statikFS,
	}))

	app.Use(logger.New(logger.Config{
		TimeZone: "Europe/Istanbul",
	}))

	if rateLimit {
		app.Use(limiter.New(limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.IP() == os.Getenv("ALLOWED_IP")
			},
			Max:        connection,
			Expiration: time.Duration(expiration) * time.Second,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(403).JSON(fiber.Map{
					"success": false,
					"message": "Limit has been reached. Please, be slow!",
				})
			},
		}))
	}

	// Register Main Routes
	routes.RegisterAPIRoutes(app)
	app.Get("/*", swagger.Handler)

	// Listen webserver
	port := "3000"
	if _, ok := os.LookupEnv("PORT"); ok {
		port = os.Getenv("PORT")
	}

	log.Fatal(app.Listen(":" + port))
}
