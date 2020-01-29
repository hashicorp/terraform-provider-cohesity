package models

import(
    "encoding/json"
)

/**
 * Type definition for HostType8Enum enum
 */
type HostType8Enum int

/**
 * Value collection for HostType8Enum enum
 */
const (
    HostType8_KLINUX HostType8Enum = 1 + iota
    HostType8_KWINDOWS
    HostType8_KAIX
    HostType8_KSOLARIS
)

func (r HostType8Enum) MarshalJSON() ([]byte, error) { 
    s := HostType8EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *HostType8Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  HostType8EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts HostType8Enum to its string representation
 */
func HostType8EnumToValue(hostType8Enum HostType8Enum) string {
    switch hostType8Enum {
        case HostType8_KLINUX:
    		return "kLinux"		
        case HostType8_KWINDOWS:
    		return "kWindows"		
        case HostType8_KAIX:
    		return "kAix"		
        case HostType8_KSOLARIS:
    		return "kSolaris"		
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostType8Enum Array to its string Array representation
*/
func HostType8EnumArrayToValue(hostType8Enum []HostType8Enum) []string {
    convArray := make([]string,len( hostType8Enum))
    for i:=0; i<len(hostType8Enum);i++ {
        convArray[i] = HostType8EnumToValue(hostType8Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostType8EnumFromValue(value string) HostType8Enum {
    switch value {
        case "kLinux":
            return HostType8_KLINUX
        case "kWindows":
            return HostType8_KWINDOWS
        case "kAix":
            return HostType8_KAIX
        case "kSolaris":
            return HostType8_KSOLARIS
        default:
            return HostType8_KLINUX
    }
}
