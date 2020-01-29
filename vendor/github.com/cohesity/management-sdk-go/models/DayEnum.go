// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for DayEnum enum
 */
type DayEnum int

/**
 * Value collection for DayEnum enum
 */
const (
    Day_KSUNDAY DayEnum = 1 + iota
    Day_KMONDAY
    Day_KTUESDAY
    Day_KWEDNESDAY
    Day_KTHURSDAY
    Day_KFRIDAY
    Day_KSATURDAY
)

func (r DayEnum) MarshalJSON() ([]byte, error) {
    s := DayEnumToValue(r)
    return json.Marshal(s)
}

func (r *DayEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  DayEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts DayEnum to its string representation
 */
func DayEnumToValue(dayEnum DayEnum) string {
    switch dayEnum {
        case Day_KSUNDAY:
    		return "kSunday"
        case Day_KMONDAY:
    		return "kMonday"
        case Day_KTUESDAY:
    		return "kTuesday"
        case Day_KWEDNESDAY:
    		return "kWednesday"
        case Day_KTHURSDAY:
    		return "kThursday"
        case Day_KFRIDAY:
    		return "kFriday"
        case Day_KSATURDAY:
    		return "kSaturday"
        default:
        	return "kSunday"
    }
}

/**
 * Converts DayEnum Array to its string Array representation
*/
func DayEnumArrayToValue(dayEnum []DayEnum) []string {
    convArray := make([]string,len( dayEnum))
    for i:=0; i<len(dayEnum);i++ {
        convArray[i] = DayEnumToValue(dayEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DayEnumFromValue(value string) DayEnum {
    switch value {
        case "kSunday":
            return Day_KSUNDAY
        case "kMonday":
            return Day_KMONDAY
        case "kTuesday":
            return Day_KTUESDAY
        case "kWednesday":
            return Day_KWEDNESDAY
        case "kThursday":
            return Day_KTHURSDAY
        case "kFriday":
            return Day_KFRIDAY
        case "kSaturday":
            return Day_KSATURDAY
        default:
            return Day_KSUNDAY
    }
}
