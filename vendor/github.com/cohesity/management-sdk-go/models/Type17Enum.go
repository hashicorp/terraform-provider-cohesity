package models

import(
    "encoding/json"
)

/**
 * Type definition for Type17Enum enum
 */
type Type17Enum int

/**
 * Value collection for Type17Enum enum
 */
const (
    Type17_KINSTANCE Type17Enum = 1 + iota
    Type17_KDATABASE
    Type17_KAAG
    Type17_KAAGROOTCONTAINER
    Type17_KROOTCONTAINER
)

func (r Type17Enum) MarshalJSON() ([]byte, error) { 
    s := Type17EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type17Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type17EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type17Enum to its string representation
 */
func Type17EnumToValue(type17Enum Type17Enum) string {
    switch type17Enum {
        case Type17_KINSTANCE:
    		return "kInstance"		
        case Type17_KDATABASE:
    		return "kDatabase"		
        case Type17_KAAG:
    		return "kAAG"		
        case Type17_KAAGROOTCONTAINER:
    		return "kAAGRootContainer"		
        case Type17_KROOTCONTAINER:
    		return "kRootContainer"		
        default:
        	return "kInstance"
    }
}

/**
 * Converts Type17Enum Array to its string Array representation
*/
func Type17EnumArrayToValue(type17Enum []Type17Enum) []string {
    convArray := make([]string,len( type17Enum))
    for i:=0; i<len(type17Enum);i++ {
        convArray[i] = Type17EnumToValue(type17Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type17EnumFromValue(value string) Type17Enum {
    switch value {
        case "kInstance":
            return Type17_KINSTANCE
        case "kDatabase":
            return Type17_KDATABASE
        case "kAAG":
            return Type17_KAAG
        case "kAAGRootContainer":
            return Type17_KAAGROOTCONTAINER
        case "kRootContainer":
            return Type17_KROOTCONTAINER
        default:
            return Type17_KINSTANCE
    }
}
