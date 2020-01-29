// Copyright 2019 Cohesity Inc.
package statistics

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the STATISTICS_IMPL
 */
type STATISTICS interface {
    GetEntities (string, *bool, []string) ([]*models.EntityProto, error)

    GetEntitiesSchema ([]string, []string) ([]*models.EntitySchemaProto, error)

    GetEntitySchemaByName (string) ([]*models.EntitySchemaProto, error)

    GetTimeSeriesSchema (models.EntityTypeEnum, int64, string) (*models.TimeSeriesSchemaResponse, error)

    GetTimeSeriesStats (string, string, int64, *string, []string, *int64, *string, *int64) (*models.MetricDataBlock, error)

    GetTasks ([]string, *bool, *int64, *int64, *int64, *bool, []string) ([]*models.Task, error)
}

/*
 * Factory for the STATISTICS interaface returning STATISTICS_IMPL
 */
func NewSTATISTICS(config configuration.CONFIGURATION) *STATISTICS_IMPL {
    client := new(STATISTICS_IMPL)
    client.config = config
    return client
}