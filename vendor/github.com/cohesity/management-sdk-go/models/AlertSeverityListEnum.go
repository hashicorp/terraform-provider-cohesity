// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AlertSeverityListEnum enum
 */
type AlertSeverityListEnum int

/**
 * Value collection for AlertSeverityListEnum enum
 */
const (
    AlertSeverityList_KCRITICAL AlertSeverityListEnum = 1 + iota
    AlertSeverityList_KWARNING
    AlertSeverityList_KINFO
)

func (r AlertSeverityListEnum) MarshalJSON() ([]byte, error) {
    s := AlertSeverityListEnumToValue(r)
    return json.Marshal(s)
}

func (r *AlertSeverityListEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AlertSeverityListEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AlertSeverityListEnum to its string representation
 */
func AlertSeverityListEnumToValue(alertSeverityListEnum AlertSeverityListEnum) string {
    switch alertSeverityListEnum {
        case AlertSeverityList_KCRITICAL:
    		return "kCritical"
        case AlertSeverityList_KWARNING:
    		return "kWarning"
        case AlertSeverityList_KINFO:
    		return "kInfo"
        default:
        	return "kCritical"
    }
}

/**
 * Converts AlertSeverityListEnum Array to its string Array representation
*/
func AlertSeverityListEnumArrayToValue(alertSeverityListEnum []AlertSeverityListEnum) []string {
    convArray := make([]string,len( alertSeverityListEnum))
    for i:=0; i<len(alertSeverityListEnum);i++ {
        convArray[i] = AlertSeverityListEnumToValue(alertSeverityListEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlertSeverityListEnumFromValue(value string) AlertSeverityListEnum {
    switch value {
        case "kCritical":
            return AlertSeverityList_KCRITICAL
        case "kWarning":
            return AlertSeverityList_KWARNING
        case "kInfo":
            return AlertSeverityList_KINFO
        default:
            return AlertSeverityList_KCRITICAL
    }
}
