package models

import(
    "encoding/json"
)

/**
 * Type definition for Type10Enum enum
 */
type Type10Enum int

/**
 * Value collection for Type10Enum enum
 */
const (
    Type10_KGROUP Type10Enum = 1 + iota
    Type10_KHOST
    Type10_KDFSGROUP
    Type10_KDFSTOPDIR
)

func (r Type10Enum) MarshalJSON() ([]byte, error) { 
    s := Type10EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type10Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type10EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type10Enum to its string representation
 */
func Type10EnumToValue(type10Enum Type10Enum) string {
    switch type10Enum {
        case Type10_KGROUP:
    		return "kGroup"		
        case Type10_KHOST:
    		return "kHost"		
        case Type10_KDFSGROUP:
    		return "kDfsGroup"		
        case Type10_KDFSTOPDIR:
    		return "kDfsTopDir"		
        default:
        	return "kGroup"
    }
}

/**
 * Converts Type10Enum Array to its string Array representation
*/
func Type10EnumArrayToValue(type10Enum []Type10Enum) []string {
    convArray := make([]string,len( type10Enum))
    for i:=0; i<len(type10Enum);i++ {
        convArray[i] = Type10EnumToValue(type10Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type10EnumFromValue(value string) Type10Enum {
    switch value {
        case "kGroup":
            return Type10_KGROUP
        case "kHost":
            return Type10_KHOST
        case "kDfsGroup":
            return Type10_KDFSGROUP
        case "kDfsTopDir":
            return Type10_KDFSTOPDIR
        default:
            return Type10_KGROUP
    }
}
