package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/talithaalda/go-middleware/internal/infrastructure/mocks"
)

func TestGetConnection(t *testing.T) {
	// Setup
	mockDB := new(mocks.GormPostgres)
	expectedDB := &gorm.DB{} // Expected return value

	mockDB.On("GetConnection").Return(expectedDB)

	// Execution
	db := mockDB.GetConnection()

	// Assertion
	assert.Equal(t, expectedDB, db)

	// Verify
	mockDB.AssertExpectations(t)
}
