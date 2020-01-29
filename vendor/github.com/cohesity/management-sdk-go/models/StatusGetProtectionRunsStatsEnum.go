// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusGetProtectionRunsStatsEnum enum
 */
type StatusGetProtectionRunsStatsEnum int

/**
 * Value collection for StatusGetProtectionRunsStatsEnum enum
 */
const (
    StatusGetProtectionRunsStats_KSUCCESS StatusGetProtectionRunsStatsEnum = 1 + iota
    StatusGetProtectionRunsStats_KFAILURE
    StatusGetProtectionRunsStats_KCANCELED
    StatusGetProtectionRunsStats_KWARNING
)

func (r StatusGetProtectionRunsStatsEnum) MarshalJSON() ([]byte, error) {
    s := StatusGetProtectionRunsStatsEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusGetProtectionRunsStatsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusGetProtectionRunsStatsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusGetProtectionRunsStatsEnum to its string representation
 */
func StatusGetProtectionRunsStatsEnumToValue(statusGetProtectionRunsStatsEnum StatusGetProtectionRunsStatsEnum) string {
    switch statusGetProtectionRunsStatsEnum {
        case StatusGetProtectionRunsStats_KSUCCESS:
    		return "kSuccess"
        case StatusGetProtectionRunsStats_KFAILURE:
    		return "kFailure"
        case StatusGetProtectionRunsStats_KCANCELED:
    		return "kCanceled"
        case StatusGetProtectionRunsStats_KWARNING:
    		return "kWarning"
        default:
        	return "kSuccess"
    }
}

/**
 * Converts StatusGetProtectionRunsStatsEnum Array to its string Array representation
*/
func StatusGetProtectionRunsStatsEnumArrayToValue(statusGetProtectionRunsStatsEnum []StatusGetProtectionRunsStatsEnum) []string {
    convArray := make([]string,len( statusGetProtectionRunsStatsEnum))
    for i:=0; i<len(statusGetProtectionRunsStatsEnum);i++ {
        convArray[i] = StatusGetProtectionRunsStatsEnumToValue(statusGetProtectionRunsStatsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusGetProtectionRunsStatsEnumFromValue(value string) StatusGetProtectionRunsStatsEnum {
    switch value {
        case "kSuccess":
            return StatusGetProtectionRunsStats_KSUCCESS
        case "kFailure":
            return StatusGetProtectionRunsStats_KFAILURE
        case "kCanceled":
            return StatusGetProtectionRunsStats_KCANCELED
        case "kWarning":
            return StatusGetProtectionRunsStats_KWARNING
        default:
            return StatusGetProtectionRunsStats_KSUCCESS
    }
}
