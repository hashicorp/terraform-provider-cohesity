// Copyright 2019 Cohesity Inc.
package audit

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the AUDIT_IMPL
 */
type AUDIT interface {
    GetAuditLogsActions () ([]string, error)

    GetAuditLogsCategories () ([]string, error)

    SearchClusterAuditLogs ([]string, []string, []string, []string, *string, *int64, *int64, *int64, *int64, *string, *string, *bool) (*models.ClusterAuditLogsSearchResult, error)
}

/*
 * Factory for the AUDIT interaface returning AUDIT_IMPL
 */
func NewAUDIT(config configuration.CONFIGURATION) *AUDIT_IMPL {
    client := new(AUDIT_IMPL)
    client.config = config
    return client
}