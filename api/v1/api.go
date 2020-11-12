package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/usernamesalah/game-currency/pkg/services"
)

// API can register a set of endpoints in a router and handle
// them using the provided storage.
type API struct {
	currencyService    services.CurrencyService
	conversionsService services.ConversionsService

	adminUsername string
	adminPassword string
}

// NewAPI returns an initialized API type.
func NewAPI(currencyService services.CurrencyService,
	conversionsService services.ConversionsService, adminUsername, adminPassword string) *API {
	return &API{
		currencyService:    currencyService,
		conversionsService: conversionsService,

		adminUsername: adminUsername,
		adminPassword: adminPassword,
	}
}

// Register the API's endpoints in the given router.
func (api *API) Register(g *echo.Group) {
	//Currency API
	g.GET("/currency", api.listCurrency)
	g.GET("/convert", api.convert)

	g.POST("/currency", api.createCurrency, middleware.BasicAuth(api.adminValidator))
	g.POST("/conversions", api.createConversion, middleware.BasicAuth(api.adminValidator))
}

func (api *API) adminValidator(username, password string, c echo.Context) (bool, error) {
	if username == api.adminUsername && password == api.adminPassword {
		return true, nil
	}
	return false, nil
}
