package controllers

import (
	"net/url"
	"strings"

	"github.com/efectn/rest-countries/helpers"
	"github.com/gofiber/fiber/v2"
)

// GetAll godoc
// @Summary Get all countries.
// @Tags Main
// @Accept */*
// @Produce json
// @Success 200 {object} map[uint]interface{}
// @Router /all [get]
func GetAll(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return c.SendString(helpers.GetCountries())
}

// GetByName godoc
// @Summary Search countries by name or native name.
// @Tags Main
// @Accept */*
// @Produce json
// @Success 200 {object} map[uint]interface{}
// @Failure 404 {object} map[string]interface{}
// @Param name path string true "Name, Native Name"
// @Param fullText query bool false "Fulltext Search"
// @Router /name/{name} [get]
func GetByName(c *fiber.Ctx) error {
	parsedCountries := make(map[uint]interface{})

	var count uint = 0
	for _, value := range helpers.GetParsedCountries() {
		decodedValue, err := url.QueryUnescape(c.Params("name"))
		helpers.ThrowError(err)

		input, name, nativeName := decodedValue, value.Path("name").String(), value.Path("nativeName").String()
		if input := `"` + input + `"`; c.FormValue("fullText") == "true" && (strings.EqualFold(input, name) || strings.EqualFold(input, nativeName)) {
			parsedCountries[0] = value.Data()
		} else if helpers.ContainsInsensitive(name, decodedValue) || helpers.ContainsInsensitive(nativeName, decodedValue) {
			parsedCountries[count] = value.Data()
			count++
		}
	}

	if len(parsedCountries) <= 0 {
		return helpers.NotFound(c)
	}

	return c.JSON(parsedCountries)
}

// GetByCode godoc
// @Summary Search countries by ISO 3166-1 code.
// @Description Single (tr), Plural (tr;gb;usa).
// @Tags Main
// @Accept */*
// @Produce json
// @Success 200 {object} map[uint]interface{}
// @Failure 404 {object} map[string]interface{}
// @Param code path string true "ISO 3166-1 Code"
// @Router /code/{code} [get]
func GetByCode(c *fiber.Ctx) error {
	parsedCountries := make(map[uint]interface{})
	decodedValue, err := url.QueryUnescape(c.Params("code"))
	helpers.ThrowError(err)

	var codes []string = strings.Split(decodedValue, ";")

	var count uint = 0
	for _, code := range codes {
		for _, value := range helpers.GetParsedCountries() {
			if input := `"` + code + `"`; strings.EqualFold(input, value.Path("alpha2Code").String()) || strings.EqualFold(input, value.Path("alpha3Code").String()) {
				parsedCountries[count] = value.Data()
			}
		}
		count++
	}

	if len(parsedCountries) <= 0 {
		return helpers.NotFound(c)
	}

	return c.JSON(parsedCountries)
}

// GetByCurrrency godoc
// @Summary Search countries by ISO 4217 currency code.
// @Tags Main
// @Accept */*
// @Produce json
// @Success 200 {object} map[uint]interface{}
// @Failure 404 {object} map[string]interface{}
// @Param currency path string true "ISO 4217 Currency Code"
// @Router /currency/{currency} [get]
func GetByCurrrency(c *fiber.Ctx) error {
	parsedCountries := make(map[uint]interface{})

	var count uint = 0
	for _, value := range helpers.GetParsedCountries() {
		var currencies []string
		for _, cur := range value.Path("currencies").Children() {
			currencies = append(currencies, cur.Path("code").String())
		}

		if input := `"` + strings.ToUpper(c.Params("currency")) + `"`; helpers.Contains(currencies, input) {
			parsedCountries[count] = value.Data()
			count++
		}
	}

	if len(parsedCountries) <= 0 {
		return helpers.NotFound(c)
	}

	return c.JSON(parsedCountries)
}

// GetByLanguage godoc
// @Summary Search countries by ISO 639-1 language code.
// @Tags Main
// @Accept */*
// @Produce json
// @Success 200 {object} map[uint]interface{}
// @Failure 404 {object} map[string]interface{}
// @Param language path string true "ISO 639-1 Language Code"
// @Router /language/{language} [get]
func GetByLanguage(c *fiber.Ctx) error {
	parsedCountries := make(map[uint]interface{})

	var count uint = 0
	for _, value := range helpers.GetParsedCountries() {
		var languages []string
		for _, lang := range value.Path("languages").Children() {
			languages = append(languages, lang.Path("iso639_1").String())
		}

		if input := `"` + strings.ToLower(c.Params("language")) + `"`; helpers.Contains(languages, input) {
			parsedCountries[count] = value.Data()
			count++
		}
	}

	if len(parsedCountries) <= 0 {
		return helpers.NotFound(c)
	}

	return c.JSON(parsedCountries)
}

// GetByCapital godoc
// @Summary Search countries by capital city.
// @Tags Main
// @Accept */*
// @Produce json
// @Success 200 {object} map[uint]interface{}
// @Failure 404 {object} map[string]interface{}
// @Param capital path string true "Capital city"
// @Router /capital/{capital} [get]
func GetByCapital(c *fiber.Ctx) error {
	parsedCountries := make(map[uint]interface{})

	var count uint = 0
	for _, value := range helpers.GetParsedCountries() {
		decodedValue, err := url.QueryUnescape(c.Params("capital"))
		helpers.ThrowError(err)

		if helpers.ContainsInsensitive(value.Path("capital").String(), decodedValue) {
			parsedCountries[count] = value.Data()
			count++
		}
	}

	if len(parsedCountries) <= 0 {
		return helpers.NotFound(c)
	}

	return c.JSON(parsedCountries)
}

// GetByCallingCode godoc
// @Summary Search countries by calling code.
// @Tags Main
// @Accept */*
// @Produce json
// @Success 200 {object} map[uint]interface{}
// @Failure 404 {object} map[string]interface{}
// @Param code path string true "Calling Code"
// @Router /calling/{code} [get]
func GetByCallingCode(c *fiber.Ctx) error {
	parsedCountries := make(map[uint]interface{})

	var count uint = 0
	for _, value := range helpers.GetParsedCountries() {
		var codes []string
		for _, code := range value.Path("callingCodes").Children() {
			codes = append(codes, code.String())
		}

		if input := `"` + c.Params("code") + `"`; helpers.Contains(codes, input) {
			parsedCountries[count] = value.Data()
			count++
		}
	}

	if len(parsedCountries) <= 0 {
		return helpers.NotFound(c)
	}

	return c.JSON(parsedCountries)
}

// GetByRegion godoc
// @Summary Search countries by region.
// @Description Regions: Africa, Americas, Asia, Europe, Oceania.
// @Tags Main
// @Accept */*
// @Produce json
// @Success 200 {object} map[uint]interface{}
// @Failure 404 {object} map[string]interface{}
// @Param region path string true "Region"
// @Router /region/{region} [get]
func GetByRegion(c *fiber.Ctx) error {
	parsedCountries := make(map[uint]interface{})

	var count uint = 0
	for _, value := range helpers.GetParsedCountries() {
		if helpers.ContainsInsensitive(value.Path("region").String(), `"`+c.Params("region")+`"`) {
			parsedCountries[count] = value.Data()
			count++
		}
	}

	if len(parsedCountries) <= 0 {
		return helpers.NotFound(c)
	}

	return c.JSON(parsedCountries)
}

// GetByRegionalBlock godoc
// @Summary Search countries by regional block.
// @Description Blocks:
// @Description    - EU (European Union)
// @Description    - EFTA (European Free Trade Association)
// @Description    - CARICOM (Caribbean Community)
// @Description    - PA (Pacific Alliance)
// @Description    - AU (African Union)
// @Description    - USAN (Union of South American Nations)
// @Description    - EEU (Eurasian Economic Union)
// @Description    - AL (Arab League)
// @Description    - ASEAN (Association of Southeast Asian Nations)
// @Description    - CAIS (Central American Integration System)
// @Description    - CEFTA (Central European Free Trade Agreement)
// @Description    - NAFTA (North American Free Trade Agreement)
// @Description    - SAARC (South Asian Association for Regional Cooperation)
// @Tags Main
// @Accept */*
// @Produce json
// @Success 200 {object} map[uint]interface{}
// @Failure 404 {object} map[string]interface{}
// @Param block path string true "Regional Block"
// @Router /regional-block/{block} [get]
func GetByRegionalBlock(c *fiber.Ctx) error {
	parsedCountries := make(map[uint]interface{})

	var count uint = 0
	for _, value := range helpers.GetParsedCountries() {
		var blocks []string
		for _, block := range value.Path("regionalBlocs").Children() {
			blocks = append(blocks, block.Path("acronym").String())
		}

		if input := `"` + strings.ToUpper(c.Params("block")) + `"`; helpers.Contains(blocks, input) {
			parsedCountries[count] = value.Data()
			count++
		}

	}

	if len(parsedCountries) <= 0 {
		return helpers.NotFound(c)
	}

	return c.JSON(parsedCountries)
}
