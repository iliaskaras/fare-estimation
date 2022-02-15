/*
Package rides
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package rides

import (
	"github.com/iliaskaras/fare-estimation/app/distances"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests the RidePositionService.FilterOnSegmentSpeed filters the RidePosition by their RideSegment speed.
// This test case, tests all the possible scenarios that might happen in the method.
// More specifically:
// 1. The first RideID's RideSegment list, cover the case that one of the segments is filtered out due to speed.
// 2. The second RideID have odd number of RidePosition.
// 3. The third RideID have only two RidePosition and one generated RideSegment.
// 4. The forth RideID have only one RidePosition and thus zero generated RideSegment.
func TestFilterOnSegmentSpeedSuccessfulExecution(t *testing.T) {
	distanceCalculatorMethod, _ := distances.GetDistanceCalculatorService(distances.HaversineMethod)
	ridePositionService, _ := GetRidePositionService(
		distanceCalculatorMethod,
	)
	var expectedRideSegments = [][]RideSegment{
		{
			RideSegment{
				RideID: 1,
				RidePositions: [2]RidePosition{
					{
						Id:        1,
						Lat:       37.966660,
						Lng:       23.728308,
						Timestamp: 1405594957,
					},
					{
						Id:        1,
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					},
				},
				Speed:           2.15504358006091,
				DistanceCovered: 0.005387608950152276,
			},
			RideSegment{
				RideID: 1,
				RidePositions: [2]RidePosition{
					{
						Id:        1,
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					},
					{
						Id:        1,
						Lat:       37.966625,
						Lng:       23.728263,
						Timestamp: 1405594974,
					},
				},
				Speed:           0.10007543437580146,
				DistanceCovered: 0.0002223898541684477,
			},
			RideSegment{
				RideID: 1,
				RidePositions: [2]RidePosition{
					{
						Id:        1,
						Lat:       37.966625,
						Lng:       23.728263,
						Timestamp: 1405594974,
					},
					{
						Id:        1,
						Lat:       37.966613,
						Lng:       23.728375,
						Timestamp: 1405594984,
					},
				},
				Speed:           3.5670510679208447,
				DistanceCovered: 0.009908475188669013,
			},
			RideSegment{
				RideID: 1,
				RidePositions: [2]RidePosition{
					{
						Id:        1,
						Lat:       37.966613,
						Lng:       23.728375,
						Timestamp: 1405594984,
					},
					{
						Id:        1,
						Lat:       37.954302,
						Lng:       23.713370,
						Timestamp: 1405595284,
					},
				},
				Speed:           22.782481162615245,
				DistanceCovered: 1.8985400968846038,
			},
		},
		{
			RideSegment{
				RideID: 2,
				RidePositions: [2]RidePosition{
					{
						Id:        2,
						Lat:       37.946545,
						Lng:       23.754918,
						Timestamp: 1405591065,
					},
					{
						Id:        2,
						Lat:       37.946545,
						Lng:       23.754918,
						Timestamp: 1405591073,
					},
				},
				Speed:           0,
				DistanceCovered: 0,
			},
			RideSegment{
				RideID: 2,
				RidePositions: [2]RidePosition{
					{
						Id:        2,
						Lat:       37.946545,
						Lng:       23.754918,
						Timestamp: 1405591073,
					},
					{
						Id:        2,
						Lat:       37.946545,
						Lng:       23.754918,
						Timestamp: 1405591084,
					},
				},
				Speed:           0,
				DistanceCovered: 0,
			},
		},
		{
			RideSegment{
				RideID: 3,
				RidePositions: [2]RidePosition{
					{
						Id:        3,
						Lat:       37.926738,
						Lng:       23.935701,
						Timestamp: 1405591810,
					},
					{
						Id:        3,
						Lat:       37.927245,
						Lng:       23.935,
						Timestamp: 1405591818,
					},
				},
				Speed:           37.538200556271114,
				DistanceCovered: 0.08341822345838025,
			},
		},
		nil,
	}
	ridePositionsChan := make(chan []RidePosition)
	rideSegmentsChan := make(chan []RideSegment)

	go func() {
		ridePositionsChan <- []RidePosition{
			{
				Id:        1,
				Lat:       37.966660,
				Lng:       23.728308,
				Timestamp: 1405594957,
			},
			{
				Id:        1,
				Lat:       37.966627,
				Lng:       23.728263,
				Timestamp: 1405594966,
			},
			{
				Id:        1,
				Lat:       37.966625,
				Lng:       23.728263,
				Timestamp: 1405594974,
			},
			{
				Id:        1,
				Lat:       37.966613,
				Lng:       23.728375,
				Timestamp: 1405594984,
			},
			{
				Id:        1,
				Lat:       37.954302,
				Lng:       23.713370,
				Timestamp: 1405595284,
			},
			{
				Id:        1,
				Lat:       37.938042,
				Lng:       23.692308,
				Timestamp: 1405595362,
			},
		}
		ridePositionsChan <- []RidePosition{
			{
				Id:        2,
				Lat:       37.946545,
				Lng:       23.754918,
				Timestamp: 1405591065,
			},
			{
				Id:        2,
				Lat:       37.946545,
				Lng:       23.754918,
				Timestamp: 1405591073,
			},
			{
				Id:        2,
				Lat:       37.946545,
				Lng:       23.754918,
				Timestamp: 1405591084,
			},
		}
		ridePositionsChan <- []RidePosition{
			{
				Id:        3,
				Lat:       37.926738,
				Lng:       23.935701,
				Timestamp: 1405591810,
			},
			{
				Id:        3,
				Lat:       37.927245,
				Lng:       23.935000,
				Timestamp: 1405591818,
			},
		}
		ridePositionsChan <- []RidePosition{
			{
				Id:        4,
				Lat:       37.926738,
				Lng:       23.935701,
				Timestamp: 1405591810,
			},
		}
		close(ridePositionsChan)
	}()

	go func() {
		ridePositionService.FilterOnSegmentSpeed(
			ridePositionsChan,
			rideSegmentsChan,
		)
		close(rideSegmentsChan)
	}()

	i := 0
	for rideSegmentsResult := range rideSegmentsChan {
		assert.Equal(t, expectedRideSegments[i], rideSegmentsResult)
		i += 1
	}

}
