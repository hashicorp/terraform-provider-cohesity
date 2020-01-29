// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PriorityEnum enum
 */
type PriorityEnum int

/**
 * Value collection for PriorityEnum enum
 */
const (
    Priority_KLOW PriorityEnum = 1 + iota
    Priority_KMEDIUM
    Priority_KHIGH
)

func (r PriorityEnum) MarshalJSON() ([]byte, error) {
    s := PriorityEnumToValue(r)
    return json.Marshal(s)
}

func (r *PriorityEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PriorityEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PriorityEnum to its string representation
 */
func PriorityEnumToValue(priorityEnum PriorityEnum) string {
    switch priorityEnum {
        case Priority_KLOW:
    		return "kLow"
        case Priority_KMEDIUM:
    		return "kMedium"
        case Priority_KHIGH:
    		return "kHigh"
        default:
        	return "kLow"
    }
}

/**
 * Converts PriorityEnum Array to its string Array representation
*/
func PriorityEnumArrayToValue(priorityEnum []PriorityEnum) []string {
    convArray := make([]string,len( priorityEnum))
    for i:=0; i<len(priorityEnum);i++ {
        convArray[i] = PriorityEnumToValue(priorityEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PriorityEnumFromValue(value string) PriorityEnum {
    switch value {
        case "kLow":
            return Priority_KLOW
        case "kMedium":
            return Priority_KMEDIUM
        case "kHigh":
            return Priority_KHIGH
        default:
            return Priority_KLOW
    }
}
