// Copyright 2019 Cohesity Inc.
package principals

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the PRINCIPALS_IMPL
 */
type PRINCIPALS interface {
    ListSourcesForPrincipals ([]string) ([]*models.SourcesForSid, error)

    UpdateSourcesForPrincipals (*models.UpdateSourcesForPrincipalsParams) (error)

    SearchPrincipals (*string, models.ObjectClassSearchPrincipalsEnum, *string, []string, *bool) ([]*models.Principal, error)

    GetSessionUser () (*models.User, error)

    DeleteUsers (*models.UserDeleteParameters) (error)

    GetUsers ([]string, *bool, []string, []string, *string, *bool) ([]*models.User, error)

    CreateUser (*models.UserParameters) (*models.User, error)

    UpdateUser (*models.UserParameters) (*models.User, error)

    GetUserPrivileges () ([]string, error)

    CreateResetS3SecretKey (*models.ResetS3SecretKeyParameters) (*models.NewS3SecretAccessKey, error)
}

/*
 * Factory for the PRINCIPALS interaface returning PRINCIPALS_IMPL
 */
func NewPRINCIPALS(config configuration.CONFIGURATION) *PRINCIPALS_IMPL {
    client := new(PRINCIPALS_IMPL)
    client.config = config
    return client
}