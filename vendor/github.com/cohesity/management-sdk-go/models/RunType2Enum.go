package models

import(
    "encoding/json"
)

/**
 * Type definition for RunType2Enum enum
 */
type RunType2Enum int

/**
 * Value collection for RunType2Enum enum
 */
const (
    RunType2_KREGULAR RunType2Enum = 1 + iota
    RunType2_KFULL
    RunType2_KLOG
    RunType2_KSYSTEM
)

func (r RunType2Enum) MarshalJSON() ([]byte, error) { 
    s := RunType2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *RunType2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  RunType2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts RunType2Enum to its string representation
 */
func RunType2EnumToValue(runType2Enum RunType2Enum) string {
    switch runType2Enum {
        case RunType2_KREGULAR:
    		return "kRegular"		
        case RunType2_KFULL:
    		return "kFull"		
        case RunType2_KLOG:
    		return "kLog"		
        case RunType2_KSYSTEM:
    		return "kSystem"		
        default:
        	return "kRegular"
    }
}

/**
 * Converts RunType2Enum Array to its string Array representation
*/
func RunType2EnumArrayToValue(runType2Enum []RunType2Enum) []string {
    convArray := make([]string,len( runType2Enum))
    for i:=0; i<len(runType2Enum);i++ {
        convArray[i] = RunType2EnumToValue(runType2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RunType2EnumFromValue(value string) RunType2Enum {
    switch value {
        case "kRegular":
            return RunType2_KREGULAR
        case "kFull":
            return RunType2_KFULL
        case "kLog":
            return RunType2_KLOG
        case "kSystem":
            return RunType2_KSYSTEM
        default:
            return RunType2_KREGULAR
    }
}
