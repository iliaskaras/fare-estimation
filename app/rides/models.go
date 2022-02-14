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
func Unmarshal(body []string) RidePosition {
	id, _ := strconv.ParseInt(body[0], 0, 0)
	lat, _ := strconv.ParseFloat(body[1], 64)
	lng, _ := strconv.ParseFloat(body[2], 64)
	timestamp, _ := strconv.ParseInt(body[3], 0, 0)

	ridePosition := NewRidePosition(
		int(id),
		lat,
		lng,
		timestamp,
	)

	return *ridePosition
}
