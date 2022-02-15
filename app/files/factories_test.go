/*
Package files
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package files

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

type GetFileServiceTestCase struct {
	filePath       string
	expectedResult FileService
	expectedError  FileError
}

// Tests the GetFileService return a csvFileService in case a .csv type of file is provided.
func TestGetFileServiceReturnReturnCSV(t *testing.T) {
	fileService, err := GetFileService("test.csv")
	assert.NoError(t, err)

	_, ok := fileService.(FileService)
	assert.Equal(t, true, ok)
	returnedFileServiceType := reflect.TypeOf(fileService).String()
	expectedFileServiceType := "*files.csvFileService"

	assert.Equal(t, expectedFileServiceType, returnedFileServiceType)

}

// Tests the GetFileService return a csvFileService in case a .csv type of file is provided.
func TestGetFileServiceReturnNewFileErrorWhenFileTypeIsInvalid(t *testing.T) {
	testCases := []GetFileServiceTestCase{
		{
			filePath:       "test.unsupported",
			expectedResult: nil,
			expectedError: NewFileError(
				UnsupportedFileType,
				"provided file type: .unsupported, "+
					"must be one of the: "+strings.Join(supportedFileTypes[:], ",")+" \n",
			),
		},
		{
			filePath:       "test",
			expectedResult: nil,
			expectedError: NewFileError(
				UnsupportedFileType,
				"provided file type: , "+
					"must be one of the: "+strings.Join(supportedFileTypes[:], ",")+" \n",
			),
		},
	}

	for _, testCase := range testCases {
		fileService, err := GetFileService(testCase.filePath)
		assert.Error(t, err)

		_, ok := err.(FileError)
		assert.Equal(t, true, ok)
		assert.Equal(t, testCase.expectedResult, fileService)
		assert.Equal(t, testCase.expectedError, err)
	}

}
