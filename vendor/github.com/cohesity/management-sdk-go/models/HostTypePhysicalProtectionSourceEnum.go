// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HostTypePhysicalProtectionSourceEnum enum
 */
type HostTypePhysicalProtectionSourceEnum int

/**
 * Value collection for HostTypePhysicalProtectionSourceEnum enum
 */
const (
    HostTypePhysicalProtectionSource_KLINUX HostTypePhysicalProtectionSourceEnum = 1 + iota
    HostTypePhysicalProtectionSource_KWINDOWS
    HostTypePhysicalProtectionSource_KAIX
    HostTypePhysicalProtectionSource_KSOLARIS
    HostTypePhysicalProtectionSource_KSAPHANA
    HostTypePhysicalProtectionSource_KOTHER
)

func (r HostTypePhysicalProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := HostTypePhysicalProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *HostTypePhysicalProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HostTypePhysicalProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HostTypePhysicalProtectionSourceEnum to its string representation
 */
func HostTypePhysicalProtectionSourceEnumToValue(hostTypePhysicalProtectionSourceEnum HostTypePhysicalProtectionSourceEnum) string {
    switch hostTypePhysicalProtectionSourceEnum {
        case HostTypePhysicalProtectionSource_KLINUX:
    		return "kLinux"
        case HostTypePhysicalProtectionSource_KWINDOWS:
    		return "kWindows"
        case HostTypePhysicalProtectionSource_KAIX:
    		return "kAix"
        case HostTypePhysicalProtectionSource_KSOLARIS:
    		return "kSolaris"
        case HostTypePhysicalProtectionSource_KSAPHANA:
    		return "kSapHana"
        case HostTypePhysicalProtectionSource_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostTypePhysicalProtectionSourceEnum Array to its string Array representation
*/
func HostTypePhysicalProtectionSourceEnumArrayToValue(hostTypePhysicalProtectionSourceEnum []HostTypePhysicalProtectionSourceEnum) []string {
    convArray := make([]string,len( hostTypePhysicalProtectionSourceEnum))
    for i:=0; i<len(hostTypePhysicalProtectionSourceEnum);i++ {
        convArray[i] = HostTypePhysicalProtectionSourceEnumToValue(hostTypePhysicalProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostTypePhysicalProtectionSourceEnumFromValue(value string) HostTypePhysicalProtectionSourceEnum {
    switch value {
        case "kLinux":
            return HostTypePhysicalProtectionSource_KLINUX
        case "kWindows":
            return HostTypePhysicalProtectionSource_KWINDOWS
        case "kAix":
            return HostTypePhysicalProtectionSource_KAIX
        case "kSolaris":
            return HostTypePhysicalProtectionSource_KSOLARIS
        case "kSapHana":
            return HostTypePhysicalProtectionSource_KSAPHANA
        case "kOther":
            return HostTypePhysicalProtectionSource_KOTHER
        default:
            return HostTypePhysicalProtectionSource_KLINUX
    }
}
