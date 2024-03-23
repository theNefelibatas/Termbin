package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config 配置
type Config struct {
	MySQL MySQLConfig `json:"MySQL"`
}

// MySQLConfig MySQL 配置
type MySQLConfig struct {
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
	Database string `json:"Database"`
}

// Conf 全局配置变量
var Conf Config

// InitConfig 初始化配置
//
// 从 config/config.json 文件中加载配置信息, 并解析到 Conf 变量中
func InitConfig() {
	// 加载配置文件
	configFile, err := os.Open("config/config.json")
	if err != nil {
		log.Println("config: failed to open config file.")
		panic(err)
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.Println("config: failed to close config file.")
			panic(err)
		}
	}(configFile)

	// 解析配置文件, 并加载到 Conf 变量中
	_ = json.NewDecoder(configFile).Decode(&Conf)
}
