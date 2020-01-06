package cohesity

import (
	"errors"
	"fmt"
	"log"
	"testing"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccPhysicalEditionCluster(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPhysicalEditionClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPhysicalEditionClusterConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccPhysicalEditionClusterExists(),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "cluster_name",
						"AcceptanceTestTerraformPhysicalEditionCluster"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "enable_encryption", "true"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "cluster_gateway", "10.2.32.1"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "enable_fips_mode", "true"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "metadata_fault_tolerance", "0"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "virtual_ip_hostname", "test"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "cluster_subnet_mask", "255.255.240.0"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "encryption_keys_rotation_period", "1"),
				),
			},
		},
	})
}

func testAccCheckPhysicalEditionClusterDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to authenticate with Cohesity")
	}
	result, infoErr := client.Cluster().GetBasicClusterInfo()
	if infoErr == nil && result.Name != nil {
		return fmt.Errorf("Test physical edition cluster (%s) still exists", *result.Name)
	}
	return nil
}

func testAccPhysicalEditionClusterExists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
			cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}
		result, infoErr := client.Cluster().GetBasicClusterInfo()
		if infoErr == nil && result.Name == nil {
			return fmt.Errorf("Failed to create the test physical edition cluster")
		}
		var fetchStats, fetchTimeSeriesSchema = false, false
		clusterDetails, clusterDetailsErr := client.Cluster().GetCluster(&fetchStats, &fetchTimeSeriesSchema)
		if clusterDetailsErr == nil {
			if *clusterDetails.Name != "AcceptanceTestTerraformPhysicalEditionCluster" {

				return fmt.Errorf("Failed to validate test physical edition cluster name")
			}
		}
		return nil
	}
}

const testAccPhysicalEditionClusterConfig = `
provider "cohesity" {
	cluster_vip = "10.9.33.113"
	cluster_username = "admin"
	cluster_domain = "LOCAL"
}

resource "cohesity_physical_edition_cluster" "physical"{
		cluster_name = "AcceptanceTestTerraformPhysicalEditionCluster"
		dns_servers = ["10.2.145.14"]
		ntp_servers = ["time.google.com"]
		domain_names = ["eng.cohesity.com"]
		cluster_subnet_mask = "255.255.240.0"
		cluster_gateway = "10.2.32.1"
		enable_encryption = true
		enable_fips_mode = true
		encryption_keys_rotation_period = 1
		metadata_fault_tolerance = 0
		virtual_ips = ["10.2.33.137"]
		ipmi_gateway = "10.2.144.54"
		ipmi_subnet_mask = "255.255.240.0"
		virtual_ip_hostname = "test"
		ipmi_username="cohesity"
		node_configs {
					node_ip="10.9.33.113"
					node_ipmi_ip="10.9.33.113"
					}
		node_configs {
					node_ip="10.9.33.114"
					node_ipmi_ip="10.9.33.114"
					}
		node_configs {
					node_ip="10.9.33.115"
					node_ipmi_ip="10.9.33.115"
					}
}
`
