package main

import (
	"fmt"
	"github.com/happy3014/happybase/config"
	"github.com/spf13/cobra"
)

func GetConfigMigrateCmd() *cobra.Command {

	confMigrate := &cobra.Command{
		Use:   "config_migrate",
		Short: "config_migrate",
		Long:  "配置迁移，将旧版本的配置文件迁移到新版本的配置文件",
		Run: func(cmd *cobra.Command, args []string) {
			oldPath, err := cmd.Flags().GetString("old")
			if err != nil {
				fmt.Printf("failed to get old path: %v", err)
				return
			}
			newPath, err := cmd.Flags().GetString("new")
			if err != nil {
				fmt.Printf("failed to get new path: %v", err)
				return
			}
			err = config.GenerateNewConfig(oldPath, newPath)
			if err != nil {
				fmt.Printf("failed to generate new config: %v", err)
				return
			}
		},
	}

	confMigrate.Flags().StringP("new", "n", "", "生成新的配置文件的路径")
	confMigrate.Flags().StringP("old", "o", "", "旧版本的配置文件的路径")

	return confMigrate
}
