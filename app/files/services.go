/*
Package files
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package files

import (
	"fmt"
	baseAppErrors "github.com/iliaskaras/fare-estimation/app/infrastructure/errors"
	"os"
)

type FileService interface {
	Read(filePath string) error
	Write()
}

// csvFileService is the FileService implementor responsible for operating on .csv type of files.
type csvFileService struct{}

func newCSVFileService() FileService {
	return &csvFileService{}
}

func (fs *csvFileService) Read(filePath string) error {
	if filePath == "" {
		return baseAppErrors.NewBaseAppError(
			baseAppErrors.InvalidInputError,
			"file path is missing",
		)
	}

	_, err := os.Open(filePath)
	if err != nil {
		return NewFileError(err, "unable to open the file")
	}

	return nil
}

func (fs *csvFileService) Write() {
	fmt.Printf("CSV file service Write() called!")
}
