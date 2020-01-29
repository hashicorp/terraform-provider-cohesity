// Copyright 2019 Cohesity Inc.
package remotecluster

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the REMOTECLUSTER_IMPL
 */
type REMOTECLUSTER interface {
    GetRemoteClusters ([]int64, []string, *bool, *bool) ([]*models.RemoteCluster, error)

    CreateRemoteCluster (*models.RegisterRemoteCluster) (*models.RemoteCluster, error)

    DeleteRemoteCluster (int64) (error)

    GetRemoteClusterById (int64) ([]*models.RemoteCluster, error)

    UpdateRemoteCluster (int64, *models.RegisterRemoteCluster) (*models.RemoteCluster, error)

    GetReplicationEncryptionKey () (*models.ReplicationEncryptionKeyReponse, error)
}

/*
 * Factory for the REMOTECLUSTER interaface returning REMOTECLUSTER_IMPL
 */
func NewREMOTECLUSTER(config configuration.CONFIGURATION) *REMOTECLUSTER_IMPL {
    client := new(REMOTECLUSTER_IMPL)
    client.config = config
    return client
}