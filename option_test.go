package zconfig

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/zonewave/zconfig/serialization"
	"testing"
)

func TestCfgFilePathOption_applyProvideOption(t *testing.T) {
	c := New(
		WithCfgFilePath([]string{"test"}),
		WithSerializerOption(serialization.NewSerialization(map[string]serialization.UnmarshallFunc{
			"json2": json.Unmarshal,
		})))
	require.Contains(t, c.configPaths, "test")
	require.True(t, c.unmarshalMgr.IsSupportType("json2"))
}
