// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for DayMonthlyScheduleEnum enum
 */
type DayMonthlyScheduleEnum int

/**
 * Value collection for DayMonthlyScheduleEnum enum
 */
const (
    DayMonthlySchedule_KSUNDAY DayMonthlyScheduleEnum = 1 + iota
    DayMonthlySchedule_KMONDAY
    DayMonthlySchedule_KTUESDAY
    DayMonthlySchedule_KWEDNESDAY
    DayMonthlySchedule_KTHURSDAY
    DayMonthlySchedule_KFRIDAY
    DayMonthlySchedule_KSATURDAY
)

func (r DayMonthlyScheduleEnum) MarshalJSON() ([]byte, error) {
    s := DayMonthlyScheduleEnumToValue(r)
    return json.Marshal(s)
}

func (r *DayMonthlyScheduleEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  DayMonthlyScheduleEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts DayMonthlyScheduleEnum to its string representation
 */
func DayMonthlyScheduleEnumToValue(dayMonthlyScheduleEnum DayMonthlyScheduleEnum) string {
    switch dayMonthlyScheduleEnum {
        case DayMonthlySchedule_KSUNDAY:
    		return "kSunday"
        case DayMonthlySchedule_KMONDAY:
    		return "kMonday"
        case DayMonthlySchedule_KTUESDAY:
    		return "kTuesday"
        case DayMonthlySchedule_KWEDNESDAY:
    		return "kWednesday"
        case DayMonthlySchedule_KTHURSDAY:
    		return "kThursday"
        case DayMonthlySchedule_KFRIDAY:
    		return "kFriday"
        case DayMonthlySchedule_KSATURDAY:
    		return "kSaturday"
        default:
        	return "kSunday"
    }
}

/**
 * Converts DayMonthlyScheduleEnum Array to its string Array representation
*/
func DayMonthlyScheduleEnumArrayToValue(dayMonthlyScheduleEnum []DayMonthlyScheduleEnum) []string {
    convArray := make([]string,len( dayMonthlyScheduleEnum))
    for i:=0; i<len(dayMonthlyScheduleEnum);i++ {
        convArray[i] = DayMonthlyScheduleEnumToValue(dayMonthlyScheduleEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DayMonthlyScheduleEnumFromValue(value string) DayMonthlyScheduleEnum {
    switch value {
        case "kSunday":
            return DayMonthlySchedule_KSUNDAY
        case "kMonday":
            return DayMonthlySchedule_KMONDAY
        case "kTuesday":
            return DayMonthlySchedule_KTUESDAY
        case "kWednesday":
            return DayMonthlySchedule_KWEDNESDAY
        case "kThursday":
            return DayMonthlySchedule_KTHURSDAY
        case "kFriday":
            return DayMonthlySchedule_KFRIDAY
        case "kSaturday":
            return DayMonthlySchedule_KSATURDAY
        default:
            return DayMonthlySchedule_KSUNDAY
    }
}
