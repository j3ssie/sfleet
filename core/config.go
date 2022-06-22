package core

import (
	"fmt"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"

	"github.com/j3ssie/sfleet/libs"
	"github.com/j3ssie/sfleet/utils"
)

type WorkerPool struct {
	Workers []Worker `yaml:"workers" json:"workers"`
}

type Worker struct {
	// Status        string `yaml:"status"`
	// Protocol      string `yaml:"protocol"`
	// Credential    string `yaml:"credential"`
	// Jwt           string `yaml:"jwt"`

	Name string `yaml:"name" json:"name"`
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
	User string `yaml:"user" json:"user"`

	Key        string `yaml:"key" json:"key"`
	PassPhrase string `yaml:"passphrase" json:"passphrase"`

	// Stat          struct {
	// 	Raw string `yaml:"raw"`
	// } `yaml:"stat"`
	// CPUIdle float64 `yaml:"cpuidle"`
	// MemIdle float64 `yaml:"memidle"`
}

// InitConfig Init the config
func InitConfig(options *libs.Options) {
	options.ConfigFile = utils.NormalizePath(options.ConfigFile)
	options.RootFolder = filepath.Dir(options.ConfigFile)
	utils.MakeDir(options.RootFolder)

	if !utils.FileExists(options.ConfigFile) {
		var wks WorkerPool
		if data, err := jsoniter.MarshalToString(&wks); err == nil {
			utils.WriteToFile(options.ConfigFile, data)
			return
		}
	}
}

// GetConfig load config
func GetConfig(options libs.Options) (WorkerPool, error) {
	raw := utils.GetFileContent(options.ConfigFile)
	var wks WorkerPool
	err := jsoniter.UnmarshalFromString(raw, &wks)
	if err != nil {
		utils.ErrorF("Failed to read the configuration file: %v\n", err)
		return wks, fmt.Errorf("error reading config")
	}
	return wks, nil
}

func WriteConfig(wks WorkerPool, options libs.Options) error {
	utils.InforF("Writing %v worker to %v", len(wks.Workers), options.ConfigFile)
	data, err := jsoniter.MarshalToString(&wks)
	if err == nil {
		utils.WriteToFile(options.ConfigFile, data)
		return nil
	}
	utils.ErrorF("%v", err)
	return fmt.Errorf("error writing config")
}
