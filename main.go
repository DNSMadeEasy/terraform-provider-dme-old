package main

import (
	"github.com/DNSMadeEasy/terraform-provider-dme/dme"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return dme.Provider()
		},
	})
}
