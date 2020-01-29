// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AuthenticationTypeUserEnum enum
 */
type AuthenticationTypeUserEnum int

/**
 * Value collection for AuthenticationTypeUserEnum enum
 */
const (
    AuthenticationTypeUser_KAUTHLOCAL AuthenticationTypeUserEnum = 1 + iota
    AuthenticationTypeUser_KAUTHAD
    AuthenticationTypeUser_KAUTHSALESFORCE
    AuthenticationTypeUser_KAUTHGOOGLE
    AuthenticationTypeUser_KAUTHSSO
)

func (r AuthenticationTypeUserEnum) MarshalJSON() ([]byte, error) {
    s := AuthenticationTypeUserEnumToValue(r)
    return json.Marshal(s)
}

func (r *AuthenticationTypeUserEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AuthenticationTypeUserEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AuthenticationTypeUserEnum to its string representation
 */
func AuthenticationTypeUserEnumToValue(authenticationTypeUserEnum AuthenticationTypeUserEnum) string {
    switch authenticationTypeUserEnum {
        case AuthenticationTypeUser_KAUTHLOCAL:
    		return "kAuthLocal"
        case AuthenticationTypeUser_KAUTHAD:
    		return "kAuthAd"
        case AuthenticationTypeUser_KAUTHSALESFORCE:
    		return "kAuthSalesforce"
        case AuthenticationTypeUser_KAUTHGOOGLE:
    		return "kAuthGoogle"
        case AuthenticationTypeUser_KAUTHSSO:
    		return "kAuthSso"
        default:
        	return "kAuthLocal"
    }
}

/**
 * Converts AuthenticationTypeUserEnum Array to its string Array representation
*/
func AuthenticationTypeUserEnumArrayToValue(authenticationTypeUserEnum []AuthenticationTypeUserEnum) []string {
    convArray := make([]string,len( authenticationTypeUserEnum))
    for i:=0; i<len(authenticationTypeUserEnum);i++ {
        convArray[i] = AuthenticationTypeUserEnumToValue(authenticationTypeUserEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AuthenticationTypeUserEnumFromValue(value string) AuthenticationTypeUserEnum {
    switch value {
        case "kAuthLocal":
            return AuthenticationTypeUser_KAUTHLOCAL
        case "kAuthAd":
            return AuthenticationTypeUser_KAUTHAD
        case "kAuthSalesforce":
            return AuthenticationTypeUser_KAUTHSALESFORCE
        case "kAuthGoogle":
            return AuthenticationTypeUser_KAUTHGOOGLE
        case "kAuthSso":
            return AuthenticationTypeUser_KAUTHSSO
        default:
            return AuthenticationTypeUser_KAUTHLOCAL
    }
}
