package zconfig

import (
	"github.com/stretchr/testify/require"
	"github.com/zonewave/pkgs/fileutil"
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
			name: "File not found",
			err:  newErrFileNotFound("test", "", fileutil.ErrNotFound),
			want: "test",
		},
		{
			name: "unsupported config type",
			err:  newErrUnsupportedCfgType("test"),
			want: "Unsupported cfg type:string,should be pointer to struct",
		},
		{
			name: "invalid config ext",
			err:  newErrInvalidCfgExt("test"),
			want: "Unsupported Config File ext \"test\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.ErrorContains(t, tt.err, tt.want)
		})
	}

}
