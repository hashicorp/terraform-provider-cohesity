package models

import(
    "encoding/json"
)

/**
 * Type definition for Status5Enum enum
 */
type Status5Enum int

/**
 * Value collection for Status5Enum enum
 */
const (
    Status5_KSUCCESS Status5Enum = 1 + iota
    Status5_KERROR
)

func (r Status5Enum) MarshalJSON() ([]byte, error) { 
    s := Status5EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Status5Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Status5EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Status5Enum to its string representation
 */
func Status5EnumToValue(status5Enum Status5Enum) string {
    switch status5Enum {
        case Status5_KSUCCESS:
    		return "kSuccess"		
        case Status5_KERROR:
    		return "kError"		
        default:
        	return "kSuccess"
    }
}

/**
 * Converts Status5Enum Array to its string Array representation
*/
func Status5EnumArrayToValue(status5Enum []Status5Enum) []string {
    convArray := make([]string,len( status5Enum))
    for i:=0; i<len(status5Enum);i++ {
        convArray[i] = Status5EnumToValue(status5Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Status5EnumFromValue(value string) Status5Enum {
    switch value {
        case "kSuccess":
            return Status5_KSUCCESS
        case "kError":
            return Status5_KERROR
        default:
            return Status5_KSUCCESS
    }
}
