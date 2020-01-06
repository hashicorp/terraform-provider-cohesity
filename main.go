package main

import (
	"github.com/cohesity/terraform-provider-cohesity/cohesity"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cohesity.Provider})
}
