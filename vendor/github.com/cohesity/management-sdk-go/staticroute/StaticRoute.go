// Copyright 2019 Cohesity Inc.
package staticroute

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the STATICROUTE_IMPL
 */
type STATICROUTE interface {
    GetStaticRoutes () ([]*models.StaticRoute, error)

    RemoveStaticRoute (string) (error)

    UpdateStaticRoute (string, *models.StaticRoute) (*models.StaticRoute, error)
}

/*
 * Factory for the STATICROUTE interaface returning STATICROUTE_IMPL
 */
func NewSTATICROUTE(config configuration.CONFIGURATION) *STATICROUTE_IMPL {
    client := new(STATICROUTE_IMPL)
    client.config = config
    return client
}