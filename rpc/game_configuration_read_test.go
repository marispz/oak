package rpc

import (
	"context"
	"fmt"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"oak/common"
	"oak/mocks"
	"testing"
)

func TestReadGameConfigurationFromFile_Success(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "ReturnGameConfigurationFromFile RPC called").Once()

	// Context with a valid user ID
	userID := "user123"
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, userID)

	// Mock the LoadGameConfig function to return a predefined game config
	expectedConfig := `{
	  "game_name": "Super Fun Game",
	  "max_players": 100
	}`
	mockLoadGameConfig := func(logger runtime.Logger) (string, error) {
		return expectedConfig, nil
	}
	LoadGameConfig = mockLoadGameConfig

	// Call the function
	result, err := ReadGameConfigurationFromFile(ctx, mockLogger, nil, nil, "")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, result)
	mockLogger.AssertExpectations(t)
}

func TestReadGameConfigurationFromStorage_Success(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "ReadGameConfigurationFromStorage RPC called").Once()

	// Context with a valid user ID
	userID := "user123"
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, userID)

	// Mock the Nakama module to simulate a successful storage read
	expectedConfig := `{
  "game_name": "Super Fun Game",
  "max_players": 100
}`
	nk := new(mocks.NakamaModule)
	nk.On("StorageRead", ctx, []*runtime.StorageRead{
		{
			Collection: common.StorageConfiguration,
			Key:        common.StorageGameConfigKey,
			UserID:     userID,
		},
	}).Return([]*api.StorageObject{
		{Value: expectedConfig},
	}, nil)

	// Call the function
	result, err := ReadGameConfigurationFromStorage(ctx, mockLogger, nil, nk, "")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, result)
	mockLogger.AssertExpectations(t)
	nk.AssertExpectations(t)
}

func TestReadGameConfigurationFromFile_MissingUserID(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "ReturnGameConfigurationFromFile RPC called").Once()
	mockLogger.On("Error", mock.Anything, mock.Anything).Once()

	// Context without user ID
	ctx := context.Background()

	expectedConfig := `{
	  "game_name": "Super Fun Game",
	  "max_players": 100
	}`
	// Mock the LoadConfig function (though it won't be called in this case)
	mockLoadGameConfig := func(logger runtime.Logger) (string, error) {
		return expectedConfig, nil
	}
	LoadGameConfig = mockLoadGameConfig

	// Call the function
	result, err := ReadGameConfigurationFromFile(ctx, mockLogger, nil, nil, "")

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrUserNotFound, err)
	mockLogger.AssertExpectations(t)
}

func TestReadGameConfigurationFromFile_LoadGameConfigError(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "ReturnGameConfigurationFromFile RPC called").Once()

	// Context with a valid user ID
	userID := "user123"
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, userID)

	// Mock the LoadGameConfig function to return an error
	mockLoadGameConfig := func(logger runtime.Logger) (string, error) {
		return common.EmptyString, common.ErrUnMarshallingError
	}
	LoadGameConfig = mockLoadGameConfig

	// Call the function
	result, err := ReadGameConfigurationFromFile(ctx, mockLogger, nil, nil, "")

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrUnMarshallingError, err)
	mockLogger.AssertExpectations(t)
}

func TestReadGameConfigurationFromStorage_StorageReadError(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "ReadGameConfigurationFromStorage RPC called").Once()
	mockLogger.On("Error", mock.Anything, mock.Anything).Once()

	// Context with a valid user ID
	userID := "user123"
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, userID)

	// Mock the Nakama module to simulate a storage read error
	nk := new(mocks.NakamaModule)
	nk.On("StorageRead", ctx, []*runtime.StorageRead{
		{
			Collection: common.StorageConfiguration,
			Key:        common.StorageGameConfigKey,
			UserID:     userID,
		},
	}).Return(nil, fmt.Errorf("storage read error"))

	// Call the function
	result, err := ReadGameConfigurationFromStorage(ctx, mockLogger, nil, nk, "")

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrInternalError, err)
	mockLogger.AssertExpectations(t)
	nk.AssertExpectations(t)
}

func TestReadGameConfigurationFromStorage_GameConfigNotFound(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "ReadGameConfigurationFromStorage RPC called").Once()
	mockLogger.On("Error", mock.Anything, mock.Anything).Once()

	// Context with a valid user ID
	userID := "user123"
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, userID)

	// Mock the Nakama module to simulate that no configuration was found in storage
	nk := new(mocks.NakamaModule)
	nk.On("StorageRead", ctx, []*runtime.StorageRead{
		{
			Collection: common.StorageConfiguration,
			Key:        common.StorageGameConfigKey,
			UserID:     userID,
		},
	}).Return([]*api.StorageObject{}, nil)

	// Call the function
	result, err := ReadGameConfigurationFromStorage(ctx, mockLogger, nil, nk, "")

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrNotFound, err)
	mockLogger.AssertExpectations(t)
	nk.AssertExpectations(t)
}

func TestReadGameConfigurationFromFile_InvalidUserID(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "ReturnGameConfigurationFromFile RPC called").Once()
	mockLogger.On("Error", mock.Anything, mock.Anything).Once()

	// Context with an invalid user ID (non-string type)
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, 12345) // Invalid type

	expectedConfig := `{
	  "game_name": "Super Fun Game",
	  "max_players": 100
	}`
	// Mock the LoadConfig function (though it won't be called in this case)
	mockLoadGameConfig := func(logger runtime.Logger) (string, error) {
		return expectedConfig, nil
	}
	LoadGameConfig = mockLoadGameConfig

	// Call the function
	result, err := ReadGameConfigurationFromFile(ctx, mockLogger, nil, nil, "")

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrUserNotFound, err)
	mockLogger.AssertExpectations(t)
}

func TestReadGameConfigurationFromFile_ContextErrorHandling(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "ReturnGameConfigurationFromFile RPC called").Once()
	mockLogger.On("Error", mock.Anything, mock.Anything).Once()

	// Context without user ID
	ctx := context.Background()

	// Expected config to be returned
	expectedConfig := `{
	  "game_name": "Super Fun Game",
	  "max_players": 100
	}`

	// Mock the LoadConfig function (though it won't be called in this case)
	mockLoadGameConfig := func(logger runtime.Logger) (string, error) {
		return expectedConfig, nil
	}
	LoadGameConfig = mockLoadGameConfig

	// Call the function
	result, err := ReadGameConfigurationFromFile(ctx, mockLogger, nil, nil, "")

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrUserNotFound, err)
	mockLogger.AssertExpectations(t)
}
