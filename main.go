package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/bsquare-corp/terraform-provider-papertrail/papertrail"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: papertrail.Provider,
	})
}
