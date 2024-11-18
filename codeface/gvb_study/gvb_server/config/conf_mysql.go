package config

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Config   string `yaml:"config"` //高级配置
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_Level"`
}

func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DB + "?" + m.Config
}
