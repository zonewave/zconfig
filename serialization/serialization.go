package serialization

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/cockroachdb/errors"
	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v3"
)

// UnmarshallFunc is a function that unmarshals a byte slice into a value.
type UnmarshallFunc func([]byte, interface{}) error

// Serialization is a serialization manager
type Serialization struct {
	unmarshalls map[string]UnmarshallFunc
}

// NewDefaultSerialization returns a default serialization manager
func NewDefaultSerialization() *Serialization {
	return &Serialization{
		unmarshalls: map[string]UnmarshallFunc{
			"json": json.Unmarshal,
			"yaml": yaml.Unmarshal,
			"yml":  yaml.Unmarshal,
			"toml": toml.Unmarshal,
			"csv":  gocsv.UnmarshalBytes,
		},
	}

}

// NewSerialization returns a serialization manager with custom unmarshalls
func NewSerialization(fns map[string]UnmarshallFunc) *Serialization {
	s := NewDefaultSerialization()
	for k, v := range fns {
		s.unmarshalls[k] = v
	}
	return s
}

// Unmarshal  a byte slice into a value.
func (s *Serialization) Unmarshal(fType string, data []byte, v interface{}) error {
	unmarshall, ok := s.unmarshalls[fType]
	if !ok {
		return errors.WithStack(newErrUnsupportedUnmarshal(fType))
	}
	if err := unmarshall(data, v); err != nil {
		return errors.WithStack(newErrUnmarshal(err, data, fType, v))
	}
	return nil
}

// IsSupportType returns true if the serialization manager supports the file type
func (s *Serialization) IsSupportType(fType string) bool {
	_, ok := s.unmarshalls[fType]
	return ok
}
