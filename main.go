package main

import (
	"encoding/json"
	"fmt"
	"github.com/happy3014/happybase/config"
	"github.com/happy3014/happybase/log"
	"go.uber.org/zap"
)

func main() {
	configFilePath := "D:\\tmp\\happybase.toml"
	//conf := config.Config{}
	//err := godefault.SetDefaults(&conf)
	//if err != nil {
	//	fmt.Printf("failed to set defaults: %v", err)
	//	return
	//}
	//err = config.GenerateConfig(conf, configFilePath)
	//if err != nil {
	//	fmt.Printf("failed to generate config: %v", err)
	//	return
	//}

	// 初始化配置
	err := config.InitConfig(configFilePath)
	if err != nil {
		fmt.Printf("failed to parse config: %v", err)
		return
	}
	d, err := json.MarshalIndent(config.GlobalConfig(), "", "    ")
	if err != nil {
		fmt.Printf("failed to marshal config: %v", err)
		return
	}
	fmt.Printf("config:\n%s\n", string(d))

	// 初始化日志
	err = log.InitLog(config.GlobalConfig().Log)
	if err != nil {
		fmt.Printf("failed to init log: %v", err)
		return
	}

	log.Logger().Info("my name", zap.String("addr", "local"))
}
