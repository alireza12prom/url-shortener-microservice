package exceptions

type ErrorCode string

const (
	ShortURLNotFound ErrorCode = "SHORT_URL_NOT_FOUND"
	DatabaseFailer   ErrorCode = "DATABASE_FAILER"
)
