/*
Package files
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package files

import (
	"encoding/csv"
	"fmt"
	baseAppErrors "github.com/iliaskaras/fare-estimation/app/infrastructure/errors"
	"github.com/iliaskaras/fare-estimation/app/rides"
	"io"
	"os"
)

type FileService interface {
	Read(filePath string, ridePositions chan<- []rides.RidePosition) error
	Write()
}

// csvFileService is the FileService implementor responsible for operating on .csv type of files.
type csvFileService struct{}

func newCSVFileService() FileService {
	return &csvFileService{}
}

// Read parses a file that contain rows of ride positions, unmarshal the entries and
// pushes the ride positions to the ridePositions channel for further processing by
// its receivers. The file is already sorted by RideID, and it makes a single push
// to the channel for each RideID encountered through the file parsing.
func (fs *csvFileService) Read(
	filePath string,
	ridePositions chan<- []rides.RidePosition,
) error {
	if filePath == "" {
		return baseAppErrors.NewBaseAppError(
			baseAppErrors.InvalidInputError,
			"file path is missing",
		)
	}

	file, err := os.Open(filePath)

	if err != nil {
		return NewFileError(err, "unable to open the file")
	}
	defer file.Close()

	parser := csv.NewReader(file)

	positionsInRide := make(map[int][]rides.RidePosition)
	currentRideID := -1
	previousRideID := -1

	for {
		fileRecord, err := parser.Read()

		if err == io.EOF {
			// Makes sure to push the last RideID's RidePositions to the channel.
			if currentRidePositions, ok := positionsInRide[currentRideID]; ok {
				ridePositions <- currentRidePositions
				delete(positionsInRide, currentRideID)
			}
			break
		}
		if err != nil {
			return NewFileError(err, "failure on reading file records")
		}

		ridePos := rides.Unmarshal(fileRecord)

		// Initialize the current RideID in the first iteration.
		if currentRideID != ridePos.Id {
			// Keep the previous RideID for pushing to the channel its RidePositions.
			previousRideID = currentRideID
			// Change the current RideID to the new one.
			currentRideID = ridePos.Id
		}

		if seenPositions, ok := positionsInRide[currentRideID]; ok {
			// If the RideID already exist in the positionsInRide map, update its seenPositions.
			positionsInRide[currentRideID] = append(seenPositions, ridePos)
		} else {
			// Means we have a new RideID, and due to the fact that the file is sorted already by
			// RideID, we can safely push the positions found for the previously encountered RideID
			// to the ridePositions channel for further processing by the receivers.
			if previousRidePositions, ok := positionsInRide[previousRideID]; ok {
				ridePositions <- previousRidePositions
				delete(positionsInRide, previousRideID)
			}
			positionsInRide[currentRideID] = append(seenPositions, ridePos)
		}

	}

	// Since Read is the sender function of the ridePositions channel, we close it here.
	close(ridePositions)

	return nil
}

func (fs *csvFileService) Write() {
	fmt.Printf("CSV file service Write() called!")
}
