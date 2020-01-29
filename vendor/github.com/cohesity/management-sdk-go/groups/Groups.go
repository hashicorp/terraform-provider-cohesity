// Copyright 2019 Cohesity Inc.
package groups

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the GROUPS_IMPL
 */
type GROUPS interface {
    DeleteGroups (*models.GroupDeleteParameters) (error)

    GetGroups (*string, *string, []string, *bool) ([]*models.Group, error)

    CreateGroup (*models.GroupParameters) (*models.Group, error)

    UpdateGroup (*models.GroupParameters) (*models.Group, error)
}

/*
 * Factory for the GROUPS interaface returning GROUPS_IMPL
 */
func NewGROUPS(config configuration.CONFIGURATION) *GROUPS_IMPL {
    client := new(GROUPS_IMPL)
    client.config = config
    return client
}