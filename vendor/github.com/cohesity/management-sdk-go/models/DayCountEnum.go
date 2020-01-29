// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for DayCountEnum enum
 */
type DayCountEnum int

/**
 * Value collection for DayCountEnum enum
 */
const (
    DayCount_KFIRST DayCountEnum = 1 + iota
    DayCount_KSECOND
    DayCount_KTHIRD
    DayCount_KFOURTH
    DayCount_KLAST
)

func (r DayCountEnum) MarshalJSON() ([]byte, error) {
    s := DayCountEnumToValue(r)
    return json.Marshal(s)
}

func (r *DayCountEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  DayCountEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts DayCountEnum to its string representation
 */
func DayCountEnumToValue(dayCountEnum DayCountEnum) string {
    switch dayCountEnum {
        case DayCount_KFIRST:
    		return "kFirst"
        case DayCount_KSECOND:
    		return "kSecond"
        case DayCount_KTHIRD:
    		return "kThird"
        case DayCount_KFOURTH:
    		return "kFourth"
        case DayCount_KLAST:
    		return "kLast"
        default:
        	return "kFirst"
    }
}

/**
 * Converts DayCountEnum Array to its string Array representation
*/
func DayCountEnumArrayToValue(dayCountEnum []DayCountEnum) []string {
    convArray := make([]string,len( dayCountEnum))
    for i:=0; i<len(dayCountEnum);i++ {
        convArray[i] = DayCountEnumToValue(dayCountEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DayCountEnumFromValue(value string) DayCountEnum {
    switch value {
        case "kFirst":
            return DayCount_KFIRST
        case "kSecond":
            return DayCount_KSECOND
        case "kThird":
            return DayCount_KTHIRD
        case "kFourth":
            return DayCount_KFOURTH
        case "kLast":
            return DayCount_KLAST
        default:
            return DayCount_KFIRST
    }
}
