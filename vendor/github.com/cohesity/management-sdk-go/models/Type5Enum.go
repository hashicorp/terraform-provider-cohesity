package models

import(
    "encoding/json"
)

/**
 * Type definition for Type5Enum enum
 */
type Type5Enum int

/**
 * Value collection for Type5Enum enum
 */
const (
    Type5_KSERVER Type5Enum = 1 + iota
)

func (r Type5Enum) MarshalJSON() ([]byte, error) { 
    s := Type5EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type5Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type5EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type5Enum to its string representation
 */
func Type5EnumToValue(type5Enum Type5Enum) string {
    switch type5Enum {
        case Type5_KSERVER:
    		return "kServer"		
        default:
        	return "kServer"
    }
}

/**
 * Converts Type5Enum Array to its string Array representation
*/
func Type5EnumArrayToValue(type5Enum []Type5Enum) []string {
    convArray := make([]string,len( type5Enum))
    for i:=0; i<len(type5Enum);i++ {
        convArray[i] = Type5EnumToValue(type5Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type5EnumFromValue(value string) Type5Enum {
    switch value {
        case "kServer":
            return Type5_KSERVER
        default:
            return Type5_KSERVER
    }
}
