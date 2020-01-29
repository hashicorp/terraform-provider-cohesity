package models

import(
    "encoding/json"
)

/**
 * Type definition for LastRunStatusEnum enum
 */
type LastRunStatusEnum int

/**
 * Value collection for LastRunStatusEnum enum
 */
const (
    LastRunStatus_KSUCCESS LastRunStatusEnum = 1 + iota
    LastRunStatus_KRUNNING
    LastRunStatus_KWARNING
    LastRunStatus_KCANCELLED
    LastRunStatus_KERROR
)

func (r LastRunStatusEnum) MarshalJSON() ([]byte, error) { 
    s := LastRunStatusEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *LastRunStatusEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  LastRunStatusEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts LastRunStatusEnum to its string representation
 */
func LastRunStatusEnumToValue(lastRunStatusEnum LastRunStatusEnum) string {
    switch lastRunStatusEnum {
        case LastRunStatus_KSUCCESS:
    		return "kSuccess"		
        case LastRunStatus_KRUNNING:
    		return "kRunning"		
        case LastRunStatus_KWARNING:
    		return "kWarning"		
        case LastRunStatus_KCANCELLED:
    		return "kCancelled"		
        case LastRunStatus_KERROR:
    		return "kError"		
        default:
        	return "kSuccess"
    }
}

/**
 * Converts LastRunStatusEnum Array to its string Array representation
*/
func LastRunStatusEnumArrayToValue(lastRunStatusEnum []LastRunStatusEnum) []string {
    convArray := make([]string,len( lastRunStatusEnum))
    for i:=0; i<len(lastRunStatusEnum);i++ {
        convArray[i] = LastRunStatusEnumToValue(lastRunStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func LastRunStatusEnumFromValue(value string) LastRunStatusEnum {
    switch value {
        case "kSuccess":
            return LastRunStatus_KSUCCESS
        case "kRunning":
            return LastRunStatus_KRUNNING
        case "kWarning":
            return LastRunStatus_KWARNING
        case "kCancelled":
            return LastRunStatus_KCANCELLED
        case "kError":
            return LastRunStatus_KERROR
        default:
            return LastRunStatus_KSUCCESS
    }
}
