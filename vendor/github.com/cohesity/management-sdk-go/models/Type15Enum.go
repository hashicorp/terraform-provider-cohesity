package models

import(
    "encoding/json"
)

/**
 * Type definition for Type15Enum enum
 */
type Type15Enum int

/**
 * Value collection for Type15Enum enum
 */
const (
    Type15_KHOST Type15Enum = 1 + iota
    Type15_KWINDOWSCLUSTER
)

func (r Type15Enum) MarshalJSON() ([]byte, error) { 
    s := Type15EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type15Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type15EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type15Enum to its string representation
 */
func Type15EnumToValue(type15Enum Type15Enum) string {
    switch type15Enum {
        case Type15_KHOST:
    		return "kHost"		
        case Type15_KWINDOWSCLUSTER:
    		return "kWindowsCluster"		
        default:
        	return "kHost"
    }
}

/**
 * Converts Type15Enum Array to its string Array representation
*/
func Type15EnumArrayToValue(type15Enum []Type15Enum) []string {
    convArray := make([]string,len( type15Enum))
    for i:=0; i<len(type15Enum);i++ {
        convArray[i] = Type15EnumToValue(type15Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type15EnumFromValue(value string) Type15Enum {
    switch value {
        case "kHost":
            return Type15_KHOST
        case "kWindowsCluster":
            return Type15_KWINDOWSCLUSTER
        default:
            return Type15_KHOST
    }
}
