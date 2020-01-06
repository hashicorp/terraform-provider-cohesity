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

func TestAccCloudEditionCluster(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudEditionClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudEditionClusterConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCloudEditionClusterExists(),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "cluster_name",
						"AcceptanceTestTerraformCloudEditionCluster"),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "enable_encryption", "true"),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "cluster_gateway", "10.2.32.1"),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "enable_fips_mode", "true"),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "metadata_fault_tolerance", "0"),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "cluster_subnet_mask", "255.255.240.0"),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "encryption_keys_rotation_period", "1"),
				),
			},
		},
	})
}

func testAccCheckCloudEditionClusterDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to authenticate with Cohesity")
	}
	result, infoErr := client.Cluster().GetBasicClusterInfo()
	if infoErr == nil && result.Name != nil {
		return fmt.Errorf("Test cloud edition cluster (%s) still exists", *result.Name)
	}
	return nil
}

func testAccCloudEditionClusterExists() resource.TestCheckFunc {
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
			return fmt.Errorf("Failed to create the test cloud edition cluster")
		}
		var fetchStats, fetchTimeSeriesSchema = false, false
		clusterDetails, clusterDetailsErr := client.Cluster().GetCluster(&fetchStats, &fetchTimeSeriesSchema)
		if clusterDetailsErr == nil {
			if *clusterDetails.Name != "AcceptanceTestTerraformCloudEditionCluster" {

				return fmt.Errorf("Failed to validate test cloud edition cluster name")
			}
		}
		return nil
	}
}

const testAccCloudEditionClusterConfig = `
provider "cohesity" {
	cluster_vip = "10.2.45.143"
	cluster_username = "admin"
	cluster_domain = "LOCAL"
}

resource "cohesity_cloud_edition_cluster" "cloud"{
		cluster_name = "AcceptanceTestTerraformCloudEditionCluster"
		dns_servers = ["10.2.145.14"]
		ntp_servers = ["time.google.com"]
		domain_names = ["eng.cohesity.com"]
		cluster_subnet_mask = "255.255.240.0"
		cluster_gateway = "10.2.32.1"
		enable_encryption = true
		enable_fips_mode = true
		encryption_keys_rotation_period = 1
		metadata_fault_tolerance = 0
		node_ips = ["10.2.45.143", "10.2.45.144", "10.2.45.145"]        
}
`
