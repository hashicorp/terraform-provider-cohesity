// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for DiskFormatEnum enum
 */
type DiskFormatEnum int

/**
 * Value collection for DiskFormatEnum enum
 */
const (
    DiskFormat_KVMDK DiskFormatEnum = 1 + iota
    DiskFormat_KVHD
    DiskFormat_KVHDX
    DiskFormat_KRAW
    DiskFormat_KUNKNOWN
)

func (r DiskFormatEnum) MarshalJSON() ([]byte, error) {
    s := DiskFormatEnumToValue(r)
    return json.Marshal(s)
}

func (r *DiskFormatEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  DiskFormatEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts DiskFormatEnum to its string representation
 */
func DiskFormatEnumToValue(diskFormatEnum DiskFormatEnum) string {
    switch diskFormatEnum {
        case DiskFormat_KVMDK:
    		return "kVMDK"
        case DiskFormat_KVHD:
    		return "kVHD"
        case DiskFormat_KVHDX:
    		return "kVHDx"
        case DiskFormat_KRAW:
    		return "kRaw"
        case DiskFormat_KUNKNOWN:
    		return "kUnknown"
        default:
        	return "kVMDK"
    }
}

/**
 * Converts DiskFormatEnum Array to its string Array representation
*/
func DiskFormatEnumArrayToValue(diskFormatEnum []DiskFormatEnum) []string {
    convArray := make([]string,len( diskFormatEnum))
    for i:=0; i<len(diskFormatEnum);i++ {
        convArray[i] = DiskFormatEnumToValue(diskFormatEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DiskFormatEnumFromValue(value string) DiskFormatEnum {
    switch value {
        case "kVMDK":
            return DiskFormat_KVMDK
        case "kVHD":
            return DiskFormat_KVHD
        case "kVHDx":
            return DiskFormat_KVHDX
        case "kRaw":
            return DiskFormat_KRAW
        case "kUnknown":
            return DiskFormat_KUNKNOWN
        default:
            return DiskFormat_KVMDK
    }
}
