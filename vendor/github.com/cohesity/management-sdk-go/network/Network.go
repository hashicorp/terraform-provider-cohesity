// Copyright 2019 Cohesity Inc.
package network

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the NETWORK_IMPL
 */
type NETWORK interface {
    CreateBond (*models.CreateBondParameters) (*models.CreateBondResult, error)

    UpdateBond (*models.UpdateBondParameters) (*models.UpdateBondResult, error)

    DeleteBond (string) (error)

    DeleteHosts ([]string) (*models.HostResult, error)

    ListHosts () ([]*models.HostEntry, error)

    CreateAppendHosts (*models.AppendHostsParameters) (*models.HostResult, error)

    UpdateEditHosts (*models.EditHostsParameters) (*models.HostResult, error)

    ListNetworkInterfaces () ([]*models.NodeNetworkInterfaces, error)
}

/*
 * Factory for the NETWORK interaface returning NETWORK_IMPL
 */
func NewNETWORK(config configuration.CONFIGURATION) *NETWORK_IMPL {
    client := new(NETWORK_IMPL)
    client.config = config
    return client
}