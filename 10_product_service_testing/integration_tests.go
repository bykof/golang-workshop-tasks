package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProducts(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `[
		{"ID": 2,"CreatedAt": "0001-01-01T00:00:00Z","UpdatedAt": "2023-02-10T13:32:15.066438+01:00","Name": "Mentos Fruit","Price": "2"}]`,
		w.Body.String(),
	)
}
