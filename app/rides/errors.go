/*
Package rides
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package rides

import (
	"errors"
)

var (
	ErrorParsingRidePosition = errors.New("error while parsing ride position")
	InvalidLPosition         = errors.New("error while parsing ride position")
)
