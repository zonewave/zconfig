package serialization

import "fmt"

// ErrUnsupportedUnmarshal denotes an unsupported unmarshal type
type ErrUnsupportedUnmarshal string

// Error returns the formatted configuration error.
func (u ErrUnsupportedUnmarshal) Error() string {
	return fmt.Sprintf("Unsupported Unmarshal %q", string(u))
}

// newErrUnsupportedUnmarshal returns a new ErrUnsupportedUnmarshal
func newErrUnsupportedUnmarshal(str string) error {
	return ErrUnsupportedUnmarshal(str)
}

// ErrUnmarshal denotes failing to unmarshal configuration file.
type ErrUnmarshal struct {
	Err   error
	Bs    []byte
	FType string
	FCtr  interface{}
}

// newErrUnmarshal returns a new ErrUnmarshal
func newErrUnmarshal(err error, bs []byte, fType string, fCtr interface{}) *ErrUnmarshal {
	return &ErrUnmarshal{Err: err, Bs: bs, FType: fType, FCtr: fCtr}
}

// Error returns the formatted ErrUnmarshal.
func (e *ErrUnmarshal) Error() string {
	return fmt.Sprintf("unmarshal Bs:%s, %s:%T failed:%s", string(e.Bs), e.FType, e.FCtr, e.Err.Error())
}
