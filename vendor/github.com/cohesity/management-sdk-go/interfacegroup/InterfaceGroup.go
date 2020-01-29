// Copyright 2019 Cohesity Inc.
package interfacegroup

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the INTERFACEGROUP_IMPL
 */
type INTERFACEGROUP interface {
    GetInterfaceGroups () ([]*models.InterfaceGroup, error)

    CreateInterfaceGroup (*models.InterfaceGroup) (*models.InterfaceGroup, error)

    UpdateInterfaceGroup (*models.InterfaceGroup) (*models.InterfaceGroup, error)

    DeleteInterfaceGroup (string) (error)
}

/*
 * Factory for the INTERFACEGROUP interaface returning INTERFACEGROUP_IMPL
 */
func NewINTERFACEGROUP(config configuration.CONFIGURATION) *INTERFACEGROUP_IMPL {
    client := new(INTERFACEGROUP_IMPL)
    client.config = config
    return client
}