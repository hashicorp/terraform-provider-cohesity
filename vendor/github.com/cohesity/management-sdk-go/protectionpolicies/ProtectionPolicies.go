// Copyright 2019 Cohesity Inc.
package protectionpolicies

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the PROTECTIONPOLICIES_IMPL
 */
type PROTECTIONPOLICIES interface {
    GetProtectionPolicies ([]string, []string, []models.EnvironmentGetProtectionPoliciesEnum, []int64, []string, *bool) ([]*models.ProtectionPolicy, error)

    CreateProtectionPolicy (*models.ProtectionPolicyRequest) (*models.ProtectionPolicy, error)

    DeleteProtectionPolicy (string) (error)

    GetProtectionPolicyById (string) (*models.ProtectionPolicy, error)

    UpdateProtectionPolicy (*models.ProtectionPolicyRequest, string) (*models.ProtectionPolicy, error)

    GetProtectionPolicySummary (string, *bool, *bool, *int64, *int64, *int64, *string) (*models.ProtectionPolicySummary, error)
}

/*
 * Factory for the PROTECTIONPOLICIES interaface returning PROTECTIONPOLICIES_IMPL
 */
func NewPROTECTIONPOLICIES(config configuration.CONFIGURATION) *PROTECTIONPOLICIES_IMPL {
    client := new(PROTECTIONPOLICIES_IMPL)
    client.config = config
    return client
}