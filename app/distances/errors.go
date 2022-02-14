/*
Package distances
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package distances

import (
	"errors"
	baseAppErrors "github.com/iliaskaras/fare-estimation/app/infrastructure/errors"
)

type DistanceMethodError struct {
	baseAppErrors.BaseAppError
}

func NewDistanceMethodError(err error, additionalInfo string) DistanceMethodError {
	return DistanceMethodError{
		BaseAppError: baseAppErrors.NewBaseAppError(err, additionalInfo),
	}
}

var (
	UnsupportedDistanceMethod = errors.New("unsupported distance method")
)
