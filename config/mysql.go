package config

type Mysql struct {
	Username string `json:"username" yaml:"username" mapstructure:"username"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Port     string `json:"port" yaml:"port" mapstructure:"port"`
	Database string `json:"database" yaml:"database" mapstructure:"database"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database
}
