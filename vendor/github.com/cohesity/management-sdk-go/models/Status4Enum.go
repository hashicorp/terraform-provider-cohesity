package models

import(
    "encoding/json"
)

/**
 * Type definition for Status4Enum enum
 */
type Status4Enum int

/**
 * Value collection for Status4Enum enum
 */
const (
    Status4_KACCEPTED Status4Enum = 1 + iota
    Status4_KRUNNING
    Status4_KCANCELING
    Status4_KCANCELED
    Status4_KSUCCESS
    Status4_KFAILURE
)

func (r Status4Enum) MarshalJSON() ([]byte, error) { 
    s := Status4EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Status4Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Status4EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Status4Enum to its string representation
 */
func Status4EnumToValue(status4Enum Status4Enum) string {
    switch status4Enum {
        case Status4_KACCEPTED:
    		return "kAccepted"		
        case Status4_KRUNNING:
    		return "kRunning"		
        case Status4_KCANCELING:
    		return "kCanceling"		
        case Status4_KCANCELED:
    		return "kCanceled"		
        case Status4_KSUCCESS:
    		return "kSuccess"		
        case Status4_KFAILURE:
    		return "kFailure"		
        default:
        	return "kAccepted"
    }
}

/**
 * Converts Status4Enum Array to its string Array representation
*/
func Status4EnumArrayToValue(status4Enum []Status4Enum) []string {
    convArray := make([]string,len( status4Enum))
    for i:=0; i<len(status4Enum);i++ {
        convArray[i] = Status4EnumToValue(status4Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Status4EnumFromValue(value string) Status4Enum {
    switch value {
        case "kAccepted":
            return Status4_KACCEPTED
        case "kRunning":
            return Status4_KRUNNING
        case "kCanceling":
            return Status4_KCANCELING
        case "kCanceled":
            return Status4_KCANCELED
        case "kSuccess":
            return Status4_KSUCCESS
        case "kFailure":
            return Status4_KFAILURE
        default:
            return Status4_KACCEPTED
    }
}
