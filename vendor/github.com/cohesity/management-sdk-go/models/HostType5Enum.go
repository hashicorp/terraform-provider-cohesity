package models

import(
    "encoding/json"
)

/**
 * Type definition for HostType5Enum enum
 */
type HostType5Enum int

/**
 * Value collection for HostType5Enum enum
 */
const (
    HostType5_KLINUX HostType5Enum = 1 + iota
    HostType5_KWINDOWS
    HostType5_KAIX
    HostType5_KSOLARIS
)

func (r HostType5Enum) MarshalJSON() ([]byte, error) { 
    s := HostType5EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *HostType5Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  HostType5EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts HostType5Enum to its string representation
 */
func HostType5EnumToValue(hostType5Enum HostType5Enum) string {
    switch hostType5Enum {
        case HostType5_KLINUX:
    		return "kLinux"		
        case HostType5_KWINDOWS:
    		return "kWindows"		
        case HostType5_KAIX:
    		return "kAix"		
        case HostType5_KSOLARIS:
    		return "kSolaris"		
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostType5Enum Array to its string Array representation
*/
func HostType5EnumArrayToValue(hostType5Enum []HostType5Enum) []string {
    convArray := make([]string,len( hostType5Enum))
    for i:=0; i<len(hostType5Enum);i++ {
        convArray[i] = HostType5EnumToValue(hostType5Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostType5EnumFromValue(value string) HostType5Enum {
    switch value {
        case "kLinux":
            return HostType5_KLINUX
        case "kWindows":
            return HostType5_KWINDOWS
        case "kAix":
            return HostType5_KAIX
        case "kSolaris":
            return HostType5_KSOLARIS
        default:
            return HostType5_KLINUX
    }
}
