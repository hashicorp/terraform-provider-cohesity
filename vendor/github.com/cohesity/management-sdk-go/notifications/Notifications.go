// Copyright 2019 Cohesity Inc.
package notifications

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the NOTIFICATIONS_IMPL
 */
type NOTIFICATIONS interface {
    GetNotifications () (*models.Notifications, error)

    UpdateNotifications () (error)
}

/*
 * Factory for the NOTIFICATIONS interaface returning NOTIFICATIONS_IMPL
 */
func NewNOTIFICATIONS(config configuration.CONFIGURATION) *NOTIFICATIONS_IMPL {
    client := new(NOTIFICATIONS_IMPL)
    client.config = config
    return client
}