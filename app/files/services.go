/*
Package files
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package files

import (
	"encoding/csv"
	"github.com/iliaskaras/fare-estimation/app/fares"
	baseAppErrors "github.com/iliaskaras/fare-estimation/app/infrastructure/errors"
	"github.com/iliaskaras/fare-estimation/app/rides"
	"io"
	"log"
	"os"
)

type FileService interface {
	Read(filePath string, ridePositionsChan chan<- []rides.RidePosition) error
	Write(output string, faresChan <-chan fares.Fare) (bool, error)
}

// csvFileService is the FileService implementor responsible for operating on .csv type of files.
type csvFileService struct{}

func newCSVFileService() FileService {
	return &csvFileService{}
}

// Read parses a file that contain rows of ride positions, unmarshal the entries and
// pushes the ride positions to the ridePositionsChan channel for further processing by
// its receivers. The file is already sorted by RideID, and it makes a single push
// to the channel for each RideID encountered through the file parsing.
// - Pusher to the channel ridePositionsChan, where all the encountered RidePosition are pushed.
func (fs *csvFileService) Read(
	filePath string,
	ridePositionsChan chan<- []rides.RidePosition,
) error {
	// Since Read is the sender function of the ridePositionsChan channel, we close it here.
	defer close(ridePositionsChan)

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

	reader := csv.NewReader(file)

	positionsInRide := make(map[int][]rides.RidePosition)
	currentRideID := -1
	previousRideID := -1

	for {
		fileRecord, err := reader.Read()

		if err == io.EOF {
			// Makes sure to push the last RideID's RidePositions to the channel.
			if currentRidePositions, ok := positionsInRide[currentRideID]; ok {
				ridePositionsChan <- currentRidePositions
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
			// to the ridePositionsChan channel for further processing by the receivers.
			if previousRidePositions, ok := positionsInRide[previousRideID]; ok {
				ridePositionsChan <- previousRidePositions
				delete(positionsInRide, previousRideID)
			}
			positionsInRide[currentRideID] = append(seenPositions, ridePos)
		}

	}

	return nil
}

// Write writes line by line, to the output file the fare estimates, each line represents
// the Fare Estimation of a single RideID.
// - Pusher to the channel fileWriteFinishChan, a flag channel indicating when the writing is finish.
//	 Used to block the main goroutine and force it wait Write to finish.
// - Receiver to the channel faresChan, where all the estimated Fares are pushed.
func (fs *csvFileService) Write(
	output string,
	faresChan <-chan fares.Fare,
) (bool, error) {
	file, err := os.Create(output)
	if err != nil {
		return false, NewFileError(err, "unable to create the file")
	}

	writer := csv.NewWriter(file)

	for fare := range faresChan {
		err := writer.Write(fare.ToStrings())
		if err != nil {
			log.Println("failure while writing fare estimation with rideID: ", fare.RideID)
		}
	}

	// Flush the writer and close the file.
	writer.Flush()
	file.Close()

	return true, nil
}
