// Copyright 2019 Cohesity Inc.
package roles

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the ROLES_IMPL
 */
type ROLES interface {
    DeleteRoles (*models.RoleDeleteParameters) (error)

    GetRoles (*string, []string, *bool) ([]*models.Role, error)

    CreateRole (*models.RoleCreateParameters) (*models.Role, error)

    UpdateRole (string, *models.RoleUpdateParameters) (*models.Role, error)
}

/*
 * Factory for the ROLES interaface returning ROLES_IMPL
 */
func NewROLES(config configuration.CONFIGURATION) *ROLES_IMPL {
    client := new(ROLES_IMPL)
    client.config = config
    return client
}