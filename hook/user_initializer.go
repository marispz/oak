package hook

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"oak/common"
	"oak/rpc"
)

// InitializeUser is invoked when user is created for the first time to initialize the user's data.
func InitializeUser(ctx context.Context, logger runtime.Logger, _ *sql.DB, nk runtime.NakamaModule, out *api.Session, _ *api.AuthenticateDeviceRequest) error {
	if out.Created {
		// Get the user ID from the context
		userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
		if !ok {
			logger.Error("Context did not contain user ID.")
			return common.ErrUserNotFound
		}

		// Load the game configuration from the file
		gameConfiguration, err := rpc.LoadGameConfig(logger)
		if err != nil {
			return err
		}

		// Write the game configuration to the storage
		_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{{
			Collection:      common.StorageConfiguration,
			Key:             common.StorageGameConfigKey,
			PermissionRead:  1,
			PermissionWrite: 1,
			Value:           gameConfiguration,
			UserID:          userID,
		}})
		if err != nil {
			logger.Error("StorageWrite error: %+v", err)
			return common.ErrInternalError
		}
	}

	return nil
}
