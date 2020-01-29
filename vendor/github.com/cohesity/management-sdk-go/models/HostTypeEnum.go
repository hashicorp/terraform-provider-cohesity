// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HostTypeEnum enum
 */
type HostTypeEnum int

/**
 * Value collection for HostTypeEnum enum
 */
const (
    HostType_KLINUX HostTypeEnum = 1 + iota
    HostType_KWINDOWS
    HostType_KAIX
    HostType_KSOLARIS
    HostType_KSAPHANA
    HostType_KOTHER
)

func (r HostTypeEnum) MarshalJSON() ([]byte, error) {
    s := HostTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *HostTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HostTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HostTypeEnum to its string representation
 */
func HostTypeEnumToValue(hostTypeEnum HostTypeEnum) string {
    switch hostTypeEnum {
        case HostType_KLINUX:
    		return "kLinux"
        case HostType_KWINDOWS:
    		return "kWindows"
        case HostType_KAIX:
    		return "kAix"
        case HostType_KSOLARIS:
    		return "kSolaris"
        case HostType_KSAPHANA:
    		return "kSapHana"
        case HostType_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostTypeEnum Array to its string Array representation
*/
func HostTypeEnumArrayToValue(hostTypeEnum []HostTypeEnum) []string {
    convArray := make([]string,len( hostTypeEnum))
    for i:=0; i<len(hostTypeEnum);i++ {
        convArray[i] = HostTypeEnumToValue(hostTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostTypeEnumFromValue(value string) HostTypeEnum {
    switch value {
        case "kLinux":
            return HostType_KLINUX
        case "kWindows":
            return HostType_KWINDOWS
        case "kAix":
            return HostType_KAIX
        case "kSolaris":
            return HostType_KSOLARIS
        case "kSapHana":
            return HostType_KSAPHANA
        case "kOther":
            return HostType_KOTHER
        default:
            return HostType_KLINUX
    }
}
