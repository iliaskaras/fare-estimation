/*
Package fares
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package fares

import (
	"fmt"
	"strconv"
)

const (
	StandardFare float64 = 1.30
	MinimumFare  float64 = 3.47
	Idle         float64 = 11.90
	MovingDay    float64 = 0.74
	MovingNight  float64 = 1.30
)

type Fare struct {
	RideID     int
	estimation float64
}

func NewFare(rideID int, estimation float64) *Fare {
	return &Fare{
		rideID,
		estimation,
	}
}

func (f Fare) ToStrings() []string {
	return []string{strconv.Itoa(f.RideID), fmt.Sprintf("%v", f.estimation)}
}
