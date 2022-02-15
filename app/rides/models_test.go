/*
Package rides
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package rides

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests the RidePosition Unmarshal method successfully unmarshal the string.
func TestRidePositionUnmarshalMethod(t *testing.T) {
	rawRidePositionEntry := []string{"1", "37.938598", "23.630322", "1405596152"}

	expectedServiceType := NewRidePosition(
		1,
		37.938598,
		23.630322,
		1405596152,
	)

	result, _ := Unmarshal(rawRidePositionEntry)

	assert.Equal(t, expectedServiceType, result)
}

// Tests the RidePosition Unmarshal method raise error cases.
func TestRidePositionUnmarshalRaiseErr(t *testing.T) {
	rawRidePositionEntries := [][]string{
		{"invalidInt", "37.938598", "23.630322", "1405596152"},
		{"1", "invalidFloat", "23.630322", "1405596152"},
		{"1", "37.938598", "invalidFloat", "1405596152"},
		{"1", "37.938598", "23.630322", "invalidInt"},
	}

	i := 0
	for i < len(rawRidePositionEntries) {
		test := rawRidePositionEntries[i]
		_, err := Unmarshal(test)
		assert.Error(t, err)

		assert.Equal(t, ErrorParsingRidePosition, err)
		i += 1
	}

}
