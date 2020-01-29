// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PeriodicitySnapshotArchivalCopyPolicyEnum enum
 */
type PeriodicitySnapshotArchivalCopyPolicyEnum int

/**
 * Value collection for PeriodicitySnapshotArchivalCopyPolicyEnum enum
 */
const (
    PeriodicitySnapshotArchivalCopyPolicy_KEVERY PeriodicitySnapshotArchivalCopyPolicyEnum = 1 + iota
    PeriodicitySnapshotArchivalCopyPolicy_KHOUR
    PeriodicitySnapshotArchivalCopyPolicy_KDAY
    PeriodicitySnapshotArchivalCopyPolicy_KWEEK
    PeriodicitySnapshotArchivalCopyPolicy_KMONTH
    PeriodicitySnapshotArchivalCopyPolicy_KYEAR
)

func (r PeriodicitySnapshotArchivalCopyPolicyEnum) MarshalJSON() ([]byte, error) {
    s := PeriodicitySnapshotArchivalCopyPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *PeriodicitySnapshotArchivalCopyPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PeriodicitySnapshotArchivalCopyPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PeriodicitySnapshotArchivalCopyPolicyEnum to its string representation
 */
func PeriodicitySnapshotArchivalCopyPolicyEnumToValue(periodicitySnapshotArchivalCopyPolicyEnum PeriodicitySnapshotArchivalCopyPolicyEnum) string {
    switch periodicitySnapshotArchivalCopyPolicyEnum {
        case PeriodicitySnapshotArchivalCopyPolicy_KEVERY:
    		return "kEvery"
        case PeriodicitySnapshotArchivalCopyPolicy_KHOUR:
    		return "kHour"
        case PeriodicitySnapshotArchivalCopyPolicy_KDAY:
    		return "kDay"
        case PeriodicitySnapshotArchivalCopyPolicy_KWEEK:
    		return "kWeek"
        case PeriodicitySnapshotArchivalCopyPolicy_KMONTH:
    		return "kMonth"
        case PeriodicitySnapshotArchivalCopyPolicy_KYEAR:
    		return "kYear"
        default:
        	return "kEvery"
    }
}

/**
 * Converts PeriodicitySnapshotArchivalCopyPolicyEnum Array to its string Array representation
*/
func PeriodicitySnapshotArchivalCopyPolicyEnumArrayToValue(periodicitySnapshotArchivalCopyPolicyEnum []PeriodicitySnapshotArchivalCopyPolicyEnum) []string {
    convArray := make([]string,len( periodicitySnapshotArchivalCopyPolicyEnum))
    for i:=0; i<len(periodicitySnapshotArchivalCopyPolicyEnum);i++ {
        convArray[i] = PeriodicitySnapshotArchivalCopyPolicyEnumToValue(periodicitySnapshotArchivalCopyPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PeriodicitySnapshotArchivalCopyPolicyEnumFromValue(value string) PeriodicitySnapshotArchivalCopyPolicyEnum {
    switch value {
        case "kEvery":
            return PeriodicitySnapshotArchivalCopyPolicy_KEVERY
        case "kHour":
            return PeriodicitySnapshotArchivalCopyPolicy_KHOUR
        case "kDay":
            return PeriodicitySnapshotArchivalCopyPolicy_KDAY
        case "kWeek":
            return PeriodicitySnapshotArchivalCopyPolicy_KWEEK
        case "kMonth":
            return PeriodicitySnapshotArchivalCopyPolicy_KMONTH
        case "kYear":
            return PeriodicitySnapshotArchivalCopyPolicy_KYEAR
        default:
            return PeriodicitySnapshotArchivalCopyPolicy_KEVERY
    }
}
