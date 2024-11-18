package config

type QiNiu struct {
	Enable    bool    `yaml:"enable" json:"enable"` // 是否启用七牛
	AccessKey string  `yaml:"access_key" json:"accessKey"`
	SecretKey string  `yaml:"secret_key" json:"secretKey"`
	Bucket    string  `yaml:"bucket" json:"bucket"` // 存储桶的名字
	CDN       string  `yaml:"cdn" json:"cdn"`       // 访问图片地址的前缀
	Zone      string  `yaml:"zone" json:"zone"`     // 存储的地区
	Size      float64 `yaml:"size" json:"size"`     // 存储的大小限制，单位是MB
}
