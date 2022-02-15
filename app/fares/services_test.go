/*
Package fares
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package fares

import (
	"github.com/iliaskaras/fare-estimation/app/rides"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests the FareService.Estimate fare estimation on the received rides.RideSegment.
// This test case, tests a successful estimation on multiple rides.RideSegment scenarios.
func TestEstimateSuccessfulExecution(t *testing.T) {
	rideSegmentsChan := make(chan []rides.RideSegment)
	faresChan := make(chan Fare)

	fareService := NewFareService()
	var expectedFareResults = []Fare{
		*NewFare(
			1,
			3.47,
		),
		*NewFare(
			2,
			3.47,
		),
		*NewFare(
			3,
			3.47,
		),
	}
	go func() {
		rideSegmentsChan <- []rides.RideSegment{
			{
				RideID: 1,
				RidePositions: [2]rides.RidePosition{
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
			{
				RideID: 1,
				RidePositions: [2]rides.RidePosition{
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
			{
				RideID: 1,
				RidePositions: [2]rides.RidePosition{
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
			{
				RideID: 1,
				RidePositions: [2]rides.RidePosition{
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
		}

		rideSegmentsChan <- []rides.RideSegment{
			{
				RideID: 2,
				RidePositions: [2]rides.RidePosition{
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
			{
				RideID: 2,
				RidePositions: [2]rides.RidePosition{
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
		}

		rideSegmentsChan <- []rides.RideSegment{
			{
				RideID: 3,
				RidePositions: [2]rides.RidePosition{
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
		}
		rideSegmentsChan <- nil
		close(rideSegmentsChan)
	}()

	go func() {
		fareService.Estimate(
			rideSegmentsChan,
			faresChan,
		)
	}()

	i := 0
	for faresResult := range faresChan {
		assert.Equal(t, expectedFareResults[i], faresResult)
		i += 1
	}

}

// Tests the FareService.Estimate fare estimation on the case where the calculated
// fare is less than the minimum.
func TestEstimateWhenTheFareEstimateIsTheMinimum(t *testing.T) {
	rideSegmentsChan := make(chan []rides.RideSegment)
	faresChan := make(chan Fare)

	fareService := NewFareService()
	var expectedFareResults = []Fare{
		*NewFare(
			1,
			3.47,
		),
	}
	go func() {
		rideSegmentsChan <- []rides.RideSegment{
			{
				RideID: 1,
				RidePositions: [2]rides.RidePosition{
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
		}

		close(rideSegmentsChan)
	}()

	go func() {
		fareService.Estimate(
			rideSegmentsChan,
			faresChan,
		)
	}()

	i := 0
	for faresResult := range faresChan {
		assert.Equal(t, expectedFareResults[i], faresResult)
		i += 1
	}

}

// Tests the FareService.Estimate fare estimation on the case where the calculated
// fare is above the minimum and the ride speed is idle.
func TestEstimateWhenFareIsAboveMinimumAndIdle(t *testing.T) {
	rideSegmentsChan := make(chan []rides.RideSegment)
	faresChan := make(chan Fare)

	fareService := NewFareService()
	var expectedFareResults = []Fare{
		*NewFare(
			1,
			4.16,
		),
	}
	go func() {
		rideSegmentsChan <- []rides.RideSegment{
			{
				RideID: 1,
				RidePositions: [2]rides.RidePosition{
					{
						Id:        1,
						Lat:       37.966660,
						Lng:       23.728308,
						Timestamp: 1405594100,
					},
					{
						Id:        1,
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					},
				},
				Speed:           9,
				DistanceCovered: 0.005387608950152276,
			},
		}

		close(rideSegmentsChan)
	}()

	go func() {
		fareService.Estimate(
			rideSegmentsChan,
			faresChan,
		)
	}()

	i := 0
	for faresResult := range faresChan {
		assert.Equal(t, expectedFareResults[i], faresResult)
		i += 1
	}

}

// Tests the FareService.Estimate fare estimation on the case where the calculated
// fare is above the minimum and is moving at day hours.
func TestEstimateWhenFareIsAboveMinimumAndMovingDay(t *testing.T) {
	rideSegmentsChan := make(chan []rides.RideSegment)
	faresChan := make(chan Fare)

	fareService := NewFareService()
	var expectedFareResult = Fare{
		1,
		5,
	}

	go func() {
		rideSegmentsChan <- []rides.RideSegment{
			{
				RideID: 1,
				RidePositions: [2]rides.RidePosition{
					{
						Id:        1,
						Lat:       37.966660,
						Lng:       23.728308,
						Timestamp: 1405594100,
					},
					{
						Id:        1,
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1405594966,
					},
				},
				Speed: 11,
				// Not realistic for a single RideSegment to have that big distance covered,
				// but is used for emulate the test scenario.
				DistanceCovered: 5000,
			},
		}

		close(rideSegmentsChan)
	}()

	go func() {
		fareService.Estimate(
			rideSegmentsChan,
			faresChan,
		)

	}()

	i := 0
	for faresResult := range faresChan {
		assert.Equal(t, expectedFareResult, faresResult)
		i += 1
	}

}

// Tests the FareService.Estimate fare estimation on the case where the calculated
// fare is above the minimum and is moving at night hours.
func TestEstimateWhenFareIsAboveMinimumAndMovingNight(t *testing.T) {
	rideSegmentsChan := make(chan []rides.RideSegment)
	faresChan := make(chan Fare)

	fareService := NewFareService()
	var expectedFareResult = Fare{
		1,
		7.8,
	}

	go func() {
		rideSegmentsChan <- []rides.RideSegment{
			{
				RideID: 1,
				RidePositions: [2]rides.RidePosition{
					{
						Id:  1,
						Lat: 37.966660,
						Lng: 23.728308,
						// StartTime = 1 AM
						Timestamp: 1644886800,
					},
					{
						Id:        1,
						Lat:       37.966627,
						Lng:       23.728263,
						Timestamp: 1644943444,
					},
				},
				Speed: 11,
				// Not realistic for a single RideSegment to have that big distance covered,
				// but is used for emulate the test scenario.
				DistanceCovered: 5000,
			},
		}

		close(rideSegmentsChan)
	}()

	go func() {
		fareService.Estimate(
			rideSegmentsChan,
			faresChan,
		)
	}()

	i := 0
	for faresResult := range faresChan {
		assert.Equal(t, expectedFareResult, faresResult)
		i += 1
	}

}
