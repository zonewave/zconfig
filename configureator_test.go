package zconfig

import (
	"errors"
	"github.com/stretchr/testify/require"
	"github.com/zonewave/pkgs/standutil/fileutil"
	"github.com/zonewave/zconfig/test"
	"testing"
)

func TestConfigurator_checkExt(t *testing.T) {
	c := New()
	err := c.checkExt("json")
	require.NoError(t, err)

	err = c.checkExt("test")
	require.ErrorIs(t, err, ErrInvalidCfgExt("test"))

}

func TestConfigurator_checkObject(t *testing.T) {
	c := New()
	err := c.checkObject(&struct {
	}{})
	require.NoError(t, err)

	err = c.checkObject("test")
	require.ErrorIs(t, err, NewErrUnsupportedCfgType("test"))
}

func (s *Suite) TestConfigurator_set() {
	s.Run("err", func() {
		tests := []struct {
			name     string
			fileName string
			cfg      interface{}
			err      error
		}{
			// TODO: Add test cases.
			{
				name:     "ext error",
				fileName: "2.test",
				cfg:      nil,
				err:      ErrInvalidCfgExt("test"),
			},
			{
				name:     "object error",
				fileName: "2.json",
				cfg:      "test",
				err:      NewErrUnsupportedCfgType("test"),
			},
			{
				name:     "file not found",
				fileName: "2.json",
				cfg:      &struct{}{},
				err:      fileutil.ErrNotFound,
			},
		}
		c := New()
		var err error
		for _, tt := range tests {
			s.Run(tt.name, func() {
				c.Reset()
				err = c.set(tt.fileName, tt.cfg)
				s.Require().ErrorIs(err, tt.err)

			})
		}

	})
	s.Run("ok", func() {
		c := New()
		c.fs = s.mockAfero
		s.mockAfero.EXPECT().Exists("1.json").Return(true, nil).Times(1)

		cfg := &struct{}{}
		err := c.set("1.json", cfg)
		s.Require().NoError(err)
		s.Require().Equal("1.json", c.mainFile)
		s.Require().Equal(cfg, c.container)
		s.Require().Equal("json", c.mainFileType)
	})
}

func TestConfigurator_unmarshal(t *testing.T) {

}

func (s *Suite) TestConfigurator_loadConfig() {
	s.Run("err", func() {
		tests := []struct {
			name string
			file string
			err  error
			fn   func(c *Configurator)
		}{
			// TODO: Add test cases.
			{
				name: "file read failed",
				file: "load_config_err.json",
				err:  fileutil.ErrNotFound,
				fn: func(c *Configurator) {
					s.mockAfero.EXPECT().ReadFile(c.mainFile).Return(nil, fileutil.ErrNotFound).Times(1)
				},
			},
			{
				name: "unmarshal error",
				file: "unmarshal_err.yaml",
				err:  errors.New("yaml: line 1: did not find expected node content"),
				fn: func(c *Configurator) {
					c.mainFileType = "yaml"
					s.mockAfero.EXPECT().ReadFile(c.mainFile).Return([]byte("{"), nil).Times(1)
				},
			},
		}
		c := New()
		for _, tt := range tests {
			s.Run(tt.name, func() {
				c.Reset()
				c.mainFile = tt.file
				c.fs = s.mockAfero
				tt.fn(c)
				err := c.loadConfig()
				s.Require().ErrorContains(err, tt.err.Error())

			})
		}
	})
	s.Run("ok", func() {
		c := New()
		c.mainFile = "load_config_ok.json"
		c.fs = s.mockAfero
		s.mockAfero.EXPECT().ReadFile(c.mainFile).Return([]byte(test.JsonExample), nil).Times(1)
		c.container = &test.AppConfig{}
		err := c.loadConfig()
		s.Require().NoError(err)

	})

}

func (s *Suite) TestConfigurator_Initialize() {
	s.Run("err", func() {
		tests := []struct {
			name     string
			fileName string
			err      error
			cfg      interface{}
			fn       func(c *Configurator)
		}{
			// TODO: Add test cases.
			{
				name:     "set error",
				fileName: "set_error.test",
				cfg:      nil,
				err:      ErrInvalidCfgExt("test"),
				fn:       func(c *Configurator) {},
			},
			{
				name:     "load config err",
				fileName: "load_config_err.json",
				err:      errors.New("load config failed"),
				cfg:      &test.AppConfig{},
				fn: func(c *Configurator) {
					c.fs = s.mockAfero
					s.mockAfero.EXPECT().Exists("load_config_err.json").Return(true, nil).Times(1)
					s.mockAfero.EXPECT().ReadFile("load_config_err.json").Return(nil, errors.New("load config failed")).Times(1)
				},
			},
			{
				name:     "copier err",
				fileName: "copier.json",
				err:      errors.New("copier field name tag must be start upper case"),
				cfg: &struct {
					AppName string `json:"app_name" copier:"app_name"`
				}{},
				fn: func(c *Configurator) {
					c.fs = s.mockAfero
					s.mockAfero.EXPECT().Exists("copier.json").Return(true, nil).Times(1)
					s.mockAfero.EXPECT().ReadFile("copier.json").Return([]byte("{}"), nil).Times(1)
				},
			},
		}
		for _, tt := range tests {
			s.Run(tt.name, func() {
				c := New()
				tt.fn(c)
				err := c.Initialize(tt.fileName, tt.cfg)
				s.Require().EqualError(err, tt.err.Error())

			})
		}
	})
	s.Run("ok", func() {
		c := New()
		file := "Initialize_ok.json"
		c.fs = s.mockAfero
		s.mockAfero.EXPECT().Exists("Initialize_ok.json").Return(true, nil).Times(1)
		s.mockAfero.EXPECT().ReadFile(file).Return([]byte(test.JsonExample), nil).Times(1)
		cfg := &test.AppConfig{}
		err := c.Initialize(file, cfg)
		s.Require().NoError(err)
		s.Require().Equal("test-app", cfg.AppName)
	})
}
