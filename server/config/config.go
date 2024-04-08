package config

type Config struct {
	System System `mapstructure:"system" json:"system" yaml:"system"` // 系统配置
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`    // 数据库配置
}
