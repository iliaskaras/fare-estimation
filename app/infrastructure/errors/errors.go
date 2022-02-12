/*
Package errors
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package errors

import (
	"errors"
	"fmt"
)

type BaseAppError struct {
	Context string
	Err     error
}

func (bae BaseAppError) Error() string {
	return fmt.Sprintf("error details: %v, info: %s\n", bae.Err, bae.Context)
}

// Unwrap implements the errors.Unwrap interface to solve equality
// problems when we make use of BaseAppError
func (bae BaseAppError) Unwrap() error {
	// Returns inner error
	return bae.Err
}

func NewBaseAppError(err error, additionalInfo string) BaseAppError {
	return BaseAppError{
		Context: additionalInfo,
		Err:     err,
	}
}

var (
	InvalidInputError = errors.New("invalid input")
)
