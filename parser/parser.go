package parser

import "os"

type Data struct {
	Cmd        *CommandLineConfig
	Ini        *IniFileConfig
	OutputFile *os.File
}

func Do() *Data {
	p := &Data{}

	p.Cmd = CommandLineConfigNew()
	p.Cmd.CommandLineParse()
	p.Ini = IniFileConfigNew()
	p.Ini.IniFileParse(p.Cmd.ConfigPath)

	return p
}
