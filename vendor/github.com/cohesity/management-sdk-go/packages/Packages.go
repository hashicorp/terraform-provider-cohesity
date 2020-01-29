// Copyright 2019 Cohesity Inc.
package packages

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the PACKAGES_IMPL
 */
type PACKAGES interface {
    ListPackages () ([]*models.PackageDetails, error)

    CreateDownloadPackage (*models.DownloadPackageParameters) (*models.DownloadPackageResult, error)
}

/*
 * Factory for the PACKAGES interaface returning PACKAGES_IMPL
 */
func NewPACKAGES(config configuration.CONFIGURATION) *PACKAGES_IMPL {
    client := new(PACKAGES_IMPL)
    client.config = config
    return client
}