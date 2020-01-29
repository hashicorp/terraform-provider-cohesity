package models

import(
    "encoding/json"
)

/**
 * Type definition for Type25Enum enum
 */
type Type25Enum int

/**
 * Value collection for Type25Enum enum
 */
const (
    Type25_KCLONEVMS Type25Enum = 1 + iota
    Type25_KCLONEVIEW
)

func (r Type25Enum) MarshalJSON() ([]byte, error) { 
    s := Type25EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type25Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type25EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type25Enum to its string representation
 */
func Type25EnumToValue(type25Enum Type25Enum) string {
    switch type25Enum {
        case Type25_KCLONEVMS:
    		return "kCloneVMs"		
        case Type25_KCLONEVIEW:
    		return "kCloneView"		
        default:
        	return "kCloneVMs"
    }
}

/**
 * Converts Type25Enum Array to its string Array representation
*/
func Type25EnumArrayToValue(type25Enum []Type25Enum) []string {
    convArray := make([]string,len( type25Enum))
    for i:=0; i<len(type25Enum);i++ {
        convArray[i] = Type25EnumToValue(type25Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type25EnumFromValue(value string) Type25Enum {
    switch value {
        case "kCloneVMs":
            return Type25_KCLONEVMS
        case "kCloneView":
            return Type25_KCLONEVIEW
        default:
            return Type25_KCLONEVMS
    }
}
