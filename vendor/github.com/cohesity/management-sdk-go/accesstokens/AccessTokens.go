// Copyright 2019 Cohesity Inc.
package accesstokens

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the ACCESSTOKENS_IMPL
 */
type ACCESSTOKENS interface {
    CreateGenerateAccessToken (*models.AccessTokenCredential) (*models.AccessToken, error)
}

/*
 * Factory for the ACCESSTOKENS interaface returning ACCESSTOKENS_IMPL
 */
func NewACCESSTOKENS(config configuration.CONFIGURATION) *ACCESSTOKENS_IMPL {
    client := new(ACCESSTOKENS_IMPL)
    client.config = config
    return client
}