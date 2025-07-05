package exceptions

type ErrorCode string

const (
	ValidationFailed   ErrorCode = "VALIDATION_FAILED"
	DatabaseFailer     ErrorCode = "DATABASE_FAILER"
	InvalidCommandType ErrorCode = "INVALID_COMMAND_TYPE"
)
