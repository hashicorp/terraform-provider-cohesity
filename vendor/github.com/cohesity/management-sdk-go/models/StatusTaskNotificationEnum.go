// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusTaskNotificationEnum enum
 */
type StatusTaskNotificationEnum int

/**
 * Value collection for StatusTaskNotificationEnum enum
 */
const (
    StatusTaskNotification_KSUCCESS StatusTaskNotificationEnum = 1 + iota
    StatusTaskNotification_KERROR
)

func (r StatusTaskNotificationEnum) MarshalJSON() ([]byte, error) {
    s := StatusTaskNotificationEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusTaskNotificationEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusTaskNotificationEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusTaskNotificationEnum to its string representation
 */
func StatusTaskNotificationEnumToValue(statusTaskNotificationEnum StatusTaskNotificationEnum) string {
    switch statusTaskNotificationEnum {
        case StatusTaskNotification_KSUCCESS:
    		return "kSuccess"
        case StatusTaskNotification_KERROR:
    		return "kError"
        default:
        	return "kSuccess"
    }
}

/**
 * Converts StatusTaskNotificationEnum Array to its string Array representation
*/
func StatusTaskNotificationEnumArrayToValue(statusTaskNotificationEnum []StatusTaskNotificationEnum) []string {
    convArray := make([]string,len( statusTaskNotificationEnum))
    for i:=0; i<len(statusTaskNotificationEnum);i++ {
        convArray[i] = StatusTaskNotificationEnumToValue(statusTaskNotificationEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusTaskNotificationEnumFromValue(value string) StatusTaskNotificationEnum {
    switch value {
        case "kSuccess":
            return StatusTaskNotification_KSUCCESS
        case "kError":
            return StatusTaskNotification_KERROR
        default:
            return StatusTaskNotification_KSUCCESS
    }
}
