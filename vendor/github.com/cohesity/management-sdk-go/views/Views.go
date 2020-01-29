// Copyright 2019 Cohesity Inc.
package views

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the VIEWS_IMPL
 */
type VIEWS interface {
    DeleteClearNlmLocks (*models.ClearNlmLocksParameters) (error)

    ListNlmLocks (*string, *string, *int64, *string) (*models.ListNlmLocksResponse, error)

    GetQoSPolicies ([]int64, []string) ([]*models.QoSPolicy, error)

    GetViewsByShareName ([]string, *bool, *string, *int64, *string) (*models.GetViewsByShareNameResult, error)

    GetSmbConnections ([]string, []int64, *int64, *bool) ([]*models.SmbConnection, error)

    CreateViewAlias (*models.ViewAlias) (*models.ViewAlias, error)

    UpdateViewAlias (*models.UpdateViewAliasParam) (*models.ViewAlias, error)

    DeleteViewAlias (string) (error)

    CreateActivateViewAliases (string) (*models.ActivateViewAliasesResult, error)

    GetViewDirQuotaInfo (string) (*models.DirQuotaInfo, error)

    UpdateViewDirQuota (*models.UpdateDirQuotaArgs) (*models.DirQuotaInfo, error)

    DeleteViewUsersQuota (*models.DeleteViewUsersQuotaParameters) (error)

    GetViewUserQuotas (*string, *bool, *bool, *int64, *string, []int64, []string, []string, *bool, *int64, *int64, *string, *string) (*models.ViewUserQuotas, error)

    CreateViewUserQuota (*models.ViewUserQuotaParameters) (*models.UserQuotaAndUsage, error)

    UpdateViewUserQuota (*models.ViewUserQuotaParameters) (*models.UserQuotaAndUsage, error)

    UpdateUserQuotaSettings (*models.UpdateUserQuotaSettingsForView) (*models.UserQuotaSettings, error)

    GetViews ([]string, *bool, []string, []int64, []int64, []string, *bool, *int64, *bool, *int64, *bool, []int64, *bool, *bool, *bool, *bool) (*models.GetViewsResult, error)

    CreateView (*models.CreateViewRequest) (*models.View, error)

    UpdateView (*models.UpdateViewParam) (*models.View, error)

    CreateCloneView (*models.CloneViewRequest) (*models.View, error)

    CreateCloneDirectory (*models.CloneDirectoryParams) (error)

    DeleteViewById (int64) (error)

    GetViewById (int64) (*models.View, error)

    GetFileLockStatusById (int64, *string) (*models.FileLockStatus, error)

    CreateLockFileById (int64, *models.LockFileParams) (*models.FileLockStatus, error)

    CreateOverwriteView (*models.OverwriteViewParam) (*models.View, error)

    CreateRenameViewById (*models.RenameViewParam, int64) (*models.View, error)

    CreateRenameView (*models.RenameViewParam, string) (*models.View, error)

    DeleteView (string) (error)

    GetViewByName (string) (*models.View, error)

    UpdateViewByName (string, *models.UpdateViewParam) (*models.View, error)

    GetFileLockStatus (string, *string) (*models.FileLockStatus, error)

    CreateLockFile (string, *models.LockFileParams) (*models.FileLockStatus, error)
}

/*
 * Factory for the VIEWS interaface returning VIEWS_IMPL
 */
func NewVIEWS(config configuration.CONFIGURATION) *VIEWS_IMPL {
    client := new(VIEWS_IMPL)
    client.config = config
    return client
}