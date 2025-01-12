package rpc

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
	"oak/common"
)

type (
	AccountMetaDataResponse struct {
		Status common.Status `json:"status"`
	}
)

// UpdateAccountMetaData updates the metadata of the user account
func UpdateAccountMetaData(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("UpdateAccountMetaData RPC called")
	// Get the user ID from the context
	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		logger.Error("Context did not contain user ID.")
		return common.EmptyString, common.ErrUserNotFound
	}

	// Metadata is limited to 16KB per user
	if len(payload) > 16384 {
		logger.Error("Metadata size limit exceeded")
		return common.EmptyString, common.ErrMetadataSizeLimit
	}

	// Request struct to unmarshal the payload
	metaData := make(map[string]any)

	// Marshal the metadata struct to JSON
	err := json.Unmarshal([]byte(payload), &metaData)
	if err != nil {
		logger.Error("Cannot unmarshal metadata: %+v", err)
		return common.EmptyString, common.ErrUnMarshallingError
	}

	// SQL statement to update the metadata
	stmt := "UPDATE users SET metadata = $1::jsonb WHERE id = $2"

	// Execute the update with parameters
	_, err = db.ExecContext(ctx, stmt, payload, userID)
	if err != nil {
		logger.Error("Cannot update metadata: %+v", err)
		return common.EmptyString, common.ErrInternalError
	}

	resp := &AccountMetaDataResponse{
		Status: common.StatusSuccess,
	}

	// Marshal the response struct to JSON
	respJSON, err := json.Marshal(resp)
	if err != nil {
		logger.Error("Cannot marshal response %+v", err)
		return common.EmptyString, common.ErrMarshallingError
	}

	// Returns the status
	return string(respJSON), nil
}
