package models

import(
    "encoding/json"
)

/**
 * Type definition for HostType3Enum enum
 */
type HostType3Enum int

/**
 * Value collection for HostType3Enum enum
 */
const (
    HostType3_KLINUX HostType3Enum = 1 + iota
    HostType3_KWINDOWS
    HostType3_KAIX
    HostType3_KSOLARIS
)

func (r HostType3Enum) MarshalJSON() ([]byte, error) { 
    s := HostType3EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *HostType3Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  HostType3EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts HostType3Enum to its string representation
 */
func HostType3EnumToValue(hostType3Enum HostType3Enum) string {
    switch hostType3Enum {
        case HostType3_KLINUX:
    		return "kLinux"		
        case HostType3_KWINDOWS:
    		return "kWindows"		
        case HostType3_KAIX:
    		return "kAix"		
        case HostType3_KSOLARIS:
    		return "kSolaris"		
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostType3Enum Array to its string Array representation
*/
func HostType3EnumArrayToValue(hostType3Enum []HostType3Enum) []string {
    convArray := make([]string,len( hostType3Enum))
    for i:=0; i<len(hostType3Enum);i++ {
        convArray[i] = HostType3EnumToValue(hostType3Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostType3EnumFromValue(value string) HostType3Enum {
    switch value {
        case "kLinux":
            return HostType3_KLINUX
        case "kWindows":
            return HostType3_KWINDOWS
        case "kAix":
            return HostType3_KAIX
        case "kSolaris":
            return HostType3_KSOLARIS
        default:
            return HostType3_KLINUX
    }
}
