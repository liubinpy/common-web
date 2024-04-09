package config

type LogConfig struct {
	Level      string `json:"level" yaml:"level" mapstructure:"level"`
	Filename   string `json:"filename" yaml:"filename" mapstructure:"filename"`
	MaxSize    int    `json:"max_size" yaml:"max_size" mapstructure:"max_size"`
	MaxAge     int    `json:"max_age" yaml:"max_age" mapstructure:"max_age"`
	MaxBackups int    `json:"max_backups" yaml:"max_backups" mapstructure:"max_backups"`
}
