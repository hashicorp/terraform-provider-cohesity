// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for DayBlackoutPeriodEnum enum
 */
type DayBlackoutPeriodEnum int

/**
 * Value collection for DayBlackoutPeriodEnum enum
 */
const (
    DayBlackoutPeriod_KSUNDAY DayBlackoutPeriodEnum = 1 + iota
    DayBlackoutPeriod_KMONDAY
    DayBlackoutPeriod_KTUESDAY
    DayBlackoutPeriod_KWEDNESDAY
    DayBlackoutPeriod_KTHURSDAY
    DayBlackoutPeriod_KFRIDAY
    DayBlackoutPeriod_KSATURDAY
)

func (r DayBlackoutPeriodEnum) MarshalJSON() ([]byte, error) {
    s := DayBlackoutPeriodEnumToValue(r)
    return json.Marshal(s)
}

func (r *DayBlackoutPeriodEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  DayBlackoutPeriodEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts DayBlackoutPeriodEnum to its string representation
 */
func DayBlackoutPeriodEnumToValue(dayBlackoutPeriodEnum DayBlackoutPeriodEnum) string {
    switch dayBlackoutPeriodEnum {
        case DayBlackoutPeriod_KSUNDAY:
    		return "kSunday"
        case DayBlackoutPeriod_KMONDAY:
    		return "kMonday"
        case DayBlackoutPeriod_KTUESDAY:
    		return "kTuesday"
        case DayBlackoutPeriod_KWEDNESDAY:
    		return "kWednesday"
        case DayBlackoutPeriod_KTHURSDAY:
    		return "kThursday"
        case DayBlackoutPeriod_KFRIDAY:
    		return "kFriday"
        case DayBlackoutPeriod_KSATURDAY:
    		return "kSaturday"
        default:
        	return "kSunday"
    }
}

/**
 * Converts DayBlackoutPeriodEnum Array to its string Array representation
*/
func DayBlackoutPeriodEnumArrayToValue(dayBlackoutPeriodEnum []DayBlackoutPeriodEnum) []string {
    convArray := make([]string,len( dayBlackoutPeriodEnum))
    for i:=0; i<len(dayBlackoutPeriodEnum);i++ {
        convArray[i] = DayBlackoutPeriodEnumToValue(dayBlackoutPeriodEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DayBlackoutPeriodEnumFromValue(value string) DayBlackoutPeriodEnum {
    switch value {
        case "kSunday":
            return DayBlackoutPeriod_KSUNDAY
        case "kMonday":
            return DayBlackoutPeriod_KMONDAY
        case "kTuesday":
            return DayBlackoutPeriod_KTUESDAY
        case "kWednesday":
            return DayBlackoutPeriod_KWEDNESDAY
        case "kThursday":
            return DayBlackoutPeriod_KTHURSDAY
        case "kFriday":
            return DayBlackoutPeriod_KFRIDAY
        case "kSaturday":
            return DayBlackoutPeriod_KSATURDAY
        default:
            return DayBlackoutPeriod_KSUNDAY
    }
}
