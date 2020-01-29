// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AuthTypeEnum enum
 */
type AuthTypeEnum int

/**
 * Value collection for AuthTypeEnum enum
 */
const (
    AuthType_KANONYMOUS AuthTypeEnum = 1 + iota
    AuthType_KSIMPLE
)

func (r AuthTypeEnum) MarshalJSON() ([]byte, error) {
    s := AuthTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *AuthTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AuthTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AuthTypeEnum to its string representation
 */
func AuthTypeEnumToValue(authTypeEnum AuthTypeEnum) string {
    switch authTypeEnum {
        case AuthType_KANONYMOUS:
    		return "kAnonymous"
        case AuthType_KSIMPLE:
    		return "kSimple"
        default:
        	return "kAnonymous"
    }
}

/**
 * Converts AuthTypeEnum Array to its string Array representation
*/
func AuthTypeEnumArrayToValue(authTypeEnum []AuthTypeEnum) []string {
    convArray := make([]string,len( authTypeEnum))
    for i:=0; i<len(authTypeEnum);i++ {
        convArray[i] = AuthTypeEnumToValue(authTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AuthTypeEnumFromValue(value string) AuthTypeEnum {
    switch value {
        case "kAnonymous":
            return AuthType_KANONYMOUS
        case "kSimple":
            return AuthType_KSIMPLE
        default:
            return AuthType_KANONYMOUS
    }
}
