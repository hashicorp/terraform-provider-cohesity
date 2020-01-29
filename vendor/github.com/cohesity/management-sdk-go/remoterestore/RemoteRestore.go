// Copyright 2019 Cohesity Inc.
package remoterestore

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the REMOTERESTORE_IMPL
 */
type REMOTERESTORE interface {
    UploadVaultEncryptionKeys (int64, []*models.VaultEncryptionKey) (error)

    ListRemoteVaultRestoreTasks () ([]*models.RemoteVaultRestoreTaskStatus, error)

    CreateRemoteVaultRestoreTask (*models.CreateRemoteVaultRestoreTaskParameters) (*models.UniversalId, error)

    GetRemoteVaultSearchJobResults (int64, int64, int64, *int64, *string, *string) (*models.RemoteVaultSearchJobResults, error)

    DeleteStopRemoteVaultSearchJob (*models.StopRemoteVaultSearchJobParameters) (error)

    ListRemoteVaultSearchJobs () ([]*models.RemoteVaultSearchJobInformation, error)

    CreateRemoteVaultSearchJob (*models.CreateRemoteVaultSearchJobParameters) (*models.CreatedRemoteVaultSearchJobUid, error)

    ListRemoteVaultSearchJobById (int64) (*models.RemoteVaultSearchJobInformation, error)
}

/*
 * Factory for the REMOTERESTORE interaface returning REMOTERESTORE_IMPL
 */
func NewREMOTERESTORE(config configuration.CONFIGURATION) *REMOTERESTORE_IMPL {
    client := new(REMOTERESTORE_IMPL)
    client.config = config
    return client
}