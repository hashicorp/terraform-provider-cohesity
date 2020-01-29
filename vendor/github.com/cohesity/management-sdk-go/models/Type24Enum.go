package models

import(
    "encoding/json"
)

/**
 * Type definition for Type24Enum enum
 */
type Type24Enum int

/**
 * Value collection for Type24Enum enum
 */
const (
    Type24_KLINUX Type24Enum = 1 + iota
    Type24_KWINDOWS
    Type24_KAIX
    Type24_KSOLARIS
)

func (r Type24Enum) MarshalJSON() ([]byte, error) { 
    s := Type24EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type24Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type24EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type24Enum to its string representation
 */
func Type24EnumToValue(type24Enum Type24Enum) string {
    switch type24Enum {
        case Type24_KLINUX:
    		return "kLinux"		
        case Type24_KWINDOWS:
    		return "kWindows"		
        case Type24_KAIX:
    		return "kAix"		
        case Type24_KSOLARIS:
    		return "kSolaris"		
        default:
        	return "kLinux"
    }
}

/**
 * Converts Type24Enum Array to its string Array representation
*/
func Type24EnumArrayToValue(type24Enum []Type24Enum) []string {
    convArray := make([]string,len( type24Enum))
    for i:=0; i<len(type24Enum);i++ {
        convArray[i] = Type24EnumToValue(type24Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type24EnumFromValue(value string) Type24Enum {
    switch value {
        case "kLinux":
            return Type24_KLINUX
        case "kWindows":
            return Type24_KWINDOWS
        case "kAix":
            return Type24_KAIX
        case "kSolaris":
            return Type24_KSOLARIS
        default:
            return Type24_KLINUX
    }
}
