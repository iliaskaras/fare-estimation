/*
Package rides
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package rides

import (
	"github.com/iliaskaras/fare-estimation/app/distances"
)

// GetRidePositionService is responsible for initializing and injecting all the dependencies
// of the RidePositionService.
func GetRidePositionService(
	distanceCalculatorMethod distances.DistanceCalculatorService,
) (*RidePositionService, error) {

	return NewRidePositionService(
		distanceCalculatorMethod,
	), nil
}
