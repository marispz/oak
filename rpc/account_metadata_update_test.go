package rpc

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	"oak/common"
	"oak/mocks"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/assert"
)

func TestUpdateAccountMetaData_Success(t *testing.T) {
	userID := "user123"
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, userID)

	// Setup mocks
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "UpdateAccountMetaData RPC called").Once()
	// database mock
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dbMock.ExpectExec(`^UPDATE users SET metadata .*`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// NakamaModule mock
	nk := new(mocks.NakamaModule) // Mock NakamaModule

	// Call the function
	result, err := UpdateAccountMetaData(ctx, mockLogger, db, nk, "{}")

	// Assertions
	assert.Equal(t, `{"status":"success"}`, result)
	assert.NoError(t, err)
	mockLogger.AssertExpectations(t)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestUpdateAccountMetaData_MetadataSizeLimitExceeded(t *testing.T) {
	userID := "user123"
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, userID)

	// Setup mocks
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "UpdateAccountMetaData RPC called").Once()
	mockLogger.On("Error", "Metadata size limit exceeded").Once()

	// Create a payload larger than 16KB
	payload := string(make([]byte, 16385)) // 16KB + 1 byte

	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	nk := new(mocks.NakamaModule) // Mock NakamaModule

	// Call the function
	result, err := UpdateAccountMetaData(ctx, mockLogger, db, nk, payload)

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrMetadataSizeLimit, err)
	mockLogger.AssertExpectations(t)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestUpdateAccountMetaData_MissingUserID(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "UpdateAccountMetaData RPC called").Once()
	mockLogger.On("Error", "Context did not contain user ID.").Once()

	ctx := context.Background() // No user ID in context

	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	nk := new(mocks.NakamaModule) // Mock NakamaModule

	// Call the function
	result, err := UpdateAccountMetaData(ctx, mockLogger, db, nk, "{}")

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrUserNotFound, err)
	mockLogger.AssertExpectations(t)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestUpdateAccountMetaData_UnmarshalError(t *testing.T) {
	userID := "user123"

	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "UpdateAccountMetaData RPC called").Once()
	mockLogger.On("Error", mock.Anything, mock.Anything).Once()

	// Invalid JSON payload
	payload := "{invalid_json"

	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, userID)

	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	nk := new(mocks.NakamaModule) // Mock NakamaModule

	// Call the function
	result, err := UpdateAccountMetaData(ctx, mockLogger, db, nk, payload)

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrUnMarshallingError, err)
	mockLogger.AssertExpectations(t)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestUpdateAccountMetaData_SQLExecutionError(t *testing.T) {
	userID := "user123"

	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "UpdateAccountMetaData RPC called").Once()
	mockLogger.On("Error", mock.Anything, mock.Anything).Once()

	payload := "{}"
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, userID)

	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Simulate an error while executing the SQL statement
	dbMock.ExpectExec(`^UPDATE users SET metadata .*`).
		WithArgs(payload, userID).
		WillReturnError(fmt.Errorf("SQL execution error"))

	nk := new(mocks.NakamaModule) // Mock NakamaModule

	// Call the function
	result, err := UpdateAccountMetaData(ctx, mockLogger, db, nk, payload)

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrInternalError, err)
	mockLogger.AssertExpectations(t)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}
