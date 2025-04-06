package errorutils

import (
	"log"
)

// CheckFatal logs the error and exits the program.
// Useful in main or critical setup code.
func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CheckPanic panics if an error occurs.
// Useful for prototyping or unrecoverable failures.
func CheckPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// ReturnIfError returns the error directly.
// Useful for early return in functions.
func ReturnIfError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

// Must unwraps a (value, error) return.
// It panics if the error is non-nil.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// LogIfError logs the error with a custom context.
// Does not stop execution.
func LogIfError(err error, context string) {
	if err != nil {
		log.Printf("Error in %s: %v", context, err)
	}
}

// HandleWithCallback executes a custom function when an error occurs.
func HandleWithCallback(err error, handler func(error)) {
	if err != nil {
		handler(err)
	}
}
