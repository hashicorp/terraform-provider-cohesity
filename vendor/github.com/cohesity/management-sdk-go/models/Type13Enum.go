package models

import(
    "encoding/json"
)

/**
 * Type definition for Type13Enum enum
 */
type Type13Enum int

/**
 * Value collection for Type13Enum enum
 */
const (
    Type13_KDATA Type13Enum = 1 + iota
    Type13_KADMIN
    Type13_KSYSTEM
    Type13_KNODE
    Type13_KUNKNOWN
)

func (r Type13Enum) MarshalJSON() ([]byte, error) { 
    s := Type13EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type13Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type13EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type13Enum to its string representation
 */
func Type13EnumToValue(type13Enum Type13Enum) string {
    switch type13Enum {
        case Type13_KDATA:
    		return "kData"		
        case Type13_KADMIN:
    		return "kAdmin"		
        case Type13_KSYSTEM:
    		return "kSystem"		
        case Type13_KNODE:
    		return "kNode"		
        case Type13_KUNKNOWN:
    		return "kUnknown"		
        default:
        	return "kData"
    }
}

/**
 * Converts Type13Enum Array to its string Array representation
*/
func Type13EnumArrayToValue(type13Enum []Type13Enum) []string {
    convArray := make([]string,len( type13Enum))
    for i:=0; i<len(type13Enum);i++ {
        convArray[i] = Type13EnumToValue(type13Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type13EnumFromValue(value string) Type13Enum {
    switch value {
        case "kData":
            return Type13_KDATA
        case "kAdmin":
            return Type13_KADMIN
        case "kSystem":
            return Type13_KSYSTEM
        case "kNode":
            return Type13_KNODE
        case "kUnknown":
            return Type13_KUNKNOWN
        default:
            return Type13_KDATA
    }
}
