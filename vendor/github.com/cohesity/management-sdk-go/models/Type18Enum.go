package models

import(
    "encoding/json"
)

/**
 * Type definition for Type18Enum enum
 */
type Type18Enum int

/**
 * Value collection for Type18Enum enum
 */
const (
    Type18_KVIEWBOX Type18Enum = 1 + iota
    Type18_KVIEW
)

func (r Type18Enum) MarshalJSON() ([]byte, error) { 
    s := Type18EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type18Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type18EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type18Enum to its string representation
 */
func Type18EnumToValue(type18Enum Type18Enum) string {
    switch type18Enum {
        case Type18_KVIEWBOX:
    		return "kViewBox"		
        case Type18_KVIEW:
    		return "kView"		
        default:
        	return "kViewBox"
    }
}

/**
 * Converts Type18Enum Array to its string Array representation
*/
func Type18EnumArrayToValue(type18Enum []Type18Enum) []string {
    convArray := make([]string,len( type18Enum))
    for i:=0; i<len(type18Enum);i++ {
        convArray[i] = Type18EnumToValue(type18Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type18EnumFromValue(value string) Type18Enum {
    switch value {
        case "kViewBox":
            return Type18_KVIEWBOX
        case "kView":
            return Type18_KVIEW
        default:
            return Type18_KVIEWBOX
    }
}
