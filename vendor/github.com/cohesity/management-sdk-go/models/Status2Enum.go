package models

import(
    "encoding/json"
)

/**
 * Type definition for Status2Enum enum
 */
type Status2Enum int

/**
 * Value collection for Status2Enum enum
 */
const (
    Status2_KACCEPTED Status2Enum = 1 + iota
    Status2_KRUNNING
    Status2_KCANCELING
    Status2_KCANCELED
    Status2_KSUCCESS
    Status2_KFAILURE
)

func (r Status2Enum) MarshalJSON() ([]byte, error) { 
    s := Status2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Status2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Status2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Status2Enum to its string representation
 */
func Status2EnumToValue(status2Enum Status2Enum) string {
    switch status2Enum {
        case Status2_KACCEPTED:
    		return "kAccepted"		
        case Status2_KRUNNING:
    		return "kRunning"		
        case Status2_KCANCELING:
    		return "kCanceling"		
        case Status2_KCANCELED:
    		return "kCanceled"		
        case Status2_KSUCCESS:
    		return "kSuccess"		
        case Status2_KFAILURE:
    		return "kFailure"		
        default:
        	return "kAccepted"
    }
}

/**
 * Converts Status2Enum Array to its string Array representation
*/
func Status2EnumArrayToValue(status2Enum []Status2Enum) []string {
    convArray := make([]string,len( status2Enum))
    for i:=0; i<len(status2Enum);i++ {
        convArray[i] = Status2EnumToValue(status2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Status2EnumFromValue(value string) Status2Enum {
    switch value {
        case "kAccepted":
            return Status2_KACCEPTED
        case "kRunning":
            return Status2_KRUNNING
        case "kCanceling":
            return Status2_KCANCELING
        case "kCanceled":
            return Status2_KCANCELED
        case "kSuccess":
            return Status2_KSUCCESS
        case "kFailure":
            return Status2_KFAILURE
        default:
            return Status2_KACCEPTED
    }
}
