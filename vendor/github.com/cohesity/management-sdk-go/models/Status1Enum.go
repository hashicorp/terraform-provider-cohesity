package models

import(
    "encoding/json"
)

/**
 * Type definition for Status1Enum enum
 */
type Status1Enum int

/**
 * Value collection for Status1Enum enum
 */
const (
    Status1_KACCEPTED Status1Enum = 1 + iota
    Status1_KRUNNING
    Status1_KCANCELING
    Status1_KCANCELED
    Status1_KSUCCESS
    Status1_KFAILURE
)

func (r Status1Enum) MarshalJSON() ([]byte, error) { 
    s := Status1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Status1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Status1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Status1Enum to its string representation
 */
func Status1EnumToValue(status1Enum Status1Enum) string {
    switch status1Enum {
        case Status1_KACCEPTED:
    		return "kAccepted"		
        case Status1_KRUNNING:
    		return "kRunning"		
        case Status1_KCANCELING:
    		return "kCanceling"		
        case Status1_KCANCELED:
    		return "kCanceled"		
        case Status1_KSUCCESS:
    		return "kSuccess"		
        case Status1_KFAILURE:
    		return "kFailure"		
        default:
        	return "kAccepted"
    }
}

/**
 * Converts Status1Enum Array to its string Array representation
*/
func Status1EnumArrayToValue(status1Enum []Status1Enum) []string {
    convArray := make([]string,len( status1Enum))
    for i:=0; i<len(status1Enum);i++ {
        convArray[i] = Status1EnumToValue(status1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Status1EnumFromValue(value string) Status1Enum {
    switch value {
        case "kAccepted":
            return Status1_KACCEPTED
        case "kRunning":
            return Status1_KRUNNING
        case "kCanceling":
            return Status1_KCANCELING
        case "kCanceled":
            return Status1_KCANCELED
        case "kSuccess":
            return Status1_KSUCCESS
        case "kFailure":
            return Status1_KFAILURE
        default:
            return Status1_KACCEPTED
    }
}
