package gcp

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/occu-io/nmap-ip-gen/output"
	"github.com/occu-io/nmap-ip-gen/parser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

type GoogleCloudPlatformService struct {
	ctx context.Context
	svc *compute.Service
}

func Do(p *parser.Data) {
	var g GoogleCloudPlatformService
	g.getService(p)
	g.getInstancesIPs(p)
}

func (service *GoogleCloudPlatformService) getCredentials(p *parser.Data) *google.Credentials {
	serviceAccountPath := p.Ini.Cfg.Section(p.Ini.SectionName).Key("serviceAccount").String()

	data, err := ioutil.ReadFile(serviceAccountPath)
	if err != nil {
		log.Fatal(err)
	}

	creds, err := google.CredentialsFromJSON(service.ctx, data, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	return creds
}

func (service *GoogleCloudPlatformService) getService(p *parser.Data) {
	var err error
	service.ctx = context.Background()

	ts := service.getCredentials(p).TokenSource

	client := oauth2.NewClient(service.ctx, ts)
	service.svc, err = compute.New(client)
	if err != nil {
		log.Fatal(err)
	}
}

func (service *GoogleCloudPlatformService) getInstancesIPs(p *parser.Data) []string {
	s := make([]string, 0)

	projectID := p.Ini.Cfg.Section(p.Ini.SectionName).Key("project").String()

	req := service.svc.Addresses.AggregatedList(projectID)
	if err := req.Pages(service.ctx, func(page *compute.AddressAggregatedList) error {
		for _, instance := range page.Items {
			// TODO: Change code below to process each `instance` resource:
			for _, address := range instance.Addresses {
				if address.AddressType != "INTERNAL" {
					output.Append(p, address.Address+"\n")
				}
			}

		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return s
}
