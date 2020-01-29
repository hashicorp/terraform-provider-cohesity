// Copyright 2019 Cohesity Inc.
package configuration

import(
	"github.com/cohesity/management-sdk-go/apihelper"
	"github.com/cohesity/management-sdk-go/models"
)
/* Setting up enums for Environments and Servers
*/
type Environments int

type Servers int

// Environment Enums
const (
     PRODUCTION Environments = 1 + iota
)

// Servers Enums
const (
 	 DEFAULT_HOST Servers = 1 + iota
)

type CONFIGURATION_IMPL struct {
    /** Replace the value of cluster_vip with SetClusterVip function */
    cluster_vip *string
    /** Specifies the login name of the Cohesity user. */
    /** Replace the value of username with SetUsername function */
    username string
    /** Specifies the password of the Cohesity user account. */
    /** Replace the value of password with SetPassword function */
    password string
    /** Specifies the domain the user is logging in to. For a Local user model, the domain is always LOCAL. For LDAP / AD user models, the domain will map to an LDAP connection string. A user is uniquely identified by a combination of username and domain. If this is not set, LOCAL is assumed. */
    /** Replace the value of domain with SetDomain function */
    domain string
    /** Access Token to be used to make calls to the API */
    access_token *models.AccessToken
    /**Configures verification of SSL certificates**/
    skip_ssl bool
}

/*
 * Getter function returning cluster_vip
 */
func (me *CONFIGURATION_IMPL) ClusterVip() *string{
    return me.cluster_vip
}

/*
 * Setter function setting up cluster_vip
 */
func (me *CONFIGURATION_IMPL) SetClusterVip(clusterVip *string) {
    me.cluster_vip = clusterVip
}

/*
 * Getter function returning username
 */
func (me *CONFIGURATION_IMPL) Username() string{
    return me.username
}

/*
 * Setter function setting up username
 */
func (me *CONFIGURATION_IMPL) SetUsername(username string) {
    me.username = username
}

/*
 * Getter function returning password
 */
func (me *CONFIGURATION_IMPL) Password() string{
    return me.password
}

/*
 * Setter function setting up password
 */
func (me *CONFIGURATION_IMPL) SetPassword(password string) {
    me.password = password
}

/*
 * Getter function returning domain
 */
func (me *CONFIGURATION_IMPL) Domain() string{
    return me.domain
}

/*
 * Setter function setting up domain
 */
func (me *CONFIGURATION_IMPL) SetDomain(domain string) {
    me.domain = domain
}

/*
 * Getter function returning access_token
 */
func (me *CONFIGURATION_IMPL) AccessToken() *models.AccessToken{
    return me.access_token
}

/*
 * Setter function setting up access_token
 */
func (me *CONFIGURATION_IMPL) SetAccessToken(accessToken *models.AccessToken) {
    me.access_token = accessToken
}
/*
 * Getter function returning skip_ssl
 */
func (me *CONFIGURATION_IMPL) SkipSSL() bool{
    return me.skip_ssl
}

/*
 * Setter function setting up skip_ssl
 */
func (me *CONFIGURATION_IMPL) SetSkipSSL(skipSSL bool) {
    me.skip_ssl = skipSSL
}
// Setting up Default Environment
var Environment = PRODUCTION

// A map of environments and their corresponding servers/baseurls
var EnvironmentsMap = map[Environments](map[Servers]string){

    PRODUCTION : map[Servers]string{
        DEFAULT_HOST:"https://{cluster_vip}/irisservices/api/v1",
    },
}

// Make and return the map of parameters
func GetBaseURIParameters(config CONFIGURATION) map[string]interface{} {
     kvpMap := map[string]interface{}{
         "cluster_vip": config.ClusterVip(),
    }
    return kvpMap;
}

// Gets the URL for a particular alias in the current environment and appends it with template parameters
// return the baseurl
func GetBaseURI(server Servers, config CONFIGURATION) string {
    url := EnvironmentsMap[Environment][server]
    appendedURL, _ := apihelper.AppendUrlWithTemplateParameters(url, GetBaseURIParameters(config))
    return appendedURL

}
