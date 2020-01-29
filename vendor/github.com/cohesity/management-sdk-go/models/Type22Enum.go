package models

import(
    "encoding/json"
)

/**
 * Type definition for Type22Enum enum
 */
type Type22Enum int

/**
 * Value collection for Type22Enum enum
 */
const (
    Type22_KALLOW Type22Enum = 1 + iota
    Type22_KDENY
    Type22_KSPECIALTYPE
)

func (r Type22Enum) MarshalJSON() ([]byte, error) { 
    s := Type22EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type22Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type22EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type22Enum to its string representation
 */
func Type22EnumToValue(type22Enum Type22Enum) string {
    switch type22Enum {
        case Type22_KALLOW:
    		return "kAllow"		
        case Type22_KDENY:
    		return "kDeny"		
        case Type22_KSPECIALTYPE:
    		return "kSpecialType"		
        default:
        	return "kAllow"
    }
}

/**
 * Converts Type22Enum Array to its string Array representation
*/
func Type22EnumArrayToValue(type22Enum []Type22Enum) []string {
    convArray := make([]string,len( type22Enum))
    for i:=0; i<len(type22Enum);i++ {
        convArray[i] = Type22EnumToValue(type22Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type22EnumFromValue(value string) Type22Enum {
    switch value {
        case "kAllow":
            return Type22_KALLOW
        case "kDeny":
            return Type22_KDENY
        case "kSpecialType":
            return Type22_KSPECIALTYPE
        default:
            return Type22_KALLOW
    }
}
