package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

// InitConf 读取yaml文件的配置
func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}

func SetYaml() error {
	out, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, out, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("设置文件参数修改成功")
	return nil
}
