# Cohesity Terraform Provider
## Requirements

- Terraform 0.12.4+
- Go 1.12.6+ (to build the provider plugin)

## Building the Provider

- Install Go and setup the GOPATH
- create the directory $GOPATH/src/github.com/cohesity <br>
 mkdir -p $GOPATH/src/github.com/cohesity
- cd $GOPATH/src/github.com/cohesity
- clone the cohesity terraform provider repository <br>
  git clone https://github.com/cohesity/terraform-provider-cohesity.git
- clone v1.1.2 branch of cohesity management sdk for go <br>
   git clone -b v1.1.2 https://github.com/cohesity/management-sdk-go.git 
- go get github.com/hashicorp/terraform
- go get github.com/satori/go.uuid
- go get github.com/apimatic/form
- go build $GOPATH/src/github.com/cohesity/terraform-provider-cohesity
(If you want to build a binary for different OS and platform, set GOOS and GOARCH environment variables for target operating system and architecture)
- You can find the cohesity provider binary in $GOPATH/bin

## Running the acceptance tests

Change the configurations in resource_cohesity_cloud_edition_cluster_test.go, resource_cohesity_physical_edition_cluster_test.go, resource_cohesity_virtual_edition_cluster_test.go files to configure the test clusters. <br>

Before running the acceptance tests make sure the provider credentials and cluster license keys are set. They can be set directy in the provider configuration or 
in environment variables

License keys can be set as following environment variables:
 - CLOUD_COHESITY_CLUSTER_LICENCE_KEY <br>
 - PHYSICAL_COHESITY_CLUSTER_LICENCE_KEY <br>
 - VIRTUAL_COHESITY_CLUSTER_LICENCE_KEY

Password can be set as following environment variable:
- COHESITY_PASSWORD 

> **Note:**
- Make sure the password is same for the clusters when setting it as environment variable 

Terraform requires the following environment variable to be set in order to run the acceeptance tests. Can be set to any value
- TF_ACC

After setting the environment variables <br> 
`cd $GOPATH/src/github.com/cohesity/terraform-provider-cohesity` <br>
`go test -v cohesity/* -timeout 120m`

## Using the provider

### Steps to create a cohesity virtual edition cluster
- Install terraform 0.12.4 or above
- Download the compiled binary from here or build the provider following the steps above <br>
  MacOS : https://drive.google.com/drive/folders/1aX0yOcvnWFaqZ-r83FpIW7_Vo-Hs1ffH?usp=sharing<br>
  Linux : https://drive.google.com/drive/folders/1QL3aTWePQUwZj7sXbCgQyPZHpMpfIO38?usp=sharing<br>
- Place the provider binary in ~/.terraform.d/plugins/ directory on Linux or Mac and %APPDATA%\terraform.d\plugins\ directory on Windows
- create a directory for example cohesity_configuration
- cd cohesity_configuration
- Write terraform configuration in a file for example main.tf <br>

    sample configuration:

    ```
    provider "cohesity" {
            cluster_vip = "10.2.35.147"
            cluster_username = "admin"
            cluster_domain = "LOCAL"
    }

    resource "cohesity_virtual_edition_cluster" "virtual"{
                cluster_name = "TerraformVirtaulEditionCluster"
                dns_servers = ["10.2.145.14"]
                ntp_servers = ["time.google.com"]
                domain_names = ["eng.cohesity.com"]
                cluster_subnet_mask = "255.255.240.0"
                cluster_gateway = "10.2.32.1"
                enable_encryption = true
                enable_fips_mode = true
                encryption_keys_rotation_period = 1
                metadata_fault_tolerance = 0
                virtual_ips = ["10.2.35.147"]
                virtual_ip_hostname = "test"
                node_ips = ["10.2.35.147"]
    }

    ```

    In this sample configuration, the provider credentails cohesity_vip, cohesity_username, cohesity_domain are given in the configuration and cohesity_password is set in the environment variable, so it is not seen in the configuration. <br>

    To set the cohesity password environment variable<br>
    `export COHESITY_PASSWORD=abcd`

    The cohesity provider credentails and cluster license keys can be set as static variables in the configuration or set as environment variables. When they are set as environment variables, they need not be set again in the configuration<br>

    Set virtual edition cluster license key as environment variable <br>
    `export VIRTUAL_COHESITY_CLUSTER_LICENCE_KEY=abcd-abcd-adbf-kdjs`

- `terraform init` to initialize the configuration in the current directory
- `terraform apply` apply the terraform configuration. The cluster is created
- The created virtual edition cluster can be destroyed using `terraform destroy`

## Quick Start Guide

### Cohesity Terraform Provider
Cohesity terraform provider is used to interact with cohesity and automate cluster workflows using Cohesity REST API's. The provider needs to be configured with proper credentails prior to managing its resources. Currently it supports physical, cloud, virtual edition cluster creation and destruction workflows

#### Example usage

```
provider "cohesity" {
        cluster_vip = "10.2.45.143"
        cluster_username = "abc"
        cluster_password = "abc"
        cluster_domain = "LOCAL"

}
```
#### Argument Reference
- cluster_vip - (Required, string) IP or hostname of Cohesity cluster node. This can also be read from **COHESITY_IP** environment variable
- cluster_username - (Required, string) Cohesity cluster username. This can also be read from **COHESITY_USERNAME** environment variable
- cluster_password - (Required, string) Cohesity cluster password. This can also be read from **COHESITY_PASSWORD** environment variable
- cluster_domain - (Optional, string) The domain name of cohesity user. Defaults to **LOCAL**

## Resources

### cohesity_virtual_edition_cluster

Create virtual edition cluster, apply license key and destroy cluster

#### Example usage
```
provider "cohesity" {
        cluster_vip = "10.2.33.199"
        cluster_username = "abcd"
        cluster_domain = "LOCAL"
}

resource "cohesity_virtual_edition_cluster" "virtual"{
            cluster_name = "TerraformVirtaulEditionCluster"
            dns_servers = ["10.2.145.14"]
            ntp_servers = ["time.google.com"]
            domain_names = ["eng.cohesity.com"]
            cluster_subnet_mask = "255.255.240.0"
            cluster_gateway = "10.2.32.1"
            enable_encryption = true
            enable_fips_mode = true
            encryption_keys_rotation_period = 1
            metadata_fault_tolerance = 0
            virtual_ips = ["10.2.33.40"]
            virtual_ip_hostname = "test"
            node_ips = ["10.2.33.199", "10.2.33.198", "10.2.33.197"]
}

```
#### Argument Reference

The following arguments are supported:
- cluster_name - (Required, string) The name of the new Virtual edition cluster
- licence_key - (Required, string) Cohesity licence key to apply after cluster creation. This can also be read from **VIRTUAL_COHESITY_CLUSTER_LICENCE_KEY** environment variable 
- metadata_fault_tolerance - (Optional, int) The metadata fault tolerance. Default value is **0**
- enable_encryption - (Optional, bool) Specifies whether or not to enable encryption. If encryption is enabled, all data on the cluster will be encrypted. Default value is **true**
- enable_fips_mode - (Optional, bool) Specifies whether or not to enable FIPS mode. This must be set to true in order to enable FIPS. Default value is **true**
- encryption_keys_rotation_period - (Optional, int) The rotation period for encryption keys in days. The default value is **1**
- cluster_gateway - (Required, string) The default gateway IP address for the cluster network
- cluster_subnet_mask - (Required, string) The subnet mask of the cluster network
- domain_names - (Required, set of strings) The domain names to configure on the cluster
- ntp_servers - (Required, set of strings) The NTP servers to configure on the cluster
- dns_servers - (Required, set of strings) The DNS servers to configure on the cluster
- virtual_ips - (Required, set of strings) The virtauls IPs for the new cluster
- operation_timeout - (Optional, int) The time to wait in minutes for cluster creation or destruction. The default value is **120 minutes**
- virtual_ip_hostname - (Required, string) The virtual IP hostname
- node_ips - (Required, set of strings) IP addresses of the nodes in the cluster

#### Attributes Reference
The following attributes are exported:
- id - ID of the cluster

### cohesity_cloud_edition_cluster

Create cloud edition cluster, apply license key and destroy cluster

#### Example usage
```
provider "cohesity" {
        cluster_vip = "10.2.45.143"
        cluster_username = "abcd"
        cluster_domain = "LOCAL"
}

resource "cohesity_cloud_edition_cluster" "cloud"{
            cluster_name = "TerraformCloudEditionCluster"
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
```
#### Argument Reference

The following arguments are supported:
- cluster_name - (Required, string) The name of the new Cloud edition cluster
- licence_key - (Required, string) Cohesity licence key to apply after cluster creation. This can also be read from **CLOUD_COHESITY_CLUSTER_LICENCE_KEY** environment variable 
- metadata_fault_tolerance - (Optional, int) The metadata fault tolerance. Default value is **0**
- enable_encryption - (Optional, bool) Specifies whether or not to enable encryption. If encryption is enabled, all data on the cluster will be encrypted. Default value is **true**
- enable_fips_mode - (Optional, bool) Specifies whether or not to enable FIPS mode. This must be set to true in order to enable FIPS. Default value is **true**
- encryption_keys_rotation_period - (Optional, int) The rotation period for encryption keys in days. The default value is **1**
- cluster_gateway - (Required, string) The default gateway IP address for the cluster network
- cluster_subnet_mask - (Required, string) The subnet mask of the cluster network
- domain_names - (Required, set of strings) The domain names to configure on the cluster
- ntp_servers - (Required, set of strings) The NTP servers to configure on the cluster
- dns_servers - (Required, set of strings) The DNS servers to configure on the cluster
- node_ips - (Required, set of strings) IP addresses of the nodes in the cluster
- operation_timeout - (Optional, int) The time to wait in minutes for cluster creation or destruction. The default value is **120 minutes**

#### Attributes Reference
The following attributes are exported:
- id - ID of the cluster

### cohesity_physical_edition_cluster
Create physical edition cluster, apply license key and destroy cluster

#### Example usage
```
provider "cohesity" {
        cluster_vip = "10.9.33.133"
        cluster_username = "abcd"
        cluster_domain = "LOCAL"
}

resource "cohesity_virtual_edition_cluster" "physical"{
            cluster_name = "TerraformPhysicalEditionCluster"
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
                        node_ip="10.9.33.133"
                        node_ipmi_ip="10.9.33.133"
                        }
            node_configs {
                        node_ip="10.9.33.134"
                        node_ipmi_ip="10.9.33.134"
                        }
            node_configs {
                        node_ip="10.9.33.135"
                        node_ipmi_ip="10.9.33.135"
                        }
}
```

### Argument Reference
The following arguments are supported:
- cluster_name - (Required, string) The name of the new Physical edition cluster
- licence_key - (Required, string) Cohesity licence key to apply after cluster creation. This can also be read from **PHYSICAL_COHESITY_CLUSTER_LICENCE_KEY** environment variable 
- metadata_fault_tolerance - (Optional, int) The metadata fault tolerance. Default value is **0**
- enable_encryption - (Optional, bool) Specifies whether or not to enable encryption. If encryption is enabled, all data on the cluster will be encrypted. Default value is true
- enable_fips_mode - (Optional, bool) Specifies whether or not to enable FIPS mode. This must be set to true in order to enable FIPS. Default value is **true**
- encryption_keys_rotation_period - (Optional, int) The rotation period for encryption keys in days. The default value is **1**
- cluster_gateway - (Required, string) The default gateway IP address for the cluster network
- cluster_subnet_mask - (Required, string) The subnet mask of the cluster network
- domain_names - (Required, set of strings) The domain names to configure on the cluster
- ntp_servers - (Required, set of strings) The NTP servers to configure on the cluster
- dns_servers - (Required, set of strings) The DNS servers to configure on the cluster
- virtual_ips - (Required, set of strings) The virtauls IPs for the new cluster
- operation_timeout - (Optional, int) The time to wait in minutes for cluster creation or destruction. The default value is **120 minutes**
- virtual_ip_hostname - (Required, string) The virtual IP hostname
- ipmi_username - (Required, string) The IPMI username. This can also be read from **PHYSICAL_COHESITY_CLUSTER_IPMI_USERNAME** environment variable 
- ipmi_password - (Required, string) The IPMI password. This can also be read from **PHYSICAL_COHESITY_CLUSTER_IPMI_PASSWORD** environment variable 
- ipmi_gateway - (Required, string) The default gateway IP address for the IPMI network
- ipmi_subnet_mask - (Required, string) The subnet mask for the IPMI network
- node_configs - (Required, block) node_configs is a block within the configuration to configure the nodes in the cluster. The block can be repeated to configure multiple nodes in the cluster. Each block supports the following:
   - node_ip - (Required, string) IP address of the node
   - node_ipmi_ip - (Required, string) IPMI IP for this node

#### Attributes Reference
The following attributes are exported:
- id - ID of the cluster

