// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StateEnum enum
 */
type StateEnum int

/**
 * Value collection for StateEnum enum
 */
const (
    State_KONLINE StateEnum = 1 + iota
    State_KRESTRICTED
    State_KOFFLINE
    State_KMIXED
)

func (r StateEnum) MarshalJSON() ([]byte, error) {
    s := StateEnumToValue(r)
    return json.Marshal(s)
}

func (r *StateEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StateEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StateEnum to its string representation
 */
func StateEnumToValue(stateEnum StateEnum) string {
    switch stateEnum {
        case State_KONLINE:
    		return "kOnline"
        case State_KRESTRICTED:
    		return "kRestricted"
        case State_KOFFLINE:
    		return "kOffline"
        case State_KMIXED:
    		return "kMixed"
        default:
        	return "kOnline"
    }
}

/**
 * Converts StateEnum Array to its string Array representation
*/
func StateEnumArrayToValue(stateEnum []StateEnum) []string {
    convArray := make([]string,len( stateEnum))
    for i:=0; i<len(stateEnum);i++ {
        convArray[i] = StateEnumToValue(stateEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StateEnumFromValue(value string) StateEnum {
    switch value {
        case "kOnline":
            return State_KONLINE
        case "kRestricted":
            return State_KRESTRICTED
        case "kOffline":
            return State_KOFFLINE
        case "kMixed":
            return State_KMIXED
        default:
            return State_KONLINE
    }
}
