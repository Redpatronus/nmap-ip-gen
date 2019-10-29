package main

import (
	"log"

	"github.com/occu-io/nmap-ip-gen/core"
	"github.com/occu-io/nmap-ip-gen/output"
	"github.com/occu-io/nmap-ip-gen/parser"
	"gopkg.in/ini.v1"
)

func main() {
	p := parser.Do()

	output.Open(p)
	for _, section := range p.Ini.Cfg.Sections() {
		p.Ini.SectionName = section.Name()

		if section.Name() == ini.DEFAULT_SECTION || section.Name() == "general" {
			continue
		}

		output.Append(p, "# ["+section.Name()+"]\n")
		output.Append(p, "# "+section.Key("desc").String()+"\n")

		t := section.Key("type").In("manual", []string{
			"gcp",
			"asn",
		})

		core.Func[t](p)
	}
	output.Close(p)

	log.Printf("Nmap IP list was generated to: %s\n", p.Ini.Cfg.Section("general").Key("output").String())
}
