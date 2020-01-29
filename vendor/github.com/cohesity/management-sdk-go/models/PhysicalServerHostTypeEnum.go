// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for PhysicalServerHostTypeEnum enum
 */
type PhysicalServerHostTypeEnum int

/**
 * Value collection for PhysicalServerHostTypeEnum enum
 */
const (
    PhysicalServerHostType_KLINUX PhysicalServerHostTypeEnum = 1 + iota
    PhysicalServerHostType_KWINDOWS
    PhysicalServerHostType_KAIX
    PhysicalServerHostType_KSOLARIS
    PhysicalServerHostType_KSAPHANA
    PhysicalServerHostType_KOTHER
)

func (r PhysicalServerHostTypeEnum) MarshalJSON() ([]byte, error) {
    s := PhysicalServerHostTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *PhysicalServerHostTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  PhysicalServerHostTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts PhysicalServerHostTypeEnum to its string representation
 */
func PhysicalServerHostTypeEnumToValue(physicalServerHostTypeEnum PhysicalServerHostTypeEnum) string {
    switch physicalServerHostTypeEnum {
        case PhysicalServerHostType_KLINUX:
    		return "kLinux"
        case PhysicalServerHostType_KWINDOWS:
    		return "kWindows"
        case PhysicalServerHostType_KAIX:
    		return "kAix"
        case PhysicalServerHostType_KSOLARIS:
    		return "kSolaris"
        case PhysicalServerHostType_KSAPHANA:
    		return "kSapHana"
        case PhysicalServerHostType_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts PhysicalServerHostTypeEnum Array to its string Array representation
*/
func PhysicalServerHostTypeEnumArrayToValue(physicalServerHostTypeEnum []PhysicalServerHostTypeEnum) []string {
    convArray := make([]string,len( physicalServerHostTypeEnum))
    for i:=0; i<len(physicalServerHostTypeEnum);i++ {
        convArray[i] = PhysicalServerHostTypeEnumToValue(physicalServerHostTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PhysicalServerHostTypeEnumFromValue(value string) PhysicalServerHostTypeEnum {
    switch value {
        case "kLinux":
            return PhysicalServerHostType_KLINUX
        case "kWindows":
            return PhysicalServerHostType_KWINDOWS
        case "kAix":
            return PhysicalServerHostType_KAIX
        case "kSolaris":
            return PhysicalServerHostType_KSOLARIS
        case "kSapHana":
            return PhysicalServerHostType_KSAPHANA
        case "kOther":
            return PhysicalServerHostType_KOTHER
        default:
            return PhysicalServerHostType_KLINUX
    }
}
