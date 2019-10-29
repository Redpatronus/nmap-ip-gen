package manual

import (
	"github.com/occu-io/nmap-ip-gen/output"
	"github.com/occu-io/nmap-ip-gen/parser"
)

func Do(p *parser.Data) {
	for _, address := range p.Ini.Cfg.Section(p.Ini.SectionName).Key("ip").Strings(",") {
		output.Append(p, address+"\n")
	}
}
