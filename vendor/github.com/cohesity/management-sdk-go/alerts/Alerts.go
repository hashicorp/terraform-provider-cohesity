// Copyright 2019 Cohesity Inc.
package alerts

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the ALERTS_IMPL
 */
type ALERTS interface {
    GetAlertCategories () ([]*models.AlertCategoryName, error)

    GetNotificationRules () ([]*models.NotificationRule, error)

    CreateNotificationRule (*models.NotificationRule) (*models.NotificationRule, error)

    UpdateNotificationRule () (*models.NotificationRule, error)

    DeleteNotificationRule (int64) (error)

    GetResolutions (int64, []int64, []string, *int64, *int64, []string, *bool) ([]*models.AlertResolution, error)

    CreateResolution (*models.AlertResolutionRequest) (*models.AlertResolution, error)

    GetResolutionById (int64) (*models.AlertResolution, error)

    UpdateResolution (int64, *models.UpdateResolutionParams) (*models.AlertResolution, error)

    GetAlertTypes () ([]*models.AlertMetadata, error)

    GetAlerts (int64, []string, []int64, []models.AlertCategoryListGetAlertsEnum, *string, *string, *int64, *int64, []models.AlertStateListEnum, []models.AlertSeverityListEnum, []models.AlertTypeBucketListEnum, []int64, []string, *bool) ([]*models.Alert, error)

    GetAlertById (string) (*models.Alert, error)
}

/*
 * Factory for the ALERTS interaface returning ALERTS_IMPL
 */
func NewALERTS(config configuration.CONFIGURATION) *ALERTS_IMPL {
    client := new(ALERTS_IMPL)
    client.config = config
    return client
}