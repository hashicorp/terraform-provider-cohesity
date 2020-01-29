// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StateServiceStateResultEnum enum
 */
type StateServiceStateResultEnum int

/**
 * Value collection for StateServiceStateResultEnum enum
 */
const (
    StateServiceStateResult_KSERVICESTOPPED StateServiceStateResultEnum = 1 + iota
    StateServiceStateResult_KSERVICERUNNING
    StateServiceStateResult_KSERVICERESTARTING
)

func (r StateServiceStateResultEnum) MarshalJSON() ([]byte, error) {
    s := StateServiceStateResultEnumToValue(r)
    return json.Marshal(s)
}

func (r *StateServiceStateResultEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StateServiceStateResultEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StateServiceStateResultEnum to its string representation
 */
func StateServiceStateResultEnumToValue(stateServiceStateResultEnum StateServiceStateResultEnum) string {
    switch stateServiceStateResultEnum {
        case StateServiceStateResult_KSERVICESTOPPED:
    		return "kServiceStopped"
        case StateServiceStateResult_KSERVICERUNNING:
    		return "kServiceRunning"
        case StateServiceStateResult_KSERVICERESTARTING:
    		return "kServiceRestarting"
        default:
        	return "kServiceStopped"
    }
}

/**
 * Converts StateServiceStateResultEnum Array to its string Array representation
*/
func StateServiceStateResultEnumArrayToValue(stateServiceStateResultEnum []StateServiceStateResultEnum) []string {
    convArray := make([]string,len( stateServiceStateResultEnum))
    for i:=0; i<len(stateServiceStateResultEnum);i++ {
        convArray[i] = StateServiceStateResultEnumToValue(stateServiceStateResultEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StateServiceStateResultEnumFromValue(value string) StateServiceStateResultEnum {
    switch value {
        case "kServiceStopped":
            return StateServiceStateResult_KSERVICESTOPPED
        case "kServiceRunning":
            return StateServiceStateResult_KSERVICERUNNING
        case "kServiceRestarting":
            return StateServiceStateResult_KSERVICERESTARTING
        default:
            return StateServiceStateResult_KSERVICESTOPPED
    }
}
