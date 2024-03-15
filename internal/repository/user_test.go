package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/talithaalda/go-middleware/internal/infrastructure"
	"github.com/talithaalda/go-middleware/internal/model"
)

type mockGormPostgres struct {
	mock.Mock
}

func (m *mockGormPostgres) GetConnection() infrastructure.DBConnection {
	args := m.Called()
	return args.Get(0).(infrastructure.DBConnection)
}

func TestCreateUser_Success(t *testing.T) {
	mockDB := new(mockGormPostgres)
	mockDB.On("GetConnection").Return(&infrastructure.DBConnectionMock{})

	repo := NewUserQuery(mockDB)

	expectedUser := model.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}

	mockDBConnection := &infrastructure.DBConnectionMock{}
	mockDBConnection.On("WithContext", mock.Anything).Return(mockDBConnection)
	mockDBConnection.On("Table", "users").Return(mockDBConnection)
	mockDBConnection.On("Save", mock.Anything).Return(nil)

	user, err := repo.CreateUser(context.Background(), expectedUser)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestCreateUser_Error(t *testing.T) {
	mockDB := new(mockGormPostgres)
	mockDB.On("GetConnection").Return(&infrastructure.DBConnectionMock{})

	repo := NewUserQuery(mockDB)

	expectedUser := model.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}

	mockDBConnection := &infrastructure.DBConnectionMock{}
	mockDBConnection.On("WithContext", mock.Anything).Return(mockDBConnection)
	mockDBConnection.On("Table", "users").Return(mockDBConnection)
	mockDBConnection.On("Save", mock.Anything).Return(errors.New("unable to save user"))

	user, err := repo.CreateUser(context.Background(), expectedUser)

	assert.Error(t, err)
	assert.Equal(t, model.User{}, user)
}
