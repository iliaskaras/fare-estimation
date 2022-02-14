/*
Package distances
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package distances

import "math"

const (
	earthKmRadius = 6371
)

type DistanceCalculatorService interface {
	GetDistance(lat1Degree, lng1Degree, lat2Degree, lng2Degree float64) (float64, error)
}

// HaversineDistanceService is the DistanceCalculatorService implementor that calculates distance using the
// Haversine distance method.
type HaversineDistanceService struct{}

func NewHaversineDistanceService() DistanceCalculatorService {
	return &HaversineDistanceService{}
}

// GetDistance returns the Haversine distance provided degree positions.
// More details about the Haversine distance formula calculation can be found at:
// https://www.igismap.com/haversine-formula-calculate-geographic-distance-earth/.
func (hs *HaversineDistanceService) GetDistance(
	lat1Degree, lng1Degree, lat2Degree, lng2Degree float64,
) (float64, error) {

	lat1R := mapDegreesToRadians(lat1Degree)
	lng1R := mapDegreesToRadians(lng1Degree)

	lat2R := mapDegreesToRadians(lat2Degree)
	lng2R := mapDegreesToRadians(lng2Degree)

	latDiffR := lat2R - lat1R
	lngDiffR := lng2R - lng1R

	//a = sin²(ΔlatDifference/2) + cos(lat1)*cos(lt2)*sin²(ΔlonDifference/2)
	a := math.Pow(math.Sin(latDiffR/2), 2) + math.Cos(lat1R)*math.Cos(lat2R)*math.Pow(math.Sin(lngDiffR/2), 2)
	//c = 2*atan2(√a, √(1−a))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	//d = R*c
	distance := earthKmRadius * c

	return distance, nil
}
