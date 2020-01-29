package models

import(
    "encoding/json"
)

/**
 * Type definition for ObjectClass4Enum enum
 */
type ObjectClass4Enum int

/**
 * Value collection for ObjectClass4Enum enum
 */
const (
    ObjectClass4_KUSER ObjectClass4Enum = 1 + iota
    ObjectClass4_KGROUP
    ObjectClass4_KCOMPUTER
)

func (r ObjectClass4Enum) MarshalJSON() ([]byte, error) { 
    s := ObjectClass4EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *ObjectClass4Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  ObjectClass4EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ObjectClass4Enum to its string representation
 */
func ObjectClass4EnumToValue(objectClass4Enum ObjectClass4Enum) string {
    switch objectClass4Enum {
        case ObjectClass4_KUSER:
    		return "kUser"		
        case ObjectClass4_KGROUP:
    		return "kGroup"		
        case ObjectClass4_KCOMPUTER:
    		return "kComputer"		
        default:
        	return "kUser"
    }
}

/**
 * Converts ObjectClass4Enum Array to its string Array representation
*/
func ObjectClass4EnumArrayToValue(objectClass4Enum []ObjectClass4Enum) []string {
    convArray := make([]string,len( objectClass4Enum))
    for i:=0; i<len(objectClass4Enum);i++ {
        convArray[i] = ObjectClass4EnumToValue(objectClass4Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ObjectClass4EnumFromValue(value string) ObjectClass4Enum {
    switch value {
        case "kUser":
            return ObjectClass4_KUSER
        case "kGroup":
            return ObjectClass4_KGROUP
        case "kComputer":
            return ObjectClass4_KCOMPUTER
        default:
            return ObjectClass4_KUSER
    }
}
