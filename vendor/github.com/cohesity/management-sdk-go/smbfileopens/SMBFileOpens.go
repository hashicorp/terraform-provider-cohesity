// Copyright 2019 Cohesity Inc.
package smbfileopens

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the SMBFILEOPENS_IMPL
 */
type SMBFILEOPENS interface {
    GetSmbFileOpens (*string, *string, *int64, *string) (*models.SmbActiveFileOpensResponse, error)

    CreateCloseSmbFileOpen (*models.CloseSmbFileOpenParameters) (error)
}

/*
 * Factory for the SMBFILEOPENS interaface returning SMBFILEOPENS_IMPL
 */
func NewSMBFILEOPENS(config configuration.CONFIGURATION) *SMBFILEOPENS_IMPL {
    client := new(SMBFILEOPENS_IMPL)
    client.config = config
    return client
}