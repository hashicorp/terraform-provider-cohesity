// Copyright 2019 Cohesity Inc.
package cluster

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the CLUSTER_IMPL
 */
type CLUSTER interface {
    GetBasicClusterInfo () (*models.BasicClusterInfo, error)

    GetCluster (*bool, *bool) (*models.Cluster, error)

    UpdateCluster (*models.UpdateClusterParams) (*models.Cluster, error)
}

/*
 * Factory for the CLUSTER interaface returning CLUSTER_IMPL
 */
func NewCLUSTER(config configuration.CONFIGURATION) *CLUSTER_IMPL {
    client := new(CLUSTER_IMPL)
    client.config = config
    return client
}