package models

import(
    "encoding/json"
)

/**
 * Type definition for Protocol2Enum enum
 */
type Protocol2Enum int

/**
 * Value collection for Protocol2Enum enum
 */
const (
    Protocol2_KNFS3 Protocol2Enum = 1 + iota
    Protocol2_KCIFS1
)

func (r Protocol2Enum) MarshalJSON() ([]byte, error) { 
    s := Protocol2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Protocol2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Protocol2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Protocol2Enum to its string representation
 */
func Protocol2EnumToValue(protocol2Enum Protocol2Enum) string {
    switch protocol2Enum {
        case Protocol2_KNFS3:
    		return "kNfs3"		
        case Protocol2_KCIFS1:
    		return "kCifs1"		
        default:
        	return "kNfs3"
    }
}

/**
 * Converts Protocol2Enum Array to its string Array representation
*/
func Protocol2EnumArrayToValue(protocol2Enum []Protocol2Enum) []string {
    convArray := make([]string,len( protocol2Enum))
    for i:=0; i<len(protocol2Enum);i++ {
        convArray[i] = Protocol2EnumToValue(protocol2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Protocol2EnumFromValue(value string) Protocol2Enum {
    switch value {
        case "kNfs3":
            return Protocol2_KNFS3
        case "kCifs1":
            return Protocol2_KCIFS1
        default:
            return Protocol2_KNFS3
    }
}
