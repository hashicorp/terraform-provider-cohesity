// Copyright 2019 Cohesity Inc.
package viewboxes

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the VIEWBOXES_IMPL
 */
type VIEWBOXES interface {
    GetViewBoxes ([]string, *bool, []int64, []string, []int64, *bool, *bool) ([]*models.ViewBox, error)

    CreateViewBox (*models.CreateViewBoxParams) (*models.ViewBox, error)

    GetViewBoxById (int64, *bool) (*models.ViewBox, error)

    UpdateViewBox (int64, *models.CreateViewBoxParams) (*models.ViewBox, error)
}

/*
 * Factory for the VIEWBOXES interaface returning VIEWBOXES_IMPL
 */
func NewVIEWBOXES(config configuration.CONFIGURATION) *VIEWBOXES_IMPL {
    client := new(VIEWBOXES_IMPL)
    client.config = config
    return client
}