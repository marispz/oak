package common

const (
	EmptyString          = ""
	StorageConfiguration = "configuration"
	StorageGameConfigKey = "game_configuration"
)

const (
	RpcCodeOK                 = 0  // OK
	RpcCodeCancelled          = 1  // The operation was cancelled (typically by the caller).
	RpcCodeUnknown            = 2  // Unknown error; mostly used when an error is not categorized.
	RpcCodeInvalidArgument    = 3  // Client specified an invalid argument.
	RpcCodeDeadlineExceeded   = 4  // Deadline expired before the operation could complete.
	RpcCodeNotFound           = 5  // Some requested entity was not found.
	RpcCodeAlreadyExists      = 6  // The entity that a client attempted to create already exists.
	RpcCodePermissionDenied   = 7  // The caller does not have permission to execute the specified operation.
	RpcCodeResourceExhausted  = 8  // Some resource has been exhausted, like a quota or file descriptor limit.
	RpcCodeFailedPrecondition = 9  // Operation was rejected because the system is not in a state required for the operation's execution.
	RpcCodeAborted            = 10 // Operation was aborted, typically due to a concurrency issue like a transaction failure.
	RpcCodeOutOfRange         = 11 // Operation was attempted past the valid range.
	RpcCodeUnimplemented      = 12 // Operation is not implemented or not supported/enabled in this service.
	RpcCodeInternal           = 13 // Internal errors; these should be rare and typically indicate a bug.
	RpcCodeUnavailable        = 14 // Service is currently unavailable. This is typically a transient condition.
	RpcCodeDataLoss           = 15 // Unrecoverable data loss or corruption.
	RpcCodeUnauthenticated    = 16 // The request does not have valid authentication credentials.
)
