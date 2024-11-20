package config

type UpLoad struct {
	Path string `yaml:"path" json:"path"` //文件上传路径
	Size int    `yaml:"size" json:"size"` //文件上传大小
}
