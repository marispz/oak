package common

import "github.com/heroiclabs/nakama-common/runtime"

var (
	ErrUserNotFound        = runtime.NewError("user not found", RpcCodeNotFound)
	ErrNotFound            = runtime.NewError("not found", RpcCodeNotFound)
	ErrMetadataSizeLimit   = runtime.NewError("metadata size limit can not exceed 16KB", RpcCodeInvalidArgument)
	ErrMarshallingError    = runtime.NewError("marshalling error", RpcCodeInternal)
	ErrUnMarshallingError  = runtime.NewError("unmarshalling error", RpcCodeInternal)
	ErrInternalError       = runtime.NewError("internal error", RpcCodeInternal)
	ErrS2SPermissionDenied = runtime.NewError("rpc is only callable via server to server", RpcCodePermissionDenied)
)
