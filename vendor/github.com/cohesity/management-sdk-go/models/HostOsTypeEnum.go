// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HostOsTypeEnum enum
 */
type HostOsTypeEnum int

/**
 * Value collection for HostOsTypeEnum enum
 */
const (
    HostOsType_KLINUX HostOsTypeEnum = 1 + iota
    HostOsType_KWINDOWS
    HostOsType_KAIX
    HostOsType_KSOLARIS
    HostOsType_KSAPHANA
    HostOsType_KOTHER
)

func (r HostOsTypeEnum) MarshalJSON() ([]byte, error) {
    s := HostOsTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *HostOsTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HostOsTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HostOsTypeEnum to its string representation
 */
func HostOsTypeEnumToValue(hostOsTypeEnum HostOsTypeEnum) string {
    switch hostOsTypeEnum {
        case HostOsType_KLINUX:
    		return "kLinux"
        case HostOsType_KWINDOWS:
    		return "kWindows"
        case HostOsType_KAIX:
    		return "kAix"
        case HostOsType_KSOLARIS:
    		return "kSolaris"
        case HostOsType_KSAPHANA:
    		return "kSapHana"
        case HostOsType_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostOsTypeEnum Array to its string Array representation
*/
func HostOsTypeEnumArrayToValue(hostOsTypeEnum []HostOsTypeEnum) []string {
    convArray := make([]string,len( hostOsTypeEnum))
    for i:=0; i<len(hostOsTypeEnum);i++ {
        convArray[i] = HostOsTypeEnumToValue(hostOsTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostOsTypeEnumFromValue(value string) HostOsTypeEnum {
    switch value {
        case "kLinux":
            return HostOsType_KLINUX
        case "kWindows":
            return HostOsType_KWINDOWS
        case "kAix":
            return HostOsType_KAIX
        case "kSolaris":
            return HostOsType_KSOLARIS
        case "kSapHana":
            return HostOsType_KSAPHANA
        case "kOther":
            return HostOsType_KOTHER
        default:
            return HostOsType_KLINUX
    }
}
