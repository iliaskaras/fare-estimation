/*
Package rides
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package rides

import (
	"github.com/iliaskaras/fare-estimation/app/distances"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Tests the GetRidePositionService initializes and returns the RidePositionService.
func TestGetRidePositionService(t *testing.T) {
	distanceCalculatorMethod, _ := distances.GetDistanceCalculatorService(distances.HaversineMethod)
	ridePositionService, err := GetRidePositionService(distanceCalculatorMethod)
	assert.NoError(t, err)

	returnedServiceType := reflect.TypeOf(ridePositionService).String()
	expectedServiceType := "*rides.RidePositionService"

	assert.Equal(t, expectedServiceType, returnedServiceType)

}
