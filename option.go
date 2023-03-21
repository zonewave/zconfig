package zconfig

import "github.com/zonewave/zconfig/serialization"

// Option modifies the default behavior
type Option interface {
	applyProvideOption(cfg *Configurator)
}

func WithCfgFilePath(path []string) Option {
	return CfgFilePathOption(path)
}

type CfgFilePathOption []string

func (p CfgFilePathOption) applyProvideOption(opt *Configurator) {
	opt.configPaths = p
}

type SerializerOption struct {
	s *serialization.Serialization
}

func WithSerializerOption(s *serialization.Serialization) Option {
	return &SerializerOption{s: s}
}

func (s SerializerOption) applyProvideOption(cfg *Configurator) {
	cfg.unmarshalMgr = s.s
}
