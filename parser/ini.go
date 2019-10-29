package parser

import (
	"log"

	"gopkg.in/ini.v1"
)

type IniFileConfig struct {
	Cfg         *ini.File
	SectionName string
}

func IniFileConfigNew() *IniFileConfig {
	return &IniFileConfig{}
}

func (i *IniFileConfig) IniFileParse(config_path string) {
	var err error

	log.Printf("Parsing file: %s", config_path)
	i.Cfg, err = ini.Load(config_path)
	if err != nil {
		log.Fatal(err)
	}
}
