package tests

import (
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
func TestPing404Errors1Of3(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestPing404Errors2Of3(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/services/search", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestPing404Errors3Of3(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

/************** All Services Tests **************/
func TestPingAllServicesWhenLanguageIsEn(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/services", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingAllServicesWhenLanguageIsFr(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/services", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingAllServicesWhenLanguageIsUndefined(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/sp/services", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

/************** One Service Tests **************/
func TestPingOneServiceWhenServiceExistsInEn(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/service/security-police-direct-2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingOneServiceWhenServiceDoesntExistsInEn(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/service/security-police-direct-4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestPingOneServiceWhenServiceExistsInFr(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/service/security-police-direct-1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingOneServiceWhenServiceDoesntExistsInFr(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/service/security-police-direct-4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

/************** Search Services Tests **************/
func TestPingSearchServicesWhenQueryMatchesInEn1Of2(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/services/search?query=police-direct", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingSearchServicesWhenQueryMatchesInEn2Of2(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/services/search?query=pol", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingSearchServicesWhenQueryDoesntMatchInEn(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/services/search?query=poolicing", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingSearchServicesWhenQueryIsEmptyInEn(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/en/services/search?query=", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingSearchServicesWhenQueryMatchesInFr1Of2(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/services/search?query=police-direct", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingSearchServicesWhenQueryMatchesInFr2Of2(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/services/search?query=pol", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingSearchServicesWhenQueryDoesntMatchInFr(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/services/search?query=poolicing", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPingSearchServicesWhenQueryIsEmptyInFr(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mesg/fr/services/search?query=", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
