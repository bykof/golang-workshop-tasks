package controllers

import (
	"database_example/core/domain/entities"
	"database_example/core/ports/mocks"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProductController_List(t *testing.T) {
	productRepositoryMock := mocks.NewProductRepository(t)
	productController := NewProductController(productRepositoryMock)
	recorder := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(recorder)
	productRepositoryMock.On("List").Return([]entities.Product{
		{
			Name: "test",
		},
	})
	productController.List(testContext)
	productRepositoryMock.AssertExpectations(t)

	content, err := ioutil.ReadAll(recorder.Body)
	assert.NoError(t, err)
	assert.Equal(t, `[{"ID":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","Name":"test","Price":"0"}]`, string(content))
}
