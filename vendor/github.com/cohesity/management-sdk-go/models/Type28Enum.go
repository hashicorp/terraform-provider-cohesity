package models

import(
    "encoding/json"
)

/**
 * Type definition for Type28Enum enum
 */
type Type28Enum int

/**
 * Value collection for Type28Enum enum
 */
const (
    Type28_LOCAL Type28Enum = 1 + iota
    Type28_ARCHIVE
)

func (r Type28Enum) MarshalJSON() ([]byte, error) { 
    s := Type28EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type28Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type28EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type28Enum to its string representation
 */
func Type28EnumToValue(type28Enum Type28Enum) string {
    switch type28Enum {
        case Type28_LOCAL:
    		return "local"		
        case Type28_ARCHIVE:
    		return "archive"		
        default:
        	return "local"
    }
}

/**
 * Converts Type28Enum Array to its string Array representation
*/
func Type28EnumArrayToValue(type28Enum []Type28Enum) []string {
    convArray := make([]string,len( type28Enum))
    for i:=0; i<len(type28Enum);i++ {
        convArray[i] = Type28EnumToValue(type28Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type28EnumFromValue(value string) Type28Enum {
    switch value {
        case "local":
            return Type28_LOCAL
        case "archive":
            return Type28_ARCHIVE
        default:
            return Type28_LOCAL
    }
}
