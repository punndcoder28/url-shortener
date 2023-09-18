package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateURL(t *testing.T) {
	requestBody := strings.NewReader(`{"url":"https://www.google.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", requestBody)
	w := httptest.NewRecorder()

	CreateURL(w, req)

	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code should be 200")

	responseBody, err := io.ReadAll(resp.Body)

	assert.NoError(t, err, "Error should be nil while reading response body")

	responseBody = []byte(strings.TrimSuffix(string(responseBody), "\n"))

	expectedResponseBody := `{"url":"http://localhost:8080/8739bc55"}`

	assert.Equal(
		t,
		expectedResponseBody,
		string(responseBody),
		"Response body should be equal to expected response body",
	)
}

func TestCreateEmptyURL(t *testing.T) {
	requestBody := strings.NewReader(`{"url":""}`)
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", requestBody)
	w := httptest.NewRecorder()

	CreateURL(w, req)

	resp := w.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Status code should be 400")

	responseBody, err := io.ReadAll(resp.Body)

	assert.NoError(t, err, "Error should be nil while reading response body")

	responseBody = []byte(strings.TrimSuffix(string(responseBody), "\n"))

	expectedResponseBody := `{"Message":"Invalid request body sent"}`

	assert.Equal(
		t,
		expectedResponseBody,
		string(responseBody),
		"Response body should be equal to expected response body",
	)
}

func TestCreateDuplicateURL(t *testing.T) {
	requestBody := strings.NewReader(`{"url":"https://www.google.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", requestBody)
	w := httptest.NewRecorder()

	CreateURL(w, req)

	resp := w.Result()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Status code should be 500")

	responseBody, err := io.ReadAll(resp.Body)

	assert.NoError(t, err, "Error should be nil while reading response body")

	responseBody = []byte(strings.TrimSuffix(string(responseBody), "\n"))

	expectedResponseBody := `{"error":"ERROR: duplicate key value violates unique constraint \"urls_hash_key\" (SQLSTATE 23505)"}`

	assert.Equal(
		t,
		expectedResponseBody,
		string(responseBody),
		"Response body should be equal to expected response body",
	)
}

func TestGetURL(t *testing.T) {}
