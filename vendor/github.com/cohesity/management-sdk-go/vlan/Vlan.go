// Copyright 2019 Cohesity Inc.
package vlan

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the VLAN_IMPL
 */
type VLAN interface {
    GetVlans ([]string, *bool, *bool) ([]*models.Vlan, error)

    CreateVlan (*models.Vlan) (*models.Vlan, error)

    RemoveVlan (int64) (error)

    GetVlanById (int64) (*models.Vlan, error)

    UpdateVlan (int64, *models.Vlan) (*models.Vlan, error)
}

/*
 * Factory for the VLAN interaface returning VLAN_IMPL
 */
func NewVLAN(config configuration.CONFIGURATION) *VLAN_IMPL {
    client := new(VLAN_IMPL)
    client.config = config
    return client
}