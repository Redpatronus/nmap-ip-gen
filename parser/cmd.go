package parser

import "flag"

type CommandLineConfig struct {
	ConfigPath string
	OutputPath string
}

func CommandLineConfigNew() *CommandLineConfig {
	return &CommandLineConfig{}
}

func (c *CommandLineConfig) CommandLineParse() {
	flag.StringVar(&c.ConfigPath, "config", "conf/config.ini", "path to config file")
	flag.StringVar(&c.OutputPath, "output", "/tmp/nmap-ip-gen-output.txt", "path to output")
	flag.Parse()
}
