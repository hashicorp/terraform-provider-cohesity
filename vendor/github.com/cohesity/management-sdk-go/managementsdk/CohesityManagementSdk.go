// Copyright 2019 Cohesity Inc.
// Copyright 2019 Cohesity Inc.

package CohesityManagementSdk

import(
	"github.com/cohesity/management-sdk-go/configuration"
	"github.com/cohesity/management-sdk-go/accesstokens"
	"github.com/cohesity/management-sdk-go/activedirectory"
	"github.com/cohesity/management-sdk-go/alerts"
	"github.com/cohesity/management-sdk-go/antivirusservicegroup"
	"github.com/cohesity/management-sdk-go/audit"
	"github.com/cohesity/management-sdk-go/cluster"
	"github.com/cohesity/management-sdk-go/certificates"
	"github.com/cohesity/management-sdk-go/clusters"
	"github.com/cohesity/management-sdk-go/clusterpartitions"
	"github.com/cohesity/management-sdk-go/groups"
	"github.com/cohesity/management-sdk-go/idps"
	"github.com/cohesity/management-sdk-go/interfacegroup"
	"github.com/cohesity/management-sdk-go/kmsconfiguration"
	"github.com/cohesity/management-sdk-go/ldapprovider"
	"github.com/cohesity/management-sdk-go/network"
	"github.com/cohesity/management-sdk-go/views"
	"github.com/cohesity/management-sdk-go/nodes"
	"github.com/cohesity/management-sdk-go/packages"
	"github.com/cohesity/management-sdk-go/protectionsources"
	"github.com/cohesity/management-sdk-go/customreporting"
	"github.com/cohesity/management-sdk-go/principals"
	"github.com/cohesity/management-sdk-go/privileges"
	"github.com/cohesity/management-sdk-go/protectionjobs"
	"github.com/cohesity/management-sdk-go/protectionobjects"
	"github.com/cohesity/management-sdk-go/protectionpolicies"
	"github.com/cohesity/management-sdk-go/protectionruns"
	"github.com/cohesity/management-sdk-go/remotecluster"
	"github.com/cohesity/management-sdk-go/remoterestore"
	"github.com/cohesity/management-sdk-go/restoretasks"
	"github.com/cohesity/management-sdk-go/roles"
	"github.com/cohesity/management-sdk-go/routes"
	"github.com/cohesity/management-sdk-go/search"
	"github.com/cohesity/management-sdk-go/notifications"
	"github.com/cohesity/management-sdk-go/smbfileopens"
	"github.com/cohesity/management-sdk-go/staticroute"
	"github.com/cohesity/management-sdk-go/statistics"
	"github.com/cohesity/management-sdk-go/tenant"
	"github.com/cohesity/management-sdk-go/tenants"
	"github.com/cohesity/management-sdk-go/vaults"
	"github.com/cohesity/management-sdk-go/viewboxes"
	"github.com/cohesity/management-sdk-go/vlan"
	"github.com/cohesity/management-sdk-go/models"
)


/*
 * Interface for the COHESITYMANAGEMENTSDK_IMPL
 */
type COHESITYMANAGEMENTSDK interface {
    AccessTokens()          accesstokens.ACCESSTOKENS
    ActiveDirectory()       activedirectory.ACTIVEDIRECTORY
    Alerts()                alerts.ALERTS
    AntivirusServiceGroup()       antivirusservicegroup.ANTIVIRUSSERVICEGROUP
    Audit()                 audit.AUDIT
    Cluster()               cluster.CLUSTER
    Certificates()          certificates.CERTIFICATES
    Clusters()              clusters.CLUSTERS
    ClusterPartitions()       clusterpartitions.CLUSTERPARTITIONS
    Groups()                groups.GROUPS
    Idps()                  idps.IDPS
    InterfaceGroup()        interfacegroup.INTERFACEGROUP
    KmsConfiguration()       kmsconfiguration.KMSCONFIGURATION
    LdapProvider()          ldapprovider.LDAPPROVIDER
    Network()               network.NETWORK
    Views()                 views.VIEWS
    Nodes()                 nodes.NODES
    Packages()              packages.PACKAGES
    ProtectionSources()       protectionsources.PROTECTIONSOURCES
    CustomReporting()       customreporting.CUSTOMREPORTING
    Principals()            principals.PRINCIPALS
    Privileges()            privileges.PRIVILEGES
    ProtectionJobs()        protectionjobs.PROTECTIONJOBS
    ProtectionObjects()       protectionobjects.PROTECTIONOBJECTS
    ProtectionPolicies()       protectionpolicies.PROTECTIONPOLICIES
    ProtectionRuns()        protectionruns.PROTECTIONRUNS
    RemoteCluster()         remotecluster.REMOTECLUSTER
    RemoteRestore()         remoterestore.REMOTERESTORE
    RestoreTasks()          restoretasks.RESTORETASKS
    Roles()                 roles.ROLES
    Routes()                routes.ROUTES
    Search()                search.SEARCH
    Notifications()         notifications.NOTIFICATIONS
    SMBFileOpens()          smbfileopens.SMBFILEOPENS
    StaticRoute()           staticroute.STATICROUTE
    Statistics()            statistics.STATISTICS
    Tenant()                tenant.TENANT
    Tenants()               tenants.TENANTS
    Vaults()                vaults.VAULTS
    ViewBoxes()             viewboxes.VIEWBOXES
    Vlan()                  vlan.VLAN
    Configuration()         configuration.CONFIGURATION
    Authorize()				(*models.AccessToken,error)
}

/*
 * Factory for the COHESITYMANAGEMENTSDK interface returning COHESITYMANAGEMENTSDK_IMPL
 */
func NewCohesitySdkClient(clustervip string, username string, password string, domain string) (COHESITYMANAGEMENTSDK, error) {
    cohesityManagementSdkClient := new(COHESITYMANAGEMENTSDK_IMPL)
    cohesityManagementSdkClient.config = configuration.NewCONFIGURATION()

    cohesityManagementSdkClient.config.SetUsername(username)
    cohesityManagementSdkClient.config.SetPassword(password)
    cohesityManagementSdkClient.config.SetDomain(domain)
    cohesityManagementSdkClient.config.SetClusterVip(&clustervip)
    cohesityManagementSdkClient.config.SetSkipSSL(true)
    _, err := cohesityManagementSdkClient.Authorize()
    if err != nil {
       return nil, err
    }
    return cohesityManagementSdkClient, nil
}

/*
 * Factory for the COHESITYMANAGEMENTSDK interface that takes and sets the AccessToken
 */
func NewCohesitySdkClientWithToken(clustervip string, mgmtauthtoken *models.AccessToken) COHESITYMANAGEMENTSDK {
    cohesityManagementSdkClient := new(COHESITYMANAGEMENTSDK_IMPL)
    cohesityManagementSdkClient.config = configuration.NewCONFIGURATION()
    cohesityManagementSdkClient.config.SetClusterVip(&clustervip)
    cohesityManagementSdkClient.config.SetSkipSSL(true)
    cohesityManagementSdkClient.config.SetAccessToken(mgmtauthtoken)
    return cohesityManagementSdkClient
}
