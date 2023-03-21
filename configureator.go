package zconfig

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/spf13/afero"
	"github.com/zonewave/pkgs/standutil/cputil"
	"github.com/zonewave/pkgs/standutil/fileutil"
	"github.com/zonewave/pkgs/standutil/reflectutil"
	"github.com/zonewave/zconfig/serialization"
	"strings"
)

// _defaultLookupPaths find config  in those paths
var _defaultLookupPaths = []string{
	"./", "../", "../../", "../../../", "../../../../",
	"./conf", "/conf", "../conf", "../../conf", "../../../conf", "../../../../conf",
	"./config", "/config", "../config", "../../config", "../../../config", "../../../../config",
}

// _defaultSupportedExts are universally supported extensions.

// Configurator config manager
type Configurator struct {
	container   interface{}
	configPaths []string

	mainFileType string
	mainFile     string

	// The filesystem to read config from.
	fs fileutil.Afero

	unmarshalMgr *serialization.Serialization
}

// New return config manager
func New() *Configurator {
	return new(Configurator).Reset()
}

// Reset reset config manager
func (c *Configurator) Reset() *Configurator {
	c.container = nil
	c.configPaths = _defaultLookupPaths
	c.mainFileType = "json"
	c.mainFile = ""
	c.fs = &afero.Afero{Fs: afero.NewOsFs()}
	c.unmarshalMgr = serialization.NewSerialization(nil)
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
	if !c.unmarshalMgr.IsSupportType(configType) {
		return errors.WithStack(NewErrInvalidCfgExt(configType))
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

// searchFile search file in configPaths
func (c *Configurator) searchFile(file string) (string, error) {
	f, err := fileutil.SearchInPaths(c.fs, c.configPaths, file)
	if err != nil {
		return "", errors.WithStack(NewErrFileNotFound(file, fmt.Sprintf("%s", c.configPaths), err))
	}
	return f, nil
}

func (c *Configurator) loadConfig() error {
	file, err := c.fs.ReadFile(c.mainFile)
	if err != nil {
		return err
	}

	err = c.unmarshalMgr.Unmarshal(c.mainFileType, file, c.container)
	if err != nil {
		return err
	}

	return nil
}
