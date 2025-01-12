package main

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	"oak/hook"
	"oak/rpc"
)

const (
	rpcUpdateAccountMetaData            = "update_account_metaData"
	rpcReadGameConfigurationFromFile    = "read_game_config_from_file"
	rpcReadGameConfigurationFromStorage = "read_game_config_from_storage"
	rpcS2SReadGameStats                 = "read_game_stats"
)

func InitModule(_ context.Context, logger runtime.Logger, _ *sql.DB, _ runtime.NakamaModule, initializer runtime.Initializer) error {
	// Register RPCs
	err := initializer.RegisterRpc(rpcUpdateAccountMetaData, rpc.UpdateAccountMetaData)
	if err != nil {
		logger.Error("Failed to register RPC %+v", err)
	}

	err = initializer.RegisterRpc(rpcReadGameConfigurationFromFile, rpc.ReadGameConfigurationFromFile)
	if err != nil {
		logger.Error("Failed to register RPC %+v", err)
	}

	err = initializer.RegisterRpc(rpcReadGameConfigurationFromStorage, rpc.ReadGameConfigurationFromStorage)
	if err != nil {
		logger.Error("Failed to register RPC %+v", err)
	}

	err = initializer.RegisterRpc(rpcS2SReadGameStats, rpc.S2SReadStats)
	if err != nil {
		logger.Error("Error in registering RPC %+v", err)
		return err
	}

	// Register after hooks.
	if err := initializer.RegisterAfterAuthenticateDevice(hook.InitializeUser); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}

	logger.Info("Module loaded")
	return nil
}
