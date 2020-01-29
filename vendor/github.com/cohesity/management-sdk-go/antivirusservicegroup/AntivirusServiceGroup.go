// Copyright 2019 Cohesity Inc.
package antivirusservicegroup

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the ANTIVIRUSSERVICEGROUP_IMPL
 */
type ANTIVIRUSSERVICEGROUP interface {
    GetAntivirusServiceGroup () ([]*models.AntivirusServiceGroup, error)

    CreateAntivirusServiceGroup (*models.AntivirusServiceGroupParams) (*models.AntivirusServiceGroup, error)

    UpdateAntivirusServiceGroup (*models.UpdateAntivirusServiceGroupParams) (*models.AntivirusServiceGroup, error)

    UpdateAntivirusServiceGroupState (*models.AntivirusServiceGroupStateParams) (*models.AntivirusServiceGroupStateParams, error)

    DeleteAntivirusServiceGroup (int64) (error)

    GetIcapConnectionStatus ([]string) (*models.IcapConnectionStatusResponse, error)

    DeleteInfectedFiles (*models.DeleteInfectedFileParams) (*models.DeleteInfectedFileResponse, error)

    GetInfectedFiles ([]string, *bool, *bool, *string, *int64, *string) (*models.InfectedFiles, error)

    UpdateInfectedFiles (*models.UpdateInfectedFileParams) (*models.UpdateInfectedFileResponse, error)
}

/*
 * Factory for the ANTIVIRUSSERVICEGROUP interaface returning ANTIVIRUSSERVICEGROUP_IMPL
 */
func NewANTIVIRUSSERVICEGROUP(config configuration.CONFIGURATION) *ANTIVIRUSSERVICEGROUP_IMPL {
    client := new(ANTIVIRUSSERVICEGROUP_IMPL)
    client.config = config
    return client
}