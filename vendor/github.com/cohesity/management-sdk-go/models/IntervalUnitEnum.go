// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for IntervalUnitEnum enum
 */
type IntervalUnitEnum int

/**
 * Value collection for IntervalUnitEnum enum
 */
const (
    IntervalUnit_KMINUTES IntervalUnitEnum = 1 + iota
    IntervalUnit_KHOURS
    IntervalUnit_KDAYS
    IntervalUnit_KWEEKS
    IntervalUnit_KMONTHS
)

func (r IntervalUnitEnum) MarshalJSON() ([]byte, error) {
    s := IntervalUnitEnumToValue(r)
    return json.Marshal(s)
}

func (r *IntervalUnitEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  IntervalUnitEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts IntervalUnitEnum to its string representation
 */
func IntervalUnitEnumToValue(intervalUnitEnum IntervalUnitEnum) string {
    switch intervalUnitEnum {
        case IntervalUnit_KMINUTES:
    		return "kMinutes"
        case IntervalUnit_KHOURS:
    		return "kHours"
        case IntervalUnit_KDAYS:
    		return "kDays"
        case IntervalUnit_KWEEKS:
    		return "kWeeks"
        case IntervalUnit_KMONTHS:
    		return "kMonths"
        default:
        	return "kMinutes"
    }
}

/**
 * Converts IntervalUnitEnum Array to its string Array representation
*/
func IntervalUnitEnumArrayToValue(intervalUnitEnum []IntervalUnitEnum) []string {
    convArray := make([]string,len( intervalUnitEnum))
    for i:=0; i<len(intervalUnitEnum);i++ {
        convArray[i] = IntervalUnitEnumToValue(intervalUnitEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func IntervalUnitEnumFromValue(value string) IntervalUnitEnum {
    switch value {
        case "kMinutes":
            return IntervalUnit_KMINUTES
        case "kHours":
            return IntervalUnit_KHOURS
        case "kDays":
            return IntervalUnit_KDAYS
        case "kWeeks":
            return IntervalUnit_KWEEKS
        case "kMonths":
            return IntervalUnit_KMONTHS
        default:
            return IntervalUnit_KMINUTES
    }
}
