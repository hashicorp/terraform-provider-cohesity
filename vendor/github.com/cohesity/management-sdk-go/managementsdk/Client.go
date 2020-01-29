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
 * Client structure as interface implementation
 */
type COHESITYMANAGEMENTSDK_IMPL struct {
     accesstokens accesstokens.ACCESSTOKENS
     activedirectory activedirectory.ACTIVEDIRECTORY
     alerts alerts.ALERTS
     antivirusservicegroup antivirusservicegroup.ANTIVIRUSSERVICEGROUP
     audit audit.AUDIT
     cluster cluster.CLUSTER
     certificates certificates.CERTIFICATES
     clusters clusters.CLUSTERS
     clusterpartitions clusterpartitions.CLUSTERPARTITIONS
     groups groups.GROUPS
     idps idps.IDPS
     interfacegroup interfacegroup.INTERFACEGROUP
     kmsconfiguration kmsconfiguration.KMSCONFIGURATION
     ldapprovider ldapprovider.LDAPPROVIDER
     network network.NETWORK
     views views.VIEWS
     nodes nodes.NODES
     packages packages.PACKAGES
     protectionsources protectionsources.PROTECTIONSOURCES
     customreporting customreporting.CUSTOMREPORTING
     principals principals.PRINCIPALS
     privileges privileges.PRIVILEGES
     protectionjobs protectionjobs.PROTECTIONJOBS
     protectionobjects protectionobjects.PROTECTIONOBJECTS
     protectionpolicies protectionpolicies.PROTECTIONPOLICIES
     protectionruns protectionruns.PROTECTIONRUNS
     remotecluster remotecluster.REMOTECLUSTER
     remoterestore remoterestore.REMOTERESTORE
     restoretasks restoretasks.RESTORETASKS
     roles roles.ROLES
     routes routes.ROUTES
     search search.SEARCH
     notifications notifications.NOTIFICATIONS
     smbfileopens smbfileopens.SMBFILEOPENS
     staticroute staticroute.STATICROUTE
     statistics statistics.STATISTICS
     tenant tenant.TENANT
     tenants tenants.TENANTS
     vaults vaults.VAULTS
     viewboxes viewboxes.VIEWBOXES
     vlan vlan.VLAN
     config  configuration.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Configuration() configuration.CONFIGURATION {
    return me.config
}
/**
     * Access to AccessTokens controller
     * @return Returns the AccessTokens() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) AccessTokens() accesstokens.ACCESSTOKENS {
    if(me.accesstokens) == nil {
        me.accesstokens = accesstokens.NewACCESSTOKENS(me.config)
    }
    return me.accesstokens
}
/**
     * Access to ActiveDirectory controller
     * @return Returns the ActiveDirectory() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) ActiveDirectory() activedirectory.ACTIVEDIRECTORY {
    if(me.activedirectory) == nil {
        me.activedirectory = activedirectory.NewACTIVEDIRECTORY(me.config)
    }
    return me.activedirectory
}
/**
     * Access to Alerts controller
     * @return Returns the Alerts() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Alerts() alerts.ALERTS {
    if(me.alerts) == nil {
        me.alerts = alerts.NewALERTS(me.config)
    }
    return me.alerts
}
/**
     * Access to AntivirusServiceGroup controller
     * @return Returns the AntivirusServiceGroup() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) AntivirusServiceGroup() antivirusservicegroup.ANTIVIRUSSERVICEGROUP {
    if(me.antivirusservicegroup) == nil {
        me.antivirusservicegroup = antivirusservicegroup.NewANTIVIRUSSERVICEGROUP(me.config)
    }
    return me.antivirusservicegroup
}
/**
     * Access to Audit controller
     * @return Returns the Audit() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Audit() audit.AUDIT {
    if(me.audit) == nil {
        me.audit = audit.NewAUDIT(me.config)
    }
    return me.audit
}
/**
     * Access to Cluster controller
     * @return Returns the Cluster() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Cluster() cluster.CLUSTER {
    if(me.cluster) == nil {
        me.cluster = cluster.NewCLUSTER(me.config)
    }
    return me.cluster
}
/**
     * Access to Certificates controller
     * @return Returns the Certificates() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Certificates() certificates.CERTIFICATES {
    if(me.certificates) == nil {
        me.certificates = certificates.NewCERTIFICATES(me.config)
    }
    return me.certificates
}
/**
     * Access to Clusters controller
     * @return Returns the Clusters() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Clusters() clusters.CLUSTERS {
    if(me.clusters) == nil {
        me.clusters = clusters.NewCLUSTERS(me.config)
    }
    return me.clusters
}
/**
     * Access to ClusterPartitions controller
     * @return Returns the ClusterPartitions() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) ClusterPartitions() clusterpartitions.CLUSTERPARTITIONS {
    if(me.clusterpartitions) == nil {
        me.clusterpartitions = clusterpartitions.NewCLUSTERPARTITIONS(me.config)
    }
    return me.clusterpartitions
}
/**
     * Access to Groups controller
     * @return Returns the Groups() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Groups() groups.GROUPS {
    if(me.groups) == nil {
        me.groups = groups.NewGROUPS(me.config)
    }
    return me.groups
}
/**
     * Access to Idps controller
     * @return Returns the Idps() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Idps() idps.IDPS {
    if(me.idps) == nil {
        me.idps = idps.NewIDPS(me.config)
    }
    return me.idps
}
/**
     * Access to InterfaceGroup controller
     * @return Returns the InterfaceGroup() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) InterfaceGroup() interfacegroup.INTERFACEGROUP {
    if(me.interfacegroup) == nil {
        me.interfacegroup = interfacegroup.NewINTERFACEGROUP(me.config)
    }
    return me.interfacegroup
}
/**
     * Access to KmsConfiguration controller
     * @return Returns the KmsConfiguration() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) KmsConfiguration() kmsconfiguration.KMSCONFIGURATION {
    if(me.kmsconfiguration) == nil {
        me.kmsconfiguration = kmsconfiguration.NewKMSCONFIGURATION(me.config)
    }
    return me.kmsconfiguration
}
/**
     * Access to LdapProvider controller
     * @return Returns the LdapProvider() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) LdapProvider() ldapprovider.LDAPPROVIDER {
    if(me.ldapprovider) == nil {
        me.ldapprovider = ldapprovider.NewLDAPPROVIDER(me.config)
    }
    return me.ldapprovider
}
/**
     * Access to Network controller
     * @return Returns the Network() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Network() network.NETWORK {
    if(me.network) == nil {
        me.network = network.NewNETWORK(me.config)
    }
    return me.network
}
/**
     * Access to Views controller
     * @return Returns the Views() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Views() views.VIEWS {
    if(me.views) == nil {
        me.views = views.NewVIEWS(me.config)
    }
    return me.views
}
/**
     * Access to Nodes controller
     * @return Returns the Nodes() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Nodes() nodes.NODES {
    if(me.nodes) == nil {
        me.nodes = nodes.NewNODES(me.config)
    }
    return me.nodes
}
/**
     * Access to Packages controller
     * @return Returns the Packages() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Packages() packages.PACKAGES {
    if(me.packages) == nil {
        me.packages = packages.NewPACKAGES(me.config)
    }
    return me.packages
}
/**
     * Access to ProtectionSources controller
     * @return Returns the ProtectionSources() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) ProtectionSources() protectionsources.PROTECTIONSOURCES {
    if(me.protectionsources) == nil {
        me.protectionsources = protectionsources.NewPROTECTIONSOURCES(me.config)
    }
    return me.protectionsources
}
/**
     * Access to CustomReporting controller
     * @return Returns the CustomReporting() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) CustomReporting() customreporting.CUSTOMREPORTING {
    if(me.customreporting) == nil {
        me.customreporting = customreporting.NewCUSTOMREPORTING(me.config)
    }
    return me.customreporting
}
/**
     * Access to Principals controller
     * @return Returns the Principals() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Principals() principals.PRINCIPALS {
    if(me.principals) == nil {
        me.principals = principals.NewPRINCIPALS(me.config)
    }
    return me.principals
}
/**
     * Access to Privileges controller
     * @return Returns the Privileges() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Privileges() privileges.PRIVILEGES {
    if(me.privileges) == nil {
        me.privileges = privileges.NewPRIVILEGES(me.config)
    }
    return me.privileges
}
/**
     * Access to ProtectionJobs controller
     * @return Returns the ProtectionJobs() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) ProtectionJobs() protectionjobs.PROTECTIONJOBS {
    if(me.protectionjobs) == nil {
        me.protectionjobs = protectionjobs.NewPROTECTIONJOBS(me.config)
    }
    return me.protectionjobs
}
/**
     * Access to ProtectionObjects controller
     * @return Returns the ProtectionObjects() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) ProtectionObjects() protectionobjects.PROTECTIONOBJECTS {
    if(me.protectionobjects) == nil {
        me.protectionobjects = protectionobjects.NewPROTECTIONOBJECTS(me.config)
    }
    return me.protectionobjects
}
/**
     * Access to ProtectionPolicies controller
     * @return Returns the ProtectionPolicies() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) ProtectionPolicies() protectionpolicies.PROTECTIONPOLICIES {
    if(me.protectionpolicies) == nil {
        me.protectionpolicies = protectionpolicies.NewPROTECTIONPOLICIES(me.config)
    }
    return me.protectionpolicies
}
/**
     * Access to ProtectionRuns controller
     * @return Returns the ProtectionRuns() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) ProtectionRuns() protectionruns.PROTECTIONRUNS {
    if(me.protectionruns) == nil {
        me.protectionruns = protectionruns.NewPROTECTIONRUNS(me.config)
    }
    return me.protectionruns
}
/**
     * Access to RemoteCluster controller
     * @return Returns the RemoteCluster() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) RemoteCluster() remotecluster.REMOTECLUSTER {
    if(me.remotecluster) == nil {
        me.remotecluster = remotecluster.NewREMOTECLUSTER(me.config)
    }
    return me.remotecluster
}
/**
     * Access to RemoteRestore controller
     * @return Returns the RemoteRestore() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) RemoteRestore() remoterestore.REMOTERESTORE {
    if(me.remoterestore) == nil {
        me.remoterestore = remoterestore.NewREMOTERESTORE(me.config)
    }
    return me.remoterestore
}
/**
     * Access to RestoreTasks controller
     * @return Returns the RestoreTasks() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) RestoreTasks() restoretasks.RESTORETASKS {
    if(me.restoretasks) == nil {
        me.restoretasks = restoretasks.NewRESTORETASKS(me.config)
    }
    return me.restoretasks
}
/**
     * Access to Roles controller
     * @return Returns the Roles() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Roles() roles.ROLES {
    if(me.roles) == nil {
        me.roles = roles.NewROLES(me.config)
    }
    return me.roles
}
/**
     * Access to Routes controller
     * @return Returns the Routes() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Routes() routes.ROUTES {
    if(me.routes) == nil {
        me.routes = routes.NewROUTES(me.config)
    }
    return me.routes
}
/**
     * Access to Search controller
     * @return Returns the Search() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Search() search.SEARCH {
    if(me.search) == nil {
        me.search = search.NewSEARCH(me.config)
    }
    return me.search
}
/**
     * Access to Notifications controller
     * @return Returns the Notifications() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Notifications() notifications.NOTIFICATIONS {
    if(me.notifications) == nil {
        me.notifications = notifications.NewNOTIFICATIONS(me.config)
    }
    return me.notifications
}
/**
     * Access to SMBFileOpens controller
     * @return Returns the SMBFileOpens() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) SMBFileOpens() smbfileopens.SMBFILEOPENS {
    if(me.smbfileopens) == nil {
        me.smbfileopens = smbfileopens.NewSMBFILEOPENS(me.config)
    }
    return me.smbfileopens
}
/**
     * Access to StaticRoute controller
     * @return Returns the StaticRoute() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) StaticRoute() staticroute.STATICROUTE {
    if(me.staticroute) == nil {
        me.staticroute = staticroute.NewSTATICROUTE(me.config)
    }
    return me.staticroute
}
/**
     * Access to Statistics controller
     * @return Returns the Statistics() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Statistics() statistics.STATISTICS {
    if(me.statistics) == nil {
        me.statistics = statistics.NewSTATISTICS(me.config)
    }
    return me.statistics
}
/**
     * Access to Tenant controller
     * @return Returns the Tenant() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Tenant() tenant.TENANT {
    if(me.tenant) == nil {
        me.tenant = tenant.NewTENANT(me.config)
    }
    return me.tenant
}
/**
     * Access to Tenants controller
     * @return Returns the Tenants() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Tenants() tenants.TENANTS {
    if(me.tenants) == nil {
        me.tenants = tenants.NewTENANTS(me.config)
    }
    return me.tenants
}
/**
     * Access to Vaults controller
     * @return Returns the Vaults() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Vaults() vaults.VAULTS {
    if(me.vaults) == nil {
        me.vaults = vaults.NewVAULTS(me.config)
    }
    return me.vaults
}
/**
     * Access to ViewBoxes controller
     * @return Returns the ViewBoxes() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) ViewBoxes() viewboxes.VIEWBOXES {
    if(me.viewboxes) == nil {
        me.viewboxes = viewboxes.NewVIEWBOXES(me.config)
    }
    return me.viewboxes
}
/**
     * Access to Vlan controller
     * @return Returns the Vlan() instance
*/
func (me *COHESITYMANAGEMENTSDK_IMPL) Vlan() vlan.VLAN {
    if(me.vlan) == nil {
        me.vlan = vlan.NewVLAN(me.config)
    }
    return me.vlan
}

func (me *COHESITYMANAGEMENTSDK_IMPL) Authorize() (*models.AccessToken,error) {
    accessTokenCredential := &models.AccessTokenCredential{}
    accessTokenCredential.Username = me.config.Username();
    accessTokenCredential.Password = me.config.Password();

    domain := me.config.Domain();

    if domain!="" {
        accessTokenCredential.Domain =&domain;
    }

    var err error
    var result *models.AccessToken

    result, err = me.AccessTokens().CreateGenerateAccessToken(accessTokenCredential)

    if err != nil{
        return nil,err;
    }else{
        me.config.SetAccessToken(result)
        return result,nil;
    }
}
