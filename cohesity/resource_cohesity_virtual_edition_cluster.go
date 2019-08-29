package cohesity

import (
	"errors"
	"log"
	"strconv"
	"time"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCohesityVirtualEditionCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceCohesityVirtualEditionClusterCreate,
		Read:   resourceCohesityVirtualEditionClusterRead,
		Delete: resourceCohesityVirtualEditionClusterDelete,
		Update: resourceCohesityVirtualEditionClusterUpdate,
		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the new Virtual edition cluster",
			},
			"licence_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("VIRTUAL_COHESITY_CLUSTER_LICENCE_KEY", ""),
				Description: "Cohesity licence key to apply after cluster creation",
			},
			"metadata_fault_tolerance": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The metadata fault tolerance",
			},
			"enable_encryption": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Specifies whether or not to enable encryption. If encryption is enabled, all data on the cluster will be encrypted",
			},
			"enable_fips_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Specifies whether or not to enable FIPS mode. This must be set to true in order to enable FIPS",
			},
			"encryption_keys_rotation_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "The rotation period for encryption keys in days",
			},
			"cluster_gateway": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The default gateway IP address for the cluster network",
			},
			"cluster_subnet_mask": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The subnet mask of the cluster network",
			},
			"domain_names": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The domain names to configure on the cluster",
			},
			"ntp_servers": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The NTP servers to configure on the cluster",
			},
			"dns_servers": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The DNS servers to configure on the cluster",
			},
			"virtual_ips": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The virtauls IPs for the new cluster",
			},
			"operation_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     120,
				Description: "The time to wait in minutes for cluster creation or destruction",
			},
			"virtual_ip_hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The virtual IP hostname",
			},
			"node_configs": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The configuration for the nodes in the new cluster",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IP address of the node",
						},
						"node_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
							Description: "Id for this node",
						},
					},
				},
			},
		},
	}
}

func resourceCohesityVirtualEditionClusterCreate(d *schema.ResourceData, m interface{}) error {
	var cohesityConfig = m.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.cohesityVip, cohesityConfig.cohesityUserName, cohesityConfig.cohesityPassword, cohesityConfig.cohesityDomain)
	if err != nil {
		return errors.New("Failed to authenticate with Cohesity")
	}

	var clusterName = d.Get("cluster_name").(string)
	var metadataFaultTolerance = int64(d.Get("metadata_fault_tolerance").(int))
	var enableEncryption = d.Get("enable_encryption").(bool)
	var enableFipsMode = d.Get("enable_fips_mode").(bool)
	var encryptionKeysRotationPeriod = int64(d.Get("encryption_keys_rotation_period").(int))
	var clusterGateway = d.Get("cluster_gateway").(string)
	var clusterSubnetMask = d.Get("cluster_subnet_mask").(string)
	var vipHostName = d.Get("virtual_ip_hostname").(string)

	log.Printf("[INFO] Create virtual edition cluster: %s", clusterName)

	domainNames := make([]string, d.Get("domain_names").(*schema.Set).Len())
	for i, name := range d.Get("domain_names").(*schema.Set).List() {
		domainNames[i] = name.(string)
	}

	ntpServers := make([]string, d.Get("ntp_servers").(*schema.Set).Len())
	for i, ntp := range d.Get("ntp_servers").(*schema.Set).List() {
		ntpServers[i] = ntp.(string)
	}

	dnsServers := make([]string, d.Get("dns_servers").(*schema.Set).Len())
	for i, dns := range d.Get("dns_servers").(*schema.Set).List() {
		dnsServers[i] = dns.(string)
	}

	vips := make([]string, d.Get("virtual_ips").(*schema.Set).Len())
	for i, vip := range d.Get("virtual_ips").(*schema.Set).List() {
		vips[i] = vip.(string)
	}

	nodeConfigs := make([]*models.VirtualNodeConfiguration, d.Get("node_configs").(*schema.Set).Len())
	for i, config := range d.Get("node_configs").(*schema.Set).List() {
		nodeConfig := config.(map[string]interface{})
		nodeIP := nodeConfig["node_ip"].(string)
		nodeID := int64(nodeConfig["node_id"].(int))
		nodeConfigs[i] = &models.VirtualNodeConfiguration{
			NodeId: &nodeID,
			NodeIp: &nodeIP,
		}
	}

	var requestBody models.CreateVirtualClusterParameters
	requestBody.ClusterName = clusterName
	requestBody.MetadataFaultTolerance = &metadataFaultTolerance
	requestBody.NodeConfigs = nodeConfigs
	requestBody.EncryptionConfig = &models.EncryptionConfiguration{
		EnableEncryption: &enableEncryption,
		EnableFipsMode:   &enableFipsMode,
		RotationPeriod:   &encryptionKeysRotationPeriod,
	}
	requestBody.NetworkConfig = models.NetworkConfiguration{
		ClusterGateway:    &clusterGateway,
		ClusterSubnetMask: &clusterSubnetMask,
		DnsServers:        &dnsServers,
		DomainNames:       &domainNames,
		NtpServers:        &ntpServers,
		VipHostname:       &vipHostName,
		Vips:              &vips,
	}
	result, err := client.Clusters().CreateVirtualCluster(&requestBody)
	if err != nil {
		return errors.New("Failed to create Virtual edition cluster")
	}
	log.Printf("[INFO] Create request for virtual edition cluster (%s) successful", clusterName)
	var timeout = d.Get("operation_timeout").(int) * 60
	for timeout > 0 {
		result, err2 := client.Cluster().GetBasicClusterInfo()
		if err2 == nil && result.Name != nil {
			log.Printf("[INFO] Successfully created virtual edition cluster %s", clusterName)
			break
		}
		time.Sleep(30 * time.Second)
		timeout -= 30
	}

	client, err = CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.cohesityVip, cohesityConfig.cohesityUserName, cohesityConfig.cohesityPassword, cohesityConfig.cohesityDomain)
	if err != nil {
		return errors.New("Failed to authenticate with Cohesity")
	}

	var requestBodyLicence models.LicenceClusterParameters
	signedTime := int64(time.Now().Unix())
	signedVersion := int64(2)
	requestBodyLicence.LicenseKey = d.Get("licence_key").(string)
	requestBodyLicence.SignedByUser = cohesityConfig.cohesityUserName
	requestBodyLicence.SignedVersion = &signedVersion
	requestBodyLicence.SignedTime = &signedTime

	err = client.Clusters().ApplyClusterLicence(&requestBodyLicence)

	if err != nil {
		log.Printf("[WARNING] Failed to apply licence for virtaul edition cluster %s", clusterName)
		d.Set("licence_key", "")
	}
	log.Printf("[INFO] Successfully created and applied licence to  virtual edition cluster %s", clusterName)
	d.SetId(strconv.FormatInt(*result.ClusterId, 10))
	return resourceCohesityVirtualEditionClusterRead(d, m)
}

func resourceCohesityVirtualEditionClusterRead(d *schema.ResourceData, m interface{}) error {
	var clusterName = d.Get("cluster_name").(string)
	var cohesityConfig = m.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.cohesityVip, cohesityConfig.cohesityUserName, cohesityConfig.cohesityPassword, cohesityConfig.cohesityDomain)
	if err != nil {
		return errors.New("Failed to authenticate with Cohesity")
	}

	result, err := client.Cluster().GetBasicClusterInfo()
	if err == nil && result.Name == nil {
		log.Printf("[INFO] %s virtaul edition cluster doesn't exist", clusterName)
		d.SetId("")
	}
	return nil
}

func resourceCohesityVirtualEditionClusterDelete(d *schema.ResourceData, m interface{}) error {
	var clusterName = d.Get("cluster_name").(string)
	log.Printf("[INFO] Destroy virtual edition cluster: %s", clusterName)
	var cohesityConfig = m.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.cohesityVip, cohesityConfig.cohesityUserName, cohesityConfig.cohesityPassword, cohesityConfig.cohesityDomain)
	if err != nil {
		return errors.New("Failed to authenticate with Cohesity")
	}
	destroyErr := client.Clusters().DestroyCluster()
	if destroyErr != nil {
		return errors.New("Failed to destroy Virtual edition cluster")
	}
	log.Printf("[INFO] Destroy request for virtual edition cluster (%s) successful", clusterName)
	var timeout = d.Get("operation_timeout").(int) * 60
	for timeout > 0 {
		result, infoErr := client.Cluster().GetBasicClusterInfo()
		if infoErr == nil && result.Name == nil {
			log.Printf("[INFO] Destroyed Virtual edition cluster: %s", clusterName)
			break
		}
		time.Sleep(30 * time.Second)
		timeout -= 30
	}
	return nil
}

func resourceCohesityVirtualEditionClusterUpdate(d *schema.ResourceData, m interface{}) error {
	var clusterName = d.Get("cluster_name").(string)
	log.Printf("[INFO] Update virtual edition cluster: %s", clusterName)
	d.Partial(true)
	var cohesityConfig = m.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.cohesityVip, cohesityConfig.cohesityUserName, cohesityConfig.cohesityPassword, cohesityConfig.cohesityDomain)
	if err != nil {
		return errors.New("Failed to authenticate with Cohesity")
	}
	oldLicenceValue, _ := d.GetChange("licence_key")
	if d.HasChange("licence_key") && oldLicenceValue == "" {
		var requestBodyLicence models.LicenceClusterParameters
		signedTime := int64(time.Now().Unix())
		signedVersion := int64(2)
		requestBodyLicence.LicenseKey = d.Get("licence_key").(string)
		requestBodyLicence.SignedByUser = cohesityConfig.cohesityUserName
		requestBodyLicence.SignedVersion = &signedVersion
		requestBodyLicence.SignedTime = &signedTime
		err := client.Clusters().ApplyClusterLicence(&requestBodyLicence)
		if err != nil {
			return errors.New("Failed to update Virtual edition cluster")
		}
		d.SetPartial("licence_key")
		log.Printf("[INFO] Applied licence to virtual edition cluster: %s", clusterName)
		return resourceCohesityVirtualEditionClusterRead(d, m)
	}
	return errors.New("Failed to update Virtual edition cluster")
}
