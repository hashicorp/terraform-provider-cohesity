// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for IntervalEnum enum
 */
type IntervalEnum int

/**
 * Value collection for IntervalEnum enum
 */
const (
    Interval_KHOUR IntervalEnum = 1 + iota
    Interval_KDAY
    Interval_KWEEK
)

func (r IntervalEnum) MarshalJSON() ([]byte, error) {
    s := IntervalEnumToValue(r)
    return json.Marshal(s)
}

func (r *IntervalEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  IntervalEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts IntervalEnum to its string representation
 */
func IntervalEnumToValue(intervalEnum IntervalEnum) string {
    switch intervalEnum {
        case Interval_KHOUR:
    		return "kHour"
        case Interval_KDAY:
    		return "kDay"
        case Interval_KWEEK:
    		return "kWeek"
        default:
        	return "kHour"
    }
}

/**
 * Converts IntervalEnum Array to its string Array representation
*/
func IntervalEnumArrayToValue(intervalEnum []IntervalEnum) []string {
    convArray := make([]string,len( intervalEnum))
    for i:=0; i<len(intervalEnum);i++ {
        convArray[i] = IntervalEnumToValue(intervalEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func IntervalEnumFromValue(value string) IntervalEnum {
    switch value {
        case "kHour":
            return Interval_KHOUR
        case "kDay":
            return Interval_KDAY
        case "kWeek":
            return Interval_KWEEK
        default:
            return Interval_KHOUR
    }
}
