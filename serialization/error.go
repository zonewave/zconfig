package serialization

import "fmt"

// ErrUnsupportedUnmarshal denotes an unsupported unmarshal type
type ErrUnsupportedUnmarshal string

// Error returns the formatted configuration error.
func (u ErrUnsupportedUnmarshal) Error() string {
	return fmt.Sprintf("Unsupported Unmarshal %q", string(u))
}

// NewErrUnsupportedUnmarshal returns a new ErrUnsupportedUnmarshal
func NewErrUnsupportedUnmarshal(str string) error {
	return ErrUnsupportedUnmarshal(str)
}

// ErrUnmarshal denotes failing to unmarshal configuration file.
type ErrUnmarshal struct {
	err   error
	bs    []byte
	fType string
	fCtr  interface{}
}

// NewErrUnmarshal returns a new ErrUnmarshal
func NewErrUnmarshal(err error, bs []byte, fType string, fCtr interface{}) *ErrUnmarshal {
	return &ErrUnmarshal{err: err, bs: bs, fType: fType, fCtr: fCtr}
}

// Error returns the formatted ErrUnmarshal.
func (e *ErrUnmarshal) Error() string {
	return fmt.Sprintf("unmarshal bs:%s, %s:%T failed:%s", string(e.bs), e.fType, e.fCtr, e.err.Error())
}
