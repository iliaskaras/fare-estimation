/*
Package files
Copyright © 2022 Ilias Karatsin <hlias.karas.apps@gmail.com>
*/
package files

import (
	"errors"
	baseAppErrors "github.com/iliaskaras/fare-estimation/app/infrastructure/errors"
)

type FileError struct {
	baseAppErrors.BaseAppError
}

func NewFileError(err error, additionalInfo string) FileError {
	return FileError{
		BaseAppError: baseAppErrors.NewBaseAppError(err, additionalInfo),
	}
}

var (
	UnsupportedFileType = errors.New("unsupported file type")
)
