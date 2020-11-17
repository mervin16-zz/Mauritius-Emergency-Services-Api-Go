package tests

import (
	"encoding/json"
	"mes/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/******************* Setup Test Server *********************/
func setupServer() *gin.Engine {
	// Gin Mode
	gin.SetMode(gin.ReleaseMode)

	// Creates the Gin Engine
	engine := gin.New()

	// Setup the API Routes
	engine.GET("/mesg/:language/services", routes.AllServices)
	engine.GET("/mesg/:language/service/:id", routes.OneService)
	engine.GET("/mesg/:language/services/search", routes.SearchService)

	// Setup Error Routes
	engine.NoRoute(routes.Error404Handler)

	// Return engine
	return engine
}

/************** Error Handling Tests **************/
func TestResponse404Errors1Of3(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Wrong routes used. Please read the docs", response.Message)
	assert.Equal(t, false, response.Success)
	assert.Equal(t, []Service{}, response.Services)
	assert.Equal(t, 404, w.Code)
}

func TestResponse404Errors2Of3(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/services/search", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Wrong routes used. Please read the docs", response.Message)
	assert.Equal(t, false, response.Success)
	assert.Equal(t, []Service{}, response.Services)
	assert.Equal(t, 404, w.Code)
}

func TestResponse404Errors3Of3(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Wrong routes used. Please read the docs", response.Message)
	assert.Equal(t, false, response.Success)
	assert.Equal(t, []Service{}, response.Services)
	assert.Equal(t, 404, w.Code)
}

/************** All Services Tests **************/
func TestResponseAllServicesWhenLanguageIsEn(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/services", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseAllServicesWhenLanguageIsFr(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/services", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseAllServicesWhenLanguageIsUndefined(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/sp/services", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

/************** One Service Tests **************/
func TestResponseOneServiceWhenServiceExistsInEn(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/service/security-police-direct-2", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseOneServiceWhenServiceDoesntExistsInEn(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/service/security-police-direct-4", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, false, response.Success)
	assert.Equal(t, "No services found under id security-police-direct-4", response.Message)
	assert.Equal(t, 404, w.Code)
}

func TestResponseOneServiceWhenServiceExistsInFr(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/service/security-police-direct-1", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseOneServiceWhenServiceDoesntExistsInFr(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/service/security-police-direct-4", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, false, response.Success)
	assert.Equal(t, "No services found under id security-police-direct-4", response.Message)
	assert.Equal(t, 404, w.Code)
}

/************** Search Services Tests **************/
func TestResponseSearchServicesWhenQueryMatchesInEn1Of2(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/services/search?query=police-direct", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseSearchServicesWhenQueryMatchesInEn2Of2(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/services/search?query=pol", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseSearchServicesWhenQueryDoesntMatchInEn(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/services/search?query=poolicing", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseSearchServicesWhenQueryIsEmptyInEn(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/services/search?query=", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseSearchServicesWhenQueryMatchesInFr1Of2(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/services/search?query=police-direct", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseSearchServicesWhenQueryMatchesInFr2Of2(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/services/search?query=pol", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseSearchServicesWhenQueryDoesntMatchInFr(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/services/search?query=poolicing", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestResponseSearchServicesWhenQueryIsEmptyInFr(t *testing.T) {
	router := setupServer()
	response := Response{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/services/search?query=", nil)
	router.ServeHTTP(w, req)

	// Extract body
	resBody := w.Body.String()
	err := json.Unmarshal([]byte(resBody), &response)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, true, response.Success)
	assert.Equal(t, "", response.Message)
	assert.Equal(t, 200, w.Code)
}
