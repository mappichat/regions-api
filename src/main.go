package main

import (
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jmoiron/sqlx"
	"github.com/mappichat/region-api/src/utils"
)

var validate = validator.New()

func main() {
	startTime := time.Now()

	utils.ConfigureEnv()

	errorString := "run using one of these subcommands: populate-db"
	if len(os.Args) < 2 {
		log.Fatal(errorString)
	}

	switch os.Args[1] {
	case "":
	}
	// connect to db
	db, err := sqlx.Connect("postgres", utils.Env.DB_CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// define api routes
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Healthy")
	})

	// jwt middleware
	keyRefreshDuration := 5 * time.Minute
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod:       "RS256",
		KeySetURL:           utils.Env.AUTH_JWKS_URI,
		KeyRefreshInterval:  &keyRefreshDuration,
		KeyRefreshRateLimit: &keyRefreshDuration,
	}))

	app.Get("/geojson", func(c *fiber.Ctx) error {
		c.Accepts("json", "text")
		c.Accepts("application/json")
		query := struct {
			H3    string `query:"h3" json:"h3" validate:"required"`
			Level int    `query:"level" json:"level"`
			Ring  int    `query:"ring" json:"ring"`
		}{}

		if err := c.QueryParser(&query); err != nil {
			return err
		}
		if err := validate.Struct(query); err != nil {
			return err
		}

		tiles := []struct {
			H3 string `db:"tiles.h3"`
		}{}
		if err := db.Select(
			&tiles,
			`SELECT tiles.h3, neighbors.CustomerName FROM tiles
			WHERE tiles.region=$1
			INNER JOIN neighbors ON Orders.CustomerID=Customers.CustomerID;`,
			query.H3,
		); err != nil {
			return err
		}

		h3

		return c.JSON("")
	})

	// serve api
	log.Fatal(app.Listen(":" + utils.Env.PORT))
}
