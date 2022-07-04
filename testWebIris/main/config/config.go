package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Model      string `json:"model"`
}

var ServConfig AppConfig

//初始化服务器配置
func InitConfig() *AppConfig {
	file, err := os.Open("testWebIris/main/config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	config := AppConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
