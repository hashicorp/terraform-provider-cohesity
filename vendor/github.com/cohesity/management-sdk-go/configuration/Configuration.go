// Copyright 2019 Cohesity Inc.
package configuration

import(
	"github.com/cohesity/management-sdk-go/models"
)

type CONFIGURATION interface {
        ClusterVip() *string
        SetClusterVip(clusterVip *string)
        Username() string
        SetUsername(username   string)
        Password() string
        SetPassword(password   string)
        Domain() string
        SetDomain(domain   string)
        AccessToken()  *models.AccessToken
        SetAccessToken(accessToken *models.AccessToken)
        SkipSSL() bool
        SetSkipSSL(skipSSL   bool)
}

/*
 * Factory for the CONFIGURATION interface returning CONFIGURATION_IMPL
 */
func NewCONFIGURATION() CONFIGURATION{
    configuration := new(CONFIGURATION_IMPL)
    cluster_vip := "prod-cluster.eng.cohesity.com"
    configuration.SetClusterVip(&cluster_vip)
    return configuration
}