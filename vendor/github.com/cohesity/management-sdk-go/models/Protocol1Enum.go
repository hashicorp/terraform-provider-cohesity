package models

import(
    "encoding/json"
)

/**
 * Type definition for Protocol1Enum enum
 */
type Protocol1Enum int

/**
 * Value collection for Protocol1Enum enum
 */
const (
    Protocol1_KNFS Protocol1Enum = 1 + iota
    Protocol1_KSMB
)

func (r Protocol1Enum) MarshalJSON() ([]byte, error) { 
    s := Protocol1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Protocol1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Protocol1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Protocol1Enum to its string representation
 */
func Protocol1EnumToValue(protocol1Enum Protocol1Enum) string {
    switch protocol1Enum {
        case Protocol1_KNFS:
    		return "kNfs"		
        case Protocol1_KSMB:
    		return "kSmb"		
        default:
        	return "kNfs"
    }
}

/**
 * Converts Protocol1Enum Array to its string Array representation
*/
func Protocol1EnumArrayToValue(protocol1Enum []Protocol1Enum) []string {
    convArray := make([]string,len( protocol1Enum))
    for i:=0; i<len(protocol1Enum);i++ {
        convArray[i] = Protocol1EnumToValue(protocol1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Protocol1EnumFromValue(value string) Protocol1Enum {
    switch value {
        case "kNfs":
            return Protocol1_KNFS
        case "kSmb":
            return Protocol1_KSMB
        default:
            return Protocol1_KNFS
    }
}
