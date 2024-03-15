package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/talithaalda/go-middleware/internal/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/talithaalda/go-middleware/internal/handler"
	"github.com/talithaalda/go-middleware/internal/model"
)

func TestGetUsersById(t *testing.T) {
    // Prepare mock
    mockService := new(mocks.UserService)

    handler := handler.NewUserHandler(mockService)

    router := gin.Default()
    router.GET("/users/:id", handler.GetUsersById)

    expectedUser := model.User{
        ID:   123,
        
    }

    mockService.On("GetUsersById", mock.Anything, uint64(123)).Return(expectedUser, nil)
    req, err := http.NewRequest("GET", "/users/123", nil)
    if err != nil {
        t.Fatal(err)
    }
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)

    var user model.User
    if err := json.Unmarshal(resp.Body.Bytes(), &user); err != nil {
        t.Fatal(err)
    }

    assert.Equal(t, expectedUser.ID, user.ID)
}
