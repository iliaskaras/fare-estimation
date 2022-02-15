/*
Package files
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package files

import (
	"encoding/csv"
	"github.com/Flaque/filet"
	"github.com/iliaskaras/fare-estimation/app/fares"
	baseAppErrors "github.com/iliaskaras/fare-estimation/app/infrastructure/errors"
	"github.com/iliaskaras/fare-estimation/app/rides"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"syscall"
	"testing"
)

// Tests the FileService.Read reads a file and populates the ridePositionsChan.
func TestCSVFileServiceReadSuccessfulExecution(t *testing.T) {
	defer filet.CleanUp(t)
	// Creates a temporary file.
	testInputFile := filet.TmpFile(
		t,
		"",
		"1,37.955217,23.714548,1405595237\n"+
			"1,37.954302,23.713370,1405595284\n"+
			"2,37.946545,23.754918,1405591065\n"+
			"2,37.946545,23.754918,1405591073\n"+
			"2,37.946545,23.754918,1405591084\n"+
			"3,37.946545,23.754918,1405591084\n",
	)
	var expectedResults = [][]rides.RidePosition{
		{
			rides.RidePosition{
				Id:        1,
				Lat:       37.955217,
				Lng:       23.714548,
				Timestamp: 1405595237,
			},
			rides.RidePosition{
				Id:        1,
				Lat:       37.954302,
				Lng:       23.71337,
				Timestamp: 1405595284,
			},
		},
		{
			rides.RidePosition{
				Id:        2,
				Lat:       37.946545,
				Lng:       23.754918,
				Timestamp: 1405591065,
			},
			rides.RidePosition{
				Id:        2,
				Lat:       37.946545,
				Lng:       23.754918,
				Timestamp: 1405591073,
			},
			rides.RidePosition{
				Id:        2,
				Lat:       37.946545,
				Lng:       23.754918,
				Timestamp: 1405591084,
			},
		},
		{
			rides.RidePosition{
				Id:        3,
				Lat:       37.946545,
				Lng:       23.754918,
				Timestamp: 1405591084,
			},
		},
	}

	testRidePositionsChan := make(chan []rides.RidePosition)

	go newCSVFileService().Read(
		testInputFile.Name(),
		testRidePositionsChan,
	)

	i := 0
	for ridePositionsResult := range testRidePositionsChan {
		expectedRidePositionsResult := expectedResults[i]
		assert.Equal(t, len(expectedRidePositionsResult), len(ridePositionsResult))
		assert.Equal(t, expectedRidePositionsResult, ridePositionsResult)
		i += 1
	}

}

// Tests the csvFileService.Read() return an error when the FilePath is not provided.
func TestCSVFileServiceReadReturnErrorWhenFilePathNotProvided(t *testing.T) {
	testRidePositionsChan := make(chan []rides.RidePosition)

	err := newCSVFileService().Read(
		"",
		testRidePositionsChan,
	)
	assert.Error(t, err)

	expectedError := baseAppErrors.NewBaseAppError(
		baseAppErrors.InvalidInputError,
		"file path is missing",
	)

	assert.Equal(t, expectedError, err)
}

// Tests the csvFileService.Read return error when file does not exist.
func TestCSVFileServiceReadReturnErrorWhenFileDoesNotExist(t *testing.T) {
	testRidePositionsChan := make(chan []rides.RidePosition)

	err := newCSVFileService().Read(
		"filethatnotexist.csv",
		testRidePositionsChan,
	)
	assert.Error(t, err)

	expectedFileError := &fs.PathError{
		"open",
		"filethatnotexist.csv",
		syscall.ENOENT,
	}
	expectedError := FileError{
		BaseAppError: baseAppErrors.NewBaseAppError(expectedFileError, "unable to open the file"),
	}

	assert.Equal(t, expectedError, err)
}

// Tests the FileService.Read reads a file and populates the ridePositionsChan.
func TestCSVFileServiceWriteSuccessfulExecution(t *testing.T) {
	faresChan := make(chan fares.Fare)

	defer filet.CleanUp(t)

	testOutputFile := filet.TmpFile(
		t,
		"",
		"",
	)

	go func() {
		faresChan <- *fares.NewFare(
			1,
			1.0,
		)
		faresChan <- *fares.NewFare(
			2,
			2.0,
		)
		close(faresChan)
	}()

	newCSVFileService().Write(
		testOutputFile.Name(),
		faresChan,
	)

	file, _ := os.Open(testOutputFile.Name())
	defer file.Close()
	reader := csv.NewReader(file)
	fileRecord, _ := reader.ReadAll()

	assert.Equal(t, [][]string{[]string{"1", "1"}, []string{"2", "2"}}, fileRecord)

}

// Tests the csvFileService.Write return error when file does not exist.
func TestCSVFileServiceWriteReturnErrorWhenOsCreateFail(t *testing.T) {
	faresChan := make(chan fares.Fare)

	ok, err := newCSVFileService().Write(
		"",
		faresChan,
	)

	assert.Error(t, err)
	assert.Equal(t, ok, false)

	expectedFileError := &fs.PathError{
		"open",
		"",
		syscall.ENOENT,
	}
	expectedError := FileError{
		BaseAppError: baseAppErrors.NewBaseAppError(expectedFileError, "unable to create the file"),
	}

	assert.Equal(t, expectedError, err)
}
