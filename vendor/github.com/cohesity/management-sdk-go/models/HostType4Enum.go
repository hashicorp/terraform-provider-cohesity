package models

import(
    "encoding/json"
)

/**
 * Type definition for HostType4Enum enum
 */
type HostType4Enum int

/**
 * Value collection for HostType4Enum enum
 */
const (
    HostType4_KLINUX HostType4Enum = 1 + iota
    HostType4_KWINDOWS
    HostType4_KAIX
    HostType4_KSOLARIS
)

func (r HostType4Enum) MarshalJSON() ([]byte, error) { 
    s := HostType4EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *HostType4Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  HostType4EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts HostType4Enum to its string representation
 */
func HostType4EnumToValue(hostType4Enum HostType4Enum) string {
    switch hostType4Enum {
        case HostType4_KLINUX:
    		return "kLinux"		
        case HostType4_KWINDOWS:
    		return "kWindows"		
        case HostType4_KAIX:
    		return "kAix"		
        case HostType4_KSOLARIS:
    		return "kSolaris"		
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostType4Enum Array to its string Array representation
*/
func HostType4EnumArrayToValue(hostType4Enum []HostType4Enum) []string {
    convArray := make([]string,len( hostType4Enum))
    for i:=0; i<len(hostType4Enum);i++ {
        convArray[i] = HostType4EnumToValue(hostType4Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostType4EnumFromValue(value string) HostType4Enum {
    switch value {
        case "kLinux":
            return HostType4_KLINUX
        case "kWindows":
            return HostType4_KWINDOWS
        case "kAix":
            return HostType4_KAIX
        case "kSolaris":
            return HostType4_KSOLARIS
        default:
            return HostType4_KLINUX
    }
}
