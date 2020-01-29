// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AlertStateListEnum enum
 */
type AlertStateListEnum int

/**
 * Value collection for AlertStateListEnum enum
 */
const (
    AlertStateList_KOPEN AlertStateListEnum = 1 + iota
    AlertStateList_KRESOLVED
    AlertStateList_KALERTSUPPRESSED
)

func (r AlertStateListEnum) MarshalJSON() ([]byte, error) {
    s := AlertStateListEnumToValue(r)
    return json.Marshal(s)
}

func (r *AlertStateListEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AlertStateListEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AlertStateListEnum to its string representation
 */
func AlertStateListEnumToValue(alertStateListEnum AlertStateListEnum) string {
    switch alertStateListEnum {
        case AlertStateList_KOPEN:
    		return "kOpen"
        case AlertStateList_KRESOLVED:
    		return "kResolved"
        case AlertStateList_KALERTSUPPRESSED:
    		return "kAlertSuppressed"
        default:
        	return "kOpen"
    }
}

/**
 * Converts AlertStateListEnum Array to its string Array representation
*/
func AlertStateListEnumArrayToValue(alertStateListEnum []AlertStateListEnum) []string {
    convArray := make([]string,len( alertStateListEnum))
    for i:=0; i<len(alertStateListEnum);i++ {
        convArray[i] = AlertStateListEnumToValue(alertStateListEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlertStateListEnumFromValue(value string) AlertStateListEnum {
    switch value {
        case "kOpen":
            return AlertStateList_KOPEN
        case "kResolved":
            return AlertStateList_KRESOLVED
        case "kAlertSuppressed":
            return AlertStateList_KALERTSUPPRESSED
        default:
            return AlertStateList_KOPEN
    }
}
