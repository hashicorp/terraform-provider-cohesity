// Copyright 2019 Cohesity Inc.
package nodes

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the NODES_IMPL
 */
type NODES interface {
    ListFreeNodes () ([]*models.FreeNodeInformation, error)

    GetNodes () ([]*models.Node, error)

    UpdateUpgradeNode (*models.UpgradeNodeParameters) (*models.UpgradeNodeResult, error)

    GetNodeById (int64) ([]*models.Node, error)
}

/*
 * Factory for the NODES interaface returning NODES_IMPL
 */
func NewNODES(config configuration.CONFIGURATION) *NODES_IMPL {
    client := new(NODES_IMPL)
    client.config = config
    return client
}