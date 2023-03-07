package zconfig

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/cockroachdb/errors"
	"github.com/gocarina/gocsv"
	"github.com/spf13/afero"
	"github.com/zonewave/pkgs/standutil/cputil"
	"github.com/zonewave/pkgs/standutil/fileutil"
	"github.com/zonewave/pkgs/standutil/reflectutil"
	"github.com/zonewave/pkgs/standutil/sliceutil"
	"gopkg.in/yaml.v3"
	"strings"
)

// _defaultLookupPaths find config  in those paths
var _defaultLookupPaths = []string{
	"./", "../", "../../", "../../../", "../../../../",
	"./conf", "/conf", "../conf", "../../conf", "../../../conf", "../../../../conf",
	"./config", "/config", "../config", "../../config", "../../../config", "../../../../config",
}

// _defaultSupportedExts are universally supported extensions.
var _defaultSupportedExts = []string{"json", "yml", "yaml", "toml"}

// Configurator config manager
type Configurator struct {
	container    interface{}
	configPaths  []string
	supportExts  []string
	mainFileType string
	mainFile     string

	// The filesystem to read config from.
	fs fileutil.Afero

	unmarshalMgr map[string]func([]byte, interface{}) error
}

// New return config manager
func New() *Configurator {
	return new(Configurator).Reset()
}

func defaultUnmarshalFunc() map[string]func([]byte, interface{}) error {
	return map[string]func([]byte, interface{}) error{
		"json": json.Unmarshal,
		"yml":  yaml.Unmarshal,
		"yaml": yaml.Unmarshal,
		"toml": toml.Unmarshal,
		"csv":  gocsv.UnmarshalBytes,
	}
}
func (c *Configurator) Reset() *Configurator {
	c.container = nil
	c.configPaths = _defaultLookupPaths
	c.supportExts = _defaultSupportedExts
	c.mainFileType = "json"
	c.mainFile = ""
	c.fs = &afero.Afero{Fs: afero.NewOsFs()}
	c.unmarshalMgr = defaultUnmarshalFunc()
	return c
}

// Initialize your config
func (c *Configurator) Initialize(configFile string, cfgStructPtr interface{}) error {

	if err := c.set(configFile, cfgStructPtr); err != nil {
		return err
	}

	if err := c.loadConfig(); err != nil {
		return err
	}

	if err := cputil.DeepCopy(cfgStructPtr, c.container); err != nil {
		return err
	}
	return nil
}

func (c *Configurator) checkExt(ext string) error {
	if !sliceutil.Contain(ext, c.supportExts) {
		return errors.WithStack(NewErrInvalidCfgExt(ext))
	}
	return nil
}

func (c *Configurator) checkObject(obj interface{}) error {
	if reflectutil.IsStructPtr(obj) {
		return nil
	}
	return errors.WithStack(NewErrUnsupportedCfgType(obj))
}

func (c *Configurator) set(configFile string, cfgStructPtr interface{}) error {

	var (
		configType string
		mainFile   string
		err        error
	)
	// check type
	configType = fileutil.FileExtNoDot(configFile)
	if err = c.checkExt(configType); err != nil {
		return errors.WithStack(err)
	}
	// check object
	if err = c.checkObject(cfgStructPtr); err != nil {
		return errors.WithStack(err)
	}

	// check file
	if mainFile, err = c.searchFile(configFile); err != nil {
		return errors.WithStack(err)
	}

	c.mainFile = mainFile
	c.mainFileType = strings.ToLower(configType)
	c.container = cfgStructPtr
	return nil
}

// searchFile
func (c *Configurator) searchFile(file string) (string, error) {
	if f, err := fileutil.SearchInPaths(c.fs, c.configPaths, file); err != nil {
		return "", errors.WithStack(NewErrFileNotFound(file, fmt.Sprintf("%s", c.configPaths), err))
	} else {
		return f, nil
	}
}

func (c *Configurator) loadConfig() error {
	file, err := c.fs.ReadFile(c.mainFile)
	if err != nil {
		return err
	}

	err = c.unmarshal(c.mainFileType, file, c.container)
	if err != nil {
		return err
	}

	return nil
}

func (c *Configurator) unmarshal(fileType string, bs []byte, cfg interface{}) error {
	unmarshal, ok := c.unmarshalMgr[fileType]
	if !ok {
		return errors.WithStack(NewErrUnsupportedUnmarshal(fileType))
	}
	if err := unmarshal(bs, cfg); err != nil {
		return errors.WithStack(NewErrUnmarshal(err, bs, fileType, cfg))
	}

	return nil
}
