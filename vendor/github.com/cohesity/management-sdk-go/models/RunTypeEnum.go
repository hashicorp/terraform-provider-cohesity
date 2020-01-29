// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RunTypeEnum enum
 */
type RunTypeEnum int

/**
 * Value collection for RunTypeEnum enum
 */
const (
    RunType_KREGULAR RunTypeEnum = 1 + iota
    RunType_KFULL
    RunType_KLOG
    RunType_KSYSTEM
)

func (r RunTypeEnum) MarshalJSON() ([]byte, error) {
    s := RunTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *RunTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RunTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RunTypeEnum to its string representation
 */
func RunTypeEnumToValue(runTypeEnum RunTypeEnum) string {
    switch runTypeEnum {
        case RunType_KREGULAR:
    		return "kRegular"
        case RunType_KFULL:
    		return "kFull"
        case RunType_KLOG:
    		return "kLog"
        case RunType_KSYSTEM:
    		return "kSystem"
        default:
        	return "kRegular"
    }
}

/**
 * Converts RunTypeEnum Array to its string Array representation
*/
func RunTypeEnumArrayToValue(runTypeEnum []RunTypeEnum) []string {
    convArray := make([]string,len( runTypeEnum))
    for i:=0; i<len(runTypeEnum);i++ {
        convArray[i] = RunTypeEnumToValue(runTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RunTypeEnumFromValue(value string) RunTypeEnum {
    switch value {
        case "kRegular":
            return RunType_KREGULAR
        case "kFull":
            return RunType_KFULL
        case "kLog":
            return RunType_KLOG
        case "kSystem":
            return RunType_KSYSTEM
        default:
            return RunType_KREGULAR
    }
}
