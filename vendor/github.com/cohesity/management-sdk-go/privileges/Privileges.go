// Copyright 2019 Cohesity Inc.
package privileges

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the PRIVILEGES_IMPL
 */
type PRIVILEGES interface {
    GetPrivileges (*string) ([]*models.PrivilegeInfo, error)
}

/*
 * Factory for the PRIVILEGES interaface returning PRIVILEGES_IMPL
 */
func NewPRIVILEGES(config configuration.CONFIGURATION) *PRIVILEGES_IMPL {
    client := new(PRIVILEGES_IMPL)
    client.config = config
    return client
}