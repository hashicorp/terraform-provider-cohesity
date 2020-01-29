// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PeriodicitySnapshotReplicationCopyPolicyEnum enum
 */
type PeriodicitySnapshotReplicationCopyPolicyEnum int

/**
 * Value collection for PeriodicitySnapshotReplicationCopyPolicyEnum enum
 */
const (
    PeriodicitySnapshotReplicationCopyPolicy_KEVERY PeriodicitySnapshotReplicationCopyPolicyEnum = 1 + iota
    PeriodicitySnapshotReplicationCopyPolicy_KHOUR
    PeriodicitySnapshotReplicationCopyPolicy_KDAY
    PeriodicitySnapshotReplicationCopyPolicy_KWEEK
    PeriodicitySnapshotReplicationCopyPolicy_KMONTH
    PeriodicitySnapshotReplicationCopyPolicy_KYEAR
)

func (r PeriodicitySnapshotReplicationCopyPolicyEnum) MarshalJSON() ([]byte, error) {
    s := PeriodicitySnapshotReplicationCopyPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *PeriodicitySnapshotReplicationCopyPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PeriodicitySnapshotReplicationCopyPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PeriodicitySnapshotReplicationCopyPolicyEnum to its string representation
 */
func PeriodicitySnapshotReplicationCopyPolicyEnumToValue(periodicitySnapshotReplicationCopyPolicyEnum PeriodicitySnapshotReplicationCopyPolicyEnum) string {
    switch periodicitySnapshotReplicationCopyPolicyEnum {
        case PeriodicitySnapshotReplicationCopyPolicy_KEVERY:
    		return "kEvery"
        case PeriodicitySnapshotReplicationCopyPolicy_KHOUR:
    		return "kHour"
        case PeriodicitySnapshotReplicationCopyPolicy_KDAY:
    		return "kDay"
        case PeriodicitySnapshotReplicationCopyPolicy_KWEEK:
    		return "kWeek"
        case PeriodicitySnapshotReplicationCopyPolicy_KMONTH:
    		return "kMonth"
        case PeriodicitySnapshotReplicationCopyPolicy_KYEAR:
    		return "kYear"
        default:
        	return "kEvery"
    }
}

/**
 * Converts PeriodicitySnapshotReplicationCopyPolicyEnum Array to its string Array representation
*/
func PeriodicitySnapshotReplicationCopyPolicyEnumArrayToValue(periodicitySnapshotReplicationCopyPolicyEnum []PeriodicitySnapshotReplicationCopyPolicyEnum) []string {
    convArray := make([]string,len( periodicitySnapshotReplicationCopyPolicyEnum))
    for i:=0; i<len(periodicitySnapshotReplicationCopyPolicyEnum);i++ {
        convArray[i] = PeriodicitySnapshotReplicationCopyPolicyEnumToValue(periodicitySnapshotReplicationCopyPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PeriodicitySnapshotReplicationCopyPolicyEnumFromValue(value string) PeriodicitySnapshotReplicationCopyPolicyEnum {
    switch value {
        case "kEvery":
            return PeriodicitySnapshotReplicationCopyPolicy_KEVERY
        case "kHour":
            return PeriodicitySnapshotReplicationCopyPolicy_KHOUR
        case "kDay":
            return PeriodicitySnapshotReplicationCopyPolicy_KDAY
        case "kWeek":
            return PeriodicitySnapshotReplicationCopyPolicy_KWEEK
        case "kMonth":
            return PeriodicitySnapshotReplicationCopyPolicy_KMONTH
        case "kYear":
            return PeriodicitySnapshotReplicationCopyPolicy_KYEAR
        default:
            return PeriodicitySnapshotReplicationCopyPolicy_KEVERY
    }
}
