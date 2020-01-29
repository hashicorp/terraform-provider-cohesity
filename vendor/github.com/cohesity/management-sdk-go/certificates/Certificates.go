// Copyright 2019 Cohesity Inc.
package certificates

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the CERTIFICATES_IMPL
 */
type CERTIFICATES interface {
    DeleteWebServerCertificate () (error)

    GetWebServerCertificate () (*models.SslCertificateConfig, error)

    UpdateWebServerCertificate (*models.SslCertificateConfig) (*models.SslCertificateConfig, error)
}

/*
 * Factory for the CERTIFICATES interaface returning CERTIFICATES_IMPL
 */
func NewCERTIFICATES(config configuration.CONFIGURATION) *CERTIFICATES_IMPL {
    client := new(CERTIFICATES_IMPL)
    client.config = config
    return client
}