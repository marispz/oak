package rpc

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/redis/go-redis/v9"
	"oak/common"
)

// Embed the game_config.json file
//
//go:embed config/game_config.json
var gameConfigJSON []byte

// ReadGameConfigurationFromFile reads the game configuration from the embedded JSON file.
func ReadGameConfigurationFromFile(ctx context.Context, logger runtime.Logger, _ *sql.DB, _ *redis.Client, _ runtime.NakamaModule, _ string) (string, error) {
	logger.Debug("ReturnGameConfigurationFromFile RPC called")

	// Get the user ID from the context
	_, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		logger.Error("Context did not contain user ID.")
		return common.EmptyString, common.ErrUserNotFound
	}

	return LoadGameConfig(logger)
}

// ReadGameConfigurationFromStorage reads the game configuration from the storage.
func ReadGameConfigurationFromStorage(ctx context.Context, logger runtime.Logger, _ *sql.DB, _ *redis.Client, nk runtime.NakamaModule, _ string) (string, error) {
	logger.Debug("ReadGameConfigurationFromStorage RPC called")

	// Get the user ID from the context
	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		logger.Error("Context did not contain user ID.")
		return common.EmptyString, common.ErrUserNotFound
	}

	// Read the game configuration from the storage
	obj, err := nk.StorageRead(ctx, []*runtime.StorageRead{{
		Collection: common.StorageConfiguration,
		Key:        common.StorageGameConfigKey,
		UserID:     userID,
	}})
	if err != nil {
		logger.Error("StorageWrite error: %+v", err)
		return common.EmptyString, common.ErrInternalError
	}

	if len(obj) == 0 {
		logger.Error("Game configuration not found")
		return common.EmptyString, common.ErrNotFound
	}

	return obj[0].GetValue(), nil
}

// LoadGameConfig loads the game configuration from the embedded JSON file.
var LoadGameConfig = func(logger runtime.Logger) (string, error) {
	// Original function code
	var config common.GameConfig
	if err := json.Unmarshal(gameConfigJSON, &config); err != nil {
		logger.Error("Error decoding embedded JSON: %+v", err)
		return common.EmptyString, common.ErrUnMarshallingError
	}
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		logger.Error("Error encoding JSON: %+v", err)
		return common.EmptyString, common.ErrMarshallingError
	}
	return string(configJSON), nil
}
