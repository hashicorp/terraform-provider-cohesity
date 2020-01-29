// Copyright 2019 Cohesity Inc.
package activedirectory

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the ACTIVEDIRECTORY_IMPL
 */
type ACTIVEDIRECTORY interface {
    DeleteActiveDirectoryEntry (*models.ActiveDirectoryEntry) (error)

    GetActiveDirectoryEntry ([]string, []string, *bool) ([]*models.ActiveDirectoryEntry, error)

    CreateActiveDirectoryEntry (*models.CreateActiveDirectoryEntryParams) (*models.ActiveDirectoryEntry, error)

    ListCentrifyZones (*string) ([]*models.ListCentrifyZone, error)

    GetActiveDirectoryDomainControllers () (*models.DomainControllers, error)

    SearchActiveDirectoryPrincipals (*string, models.ObjectClassSearchActiveDirectoryPrincipalsEnum, *string, []string, *bool) ([]*models.ActiveDirectoryPrincipal, error)

    AddActiveDirectoryPrincipals ([]*models.ActiveDirectoryPrincipalsAddParameters) ([]*models.AddedActiveDirectoryPrincipal, error)

    CreateEnableTrustedDomainDiscovery (bool, string) (*models.ActiveDirectoryEntry, error)

    UpdateActiveDirectoryIdMapping (*models.IdMappingInfo, string) (*models.ActiveDirectoryEntry, error)

    UpdateActiveDirectoryIgnoredTrustedDomains (*models.UpdateIgnoredTrustedDomainsParams, string) (*models.ActiveDirectoryEntry, error)

    UpdateActiveDirectoryLdapProvider (*models.UpdateLdapProviderParams, string) (*models.ActiveDirectoryEntry, error)

    UpdateActiveDirectoryMachineAccounts (*models.UpdateMachineAccountsParams, string) (*models.ActiveDirectoryEntry, error)

    UpdatePreferredDomainControllers ([]*models.PreferredDomainController, string) (*models.ActiveDirectoryEntry, error)
}

/*
 * Factory for the ACTIVEDIRECTORY interaface returning ACTIVEDIRECTORY_IMPL
 */
func NewACTIVEDIRECTORY(config configuration.CONFIGURATION) *ACTIVEDIRECTORY_IMPL {
    client := new(ACTIVEDIRECTORY_IMPL)
    client.config = config
    return client
}