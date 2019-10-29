package parser

import "flag"

type CommandLineConfig struct {
	ConfigPath string
}

func CommandLineConfigNew() *CommandLineConfig {
	return &CommandLineConfig{}
}

func (c *CommandLineConfig) CommandLineParse() {
	flag.StringVar(&c.ConfigPath, "config", "conf/config.ini", "path to config file")
	flag.Parse()
}
