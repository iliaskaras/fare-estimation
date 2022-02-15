/*
Package fares
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package fares

import (
	"github.com/iliaskaras/fare-estimation/app/rides"
	"math"
	"time"
)

type FareService struct{}

func NewFareService() *FareService {
	return &FareService{}
}

// Estimate estimates the fare for each RideID.
// - Receiver of the channel rideSegmentsChan,
// - Pusher to the channel faresChan, where all the estimated Fare are pushed.
func (ss *FareService) Estimate(
	rideSegmentsChan <-chan []rides.RideSegment,
	faresChan chan<- Fare,
) {

	// Receives the rideSegmentsChan.
	for rideSegments := range rideSegmentsChan {
		fareAmount := StandardFare

		// Case where the RideID had only one RidePosition in the input file.
		if rideSegments == nil {
			continue
		}

		rideID := rideSegments[0].RideID

		for _, rideSegment := range rideSegments {
			if rideSegment.Speed > rides.MinimumHourKM {
				startHour := time.Unix(rideSegment.RidePositions[0].Timestamp, 0).UTC().Hour()

				if startHour >= 0 && startHour < 5 {
					// Night, time after 0 and before 5 the morning.
					fareAmount += (rideSegment.DistanceCovered / 1000) * MovingNight
				} else {
					fareAmount += (rideSegment.DistanceCovered / 1000) * MovingDay
					// Day, time after 5 the morning and before 24.
				}

			} else {
				elapsedTimeSecs := float64(
					rideSegment.RidePositions[1].Timestamp - rideSegment.RidePositions[0].Timestamp,
				)
				fareAmount += (elapsedTimeSecs / rides.HourInSeconds) * Idle
			}
		}

		if fareAmount <= MinimumFare {
			fareAmount = MinimumFare
		}

		faresChan <- *NewFare(
			rideID,
			math.Round(fareAmount*100)/100,
		)

	}

	close(faresChan)

}
