package core

import (
	"github.com/occu-io/nmap-ip-gen/core/asn"
	"github.com/occu-io/nmap-ip-gen/core/gcp"
	"github.com/occu-io/nmap-ip-gen/core/manual"
	"github.com/occu-io/nmap-ip-gen/parser"
)

var Func = map[string]func(*parser.Data){
	"manual": manual.Do,
	"gcp":    gcp.Do,
	"asn":    asn.Do,
}
