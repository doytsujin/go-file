package file

import "fmt"

type IOError struct {
	message string
}

func NewIOError(message string) error {
	return &IOError{
		message: message,
	}
}

func (e IOError) Error() string {
	return e.message
}

type LockError struct {
	message string
}

func NewLockError(message string) error {
	return &LockError{
		message: message,
	}
}

func (e LockError) Error() string {
	return e.message
}

type TimeoutError struct {
	message string
}

func NewTimeoutError(path string) error {
	return &TimeoutError{
		message: fmt.Sprintf("file %s: lock waiting time exceeded", path),
	}
}

func (e TimeoutError) Error() string {
	return e.message
}

type ContextIsDone struct {
	message string
}

func NewContextIsDone(message string) error {
	return &ContextIsDone{
		message: message,
	}
}

func (e ContextIsDone) Error() string {
	return e.message
}
