// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RunTypeRunProtectionJobParamEnum enum
 */
type RunTypeRunProtectionJobParamEnum int

/**
 * Value collection for RunTypeRunProtectionJobParamEnum enum
 */
const (
    RunTypeRunProtectionJobParam_KREGULAR RunTypeRunProtectionJobParamEnum = 1 + iota
    RunTypeRunProtectionJobParam_KFULL
    RunTypeRunProtectionJobParam_KLOG
    RunTypeRunProtectionJobParam_KSYSTEM
)

func (r RunTypeRunProtectionJobParamEnum) MarshalJSON() ([]byte, error) {
    s := RunTypeRunProtectionJobParamEnumToValue(r)
    return json.Marshal(s)
}

func (r *RunTypeRunProtectionJobParamEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RunTypeRunProtectionJobParamEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RunTypeRunProtectionJobParamEnum to its string representation
 */
func RunTypeRunProtectionJobParamEnumToValue(runTypeRunProtectionJobParamEnum RunTypeRunProtectionJobParamEnum) string {
    switch runTypeRunProtectionJobParamEnum {
        case RunTypeRunProtectionJobParam_KREGULAR:
    		return "kRegular"
        case RunTypeRunProtectionJobParam_KFULL:
    		return "kFull"
        case RunTypeRunProtectionJobParam_KLOG:
    		return "kLog"
        case RunTypeRunProtectionJobParam_KSYSTEM:
    		return "kSystem"
        default:
        	return "kRegular"
    }
}

/**
 * Converts RunTypeRunProtectionJobParamEnum Array to its string Array representation
*/
func RunTypeRunProtectionJobParamEnumArrayToValue(runTypeRunProtectionJobParamEnum []RunTypeRunProtectionJobParamEnum) []string {
    convArray := make([]string,len( runTypeRunProtectionJobParamEnum))
    for i:=0; i<len(runTypeRunProtectionJobParamEnum);i++ {
        convArray[i] = RunTypeRunProtectionJobParamEnumToValue(runTypeRunProtectionJobParamEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RunTypeRunProtectionJobParamEnumFromValue(value string) RunTypeRunProtectionJobParamEnum {
    switch value {
        case "kRegular":
            return RunTypeRunProtectionJobParam_KREGULAR
        case "kFull":
            return RunTypeRunProtectionJobParam_KFULL
        case "kLog":
            return RunTypeRunProtectionJobParam_KLOG
        case "kSystem":
            return RunTypeRunProtectionJobParam_KSYSTEM
        default:
            return RunTypeRunProtectionJobParam_KREGULAR
    }
}
