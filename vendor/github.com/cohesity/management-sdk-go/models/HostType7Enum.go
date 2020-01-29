package models

import(
    "encoding/json"
)

/**
 * Type definition for HostType7Enum enum
 */
type HostType7Enum int

/**
 * Value collection for HostType7Enum enum
 */
const (
    HostType7_KLINUX HostType7Enum = 1 + iota
    HostType7_KWINDOWS
    HostType7_KAIX
    HostType7_KSOLARIS
)

func (r HostType7Enum) MarshalJSON() ([]byte, error) { 
    s := HostType7EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *HostType7Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  HostType7EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts HostType7Enum to its string representation
 */
func HostType7EnumToValue(hostType7Enum HostType7Enum) string {
    switch hostType7Enum {
        case HostType7_KLINUX:
    		return "kLinux"		
        case HostType7_KWINDOWS:
    		return "kWindows"		
        case HostType7_KAIX:
    		return "kAix"		
        case HostType7_KSOLARIS:
    		return "kSolaris"		
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostType7Enum Array to its string Array representation
*/
func HostType7EnumArrayToValue(hostType7Enum []HostType7Enum) []string {
    convArray := make([]string,len( hostType7Enum))
    for i:=0; i<len(hostType7Enum);i++ {
        convArray[i] = HostType7EnumToValue(hostType7Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostType7EnumFromValue(value string) HostType7Enum {
    switch value {
        case "kLinux":
            return HostType7_KLINUX
        case "kWindows":
            return HostType7_KWINDOWS
        case "kAix":
            return HostType7_KAIX
        case "kSolaris":
            return HostType7_KSOLARIS
        default:
            return HostType7_KLINUX
    }
}
