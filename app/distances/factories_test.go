/*
Package distances
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package distances

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

type GetFileServiceTestCase struct {
	distanceMethod string
	expectedResult DistanceCalculatorService
	expectedError  DistanceMethodError
}

// Tests the GetDistanceCalculatorService return a haversine implementor based on the requested distanceMethod.
func TestGetDistanceCalculatorServiceReturnHaversine(t *testing.T) {
	distanceCalculatorService, err := GetDistanceCalculatorService(HaversineMethod)
	assert.NoError(t, err)

	_, ok := distanceCalculatorService.(DistanceCalculatorService)
	assert.Equal(t, true, ok)
	returnedDistanceServiceType := reflect.TypeOf(distanceCalculatorService).String()
	expectedDistanceServiceType := "*distances.HaversineDistanceService"

	assert.Equal(t, expectedDistanceServiceType, returnedDistanceServiceType)

}

// Tests the GetDistanceCalculatorService return a haversine when provided method is empty.
func TestGetDistanceCalculatorServiceReturnHaversineDefaultMethod(t *testing.T) {
	distanceCalculatorService, err := GetDistanceCalculatorService("")
	assert.NoError(t, err)

	_, ok := distanceCalculatorService.(DistanceCalculatorService)
	assert.Equal(t, true, ok)
	returnedDistanceServiceType := reflect.TypeOf(distanceCalculatorService).String()
	expectedDistanceServiceType := "*distances.HaversineDistanceService"

	assert.Equal(t, expectedDistanceServiceType, returnedDistanceServiceType)

}

// Tests the GetFileService return a csvFileService in case a .csv type of file is provided.
func TestGetDistanceCalculatorServiceReturnErrorWhenDistanceIsInvalid(t *testing.T) {
	distanceCalculatorService, err := GetDistanceCalculatorService("invalidDistanceMethod")
	assert.Error(t, err)

	_, ok := err.(DistanceMethodError)
	assert.Equal(t, true, ok)
	assert.Equal(t, nil, distanceCalculatorService)
	assert.Equal(
		t,
		NewDistanceMethodError(
			UnsupportedDistanceMethod,
			"provided distance method: invalidDistanceMethod, "+
				"must be one of the: "+strings.Join(supportedDistanceMethods[:], ",")+" \n",
		), err,
	)

}
