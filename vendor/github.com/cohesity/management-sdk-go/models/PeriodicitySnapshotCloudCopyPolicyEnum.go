// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PeriodicitySnapshotCloudCopyPolicyEnum enum
 */
type PeriodicitySnapshotCloudCopyPolicyEnum int

/**
 * Value collection for PeriodicitySnapshotCloudCopyPolicyEnum enum
 */
const (
    PeriodicitySnapshotCloudCopyPolicy_KEVERY PeriodicitySnapshotCloudCopyPolicyEnum = 1 + iota
    PeriodicitySnapshotCloudCopyPolicy_KHOUR
    PeriodicitySnapshotCloudCopyPolicy_KDAY
    PeriodicitySnapshotCloudCopyPolicy_KWEEK
    PeriodicitySnapshotCloudCopyPolicy_KMONTH
    PeriodicitySnapshotCloudCopyPolicy_KYEAR
)

func (r PeriodicitySnapshotCloudCopyPolicyEnum) MarshalJSON() ([]byte, error) {
    s := PeriodicitySnapshotCloudCopyPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *PeriodicitySnapshotCloudCopyPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PeriodicitySnapshotCloudCopyPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PeriodicitySnapshotCloudCopyPolicyEnum to its string representation
 */
func PeriodicitySnapshotCloudCopyPolicyEnumToValue(periodicitySnapshotCloudCopyPolicyEnum PeriodicitySnapshotCloudCopyPolicyEnum) string {
    switch periodicitySnapshotCloudCopyPolicyEnum {
        case PeriodicitySnapshotCloudCopyPolicy_KEVERY:
    		return "kEvery"
        case PeriodicitySnapshotCloudCopyPolicy_KHOUR:
    		return "kHour"
        case PeriodicitySnapshotCloudCopyPolicy_KDAY:
    		return "kDay"
        case PeriodicitySnapshotCloudCopyPolicy_KWEEK:
    		return "kWeek"
        case PeriodicitySnapshotCloudCopyPolicy_KMONTH:
    		return "kMonth"
        case PeriodicitySnapshotCloudCopyPolicy_KYEAR:
    		return "kYear"
        default:
        	return "kEvery"
    }
}

/**
 * Converts PeriodicitySnapshotCloudCopyPolicyEnum Array to its string Array representation
*/
func PeriodicitySnapshotCloudCopyPolicyEnumArrayToValue(periodicitySnapshotCloudCopyPolicyEnum []PeriodicitySnapshotCloudCopyPolicyEnum) []string {
    convArray := make([]string,len( periodicitySnapshotCloudCopyPolicyEnum))
    for i:=0; i<len(periodicitySnapshotCloudCopyPolicyEnum);i++ {
        convArray[i] = PeriodicitySnapshotCloudCopyPolicyEnumToValue(periodicitySnapshotCloudCopyPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PeriodicitySnapshotCloudCopyPolicyEnumFromValue(value string) PeriodicitySnapshotCloudCopyPolicyEnum {
    switch value {
        case "kEvery":
            return PeriodicitySnapshotCloudCopyPolicy_KEVERY
        case "kHour":
            return PeriodicitySnapshotCloudCopyPolicy_KHOUR
        case "kDay":
            return PeriodicitySnapshotCloudCopyPolicy_KDAY
        case "kWeek":
            return PeriodicitySnapshotCloudCopyPolicy_KWEEK
        case "kMonth":
            return PeriodicitySnapshotCloudCopyPolicy_KMONTH
        case "kYear":
            return PeriodicitySnapshotCloudCopyPolicy_KYEAR
        default:
            return PeriodicitySnapshotCloudCopyPolicy_KEVERY
    }
}
