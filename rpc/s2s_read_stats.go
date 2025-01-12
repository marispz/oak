package rpc

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/redis/go-redis/v9"
	"oak/common"
)

// S2SReadStats is an example of a server to server function that reads stats for a user.
func S2SReadStats(ctx context.Context, logger runtime.Logger, _ *sql.DB, _ *redis.Client, _ runtime.NakamaModule, _ string) (string, error) {
	logger.Debug("S2SReadStats RPC called")

	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if ok && userID != "" {
		logger.Error("Rpc was called by a user")
		return common.EmptyString, common.ErrS2SPermissionDenied
	}

	logger.Info("processing S2S read stats %s", userID)

	return common.EmptyString, nil
}
