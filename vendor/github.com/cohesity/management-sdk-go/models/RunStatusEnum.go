package models

import(
    "encoding/json"
)

/**
 * Type definition for RunStatusEnum enum
 */
type RunStatusEnum int

/**
 * Value collection for RunStatusEnum enum
 */
const (
    RunStatus_KSUCCESS RunStatusEnum = 1 + iota
    RunStatus_KRUNNING
    RunStatus_KWARNING
    RunStatus_KCANCELLED
    RunStatus_KERROR
)

func (r RunStatusEnum) MarshalJSON() ([]byte, error) { 
    s := RunStatusEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *RunStatusEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  RunStatusEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts RunStatusEnum to its string representation
 */
func RunStatusEnumToValue(runStatusEnum RunStatusEnum) string {
    switch runStatusEnum {
        case RunStatus_KSUCCESS:
    		return "kSuccess"		
        case RunStatus_KRUNNING:
    		return "kRunning"		
        case RunStatus_KWARNING:
    		return "kWarning"		
        case RunStatus_KCANCELLED:
    		return "kCancelled"		
        case RunStatus_KERROR:
    		return "kError"		
        default:
        	return "kSuccess"
    }
}

/**
 * Converts RunStatusEnum Array to its string Array representation
*/
func RunStatusEnumArrayToValue(runStatusEnum []RunStatusEnum) []string {
    convArray := make([]string,len( runStatusEnum))
    for i:=0; i<len(runStatusEnum);i++ {
        convArray[i] = RunStatusEnumToValue(runStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RunStatusEnumFromValue(value string) RunStatusEnum {
    switch value {
        case "kSuccess":
            return RunStatus_KSUCCESS
        case "kRunning":
            return RunStatus_KRUNNING
        case "kWarning":
            return RunStatus_KWARNING
        case "kCancelled":
            return RunStatus_KCANCELLED
        case "kError":
            return RunStatus_KERROR
        default:
            return RunStatus_KSUCCESS
    }
}
