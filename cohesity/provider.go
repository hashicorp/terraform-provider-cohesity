package cohesity

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"cohesity_vip": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COHESITY_IP", ""),
				Description: "IP or hostname of Cohesity cluster node",
			},
			"cohesity_username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COHESITY_USERNAME", ""),
				Description: "Cohesity cluster username",
			},
			"cohesity_password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COHESITY_PASSWORD", ""),
				Description: "Cohesity cluster password",
			},
			"cohesity_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "LOCAL",
				Description: "The domain name of cohesity user",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cohesity_cloud_edition_cluster":    resourceCohesityCloudEditionCluster(),
			"cohesity_virtual_edition_cluster":  resourceCohesityVirtualEditionCluster(),
			"cohesity_physical_edition_cluster": resourceCohesityPhysicalEditionCluster(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		cohesityVip:      d.Get("cohesity_vip").(string),
		cohesityUserName: d.Get("cohesity_username").(string),
		cohesityPassword: d.Get("cohesity_password").(string),
		cohesityDomain:   d.Get("cohesity_domain").(string),
	}
	return config, nil
}
