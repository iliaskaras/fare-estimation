/*
Package files
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package files

import (
	"path/filepath"
	"strings"
)

var supportedFileTypes = []string{".csv"}

// GetFileService is responsible for returning the correct FileService implementor,
// based on the file type provided.
func GetFileService(filePath string) (FileService, error) {
	fileExtension := filepath.Ext(filePath)

	if fileExtension == ".csv" {
		return newCSVFileService(), nil
	}

	return nil, NewFileError(
		UnsupportedFileType,
		"provided file type: "+fileExtension+", "+
			"must be one of the: "+strings.Join(supportedFileTypes[:], ",")+" \n",
	)
}
