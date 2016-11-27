package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/imyousuf/terraform-provider-localcmd/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.ProvideLocalCommand,
	})
}
