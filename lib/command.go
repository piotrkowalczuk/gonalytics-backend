package lib

// Executor is an interface that should implement every Command.
type Executor interface {
	Execute() (string, error)
}
