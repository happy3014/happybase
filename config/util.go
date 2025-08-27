package config

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"github.com/varunbheemaiah/godefault"
	"os"
)

var config Config

func GlobalConfig() Config {
	return config
}
func InitConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("[%s] failed to read config file: %v", configFilePath, err)
	}
	config = Config{}
	err = godefault.SetDefaults(&config)
	if err != nil {
		return fmt.Errorf("[%s] failed to set defaults: %v", configFilePath, err)
	}
	// 确保viper.Unmarshal能正确覆盖默认值
	err = viper.Unmarshal(&config)
	if err != nil {
		return fmt.Errorf("[%s] failed to unmarshal: %v", configFilePath, err)
	}
	return nil
}

// GenerateConfig 通常用于第一次初始化时，生成配置文件
func GenerateConfig(conf any, dstPath string) error {
	f, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := toml.NewEncoder(f)
	err = encoder.Encode(conf)
	if err != nil {
		return err
	}
	return nil
}

// GenerateNewConfig 用于生成新的配置文件，通常用于升级时，将旧的配置文件迁移到新的配置文件中
func GenerateNewConfig(oldConfigFilePath string, dstPath string) error {
	err := InitConfig(oldConfigFilePath)
	if err != nil {
		return fmt.Errorf("failed to load old config: %v", err)
	}
	err = GenerateConfig(GlobalConfig(), dstPath)
	if err != nil {
		return fmt.Errorf("failed to generate new config to %s: %v", dstPath, err)
	}
	return nil
}
