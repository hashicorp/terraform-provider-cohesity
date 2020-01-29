// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HostTypeHypervProtectionSourceEnum enum
 */
type HostTypeHypervProtectionSourceEnum int

/**
 * Value collection for HostTypeHypervProtectionSourceEnum enum
 */
const (
    HostTypeHypervProtectionSource_KLINUX HostTypeHypervProtectionSourceEnum = 1 + iota
    HostTypeHypervProtectionSource_KWINDOWS
    HostTypeHypervProtectionSource_KAIX
    HostTypeHypervProtectionSource_KSOLARIS
    HostTypeHypervProtectionSource_KSAPHANA
    HostTypeHypervProtectionSource_KOTHER
)

func (r HostTypeHypervProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := HostTypeHypervProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *HostTypeHypervProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HostTypeHypervProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HostTypeHypervProtectionSourceEnum to its string representation
 */
func HostTypeHypervProtectionSourceEnumToValue(hostType_hyperv_ProtectionSourceEnum HostTypeHypervProtectionSourceEnum) string {
    switch hostType_hyperv_ProtectionSourceEnum {
        case HostTypeHypervProtectionSource_KLINUX:
    		return "kLinux"
        case HostTypeHypervProtectionSource_KWINDOWS:
    		return "kWindows"
        case HostTypeHypervProtectionSource_KAIX:
    		return "kAix"
        case HostTypeHypervProtectionSource_KSOLARIS:
    		return "kSolaris"
        case HostTypeHypervProtectionSource_KSAPHANA:
    		return "kSapHana"
        case HostTypeHypervProtectionSource_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostTypeHypervProtectionSourceEnum Array to its string Array representation
*/
func HostTypeHypervProtectionSourceEnumArrayToValue(hostType_hyperv_ProtectionSourceEnum []HostTypeHypervProtectionSourceEnum) []string {
    convArray := make([]string,len( hostType_hyperv_ProtectionSourceEnum))
    for i:=0; i<len(hostType_hyperv_ProtectionSourceEnum);i++ {
        convArray[i] = HostTypeHypervProtectionSourceEnumToValue(hostType_hyperv_ProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostTypeHypervProtectionSourceEnumFromValue(value string) HostTypeHypervProtectionSourceEnum {
    switch value {
        case "kLinux":
            return HostTypeHypervProtectionSource_KLINUX
        case "kWindows":
            return HostTypeHypervProtectionSource_KWINDOWS
        case "kAix":
            return HostTypeHypervProtectionSource_KAIX
        case "kSolaris":
            return HostTypeHypervProtectionSource_KSOLARIS
        case "kSapHana":
            return HostTypeHypervProtectionSource_KSAPHANA
        case "kOther":
            return HostTypeHypervProtectionSource_KOTHER
        default:
            return HostTypeHypervProtectionSource_KLINUX
    }
}
