package zconfig

import (
	"fmt"
)

// ErrFileNotFound denotes failing to find configuration File.
type ErrFileNotFound struct {
	File, Location string

	Err error
}

// Error returns the formatted configuration error.
func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("%s,File:%q,Location:%q", e.Err.Error(), e.File, e.Location)
}

// Unwrap returns the underlying error.
func (e ErrFileNotFound) Unwrap() error {
	return e.Err
}

// newErrFileNotFound returns a new ErrFileNotFound
func newErrFileNotFound(file, location string, err error) error {
	return ErrFileNotFound{File: file, Location: location, Err: err}
}

// ErrUnsupportedCfgType denotes encountering an unsupported
// configuration File type.
type ErrUnsupportedCfgType struct {
	Obj interface{}
}

// Error returns the formatted configuration error.
func (e ErrUnsupportedCfgType) Error() string {
	return fmt.Sprintf("Unsupported cfg type:%T,should be pointer to struct", e.Obj)
}

// newErrUnsupportedCfgType returns a new ErrUnsupportedCfgType
func newErrUnsupportedCfgType(obj interface{}) error {
	return ErrUnsupportedCfgType{
		Obj: obj,
	}
}

// ErrInvalidCfgExt denotes an invalid
// configuration type
type ErrInvalidCfgExt string

// Error returns the formatted configuration error.
func (str ErrInvalidCfgExt) Error() string {
	return fmt.Sprintf("Unsupported Config File ext %q", string(str))
}

// newErrInvalidCfgExt returns a new ErrInvalidCfgExt error of not supported config File type
func newErrInvalidCfgExt(str string) error {
	return ErrInvalidCfgExt(str)
}
