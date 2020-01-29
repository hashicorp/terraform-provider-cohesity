// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TargetHostTypeEnum enum
 */
type TargetHostTypeEnum int

/**
 * Value collection for TargetHostTypeEnum enum
 */
const (
    TargetHostType_KLINUX TargetHostTypeEnum = 1 + iota
    TargetHostType_KWINDOWS
    TargetHostType_KAIX
    TargetHostType_KSOLARIS
    TargetHostType_KSAPHANA
    TargetHostType_KOTHER
)

func (r TargetHostTypeEnum) MarshalJSON() ([]byte, error) {
    s := TargetHostTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *TargetHostTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TargetHostTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TargetHostTypeEnum to its string representation
 */
func TargetHostTypeEnumToValue(targetHostTypeEnum TargetHostTypeEnum) string {
    switch targetHostTypeEnum {
        case TargetHostType_KLINUX:
    		return "kLinux"
        case TargetHostType_KWINDOWS:
    		return "kWindows"
        case TargetHostType_KAIX:
    		return "kAix"
        case TargetHostType_KSOLARIS:
    		return "kSolaris"
        case TargetHostType_KSAPHANA:
    		return "kSapHana"
        case TargetHostType_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts TargetHostTypeEnum Array to its string Array representation
*/
func TargetHostTypeEnumArrayToValue(targetHostTypeEnum []TargetHostTypeEnum) []string {
    convArray := make([]string,len( targetHostTypeEnum))
    for i:=0; i<len(targetHostTypeEnum);i++ {
        convArray[i] = TargetHostTypeEnumToValue(targetHostTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TargetHostTypeEnumFromValue(value string) TargetHostTypeEnum {
    switch value {
        case "kLinux":
            return TargetHostType_KLINUX
        case "kWindows":
            return TargetHostType_KWINDOWS
        case "kAix":
            return TargetHostType_KAIX
        case "kSolaris":
            return TargetHostType_KSOLARIS
        case "kSapHana":
            return TargetHostType_KSAPHANA
        case "kOther":
            return TargetHostType_KOTHER
        default:
            return TargetHostType_KLINUX
    }
}
