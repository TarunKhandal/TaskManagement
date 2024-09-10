package globals

import (
	"log"

	"github.com/BurntSushi/toml"
)

var RootDirPathForAppData string

type config struct {
	App      configApp
	Database configDatabase
}

type configApp struct {
	Port        string
	Environment string
}

type configDatabase struct {
	Database      string
	First         bool
	TablesCount   int
	Tables        []string
	Migration     string
	ForeignEnable bool
}

var Configuration *config

func InitConfiguration() {
	configFlPath := RootDirPathForAppData + "configs/config.toml"
	Configuration = &config{}
	if _, err := toml.DecodeFile(configFlPath, Configuration); err != nil {
		log.Panicln("Error during InitConfiguration", err)
	}
	log.Println("Configuration initialized from file : " + configFlPath)
}
