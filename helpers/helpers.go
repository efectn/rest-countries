package helpers

import (
	"log"
	"os"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/efectn/rest-countries/storage"
	"github.com/gofiber/fiber/v2"
)

// Functions
func ThrowError(err error) {
	if err != nil {
		log.Fatalf("Exception: %v\n", err)
	}
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{
		"success": false,
		"message": "Country not found!",
	})
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ContainsInsensitive(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
func GetCountries() string {
	return strings.ReplaceAll(storage.Countries, "HOST", os.Getenv("HOST"))
}

func GetParsedCountries() []*gabs.Container {
	countries, err := gabs.ParseJSON([]byte(GetCountries()))
	ThrowError(err)

	return countries.Children()
}
