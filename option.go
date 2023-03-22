package zconfig

import "github.com/zonewave/zconfig/serialization"

// Option modifies the default behavior
type Option interface {
	applyProvideOption(cfg *Configurator)
}

// WithCfgFilePath set config File path
func WithCfgFilePath(path []string) Option {
	return cfgFilePathOption(path)
}

type cfgFilePathOption []string

func (p cfgFilePathOption) applyProvideOption(opt *Configurator) {
	opt.configPaths = p
}

type serializerOption struct {
	s *serialization.Serialization
}

// WithSerializerOption set serializer
func WithSerializerOption(s *serialization.Serialization) Option {
	return &serializerOption{s: s}
}

func (s serializerOption) applyProvideOption(cfg *Configurator) {
	cfg.unmarshalMgr = s.s
}
