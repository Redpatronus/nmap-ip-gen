#!/bin/bash

TMP_FILE=$(mktemp -u)
NMAP_IP_GEN_BINARY="../nmap-ip-gen"
NMAP_OUTPUT_PARSE_BINARY="/usr/bin/nmap-parse-output"
NMAP_BINARY="/usr/bin/nmap"
OUT="output.html"

$NMAP_IP_GEN_BINARY -config /etc/nmap-ip-gen/config.ini
sudo $NMAP_BINARY -iL ./output.txt -Pn -sT -sV --script=asn-query,whois-ip,ip-geolocation-maxmind,ssl-cert,ssl-enum-ciphers,vulners --version-intensity 5 -T4 -oX $TMP_FILE
$NMAP_OUTPUT_PARSE_BINARY $TMP_FILE html-bootstrap > $OUT
sudo rm $TMP_FILE

