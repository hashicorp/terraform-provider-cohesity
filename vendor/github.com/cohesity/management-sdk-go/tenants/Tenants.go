package tenants

import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the TENANTS_IMPL
 */
type TENANTS interface {
    GetDownloadTenantsProxy (*string) ([]int64, error)
}

/*
 * Factory for the TENANTS interaface returning TENANTS_IMPL
 */
func NewTENANTS(config configuration.CONFIGURATION) *TENANTS_IMPL {
    client := new(TENANTS_IMPL)
    client.config = config
    return client
}
