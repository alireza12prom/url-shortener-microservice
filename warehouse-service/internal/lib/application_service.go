package lib

type ApplicationService interface {
	Exec(command interface{}) error
}
