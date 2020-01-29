// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PeriodicityEnum enum
 */
type PeriodicityEnum int

/**
 * Value collection for PeriodicityEnum enum
 */
const (
    Periodicity_KCONTINUOUS PeriodicityEnum = 1 + iota
    Periodicity_KDAILY
    Periodicity_KMONTHLY
    Periodicity_KCONTINUOUSRPO
)

func (r PeriodicityEnum) MarshalJSON() ([]byte, error) {
    s := PeriodicityEnumToValue(r)
    return json.Marshal(s)
}

func (r *PeriodicityEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PeriodicityEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PeriodicityEnum to its string representation
 */
func PeriodicityEnumToValue(periodicityEnum PeriodicityEnum) string {
    switch periodicityEnum {
        case Periodicity_KCONTINUOUS:
    		return "kContinuous"
        case Periodicity_KDAILY:
    		return "kDaily"
        case Periodicity_KMONTHLY:
    		return "kMonthly"
        case Periodicity_KCONTINUOUSRPO:
    		return "kContinuousRPO"
        default:
        	return "kContinuous"
    }
}

/**
 * Converts PeriodicityEnum Array to its string Array representation
*/
func PeriodicityEnumArrayToValue(periodicityEnum []PeriodicityEnum) []string {
    convArray := make([]string,len( periodicityEnum))
    for i:=0; i<len(periodicityEnum);i++ {
        convArray[i] = PeriodicityEnumToValue(periodicityEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PeriodicityEnumFromValue(value string) PeriodicityEnum {
    switch value {
        case "kContinuous":
            return Periodicity_KCONTINUOUS
        case "kDaily":
            return Periodicity_KDAILY
        case "kMonthly":
            return Periodicity_KMONTHLY
        case "kContinuousRPO":
            return Periodicity_KCONTINUOUSRPO
        default:
            return Periodicity_KCONTINUOUS
    }
}
