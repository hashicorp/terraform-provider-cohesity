// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SecurityModeEnum enum
 */
type SecurityModeEnum int

/**
 * Value collection for SecurityModeEnum enum
 */
const (
    SecurityMode_KNATIVEMODE SecurityModeEnum = 1 + iota
    SecurityMode_KUNIFIEDMODE
    SecurityMode_KNTFSMODE
)

func (r SecurityModeEnum) MarshalJSON() ([]byte, error) {
    s := SecurityModeEnumToValue(r)
    return json.Marshal(s)
}

func (r *SecurityModeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SecurityModeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SecurityModeEnum to its string representation
 */
func SecurityModeEnumToValue(securityModeEnum SecurityModeEnum) string {
    switch securityModeEnum {
        case SecurityMode_KNATIVEMODE:
    		return "kNativeMode"
        case SecurityMode_KUNIFIEDMODE:
    		return "kUnifiedMode"
        case SecurityMode_KNTFSMODE:
    		return "kNtfsMode"
        default:
        	return "kNativeMode"
    }
}

/**
 * Converts SecurityModeEnum Array to its string Array representation
*/
func SecurityModeEnumArrayToValue(securityModeEnum []SecurityModeEnum) []string {
    convArray := make([]string,len( securityModeEnum))
    for i:=0; i<len(securityModeEnum);i++ {
        convArray[i] = SecurityModeEnumToValue(securityModeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SecurityModeEnumFromValue(value string) SecurityModeEnum {
    switch value {
        case "kNativeMode":
            return SecurityMode_KNATIVEMODE
        case "kUnifiedMode":
            return SecurityMode_KUNIFIEDMODE
        case "kNtfsMode":
            return SecurityMode_KNTFSMODE
        default:
            return SecurityMode_KNATIVEMODE
    }
}
