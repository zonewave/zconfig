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

func NewErrInvalidCfgExt(str string) error {
	return ErrInvalidCfgExt(str)
}

type ErrUnsupportedUnmarshal string

func (u ErrUnsupportedUnmarshal) Error() string {
	return fmt.Sprintf("Unsupported Unmarshal %q", string(u))
}

func NewErrUnsupportedUnmarshal(str string) error {
	return ErrUnsupportedUnmarshal(str)
}

type ErrUnmarshal struct {
	err   error
	bs    []byte
	fType string
	fCtr  interface{}
}

func NewErrUnmarshal(err error, bs []byte, fType string, fCtr interface{}) *ErrUnmarshal {
	return &ErrUnmarshal{err: err, bs: bs, fType: fType, fCtr: fCtr}
}

func (e *ErrUnmarshal) Error() string {
	return fmt.Sprintf("unmarshal bs:%s, %s:%T failed:%s", string(e.bs), e.fType, e.fCtr, e.err.Error())
}
