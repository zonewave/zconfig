package zconfig

import (
	"fmt"
)

// ErrFileNotFound denotes failing to find configuration file.
type ErrFileNotFound struct {
	file, location string

	err error
}

// Error returns the formatted configuration error.
func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("%s,file:%q,location:%q", e.err.Error(), e.file, e.location)
}

// Unwrap returns the underlying error.
func (e ErrFileNotFound) Unwrap() error {
	return e.err
}

// NewErrFileNotFound returns a new ErrFileNotFound
func NewErrFileNotFound(file, location string, err error) error {
	return ErrFileNotFound{file: file, location: location, err: err}
}

// ErrUnsupportedCfgType denotes encountering an unsupported
// configuration file type.
type ErrUnsupportedCfgType struct {
	obj interface{}
}

// Error returns the formatted configuration error.
func (e ErrUnsupportedCfgType) Error() string {
	return fmt.Sprintf("Unsupported cfg type:%T,should be pointer to struct", e.obj)
}

// NewErrUnsupportedCfgType returns a new ErrUnsupportedCfgType
func NewErrUnsupportedCfgType(obj interface{}) error {
	return ErrUnsupportedCfgType{
		obj: obj,
	}
}

// ErrInvalidCfgExt denotes an invalid
// configuration type
type ErrInvalidCfgExt string

// Error returns the formatted configuration error.
func (str ErrInvalidCfgExt) Error() string {
	return fmt.Sprintf("Unsupported Config file ext %q", string(str))
}

// NewErrInvalidCfgExt returns a new ErrInvalidCfgExt error of not supported config file type
func NewErrInvalidCfgExt(str string) error {
	return ErrInvalidCfgExt(str)
}
