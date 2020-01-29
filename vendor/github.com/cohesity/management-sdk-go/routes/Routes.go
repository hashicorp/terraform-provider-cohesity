// Copyright 2019 Cohesity Inc.
package routes

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the ROUTES_IMPL
 */
type ROUTES interface {
    DeleteRoute (*models.DeleteRouteParam) (error)

    GetRoutes () ([]*models.Route, error)

    AddRoute (*models.Route) (*models.Route, error)
}

/*
 * Factory for the ROUTES interaface returning ROUTES_IMPL
 */
func NewROUTES(config configuration.CONFIGURATION) *ROUTES_IMPL {
    client := new(ROUTES_IMPL)
    client.config = config
    return client
}