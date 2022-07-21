package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var ConfigFile = flag.String("config", "", "Config file for Fusu.")

//	获取配置文件地址
func GetConfigFilePath() string {
	//	有参数返回参数值
	if len(*ConfigFile) > 0 {
		fmt.Println("1")
		return *ConfigFile
	}
	//	没有参数返回固定值
	if workingDir, err := os.Getwd(); err == nil {
		ConfigFile := filepath.Join(workingDir, "config.json")
		if fileExists(ConfigFile) {
			return ConfigFile
		}
	}

	return ""
}

//
func fileExists(file string) bool {
	info, err := os.Stat(file)
	return err == nil && !info.IsDir()
}

type configFileLoader func(string) (io.ReadCloser, error)

var (
	EffectiveConfigFileLoader configFileLoader
)

func LoadConfig(file string) (io.ReadCloser, error) {
	if EffectiveConfigFileLoader == nil {
		return os.Stdin, nil
	}
	return EffectiveConfigFileLoader(file)
}
