package models

import(
    "encoding/json"
)

/**
 * Type definition for Type14Enum enum
 */
type Type14Enum int

/**
 * Value collection for Type14Enum enum
 */
const (
    Type14_KRACROOTCONTAINER Type14Enum = 1 + iota
    Type14_KROOTCONTAINER
    Type14_KHOST
    Type14_KDATABASE
    Type14_KTABLESPACE
    Type14_KTABLE
)

func (r Type14Enum) MarshalJSON() ([]byte, error) { 
    s := Type14EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type14Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type14EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type14Enum to its string representation
 */
func Type14EnumToValue(type14Enum Type14Enum) string {
    switch type14Enum {
        case Type14_KRACROOTCONTAINER:
    		return "kRACRootContainer"		
        case Type14_KROOTCONTAINER:
    		return "kRootContainer"		
        case Type14_KHOST:
    		return "kHost"		
        case Type14_KDATABASE:
    		return "kDatabase"		
        case Type14_KTABLESPACE:
    		return "kTableSpace"		
        case Type14_KTABLE:
    		return "kTable"		
        default:
        	return "kRACRootContainer"
    }
}

/**
 * Converts Type14Enum Array to its string Array representation
*/
func Type14EnumArrayToValue(type14Enum []Type14Enum) []string {
    convArray := make([]string,len( type14Enum))
    for i:=0; i<len(type14Enum);i++ {
        convArray[i] = Type14EnumToValue(type14Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type14EnumFromValue(value string) Type14Enum {
    switch value {
        case "kRACRootContainer":
            return Type14_KRACROOTCONTAINER
        case "kRootContainer":
            return Type14_KROOTCONTAINER
        case "kHost":
            return Type14_KHOST
        case "kDatabase":
            return Type14_KDATABASE
        case "kTableSpace":
            return Type14_KTABLESPACE
        case "kTable":
            return Type14_KTABLE
        default:
            return Type14_KRACROOTCONTAINER
    }
}
