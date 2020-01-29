// Copyright 2019 Cohesity Inc.
package search

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the SEARCH_IMPL
 */
type SEARCH interface {
    SearchProtectionRuns (string) (*models.ProtectionRunResponse, error)

    SearchProtectionSources (*string, []string, []models.EnvironmentSearchProtectionSourcesEnum, []int64, []models.PhysicalServerHostTypeEnum, []string, *int64, *int64) ([]*models.ProtectionSourceResponse, error)
}

/*
 * Factory for the SEARCH interaface returning SEARCH_IMPL
 */
func NewSEARCH(config configuration.CONFIGURATION) *SEARCH_IMPL {
    client := new(SEARCH_IMPL)
    client.config = config
    return client
}