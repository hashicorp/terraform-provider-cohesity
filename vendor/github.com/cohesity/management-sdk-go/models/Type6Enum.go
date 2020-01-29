package models

import(
    "encoding/json"
)

/**
 * Type definition for Type6Enum enum
 */
type Type6Enum int

/**
 * Value collection for Type6Enum enum
 */
const (
    Type6_KFILESHARE Type6Enum = 1 + iota
    Type6_KVOLUME
)

func (r Type6Enum) MarshalJSON() ([]byte, error) { 
    s := Type6EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type6Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type6EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type6Enum to its string representation
 */
func Type6EnumToValue(type6Enum Type6Enum) string {
    switch type6Enum {
        case Type6_KFILESHARE:
    		return "kFileShare"		
        case Type6_KVOLUME:
    		return "kVolume"		
        default:
        	return "kFileShare"
    }
}

/**
 * Converts Type6Enum Array to its string Array representation
*/
func Type6EnumArrayToValue(type6Enum []Type6Enum) []string {
    convArray := make([]string,len( type6Enum))
    for i:=0; i<len(type6Enum);i++ {
        convArray[i] = Type6EnumToValue(type6Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type6EnumFromValue(value string) Type6Enum {
    switch value {
        case "kFileShare":
            return Type6_KFILESHARE
        case "kVolume":
            return Type6_KVOLUME
        default:
            return Type6_KFILESHARE
    }
}
