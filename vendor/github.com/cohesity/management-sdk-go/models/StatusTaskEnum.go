// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusTaskEnum enum
 */
type StatusTaskEnum int

/**
 * Value collection for StatusTaskEnum enum
 */
const (
    StatusTask_KACTIVE StatusTaskEnum = 1 + iota
    StatusTask_KFINISHED
    StatusTask_KFINISHEDWITHERROR
    StatusTask_KCANCELLED
    StatusTask_KFINISHEDGARBAGECOLLECTED
)

func (r StatusTaskEnum) MarshalJSON() ([]byte, error) {
    s := StatusTaskEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusTaskEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusTaskEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusTaskEnum to its string representation
 */
func StatusTaskEnumToValue(statusTaskEnum StatusTaskEnum) string {
    switch statusTaskEnum {
        case StatusTask_KACTIVE:
    		return "kActive"
        case StatusTask_KFINISHED:
    		return "kFinished"
        case StatusTask_KFINISHEDWITHERROR:
    		return "kFinishedWithError"
        case StatusTask_KCANCELLED:
    		return "kCancelled"
        case StatusTask_KFINISHEDGARBAGECOLLECTED:
    		return "kFinishedGarbageCollected"
        default:
        	return "kActive"
    }
}

/**
 * Converts StatusTaskEnum Array to its string Array representation
*/
func StatusTaskEnumArrayToValue(statusTaskEnum []StatusTaskEnum) []string {
    convArray := make([]string,len( statusTaskEnum))
    for i:=0; i<len(statusTaskEnum);i++ {
        convArray[i] = StatusTaskEnumToValue(statusTaskEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusTaskEnumFromValue(value string) StatusTaskEnum {
    switch value {
        case "kActive":
            return StatusTask_KACTIVE
        case "kFinished":
            return StatusTask_KFINISHED
        case "kFinishedWithError":
            return StatusTask_KFINISHEDWITHERROR
        case "kCancelled":
            return StatusTask_KCANCELLED
        case "kFinishedGarbageCollected":
            return StatusTask_KFINISHEDGARBAGECOLLECTED
        default:
            return StatusTask_KACTIVE
    }
}
