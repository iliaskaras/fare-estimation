/*
Package rides
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package rides

import (
	"strconv"
)

type RidePosition struct {
	Id        int
	Lat       float64
	Lng       float64
	Timestamp int
}

func NewRidePosition(id int, lat, lng float64, timestamp int) *RidePosition {
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
		int(timestamp),
	)

	return *ridePosition
}
