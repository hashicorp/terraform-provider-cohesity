// Copyright 2019 Cohesity Inc.
package protectionjobs

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the PROTECTIONJOBS_IMPL
 */
type PROTECTIONJOBS interface {
    ChangeProtectionJobState (int64, *models.ChangeProtectionJobStateParam) (error)

    GetProtectionJobs ([]int64, []string, []string, []models.EnvironmentGetProtectionJobsEnum, *bool, *bool, *bool, *bool, *bool, *bool, *bool, []string, *bool) ([]*models.ProtectionJob, error)

    CreateProtectionJob (*models.ProtectionJobRequestBody) (*models.ProtectionJob, error)

    CreateRunProtectionJob (int64, *models.RunProtectionJobParam) (error)

    UpdateProtectionJobsState (*models.UpdateProtectionJobsStateParams) (*models.UpdateProtectionJobsState, error)

    DeleteProtectionJob (int64, *models.DeleteProtectionJobParam) (error)

    GetProtectionJobById (int64) (*models.ProtectionJob, error)

    UpdateProtectionJob (*models.ProtectionJobRequestBody, int64) (*models.ProtectionJob, error)

    GetProtectionJobAudit (int64) ([]*models.ProtectionJobAuditTrail, error)
}

/*
 * Factory for the PROTECTIONJOBS interaface returning PROTECTIONJOBS_IMPL
 */
func NewPROTECTIONJOBS(config configuration.CONFIGURATION) *PROTECTIONJOBS_IMPL {
    client := new(PROTECTIONJOBS_IMPL)
    client.config = config
    return client
}