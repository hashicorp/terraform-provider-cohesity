package models

import(
    "encoding/json"
)

/**
 * Type definition for Type29Enum enum
 */
type Type29Enum int

/**
 * Value collection for Type29Enum enum
 */
const (
    Type29_KREGULAR Type29Enum = 1 + iota
    Type29_KRPO
)

func (r Type29Enum) MarshalJSON() ([]byte, error) { 
    s := Type29EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type29Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type29EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type29Enum to its string representation
 */
func Type29EnumToValue(type29Enum Type29Enum) string {
    switch type29Enum {
        case Type29_KREGULAR:
    		return "kRegular"		
        case Type29_KRPO:
    		return "kRPO"		
        default:
        	return "kRegular"
    }
}

/**
 * Converts Type29Enum Array to its string Array representation
*/
func Type29EnumArrayToValue(type29Enum []Type29Enum) []string {
    convArray := make([]string,len( type29Enum))
    for i:=0; i<len(type29Enum);i++ {
        convArray[i] = Type29EnumToValue(type29Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type29EnumFromValue(value string) Type29Enum {
    switch value {
        case "kRegular":
            return Type29_KREGULAR
        case "kRPO":
            return Type29_KRPO
        default:
            return Type29_KREGULAR
    }
}
