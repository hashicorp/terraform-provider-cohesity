// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HostTypeVmwareProtectionSourceEnum enum
 */
type HostTypeVmwareProtectionSourceEnum int

/**
 * Value collection for HostTypeVmwareProtectionSourceEnum enum
 */
const (
    HostTypeVmwareProtectionSource_KLINUX HostTypeVmwareProtectionSourceEnum = 1 + iota
    HostTypeVmwareProtectionSource_KWINDOWS
    HostTypeVmwareProtectionSource_KAIX
    HostTypeVmwareProtectionSource_KSOLARIS
    HostTypeVmwareProtectionSource_KSAPHANA
    HostTypeVmwareProtectionSource_KOTHER
)

func (r HostTypeVmwareProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := HostTypeVmwareProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *HostTypeVmwareProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HostTypeVmwareProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HostTypeVmwareProtectionSourceEnum to its string representation
 */
func HostTypeVmwareProtectionSourceEnumToValue(hostType_vmware_ProtectionSourceEnum HostTypeVmwareProtectionSourceEnum) string {
    switch hostType_vmware_ProtectionSourceEnum {
        case HostTypeVmwareProtectionSource_KLINUX:
    		return "kLinux"
        case HostTypeVmwareProtectionSource_KWINDOWS:
    		return "kWindows"
        case HostTypeVmwareProtectionSource_KAIX:
    		return "kAix"
        case HostTypeVmwareProtectionSource_KSOLARIS:
    		return "kSolaris"
        case HostTypeVmwareProtectionSource_KSAPHANA:
    		return "kSapHana"
        case HostTypeVmwareProtectionSource_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostTypeVmwareProtectionSourceEnum Array to its string Array representation
*/
func HostTypeVmwareProtectionSourceEnumArrayToValue(hostType_vmware_ProtectionSourceEnum []HostTypeVmwareProtectionSourceEnum) []string {
    convArray := make([]string,len( hostType_vmware_ProtectionSourceEnum))
    for i:=0; i<len(hostType_vmware_ProtectionSourceEnum);i++ {
        convArray[i] = HostTypeVmwareProtectionSourceEnumToValue(hostType_vmware_ProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostTypeVmwareProtectionSourceEnumFromValue(value string) HostTypeVmwareProtectionSourceEnum {
    switch value {
        case "kLinux":
            return HostTypeVmwareProtectionSource_KLINUX
        case "kWindows":
            return HostTypeVmwareProtectionSource_KWINDOWS
        case "kAix":
            return HostTypeVmwareProtectionSource_KAIX
        case "kSolaris":
            return HostTypeVmwareProtectionSource_KSOLARIS
        case "kSapHana":
            return HostTypeVmwareProtectionSource_KSAPHANA
        case "kOther":
            return HostTypeVmwareProtectionSource_KOTHER
        default:
            return HostTypeVmwareProtectionSource_KLINUX
    }
}
