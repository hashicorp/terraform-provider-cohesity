package models

import(
    "encoding/json"
)

/**
 * Type definition for HostType6Enum enum
 */
type HostType6Enum int

/**
 * Value collection for HostType6Enum enum
 */
const (
    HostType6_KLINUX HostType6Enum = 1 + iota
    HostType6_KWINDOWS
    HostType6_KAIX
    HostType6_KSOLARIS
)

func (r HostType6Enum) MarshalJSON() ([]byte, error) { 
    s := HostType6EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *HostType6Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  HostType6EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts HostType6Enum to its string representation
 */
func HostType6EnumToValue(hostType6Enum HostType6Enum) string {
    switch hostType6Enum {
        case HostType6_KLINUX:
    		return "kLinux"		
        case HostType6_KWINDOWS:
    		return "kWindows"		
        case HostType6_KAIX:
    		return "kAix"		
        case HostType6_KSOLARIS:
    		return "kSolaris"		
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostType6Enum Array to its string Array representation
*/
func HostType6EnumArrayToValue(hostType6Enum []HostType6Enum) []string {
    convArray := make([]string,len( hostType6Enum))
    for i:=0; i<len(hostType6Enum);i++ {
        convArray[i] = HostType6EnumToValue(hostType6Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostType6EnumFromValue(value string) HostType6Enum {
    switch value {
        case "kLinux":
            return HostType6_KLINUX
        case "kWindows":
            return HostType6_KWINDOWS
        case "kAix":
            return HostType6_KAIX
        case "kSolaris":
            return HostType6_KSOLARIS
        default:
            return HostType6_KLINUX
    }
}
