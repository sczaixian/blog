package config

type Server struct {
	JWT   JWT   `json:"jwt" mapstructure:"jwt" yaml:"jwt"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`

	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`

	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
}
