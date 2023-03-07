package zconfig

import (
	"github.com/stretchr/testify/require"
	"github.com/zonewave/pkgs/standutil/fileutil"
	"testing"
)

func TestErrorString(t *testing.T) {

	tests := []struct {
		name string
		err  error
		want string
	}{
		// TODO: Add test cases.
		{
			name: "file not found",
			err:  NewErrFileNotFound("test", "", fileutil.ErrNotFound),
			want: "test",
		},
		{
			name: "unsupported config type",
			err:  NewErrUnsupportedCfgType("test"),
			want: "Unsupported cfg type:string,should be pointer to struct",
		},
		{
			name: "invalid config ext",
			err:  NewErrInvalidCfgExt("test"),
			want: "Unsupported Config file ext \"test\"",
		},
		{
			name: "unsupported unmarshal",
			err:  NewErrUnsupportedUnmarshal("test"),
			want: "Unsupported Unmarshal \"test\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.ErrorContains(t, tt.err, tt.want)
		})
	}

}
