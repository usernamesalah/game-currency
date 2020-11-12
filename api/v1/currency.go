package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/usernamesalah/game-currency/pkg/models"
)

// List currency
// @Summary List currency
// @Description Get the list of currency
// @Tags currency
// @ID list-currency
// @Produce json
// @Success 200 {array} models.Currency
// @Router /currency [get]
func (api *API) listCurrency(c echo.Context) error {
	ctx := c.Request().Context()

	currency, err := api.currencyService.ListCurrency(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, currency)
}

// Create a new currency
// @Summary Create a new currency
// @Description Create a new currency
// @Tags currency
// @ID create-currency
// @Produce json
// @Param player body models.Currency true "Create currency"
// @Success 201 {object} models.Currency
// @Router /currency [post]
func (api *API) createCurrency(c echo.Context) error {
	ctx := c.Request().Context()

	currency := new(models.Currency)
	if err := c.Bind(currency); err != nil {
		return err
	}

	if err := c.Validate(currency); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newCurrency, err := api.conversionsService.CreateCurrency(ctx, *currency)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newCurrency)
}

// Create a new conversion
// @Summary Create a new conversion
// @Description Create a new conversion
// @Tags conversion
// @ID create-conversion
// @Produce json
// @Param player body models.Conversion true "Create conversion"
// @Success 201 {object} models.Conversion
// @Router /conversion [post]
func (api *API) createConversion(c echo.Context) error {
	ctx := c.Request().Context()

	conversion := new(models.Conversion)
	if err := c.Bind(conversion); err != nil {
		return err
	}

	if err := c.Validate(conversion); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newConversion, err := api.conversionsService.CreateConversion(ctx, *conversion)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newConversion)
}

// Create convert rate
// @Summary Create convert rate
// @Description Create convert rate
// @Tags conversion
// @ID convert-conversion
// @Produce json
// @Param player body models.Converted true "Create currency"
// @Success 201 {object} interface{}
// @Router /convert [post]
func (api *API) convert(c echo.Context) error {
	ctx := c.Request().Context()

	conversionFrom, _ := strconv.ParseInt(c.QueryParam("from"), 10, 64)
	conversionTo, _ := strconv.ParseInt(c.QueryParam("to"), 10, 64)
	amount, _ := strconv.ParseInt(c.QueryParam("amount"), 10, 64)

	if conversionFrom == 0 || conversionTo == 0 || amount == 0 {
		return echo.NewHTTPError(httpcode, "required params")
	}

	conversion, err := api.conversionsService.GetConversionBy(ctx, conversionFrom, conversionTo)
	if err != nil {
		return err
	}

	var rate int64
	if from := conversion.ConversionFrom; from == conversionFrom {
		rate = conversion.Rate
	} else {
		rate = int64(1) / conversion.Rate
	}

	amount = amount * rate
	return c.JSON(http.StatusCreated, map[string]int64{
		"result": amount,
	})
}
