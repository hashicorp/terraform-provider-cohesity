// Copyright 2019 Cohesity Inc.
package customreporting

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the CUSTOMREPORTING_IMPL
 */
type CUSTOMREPORTING interface {
    GetPostgres () ([]*models.PostgresNodeInfo, error)
}

/*
 * Factory for the CUSTOMREPORTING interaface returning CUSTOMREPORTING_IMPL
 */
func NewCUSTOMREPORTING(config configuration.CONFIGURATION) *CUSTOMREPORTING_IMPL {
    client := new(CUSTOMREPORTING_IMPL)
    client.config = config
    return client
}