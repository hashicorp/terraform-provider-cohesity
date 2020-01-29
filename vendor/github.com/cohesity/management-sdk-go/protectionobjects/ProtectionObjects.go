// Copyright 2019 Cohesity Inc.
package protectionobjects

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the PROTECTIONOBJECTS_IMPL
 */
type PROTECTIONOBJECTS interface {
    DeleteUnprotectObject (*models.UnprotectObjectParams) (error)

    CreateProtectObject (*models.ProtectObjectParameters) ([]*models.ProtectedObject, error)

    UpdateProtectionObject (*models.UpdateProtectionObjectParameters) (*models.ProtectionJob, error)

    GetProtectionObjectSummary (int64) (*models.ProtectionObjectSummary, error)
}

/*
 * Factory for the PROTECTIONOBJECTS interaface returning PROTECTIONOBJECTS_IMPL
 */
func NewPROTECTIONOBJECTS(config configuration.CONFIGURATION) *PROTECTIONOBJECTS_IMPL {
    client := new(PROTECTIONOBJECTS_IMPL)
    client.config = config
    return client
}