// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AlertStateEnum enum
 */
type AlertStateEnum int

/**
 * Value collection for AlertStateEnum enum
 */
const (
    AlertState_KOPEN AlertStateEnum = 1 + iota
    AlertState_KRESOLVED
    AlertState_KSUPPRESSED
)

func (r AlertStateEnum) MarshalJSON() ([]byte, error) {
    s := AlertStateEnumToValue(r)
    return json.Marshal(s)
}

func (r *AlertStateEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AlertStateEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AlertStateEnum to its string representation
 */
func AlertStateEnumToValue(alertStateEnum AlertStateEnum) string {
    switch alertStateEnum {
        case AlertState_KOPEN:
    		return "kOpen"
        case AlertState_KRESOLVED:
    		return "kResolved"
        case AlertState_KSUPPRESSED:
    		return "kSuppressed"
        default:
        	return "kOpen"
    }
}

/**
 * Converts AlertStateEnum Array to its string Array representation
*/
func AlertStateEnumArrayToValue(alertStateEnum []AlertStateEnum) []string {
    convArray := make([]string,len( alertStateEnum))
    for i:=0; i<len(alertStateEnum);i++ {
        convArray[i] = AlertStateEnumToValue(alertStateEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlertStateEnumFromValue(value string) AlertStateEnum {
    switch value {
        case "kOpen":
            return AlertState_KOPEN
        case "kResolved":
            return AlertState_KRESOLVED
        case "kSuppressed":
            return AlertState_KSUPPRESSED
        default:
            return AlertState_KOPEN
    }
}
