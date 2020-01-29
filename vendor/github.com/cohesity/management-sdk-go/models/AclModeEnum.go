// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AclModeEnum enum
 */
type AclModeEnum int

/**
 * Value collection for AclModeEnum enum
 */
const (
    AclMode_KSHARED AclModeEnum = 1 + iota
    AclMode_KNATIVE
)

func (r AclModeEnum) MarshalJSON() ([]byte, error) {
    s := AclModeEnumToValue(r)
    return json.Marshal(s)
}

func (r *AclModeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AclModeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AclModeEnum to its string representation
 */
func AclModeEnumToValue(aclModeEnum AclModeEnum) string {
    switch aclModeEnum {
        case AclMode_KSHARED:
    		return "kShared"
        case AclMode_KNATIVE:
    		return "kNative"
        default:
        	return "kShared"
    }
}

/**
 * Converts AclModeEnum Array to its string Array representation
*/
func AclModeEnumArrayToValue(aclModeEnum []AclModeEnum) []string {
    convArray := make([]string,len( aclModeEnum))
    for i:=0; i<len(aclModeEnum);i++ {
        convArray[i] = AclModeEnumToValue(aclModeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AclModeEnumFromValue(value string) AclModeEnum {
    switch value {
        case "kShared":
            return AclMode_KSHARED
        case "kNative":
            return AclMode_KNATIVE
        default:
            return AclMode_KSHARED
    }
}
