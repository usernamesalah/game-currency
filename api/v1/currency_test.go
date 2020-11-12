package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/usernamesalah/game-currency/pkg/models"
	"github.com/usernamesalah/game-currency/pkg/services/mocks"
)

// type mockRequestValidator struct{}

// func (m *mockRequestValidator) Validate(i interface{}) error {
// 	return nil
// }

func TestAPI_listCurrency(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/currency", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	mockCurrencyService := &mocks.CurrencyService{}
	mockCurrencyService.On("ListCurrency", mock.Anything).Return([]models.Currency{}, nil)

	api := NewAPI(&mocks.TeamsService{}, mockCurrencyService, "", "")
	if assert.NoError(t, api.listCurrency(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[]\n", rec.Body.String())
	}
}

func TestAPI_createCurrency(t *testing.T) {
	currency := models.Currency{
		ID:   1,
		Name: "IDR",
	}
	currencyJSON, _ := json.Marshal(currency)

	req := httptest.NewRequest(http.MethodPost, "/currency", bytes.NewReader(currencyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.Validator = &mockRequestValidator{}
	c := e.NewContext(req, rec)

	mockPlayersService := &mocks.PlayersService{}
	mockPlayersService.On("CreateCurrency", mock.Anything, currency).Return(player, nil)

	api := NewAPI(&mocks.TeamsService{}, mockPlayersService, "", "")
	if assert.NoError(t, api.createPlayer(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "{\"id\":1,\"name\":\"IDR\"}\n", rec.Body.String())
	}
}

func TestAPI_createConversion(t *testing.T) {
	conversion := models.Conversion{
		ID:             1,
		ConversionFrom: 1,
		ConversionTo:   2,
		Rate:           20,
	}
	conversionJSON, _ := json.Marshal(conversion)

	req := httptest.NewRequest(http.MethodPost, "/conversion", bytes.NewReader(conversionJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.Validator = &mockRequestValidator{}
	c := e.NewContext(req, rec)

	mockPlayersService := &mocks.PlayersService{}
	mockPlayersService.On("CreateConversion", mock.Anything, currency).Return(player, nil)

	api := NewAPI(&mocks.TeamsService{}, mockPlayersService, "", "")
	if assert.NoError(t, api.createPlayer(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "{\"id\":1,\"conversion_from\":1,\"conversion_to\":2,\"rate\":20}\n", rec.Body.String())
	}
}
