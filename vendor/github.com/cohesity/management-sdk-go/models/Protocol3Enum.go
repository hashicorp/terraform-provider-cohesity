package models

import(
    "encoding/json"
)

/**
 * Type definition for Protocol3Enum enum
 */
type Protocol3Enum int

/**
 * Value collection for Protocol3Enum enum
 */
const (
    Protocol3_KUDP Protocol3Enum = 1 + iota
    Protocol3_KTCP
)

func (r Protocol3Enum) MarshalJSON() ([]byte, error) { 
    s := Protocol3EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Protocol3Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Protocol3EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Protocol3Enum to its string representation
 */
func Protocol3EnumToValue(protocol3Enum Protocol3Enum) string {
    switch protocol3Enum {
        case Protocol3_KUDP:
    		return "kUDP"		
        case Protocol3_KTCP:
    		return "kTCP"		
        default:
        	return "kUDP"
    }
}

/**
 * Converts Protocol3Enum Array to its string Array representation
*/
func Protocol3EnumArrayToValue(protocol3Enum []Protocol3Enum) []string {
    convArray := make([]string,len( protocol3Enum))
    for i:=0; i<len(protocol3Enum);i++ {
        convArray[i] = Protocol3EnumToValue(protocol3Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Protocol3EnumFromValue(value string) Protocol3Enum {
    switch value {
        case "kUDP":
            return Protocol3_KUDP
        case "kTCP":
            return Protocol3_KTCP
        default:
            return Protocol3_KUDP
    }
}
