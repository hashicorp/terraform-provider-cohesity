// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AuthenticationStatusEnum enum
 */
type AuthenticationStatusEnum int

/**
 * Value collection for AuthenticationStatusEnum enum
 */
const (
    AuthenticationStatus_KPENDING AuthenticationStatusEnum = 1 + iota
    AuthenticationStatus_KSCHEDULED
    AuthenticationStatus_KFINISHED
    AuthenticationStatus_KREFRESHINPROGRESS
)

func (r AuthenticationStatusEnum) MarshalJSON() ([]byte, error) {
    s := AuthenticationStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *AuthenticationStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AuthenticationStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AuthenticationStatusEnum to its string representation
 */
func AuthenticationStatusEnumToValue(authenticationStatusEnum AuthenticationStatusEnum) string {
    switch authenticationStatusEnum {
        case AuthenticationStatus_KPENDING:
    		return "kPending"
        case AuthenticationStatus_KSCHEDULED:
    		return "kScheduled"
        case AuthenticationStatus_KFINISHED:
    		return "kFinished"
        case AuthenticationStatus_KREFRESHINPROGRESS:
    		return "kRefreshInProgress"
        default:
        	return "kPending"
    }
}

/**
 * Converts AuthenticationStatusEnum Array to its string Array representation
*/
func AuthenticationStatusEnumArrayToValue(authenticationStatusEnum []AuthenticationStatusEnum) []string {
    convArray := make([]string,len( authenticationStatusEnum))
    for i:=0; i<len(authenticationStatusEnum);i++ {
        convArray[i] = AuthenticationStatusEnumToValue(authenticationStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AuthenticationStatusEnumFromValue(value string) AuthenticationStatusEnum {
    switch value {
        case "kPending":
            return AuthenticationStatus_KPENDING
        case "kScheduled":
            return AuthenticationStatus_KSCHEDULED
        case "kFinished":
            return AuthenticationStatus_KFINISHED
        case "kRefreshInProgress":
            return AuthenticationStatus_KREFRESHINPROGRESS
        default:
            return AuthenticationStatus_KPENDING
    }
}
