package models

import(
    "encoding/json"
)

/**
 * Type definition for RunType1Enum enum
 */
type RunType1Enum int

/**
 * Value collection for RunType1Enum enum
 */
const (
    RunType1_KREGULAR RunType1Enum = 1 + iota
    RunType1_KFULL
    RunType1_KLOG
    RunType1_KSYSTEM
)

func (r RunType1Enum) MarshalJSON() ([]byte, error) { 
    s := RunType1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *RunType1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  RunType1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts RunType1Enum to its string representation
 */
func RunType1EnumToValue(runType1Enum RunType1Enum) string {
    switch runType1Enum {
        case RunType1_KREGULAR:
    		return "kRegular"		
        case RunType1_KFULL:
    		return "kFull"		
        case RunType1_KLOG:
    		return "kLog"		
        case RunType1_KSYSTEM:
    		return "kSystem"		
        default:
        	return "kRegular"
    }
}

/**
 * Converts RunType1Enum Array to its string Array representation
*/
func RunType1EnumArrayToValue(runType1Enum []RunType1Enum) []string {
    convArray := make([]string,len( runType1Enum))
    for i:=0; i<len(runType1Enum);i++ {
        convArray[i] = RunType1EnumToValue(runType1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RunType1EnumFromValue(value string) RunType1Enum {
    switch value {
        case "kRegular":
            return RunType1_KREGULAR
        case "kFull":
            return RunType1_KFULL
        case "kLog":
            return RunType1_KLOG
        case "kSystem":
            return RunType1_KSYSTEM
        default:
            return RunType1_KREGULAR
    }
}
