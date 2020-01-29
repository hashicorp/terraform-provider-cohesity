// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HostTypeRegisterProtectionSourceParametersEnum enum
 */
type HostTypeRegisterProtectionSourceParametersEnum int

/**
 * Value collection for HostTypeRegisterProtectionSourceParametersEnum enum
 */
const (
    HostTypeRegisterProtectionSourceParameters_KLINUX HostTypeRegisterProtectionSourceParametersEnum = 1 + iota
    HostTypeRegisterProtectionSourceParameters_KWINDOWS
    HostTypeRegisterProtectionSourceParameters_KAIX
    HostTypeRegisterProtectionSourceParameters_KSOLARIS
    HostTypeRegisterProtectionSourceParameters_KSAPHANA
    HostTypeRegisterProtectionSourceParameters_KOTHER
)

func (r HostTypeRegisterProtectionSourceParametersEnum) MarshalJSON() ([]byte, error) {
    s := HostTypeRegisterProtectionSourceParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *HostTypeRegisterProtectionSourceParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HostTypeRegisterProtectionSourceParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HostTypeRegisterProtectionSourceParametersEnum to its string representation
 */
func HostTypeRegisterProtectionSourceParametersEnumToValue(hostTypeRegisterProtectionSourceParametersEnum HostTypeRegisterProtectionSourceParametersEnum) string {
    switch hostTypeRegisterProtectionSourceParametersEnum {
        case HostTypeRegisterProtectionSourceParameters_KLINUX:
    		return "kLinux"
        case HostTypeRegisterProtectionSourceParameters_KWINDOWS:
    		return "kWindows"
        case HostTypeRegisterProtectionSourceParameters_KAIX:
    		return "kAix"
        case HostTypeRegisterProtectionSourceParameters_KSOLARIS:
    		return "kSolaris"
        case HostTypeRegisterProtectionSourceParameters_KSAPHANA:
    		return "kSapHana"
        case HostTypeRegisterProtectionSourceParameters_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostTypeRegisterProtectionSourceParametersEnum Array to its string Array representation
*/
func HostTypeRegisterProtectionSourceParametersEnumArrayToValue(hostTypeRegisterProtectionSourceParametersEnum []HostTypeRegisterProtectionSourceParametersEnum) []string {
    convArray := make([]string,len( hostTypeRegisterProtectionSourceParametersEnum))
    for i:=0; i<len(hostTypeRegisterProtectionSourceParametersEnum);i++ {
        convArray[i] = HostTypeRegisterProtectionSourceParametersEnumToValue(hostTypeRegisterProtectionSourceParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostTypeRegisterProtectionSourceParametersEnumFromValue(value string) HostTypeRegisterProtectionSourceParametersEnum {
    switch value {
        case "kLinux":
            return HostTypeRegisterProtectionSourceParameters_KLINUX
        case "kWindows":
            return HostTypeRegisterProtectionSourceParameters_KWINDOWS
        case "kAix":
            return HostTypeRegisterProtectionSourceParameters_KAIX
        case "kSolaris":
            return HostTypeRegisterProtectionSourceParameters_KSOLARIS
        case "kSapHana":
            return HostTypeRegisterProtectionSourceParameters_KSAPHANA
        case "kOther":
            return HostTypeRegisterProtectionSourceParameters_KOTHER
        default:
            return HostTypeRegisterProtectionSourceParameters_KLINUX
    }
}
