package models

import(
    "encoding/json"
)

/**
 * Type definition for Type30Enum enum
 */
type Type30Enum int

/**
 * Value collection for Type30Enum enum
 */
const (
    Type30_KRECOVERVMS Type30Enum = 1 + iota
    Type30_KMOUNTVOLUMES
)

func (r Type30Enum) MarshalJSON() ([]byte, error) { 
    s := Type30EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type30Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type30EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type30Enum to its string representation
 */
func Type30EnumToValue(type30Enum Type30Enum) string {
    switch type30Enum {
        case Type30_KRECOVERVMS:
    		return "kRecoverVMs"		
        case Type30_KMOUNTVOLUMES:
    		return "kMountVolumes"		
        default:
        	return "kRecoverVMs"
    }
}

/**
 * Converts Type30Enum Array to its string Array representation
*/
func Type30EnumArrayToValue(type30Enum []Type30Enum) []string {
    convArray := make([]string,len( type30Enum))
    for i:=0; i<len(type30Enum);i++ {
        convArray[i] = Type30EnumToValue(type30Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type30EnumFromValue(value string) Type30Enum {
    switch value {
        case "kRecoverVMs":
            return Type30_KRECOVERVMS
        case "kMountVolumes":
            return Type30_KMOUNTVOLUMES
        default:
            return Type30_KRECOVERVMS
    }
}
