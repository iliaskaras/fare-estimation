/*
Package distances
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package distances

import "strings"

const (
	HaversineMethod       = "haversine"
	defaultDistanceMethod = HaversineMethod
)

var supportedDistanceMethods = []string{"haversine"}

// GetDistanceCalculatorService is responsible for returning the correct DistanceCalculatorService implementor,
// based on the distance method provided.
func GetDistanceCalculatorService(distanceMethod string) (DistanceCalculatorService, error) {
	if distanceMethod == "" {
		distanceMethod = defaultDistanceMethod
	}

	if distanceMethod == "haversine" {
		return NewHaversineDistanceService(), nil
	}

	return nil, NewDistanceMethodError(
		UnsupportedDistanceMethod,
		"provided distance method: "+distanceMethod+", "+
			"must be one of the: "+strings.Join(supportedDistanceMethods[:], ",")+" \n",
	)
}
