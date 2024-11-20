package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Jwt      Jwt      `yaml:"jwt"`
	Email    Email    `yaml:"email"`
	Redis    Redis    `yaml:"redis"`
	QiNiu    QiNiu    `yaml:"qiNiu"`
	QQ       QQ       `yaml:"qq"`
	UpLoad   UpLoad   `yaml:"upLoad"`
}
