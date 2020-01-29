// Copyright 2019 Cohesity Inc.
package vaults

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the VAULTS_IMPL
 */
type VAULTS interface {
    GetVaults (*int64, *string, *bool) ([]*models.Vault, error)

    CreateVault (*models.Vault) (*models.Vault, error)

    GetArchiveMediaInfo (int64, int64, int64, *int64, []int64) ([]*models.TapeMediaInformation, error)

    GetBandwidthSettings () (*models.VaultBandwidthLimits, error)

    UpdateBandwidthSettings (*models.VaultBandwidthLimits) (*models.VaultBandwidthLimits, error)

    GetVaultEncryptionKey (int64) (*models.VaultEncryptionKey, error)

    GetVaultById (int64) (*models.Vault, error)

    UpdateVault (int64, *models.Vault) (*models.Vault, error)
}

/*
 * Factory for the VAULTS interaface returning VAULTS_IMPL
 */
func NewVAULTS(config configuration.CONFIGURATION) *VAULTS_IMPL {
    client := new(VAULTS_IMPL)
    client.config = config
    return client
}