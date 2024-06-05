package main

import (
	"URLSHORTNER/endpoints"
	svc "URLSHORTNER/services"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	log.Println("Test main in progress")
	urlSvc := svc.NewShortenURLServicer()
	router = gin.Default()
	endpoints.NewShortenHandler(router, urlSvc)
	rc := m.Run()
	os.Exit(rc)
}

func TestShouldShortenURL(t *testing.T) {
	req, _ := http.NewRequest("POST", "/v1/", strings.NewReader(`{"url": "http://manoj.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.NotEmpty(t, rr.Body.String())
}

func TestInvalidShortenURL(t *testing.T) {
	req, _ := http.NewRequest("POST", "/v1/", strings.NewReader(`{"url": ""}`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.NotEmpty(t, rr.Body.String())
}

func TestRedirectURL(t *testing.T) {
	req, _ := http.NewRequest("POST", "/v1/", strings.NewReader(`{"url": "http://manoj.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	urls := strings.Split(rr.Body.String(), "/")
	url := "/v1/url/" + urls[len(urls)-1]
	url = strings.Replace(url, "\"}", "", -1)
	req, _ = http.NewRequest("GET", url, nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusFound, rr.Code)
}

func TestDomainCount(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/getmetrics", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.NotContains(t, rr.Body.String(), "test.com")

	req, _ = http.NewRequest("POST", "/v1/", strings.NewReader(`{"url": "http://manoj.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	req, _ = http.NewRequest("POST", "/v1/", strings.NewReader(`{"url": "http://test.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	req, _ = http.NewRequest("GET", "/v1/getmetrics", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.Contains(t, rr.Body.String(), "manoj.com")
	require.Contains(t, rr.Body.String(), "test.com")
}
