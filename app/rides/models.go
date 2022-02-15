/*
Package rides
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package rides

import (
	"strconv"
)

const (
	MinimumHourKM float64 = 10.0
)

type RideSegment struct {
	RideID          int
	RidePositions   [2]RidePosition
	Speed           float64
	DistanceCovered float64
}

func NewRideSegment(
	rideID int,
	ridePositions [2]RidePosition,
	speed float64,
	distanceCovered float64,
) *RideSegment {
	return &RideSegment{
		rideID,
		ridePositions,
		speed,
		distanceCovered,
	}
}

type RidePosition struct {
	Id        int
	Lat       float64
	Lng       float64
	Timestamp int64
}

func NewRidePosition(id int, lat, lng float64, timestamp int64) *RidePosition {
	return &RidePosition{
		id,
		lat,
		lng,
		timestamp,
	}
}

// Unmarshal Will unmarshal the provided body which is an array of strings, to a new RidePosition
func Unmarshal(body []string) (*RidePosition, error) {
	id, errId := strconv.ParseInt(body[0], 0, 0)
	lat, errLat := strconv.ParseFloat(body[1], 64)
	lng, errLng := strconv.ParseFloat(body[2], 64)
	timestamp, errTimestamp := strconv.ParseInt(body[3], 0, 0)

	if errId != nil || errLat != nil || errLng != nil || errTimestamp != nil {
		return nil, ErrorParsingRidePosition
	}

	// There should be a sanity check for lat and lng as well, but there
	// is no information about their formats in the file input.
	ridePosition := NewRidePosition(
		int(id),
		lat,
		lng,
		timestamp,
	)

	return ridePosition, nil
}
