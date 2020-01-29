package models

import(
    "encoding/json"
)

/**
 * Type definition for Type21Enum enum
 */
type Type21Enum int

/**
 * Value collection for Type21Enum enum
 */
const (
    Type21_KAZURE Type21Enum = 1 + iota
    Type21_KAWS
    Type21_KGCP
)

func (r Type21Enum) MarshalJSON() ([]byte, error) { 
    s := Type21EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type21Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type21EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type21Enum to its string representation
 */
func Type21EnumToValue(type21Enum Type21Enum) string {
    switch type21Enum {
        case Type21_KAZURE:
    		return "kAzure"		
        case Type21_KAWS:
    		return "kAws"		
        case Type21_KGCP:
    		return "kGcp"		
        default:
        	return "kAzure"
    }
}

/**
 * Converts Type21Enum Array to its string Array representation
*/
func Type21EnumArrayToValue(type21Enum []Type21Enum) []string {
    convArray := make([]string,len( type21Enum))
    for i:=0; i<len(type21Enum);i++ {
        convArray[i] = Type21EnumToValue(type21Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type21EnumFromValue(value string) Type21Enum {
    switch value {
        case "kAzure":
            return Type21_KAZURE
        case "kAws":
            return Type21_KAWS
        case "kGcp":
            return Type21_KGCP
        default:
            return Type21_KAZURE
    }
}
