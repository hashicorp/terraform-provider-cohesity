// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PriorityQoSPolicyEnum enum
 */
type PriorityQoSPolicyEnum int

/**
 * Value collection for PriorityQoSPolicyEnum enum
 */
const (
    PriorityQoSPolicy_KLOW PriorityQoSPolicyEnum = 1 + iota
    PriorityQoSPolicy_KHIGH
)

func (r PriorityQoSPolicyEnum) MarshalJSON() ([]byte, error) {
    s := PriorityQoSPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *PriorityQoSPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PriorityQoSPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PriorityQoSPolicyEnum to its string representation
 */
func PriorityQoSPolicyEnumToValue(priorityQoSPolicyEnum PriorityQoSPolicyEnum) string {
    switch priorityQoSPolicyEnum {
        case PriorityQoSPolicy_KLOW:
    		return "kLow"
        case PriorityQoSPolicy_KHIGH:
    		return "kHigh"
        default:
        	return "kLow"
    }
}

/**
 * Converts PriorityQoSPolicyEnum Array to its string Array representation
*/
func PriorityQoSPolicyEnumArrayToValue(priorityQoSPolicyEnum []PriorityQoSPolicyEnum) []string {
    convArray := make([]string,len( priorityQoSPolicyEnum))
    for i:=0; i<len(priorityQoSPolicyEnum);i++ {
        convArray[i] = PriorityQoSPolicyEnumToValue(priorityQoSPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PriorityQoSPolicyEnumFromValue(value string) PriorityQoSPolicyEnum {
    switch value {
        case "kLow":
            return PriorityQoSPolicy_KLOW
        case "kHigh":
            return PriorityQoSPolicy_KHIGH
        default:
            return PriorityQoSPolicy_KLOW
    }
}
