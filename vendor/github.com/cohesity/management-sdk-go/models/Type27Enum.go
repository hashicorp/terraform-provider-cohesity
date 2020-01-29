package models

import(
    "encoding/json"
)

/**
 * Type definition for Type27Enum enum
 */
type Type27Enum int

/**
 * Value collection for Type27Enum enum
 */
const (
    Type27_KDIRECTORY Type27Enum = 1 + iota
    Type27_KFILE
    Type27_KEMAIL
    Type27_KSYMLINK
)

func (r Type27Enum) MarshalJSON() ([]byte, error) { 
    s := Type27EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type27Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type27EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type27Enum to its string representation
 */
func Type27EnumToValue(type27Enum Type27Enum) string {
    switch type27Enum {
        case Type27_KDIRECTORY:
    		return "kDirectory"		
        case Type27_KFILE:
    		return "kFile"		
        case Type27_KEMAIL:
    		return "kEmail"		
        case Type27_KSYMLINK:
    		return "kSymlink"		
        default:
        	return "kDirectory"
    }
}

/**
 * Converts Type27Enum Array to its string Array representation
*/
func Type27EnumArrayToValue(type27Enum []Type27Enum) []string {
    convArray := make([]string,len( type27Enum))
    for i:=0; i<len(type27Enum);i++ {
        convArray[i] = Type27EnumToValue(type27Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type27EnumFromValue(value string) Type27Enum {
    switch value {
        case "kDirectory":
            return Type27_KDIRECTORY
        case "kFile":
            return Type27_KFILE
        case "kEmail":
            return Type27_KEMAIL
        case "kSymlink":
            return Type27_KSYMLINK
        default:
            return Type27_KDIRECTORY
    }
}
