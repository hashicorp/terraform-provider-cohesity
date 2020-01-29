// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PeriodicityExtendedRetentionPolicyEnum enum
 */
type PeriodicityExtendedRetentionPolicyEnum int

/**
 * Value collection for PeriodicityExtendedRetentionPolicyEnum enum
 */
const (
    PeriodicityExtendedRetentionPolicy_KEVERY PeriodicityExtendedRetentionPolicyEnum = 1 + iota
    PeriodicityExtendedRetentionPolicy_KHOUR
    PeriodicityExtendedRetentionPolicy_KDAY
    PeriodicityExtendedRetentionPolicy_KWEEK
    PeriodicityExtendedRetentionPolicy_KMONTH
    PeriodicityExtendedRetentionPolicy_KYEAR
)

func (r PeriodicityExtendedRetentionPolicyEnum) MarshalJSON() ([]byte, error) {
    s := PeriodicityExtendedRetentionPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *PeriodicityExtendedRetentionPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PeriodicityExtendedRetentionPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PeriodicityExtendedRetentionPolicyEnum to its string representation
 */
func PeriodicityExtendedRetentionPolicyEnumToValue(periodicityExtendedRetentionPolicyEnum PeriodicityExtendedRetentionPolicyEnum) string {
    switch periodicityExtendedRetentionPolicyEnum {
        case PeriodicityExtendedRetentionPolicy_KEVERY:
    		return "kEvery"
        case PeriodicityExtendedRetentionPolicy_KHOUR:
    		return "kHour"
        case PeriodicityExtendedRetentionPolicy_KDAY:
    		return "kDay"
        case PeriodicityExtendedRetentionPolicy_KWEEK:
    		return "kWeek"
        case PeriodicityExtendedRetentionPolicy_KMONTH:
    		return "kMonth"
        case PeriodicityExtendedRetentionPolicy_KYEAR:
    		return "kYear"
        default:
        	return "kEvery"
    }
}

/**
 * Converts PeriodicityExtendedRetentionPolicyEnum Array to its string Array representation
*/
func PeriodicityExtendedRetentionPolicyEnumArrayToValue(periodicityExtendedRetentionPolicyEnum []PeriodicityExtendedRetentionPolicyEnum) []string {
    convArray := make([]string,len( periodicityExtendedRetentionPolicyEnum))
    for i:=0; i<len(periodicityExtendedRetentionPolicyEnum);i++ {
        convArray[i] = PeriodicityExtendedRetentionPolicyEnumToValue(periodicityExtendedRetentionPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PeriodicityExtendedRetentionPolicyEnumFromValue(value string) PeriodicityExtendedRetentionPolicyEnum {
    switch value {
        case "kEvery":
            return PeriodicityExtendedRetentionPolicy_KEVERY
        case "kHour":
            return PeriodicityExtendedRetentionPolicy_KHOUR
        case "kDay":
            return PeriodicityExtendedRetentionPolicy_KDAY
        case "kWeek":
            return PeriodicityExtendedRetentionPolicy_KWEEK
        case "kMonth":
            return PeriodicityExtendedRetentionPolicy_KMONTH
        case "kYear":
            return PeriodicityExtendedRetentionPolicy_KYEAR
        default:
            return PeriodicityExtendedRetentionPolicy_KEVERY
    }
}
