// Copyright 2019 Cohesity Inc.
package idps

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the IDPS_IMPL
 */
type IDPS interface {
    AddActiveIdpPrincipals () ([]*models.AddedIdpPrincipal, error)

    GetIdps ([]string, []int64, []string, []string) ([]*models.IdpServiceConfiguration, error)

    CreateIdp (*models.CreateIdpConfigurationRequest) (*models.IdpServiceConfiguration, error)

    GetIdpLogin (*string) (error)

    DeleteIdp (int64) (error)

    UpdateIdp (int64, *models.UpdateIdpConfigurationRequest) (*models.IdpServiceConfiguration, error)
}

/*
 * Factory for the IDPS interaface returning IDPS_IMPL
 */
func NewIDPS(config configuration.CONFIGURATION) *IDPS_IMPL {
    client := new(IDPS_IMPL)
    client.config = config
    return client
}