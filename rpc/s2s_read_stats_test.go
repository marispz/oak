package rpc

import (
	"context"
	"oak/mocks"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/assert"
	"oak/common"
)

// Test for S2SReadStats when the context contains a valid userID
func TestS2SReadStats_UserIDInContext(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "S2SReadStats RPC called").Once()
	mockLogger.On("Error", "Rpc was called by a user").Once()

	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, "test_user_id")

	// Call the function
	result, err := S2SReadStats(ctx, mockLogger, nil, nil, "")

	// Assertions
	assert.Equal(t, common.EmptyString, result)
	assert.Equal(t, common.ErrS2SPermissionDenied, err)
	mockLogger.AssertExpectations(t)
}

// Test for S2SReadStats when the context does not contain a userID
func TestS2SReadStats_NoUserIDInContext(t *testing.T) {
	// Setup
	mockLogger := new(mocks.Logger)
	mockLogger.On("Debug", "S2SReadStats RPC called").Once()
	mockLogger.On("Info", "processing S2S read stats %s", "").Once()

	ctx := context.Background()

	// Call the function
	result, err := S2SReadStats(ctx, mockLogger, nil, nil, "")

	// Assertions
	assert.Equal(t, "", result)
	assert.NoError(t, err)
	mockLogger.AssertExpectations(t)
}
