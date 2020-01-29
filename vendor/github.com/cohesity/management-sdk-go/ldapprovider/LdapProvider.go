// Copyright 2019 Cohesity Inc.
package ldapprovider

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the LDAPPROVIDER_IMPL
 */
type LDAPPROVIDER interface {
    GetLdapProvider ([]int64, []string, *bool) ([]*models.LdapProviderResponse, error)

    CreateLdapProvider (*models.LdapProvider) (*models.LdapProviderResponse, error)

    UpdateLdapProvider (*models.UpdateLdapProviderParam) (*models.LdapProviderResponse, error)

    DeleteLdapProvider (int64) (error)

    GetLdapProviderStatus (int64) (error)
}

/*
 * Factory for the LDAPPROVIDER interaface returning LDAPPROVIDER_IMPL
 */
func NewLDAPPROVIDER(config configuration.CONFIGURATION) *LDAPPROVIDER_IMPL {
    client := new(LDAPPROVIDER_IMPL)
    client.config = config
    return client
}