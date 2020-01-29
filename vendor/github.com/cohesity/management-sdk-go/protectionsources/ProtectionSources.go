// Copyright 2019 Cohesity Inc.
package protectionsources

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the PROTECTIONSOURCES_IMPL
 */
type PROTECTIONSOURCES interface {
    GetDownloadPhysicalAgent (models.HostTypeDownloadPhysicalAgentEnum, models.PkgTypeEnum, models.AgentTypeEnum) ([]int64, error)

    CreateUpgradePhysicalAgents (*models.UpgradePhysicalServerAgents) (*models.UpgradePhysicalAgentsMessage, error)

    ListProtectionSources (*int64, []models.ExcludeTypeEnum, []models.ExcludeOffice365TypeEnum, *bool, *bool, *bool, []models.EnvironmentListProtectionSourcesEnum, *string, *bool, []string, []string, *bool) ([]*models.ProtectionSourceNode, error)

    ListApplicationServers (*int64, models.EnvironmentListApplicationServersEnum, *int64, models.ApplicationEnum) ([]*models.RegisteredApplicationServer, error)

    CreateRegisterApplicationServers (*models.RegisterApplicationServersParameters) (*models.ProtectionSource, error)

    UpdateApplicationServers (*models.UpdateApplicationServerParameters) (*models.ProtectionSource, error)

    DeleteUnregisterApplicationServers (*models.UnRegisterApplicationServersParameters, int64) (*models.ProtectionSource, error)

    ListDataStoreInformation (int64) ([]*models.ProtectionSource, error)

    GetProtectionSourcesObjects ([]int64) ([]*models.ProtectionSource, error)

    GetProtectionSourcesObjectById (int64) (*models.ProtectionSource, error)

    ListProtectedObjects (models.EnvironmentListProtectedObjectsEnum, int64, *bool, *bool) ([]*models.ProtectedVmInfo, error)

    CreateRefreshProtectionSourceById (int64) (error)

    CreateRegisterProtectionSource (*models.RegisterProtectionSourceParameters) (*models.ProtectionSource, error)

    ListProtectionSourcesRegistrationInfo ([]models.EnvironmentListProtectionSourcesRegistrationInfoEnum, []int64, *bool, []string, []string, *bool) (*models.GetRegistrationInfoResponse, error)

    ListProtectionSourcesRootNodes (*int64, []models.EnvironmentListProtectionSourcesRootNodesEnum, *string) ([]*models.ProtectionSourceNode, error)

    ListSqlAagHostsAndDatabases ([]int64) ([]*models.SqlAagHostAndDatabases, error)

    ListVirtualMachines (*int64, []string, []string, *bool) ([]*models.ProtectionSource, error)

    DeleteUnregisterProtectionSource (int64) (error)

    UpdateProtectionSource (int64, *models.UpdateProtectionSourceParameters) (*models.ProtectionSourceNode, error)
}

/*
 * Factory for the PROTECTIONSOURCES interaface returning PROTECTIONSOURCES_IMPL
 */
func NewPROTECTIONSOURCES(config configuration.CONFIGURATION) *PROTECTIONSOURCES_IMPL {
    client := new(PROTECTIONSOURCES_IMPL)
    client.config = config
    return client
}