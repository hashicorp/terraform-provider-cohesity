// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AuthenticationTypeEnum enum
 */
type AuthenticationTypeEnum int

/**
 * Value collection for AuthenticationTypeEnum enum
 */
const (
    AuthenticationType_KPASSWORDONLY AuthenticationTypeEnum = 1 + iota
    AuthenticationType_KCERTIFICATEONLY
    AuthenticationType_KPASSWORDANDCERTIFICATE
)

func (r AuthenticationTypeEnum) MarshalJSON() ([]byte, error) {
    s := AuthenticationTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *AuthenticationTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AuthenticationTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AuthenticationTypeEnum to its string representation
 */
func AuthenticationTypeEnumToValue(authenticationTypeEnum AuthenticationTypeEnum) string {
    switch authenticationTypeEnum {
        case AuthenticationType_KPASSWORDONLY:
    		return "kPasswordOnly"
        case AuthenticationType_KCERTIFICATEONLY:
    		return "kCertificateOnly"
        case AuthenticationType_KPASSWORDANDCERTIFICATE:
    		return "kPasswordAndCertificate"
        default:
        	return "kPasswordOnly"
    }
}

/**
 * Converts AuthenticationTypeEnum Array to its string Array representation
*/
func AuthenticationTypeEnumArrayToValue(authenticationTypeEnum []AuthenticationTypeEnum) []string {
    convArray := make([]string,len( authenticationTypeEnum))
    for i:=0; i<len(authenticationTypeEnum);i++ {
        convArray[i] = AuthenticationTypeEnumToValue(authenticationTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AuthenticationTypeEnumFromValue(value string) AuthenticationTypeEnum {
    switch value {
        case "kPasswordOnly":
            return AuthenticationType_KPASSWORDONLY
        case "kCertificateOnly":
            return AuthenticationType_KCERTIFICATEONLY
        case "kPasswordAndCertificate":
            return AuthenticationType_KPASSWORDANDCERTIFICATE
        default:
            return AuthenticationType_KPASSWORDONLY
    }
}
