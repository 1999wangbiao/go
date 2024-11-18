package config

type Redis struct {
	Host         string `yaml:"host" json:"host"`
	Port         int    `yaml:"port" json:"port"`
	Password     string `yaml:"password" json:"password"`
	DB           int    `yaml:"db" json:"db"`
	PoolSize     int    `yaml:"pool_size" json:"poolSize" `
	MiniDleConns int    `yaml:"mini_dle_conns" json:"miniDleConns"`
}
