/*
Package cmd
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/iliaskaras/fare-estimation/app/distances"
	"github.com/iliaskaras/fare-estimation/app/fares"
	"github.com/iliaskaras/fare-estimation/app/files"
	baseAppErrors "github.com/iliaskaras/fare-estimation/app/infrastructure/errors"
	"github.com/iliaskaras/fare-estimation/app/rides"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var estimateCmd = &cobra.Command{
	Use:   "estimate",
	Short: "Produce a file with the fare estimations.",
	Long: `Calculating the fare estimations and printing them into a new file.

The following steps are executed:

- Filtering the provided file out of erroneous entries. Erroneous entry is the second
  part of a ride segment, that the calculated speed is greater than 100km/hour.
  The distance is calculated using the Haversine formula.
- Calculating the fare estimations out of the filtered ride segments, making a new
  file with all the ride fare estimations.
`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		filePath, _ := cmd.Flags().GetString("filepath")

		fileService, err := files.GetFileService(filePath)

		if err != nil {
			if errors.Is(err, files.UnsupportedFileType) {
				fmt.Printf(err.Error())
				os.Exit(1)
			}
		}

		ridePositions := make(chan []rides.RidePosition)
		rideSegments := make(chan []rides.RideSegment)
		faresChan := make(chan fares.Fare)

		go func() {
			err := fileService.Read(filePath, ridePositions)
			if err != nil {
				if errors.Is(err, baseAppErrors.InvalidInputError) {
					fmt.Printf(err.Error())
				}
				fmt.Printf(err.Error())
			}
		}()

		distanceCalculatorMethod, _ := distances.GetDistanceCalculatorService(distances.HaversineMethod)
		ridePositionService, err := rides.GetRidePositionService(
			distanceCalculatorMethod,
		)

		go func() {
			ridePositionService.FilterOnSegmentSpeed(ridePositions, rideSegments)
		}()

		fareService := fares.NewFareService()

		go func() {
			fareService.Estimate(rideSegments, faresChan)
		}()

		t := time.Now()
		elapsed := t.Sub(start)

		var _fares []fares.Fare
		for n := range faresChan {
			_fares = append(_fares, n)
		}

		println(len(_fares))
		println("run in: ", elapsed.Milliseconds())

	},
}

func init() {
	rootCmd.AddCommand(estimateCmd)

	estimateCmd.Flags().StringP(
		"filepath", "f", "", "The file path contains information about rides",
	)
}
