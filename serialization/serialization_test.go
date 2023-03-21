package serialization

import (
	"github.com/cockroachdb/errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"github.com/zonewave/zconfig/test"
	"testing"
)

func TestSerialization_Unmarshal(t *testing.T) {
	t.Run("err", func(t *testing.T) {

		tests := []struct {
			name     string
			bs       []byte
			fileType string
			cfg      interface{}
			err      error
		}{
			// TODO: Add test cases.
			{
				name:     "no unmarshal",
				bs:       []byte(test.JSONExample),
				fileType: "test",
				cfg:      nil,
				err:      NewErrUnsupportedUnmarshal("test"),
			},
			{
				name:     "unmarshal error",
				bs:       []byte("{"),
				fileType: "yaml",
				cfg:      &test.AppConfig{},
				err:      errors.New("unmarshal bs:{, yaml:*test.AppConfig failed:yaml: line 1: did not find expected node content"),
			},
		}
		c := NewDefaultSerialization()
		var err error
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err = c.Unmarshal(tt.fileType, tt.bs, tt.cfg)
				require.EqualError(t, err, tt.err.Error())

			})
		}
	})
	t.Run("ok", func(t *testing.T) {

		c := NewSerialization(map[string]UnmarshallFunc{
			"json": jsoniter.Unmarshal,
		})
		cfg := &test.AppConfig{}
		err := c.Unmarshal("json", []byte(test.JSONExample), cfg)
		require.NoError(t, err)
		require.Equal(t, "test-app", cfg.AppName)

	})
}

func TestSerialization_IsSupportType(t *testing.T) {
	c := NewDefaultSerialization()
	require.True(t, c.IsSupportType("json"))
	require.True(t, c.IsSupportType("yaml"))
	require.True(t, c.IsSupportType("yml"))
	require.True(t, c.IsSupportType("toml"))
	require.True(t, c.IsSupportType("csv"))
	require.False(t, c.IsSupportType("test"))
}
