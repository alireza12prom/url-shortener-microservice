package exceptions

type ErrorCode string

const (
	UsernameAlreadyTaken    ErrorCode = "USERNAME_ALREADY_TAKEN"
	WrongPasswordProvided   ErrorCode = "WRONG_PASSWORD_PROVIDED"
	WrongAuthenticationPack ErrorCode = "WRONG_AUTHENTICATION_PACK"
	LoginFailed             ErrorCode = "LOGIN_FAILED"
	ValidationFailed        ErrorCode = "VALIDATION_FAILED"
	DatabaseFailer          ErrorCode = "DATABASE_FAILER"
	HashAlreadyExists       ErrorCode = "HASH_ALREADY_EXISTS"
	ShortURLNotFound        ErrorCode = "SHORT_URL_NOT_FOUND"
)
