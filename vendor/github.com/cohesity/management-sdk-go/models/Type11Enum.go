package models

import(
    "encoding/json"
)

/**
 * Type definition for Type11Enum enum
 */
type Type11Enum int

/**
 * Value collection for Type11Enum enum
 */
const (
    Type11_KCLUSTER Type11Enum = 1 + iota
    Type11_KVSERVER
    Type11_KVOLUME
)

func (r Type11Enum) MarshalJSON() ([]byte, error) { 
    s := Type11EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type11Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type11EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type11Enum to its string representation
 */
func Type11EnumToValue(type11Enum Type11Enum) string {
    switch type11Enum {
        case Type11_KCLUSTER:
    		return "kCluster"		
        case Type11_KVSERVER:
    		return "kVserver"		
        case Type11_KVOLUME:
    		return "kVolume"		
        default:
        	return "kCluster"
    }
}

/**
 * Converts Type11Enum Array to its string Array representation
*/
func Type11EnumArrayToValue(type11Enum []Type11Enum) []string {
    convArray := make([]string,len( type11Enum))
    for i:=0; i<len(type11Enum);i++ {
        convArray[i] = Type11EnumToValue(type11Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type11EnumFromValue(value string) Type11Enum {
    switch value {
        case "kCluster":
            return Type11_KCLUSTER
        case "kVserver":
            return Type11_KVSERVER
        case "kVolume":
            return Type11_KVOLUME
        default:
            return Type11_KCLUSTER
    }
}
