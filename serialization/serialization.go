package serialization

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/cockroachdb/errors"
	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v3"
)

type UnmarshallFunc func([]byte, interface{}) error

type Serialization struct {
	Unmarshalls map[string]UnmarshallFunc
}

func NewDefaultSerialization() *Serialization {
	return &Serialization{
		Unmarshalls: map[string]UnmarshallFunc{
			"json": json.Unmarshal,
			"yaml": yaml.Unmarshal,
			"yml":  yaml.Unmarshal,
			"toml": toml.Unmarshal,
			"csv":  gocsv.UnmarshalBytes,
		},
	}

}

func NewSerialization(fns map[string]UnmarshallFunc) *Serialization {
	s := NewDefaultSerialization()
	for k, v := range fns {
		s.Unmarshalls[k] = v
	}
	return s
}

func (s *Serialization) Unmarshal(fType string, data []byte, v interface{}) error {
	unmarshall, ok := s.Unmarshalls[fType]
	if !ok {
		return errors.WithStack(NewErrUnsupportedUnmarshal(fType))
	}
	if err := unmarshall(data, v); err != nil {
		return errors.WithStack(NewErrUnmarshal(err, data, fType, v))
	}
	return nil
}
