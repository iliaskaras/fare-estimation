/*
Package rides
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package rides

import (
	"github.com/iliaskaras/fare-estimation/app/distances"
)

const (
	hourInSeconds = 3600
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

// FilterOnSegmentSpeed filters the erroneous RidePosition by using the segment speed as a filter.
// More specifically, is responsibly for filtering out the second RidePosition out of a consecutive RidePositions,
// if the calculated speed km/hour > 100km.
// - Receiver of the channel ridePositions,
// - Pusher to the channel the filteredRidePositions, where all the filtered RidePosition are pushed.
func (ss *RidePositionService) FilterOnSegmentSpeed(
	ridePositions <-chan []RidePosition,
	filteredRidePositions chan<- []RidePosition,
) error {

	// Receives the ridePositions.
	for unfilteredRidePositions := range ridePositions {
		ridePositionsSize := len(unfilteredRidePositions)
		var _filteredRidePositions []RidePosition

		i := 0
		j := 1
		for i < ridePositionsSize {

			currentRidePosition := unfilteredRidePositions[i]

			if j >= ridePositionsSize {
				// Case where we are at the end of the ridePositions,
				// and there is nothing to evaluate the current RidePosition with.
				_filteredRidePositions = append(_filteredRidePositions, currentRidePosition)
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

			segmentSpeed := (distanceCovered / float64(elapsedTimeSecs)) * hourInSeconds

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
			// 1. We are adding the currentRidePosition to the _filteredRidePositions.
			// 2. Changing the list indexes in such way that the current nextRidePosition will
			//    become the currentRidePosition in the next loop, by changing current index
			//    i to be equal to this loop's next index j, and the next iteration's
			//	  next index j, to show on the immediate next item in the list.
			_filteredRidePositions = append(_filteredRidePositions, currentRidePosition)
			i = j
			j = i + 1
		}

		filteredRidePositions <- _filteredRidePositions

	}

	close(filteredRidePositions)

	return nil
}
