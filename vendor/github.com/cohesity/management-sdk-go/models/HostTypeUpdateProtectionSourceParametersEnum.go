// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HostTypeUpdateProtectionSourceParametersEnum enum
 */
type HostTypeUpdateProtectionSourceParametersEnum int

/**
 * Value collection for HostTypeUpdateProtectionSourceParametersEnum enum
 */
const (
    HostTypeUpdateProtectionSourceParameters_KLINUX HostTypeUpdateProtectionSourceParametersEnum = 1 + iota
    HostTypeUpdateProtectionSourceParameters_KWINDOWS
    HostTypeUpdateProtectionSourceParameters_KAIX
    HostTypeUpdateProtectionSourceParameters_KSOLARIS
    HostTypeUpdateProtectionSourceParameters_KSAPHANA
    HostTypeUpdateProtectionSourceParameters_KOTHER
)

func (r HostTypeUpdateProtectionSourceParametersEnum) MarshalJSON() ([]byte, error) {
    s := HostTypeUpdateProtectionSourceParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *HostTypeUpdateProtectionSourceParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HostTypeUpdateProtectionSourceParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HostTypeUpdateProtectionSourceParametersEnum to its string representation
 */
func HostTypeUpdateProtectionSourceParametersEnumToValue(hostTypeUpdateProtectionSourceParametersEnum HostTypeUpdateProtectionSourceParametersEnum) string {
    switch hostTypeUpdateProtectionSourceParametersEnum {
        case HostTypeUpdateProtectionSourceParameters_KLINUX:
    		return "kLinux"
        case HostTypeUpdateProtectionSourceParameters_KWINDOWS:
    		return "kWindows"
        case HostTypeUpdateProtectionSourceParameters_KAIX:
    		return "kAix"
        case HostTypeUpdateProtectionSourceParameters_KSOLARIS:
    		return "kSolaris"
        case HostTypeUpdateProtectionSourceParameters_KSAPHANA:
    		return "kSapHana"
        case HostTypeUpdateProtectionSourceParameters_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostTypeUpdateProtectionSourceParametersEnum Array to its string Array representation
*/
func HostTypeUpdateProtectionSourceParametersEnumArrayToValue(hostTypeUpdateProtectionSourceParametersEnum []HostTypeUpdateProtectionSourceParametersEnum) []string {
    convArray := make([]string,len( hostTypeUpdateProtectionSourceParametersEnum))
    for i:=0; i<len(hostTypeUpdateProtectionSourceParametersEnum);i++ {
        convArray[i] = HostTypeUpdateProtectionSourceParametersEnumToValue(hostTypeUpdateProtectionSourceParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostTypeUpdateProtectionSourceParametersEnumFromValue(value string) HostTypeUpdateProtectionSourceParametersEnum {
    switch value {
        case "kLinux":
            return HostTypeUpdateProtectionSourceParameters_KLINUX
        case "kWindows":
            return HostTypeUpdateProtectionSourceParameters_KWINDOWS
        case "kAix":
            return HostTypeUpdateProtectionSourceParameters_KAIX
        case "kSolaris":
            return HostTypeUpdateProtectionSourceParameters_KSOLARIS
        case "kSapHana":
            return HostTypeUpdateProtectionSourceParameters_KSAPHANA
        case "kOther":
            return HostTypeUpdateProtectionSourceParameters_KOTHER
        default:
            return HostTypeUpdateProtectionSourceParameters_KLINUX
    }
}
