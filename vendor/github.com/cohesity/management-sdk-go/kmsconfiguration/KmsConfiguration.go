// Copyright 2019 Cohesity Inc.
package kmsconfiguration

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the KMSCONFIGURATION_IMPL
 */
type KMSCONFIGURATION interface {
    GetKmsConfig (*string) ([]*models.KmsConfigurationResponse, error)

    CreateKmsConfig (*models.KmsConfiguration) (*models.KmsConfigurationResponse, error)

    UpdateKmsConfig (*models.KmsConfiguration) (*models.KmsConfigurationResponse, error)
}

/*
 * Factory for the KMSCONFIGURATION interaface returning KMSCONFIGURATION_IMPL
 */
func NewKMSCONFIGURATION(config configuration.CONFIGURATION) *KMSCONFIGURATION_IMPL {
    client := new(KMSCONFIGURATION_IMPL)
    client.config = config
    return client
}