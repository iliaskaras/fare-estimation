/*
Package distances
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package distances

import "math"

// mapDegreesToRadians maps a degree position to radians.
func mapDegreesToRadians(degreePos float64) float64 {
	return degreePos * math.Pi / 180
}
