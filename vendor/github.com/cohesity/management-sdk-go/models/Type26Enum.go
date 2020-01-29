package models

import(
    "encoding/json"
)

/**
 * Type definition for Type26Enum enum
 */
type Type26Enum int

/**
 * Value collection for Type26Enum enum
 */
const (
    Type26_KLOCAL Type26Enum = 1 + iota
    Type26_KREMOTE
    Type26_KARCHIVAL
    Type26_KCLOUDDEPLOY
)

func (r Type26Enum) MarshalJSON() ([]byte, error) { 
    s := Type26EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type26Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type26EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type26Enum to its string representation
 */
func Type26EnumToValue(type26Enum Type26Enum) string {
    switch type26Enum {
        case Type26_KLOCAL:
    		return "kLocal"		
        case Type26_KREMOTE:
    		return "kRemote"		
        case Type26_KARCHIVAL:
    		return "kArchival"		
        case Type26_KCLOUDDEPLOY:
    		return "kCloudDeploy"		
        default:
        	return "kLocal"
    }
}

/**
 * Converts Type26Enum Array to its string Array representation
*/
func Type26EnumArrayToValue(type26Enum []Type26Enum) []string {
    convArray := make([]string,len( type26Enum))
    for i:=0; i<len(type26Enum);i++ {
        convArray[i] = Type26EnumToValue(type26Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type26EnumFromValue(value string) Type26Enum {
    switch value {
        case "kLocal":
            return Type26_KLOCAL
        case "kRemote":
            return Type26_KREMOTE
        case "kArchival":
            return Type26_KARCHIVAL
        case "kCloudDeploy":
            return Type26_KCLOUDDEPLOY
        default:
            return Type26_KLOCAL
    }
}
