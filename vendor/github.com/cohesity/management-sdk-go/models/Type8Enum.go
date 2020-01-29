package models

import(
    "encoding/json"
)

/**
 * Type definition for Type8Enum enum
 */
type Type8Enum int

/**
 * Value collection for Type8Enum enum
 */
const (
    Type8_KCLUSTER Type8Enum = 1 + iota
    Type8_KZONE
    Type8_KMOUNTPOINT
)

func (r Type8Enum) MarshalJSON() ([]byte, error) { 
    s := Type8EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type8Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type8EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type8Enum to its string representation
 */
func Type8EnumToValue(type8Enum Type8Enum) string {
    switch type8Enum {
        case Type8_KCLUSTER:
    		return "kCluster"		
        case Type8_KZONE:
    		return "kZone"		
        case Type8_KMOUNTPOINT:
    		return "kMountPoint"		
        default:
        	return "kCluster"
    }
}

/**
 * Converts Type8Enum Array to its string Array representation
*/
func Type8EnumArrayToValue(type8Enum []Type8Enum) []string {
    convArray := make([]string,len( type8Enum))
    for i:=0; i<len(type8Enum);i++ {
        convArray[i] = Type8EnumToValue(type8Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type8EnumFromValue(value string) Type8Enum {
    switch value {
        case "kCluster":
            return Type8_KCLUSTER
        case "kZone":
            return Type8_KZONE
        case "kMountPoint":
            return Type8_KMOUNTPOINT
        default:
            return Type8_KCLUSTER
    }
}
