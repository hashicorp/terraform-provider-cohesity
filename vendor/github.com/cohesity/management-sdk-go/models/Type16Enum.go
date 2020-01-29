package models

import(
    "encoding/json"
)

/**
 * Type definition for Type16Enum enum
 */
type Type16Enum int

/**
 * Value collection for Type16Enum enum
 */
const (
    Type16_KSTORAGEARRAY Type16Enum = 1 + iota
    Type16_KVOLUME
)

func (r Type16Enum) MarshalJSON() ([]byte, error) { 
    s := Type16EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type16Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type16EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type16Enum to its string representation
 */
func Type16EnumToValue(type16Enum Type16Enum) string {
    switch type16Enum {
        case Type16_KSTORAGEARRAY:
    		return "kStorageArray"		
        case Type16_KVOLUME:
    		return "kVolume"		
        default:
        	return "kStorageArray"
    }
}

/**
 * Converts Type16Enum Array to its string Array representation
*/
func Type16EnumArrayToValue(type16Enum []Type16Enum) []string {
    convArray := make([]string,len( type16Enum))
    for i:=0; i<len(type16Enum);i++ {
        convArray[i] = Type16EnumToValue(type16Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type16EnumFromValue(value string) Type16Enum {
    switch value {
        case "kStorageArray":
            return Type16_KSTORAGEARRAY
        case "kVolume":
            return Type16_KVOLUME
        default:
            return Type16_KSTORAGEARRAY
    }
}
