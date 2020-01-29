// Copyright 2019 Cohesity Inc.
package tenant

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the TENANT_IMPL
 */
type TENANT interface {
    DeleteTenant (*string) ([]*models.Tenant, error)

    GetTenants ([]string, []models.PropertyEnum, *bool, *bool, []models.StatusGetTenantsEnum) ([]*models.Tenant, error)

    CreateTenant (*models.TenantCreateParameters) (*models.Tenant, error)

    UpdateTenant (*models.TenantUpdate) (*models.Tenant, error)

    UpdateTenantActiveDirectory (*models.TenantActiveDirectoryUpdateParameters) (*models.TenantActiveDirectoryUpdate, error)

    UpdateTenantEntity (*models.TenantEntityUpdateParameters) (*models.TenantEntityUpdate, error)

    UpdateTenantGroups () ([]*models.Group, error)

    UpdateTenantLdapProvider (*models.TenantLdapProviderUpdateParameters) (*models.TenantLdapProviderUpdate, error)

    UpdateTenantProtectionPolicy (*models.TenantProtectionPolicyUpdateParameters) (*models.TenantProtectionPolicyUpdate, error)

    UpdateTenantProtectionJob (*models.TenantProtectionJobUpdateParameters) (*models.TenantProtectionJobUpdate, error)

    GetTenantsProxies ([]string) ([]*models.TenantProxy, error)

    GetTenantsProxyConfigRequest () ([]int64, error)

    GetDownloadTenantsProxy (*string) ([]int64, error)

    UpdateTenantUsers (*models.TenantUserUpdateParameters) ([]*models.User, error)

    UpdateTenantView (*models.TenantViewUpdateParameters) (*models.TenantViewUpdate, error)

    UpdateTenantViewBox (*models.TenantViewBoxUpdateParameters) (*models.TenantViewBoxUpdate, error)

    UpdateTenantVlan (*models.TenantVlanUpdateParameters) (*models.TenantVlanUpdate, error)
}

/*
 * Factory for the TENANT interaface returning TENANT_IMPL
 */
func NewTENANT(config configuration.CONFIGURATION) *TENANT_IMPL {
    client := new(TENANT_IMPL)
    client.config = config
    return client
}