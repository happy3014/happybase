package main

import (
	"context"
	"fmt"
	"github.com/happy3014/happybase/config"
	"github.com/happy3014/happybase/log"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go/types"
	"os"
	"path"
)

type HiParams struct {
	Name string `json:"name" jsonschema:"the name of the person to greet"`
}

func SayHi(ctx context.Context, req *mcp.CallToolRequest, args HiParams) (*mcp.CallToolResult, any, error) {
	log.SugarLogger().Infof("SayHi: req=%v, args=%v", *req, args)
	return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Hello, %s!", args.Name)}}}, nil, nil
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("failed to get exe path: %v\n", err)
		os.Exit(1)
	}
	configPath := path.Join(exePath, "config.toml")
	err = config.InitConfig(configPath)
	if err != nil {
		fmt.Printf("failed to init config: %v\n", err)
		os.Exit(1)
	}
	err = log.InitLog(config.GlobalConfig().Log)
	if err != nil {
		fmt.Printf("failed to init log: %v\n", err)
		os.Exit(1)
	}

	server := mcp.NewServer(&mcp.Implementation{Name: "greeter"}, nil)
	mcp.AddTool[HiParams, types.Nil](server, &mcp.Tool{
		Description: "say hi",
		Name:        "greet",
		Title:       "",
	}, SayHi)

}
