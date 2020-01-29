package models

import(
    "encoding/json"
)

/**
 * Type definition for ObjectClass1Enum enum
 */
type ObjectClass1Enum int

/**
 * Value collection for ObjectClass1Enum enum
 */
const (
    ObjectClass1_KUSER ObjectClass1Enum = 1 + iota
    ObjectClass1_KGROUP
    ObjectClass1_KCOMPUTER
)

func (r ObjectClass1Enum) MarshalJSON() ([]byte, error) { 
    s := ObjectClass1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *ObjectClass1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  ObjectClass1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ObjectClass1Enum to its string representation
 */
func ObjectClass1EnumToValue(objectClass1Enum ObjectClass1Enum) string {
    switch objectClass1Enum {
        case ObjectClass1_KUSER:
    		return "kUser"		
        case ObjectClass1_KGROUP:
    		return "kGroup"		
        case ObjectClass1_KCOMPUTER:
    		return "kComputer"		
        default:
        	return "kUser"
    }
}

/**
 * Converts ObjectClass1Enum Array to its string Array representation
*/
func ObjectClass1EnumArrayToValue(objectClass1Enum []ObjectClass1Enum) []string {
    convArray := make([]string,len( objectClass1Enum))
    for i:=0; i<len(objectClass1Enum);i++ {
        convArray[i] = ObjectClass1EnumToValue(objectClass1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ObjectClass1EnumFromValue(value string) ObjectClass1Enum {
    switch value {
        case "kUser":
            return ObjectClass1_KUSER
        case "kGroup":
            return ObjectClass1_KGROUP
        case "kComputer":
            return ObjectClass1_KCOMPUTER
        default:
            return ObjectClass1_KUSER
    }
}
