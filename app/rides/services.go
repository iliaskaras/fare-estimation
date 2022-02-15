/*
Package rides
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package rides

import (
	"github.com/iliaskaras/fare-estimation/app/distances"
)

const (
	HourInSeconds = 3600
	maxKMPerHour  = 100
)

type RidePositionService struct {
	distanceCalculator distances.DistanceCalculatorService
}

func NewRidePositionService(distanceCalculator distances.DistanceCalculatorService) *RidePositionService {
	return &RidePositionService{
		distanceCalculator: distanceCalculator,
	}
}

// FilterOnSegmentSpeed filters the erroneous RidePosition by using the segment Speed as a filter.
// More specifically, is responsibly for filtering out the second RidePosition out of a RideSegment,
// if the calculated Speed km/hour > 100km.
// - Receiver of the channel ridePositionsChan,
// - Pusher to the channel the rideSegmentsChan, where all the filtered RideSegment are pushed.
func (ss *RidePositionService) FilterOnSegmentSpeed(
	ridePositionsChan <-chan []RidePosition,
	rideSegmentsChan chan<- []RideSegment,
) {

	// Receives the RidePositions.
	for unfilteredRidePositions := range ridePositionsChan {
		ridePositionsSize := len(unfilteredRidePositions)
		var filteredRideSegments []RideSegment

		i := 0
		j := 1
		for i < ridePositionsSize {

			currentRidePosition := unfilteredRidePositions[i]

			if j >= ridePositionsSize {
				// Case where we are at the end of the RidePositions,
				// and there is nothing to evaluate the current RidePosition with.
				break
			}

			nextRidePosition := unfilteredRidePositions[j]

			// Calculate the elapsed time given the two ride position timestamps in seconds.
			elapsedTimeSecs := nextRidePosition.Timestamp - currentRidePosition.Timestamp

			// Calculate the distance covered.
			distanceCovered, _ := ss.distanceCalculator.GetDistance(
				currentRidePosition.Lat,
				currentRidePosition.Lng,
				nextRidePosition.Lat,
				nextRidePosition.Lng,
			)

			segmentSpeed := (distanceCovered / float64(elapsedTimeSecs)) * HourInSeconds

			// Sanity check on the segmentSpeed, if is greater than the maxKMPerHour,
			// then this means that the check failed and the second part of the
			// segment, which is the nextRidePosition, needs to be skipped because
			// is found to be erroneous. The skip happen by just increasing the next
			// index j.
			if segmentSpeed > maxKMPerHour {
				j += 1
				continue
			}

			// The two RidePositions are valid entries, thus:
			// 1. We are adding the RideSegment of the current and previous RidePositions.
			// 2. Changing the list indexes in such way that the current nextRidePosition will
			//    become the currentRidePosition in the next loop, by changing current index
			//    i to be equal to this loop's next index j, and the next iteration's
			//	  next index j, to show on the immediate next item in the list.
			filteredRideSegments = append(
				filteredRideSegments,
				*NewRideSegment(
					currentRidePosition.Id,
					[2]RidePosition{
						currentRidePosition,
						nextRidePosition,
					},
					segmentSpeed,
					distanceCovered,
				),
			)
			i = j
			j = i + 1
		}

		rideSegmentsChan <- filteredRideSegments

	}

}
