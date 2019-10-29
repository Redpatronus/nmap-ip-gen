package output

import (
	"log"
	"os"

	"github.com/occu-io/nmap-ip-gen/parser"
)

func Open(p *parser.Data) {
	var err error
	filePath := p.Ini.Cfg.Section("general").Key("output").String()
	p.OutputFile, err = os.Create(filePath)

	if err != nil {
		log.Fatalf("could not create output file: %s\n", err)
	}
}

func Append(p *parser.Data, s string) {
	_, err := p.OutputFile.Write([]byte(s))
	if err != nil {
		log.Panicf("could not write to file: %s\n", err)
	}
}

func Close(p *parser.Data) {
	err := p.OutputFile.Close()
	if err != nil {
		log.Panicf("could not close file: %s\n", err)
	}
}
