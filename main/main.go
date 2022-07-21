package main

import (
	"flag"
	"fmt"
	"fusu"
	"fusu/cmd"
	_ "fusu/cmd"
)

var (
	version = flag.Bool("version", false, "Show current version of Fusu.")
)

func printVersion() {
	version := fusu.VersionStatement()
	for _, info := range version {
		fmt.Println(info)
	}
}

func main() {
	flag.Parse()
	printVersion()
	//	获取配置文件
	config := cmd.GetConfigFilePath()
	d, err := cmd.LoadConfig(config)
	if err != nil {
		fmt.Println("failed to load config:", config)
	}
	fmt.Println(d)
}
