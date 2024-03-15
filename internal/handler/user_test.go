package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"tidy/internal/service/mocks"

	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/handler"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/model"
	"github.com/gin-gonic/gin"
)

func TestGetUsersById(t *testing.T) {
    // Prepare mock
    mockService := new(mocks.UserService)

    // Create handler with mock service
    handler := handler.NewUserHandler(mockService)

    // Create router
    router := gin.Default()
    router.GET("/users/:id", handler.GetUsersById)

    // Prepare expected user
    expectedUser := model.User{
        ID:   123,
        Name: "John Doe",
        // Add other user data as needed
    }

    // Configure mock service
    mockService.On("GetUsersById", mock.Anything, uint64(123)).Return(expectedUser, nil)

    // Send request
    req, err := http.NewRequest("GET", "/users/123", nil)
    if err != nil {
        t.Fatal(err)
    }
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    // Verify response
    assert.Equal(t, http.StatusOK, resp.Code)

    // Decode response
    var user model.User
    if err := json.Unmarshal(resp.Body.Bytes(), &user); err != nil {
        t.Fatal(err)
    }

    // Verify user data
    assert.Equal(t, expectedUser.ID, user.ID)
    assert.Equal(t, expectedUser.Name, user.Name)
    // Add other assertions as needed
}
