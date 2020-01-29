// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for VmBackupStatusEnum enum
 */
type VmBackupStatusEnum int

/**
 * Value collection for VmBackupStatusEnum enum
 */
const (
    VmBackupStatus_KSUPPORTED VmBackupStatusEnum = 1 + iota
    VmBackupStatus_KUNSUPPORTEDCONFIG
    VmBackupStatus_KMISSING
)

func (r VmBackupStatusEnum) MarshalJSON() ([]byte, error) {
    s := VmBackupStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *VmBackupStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  VmBackupStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts VmBackupStatusEnum to its string representation
 */
func VmBackupStatusEnumToValue(vmBackupStatusEnum VmBackupStatusEnum) string {
    switch vmBackupStatusEnum {
        case VmBackupStatus_KSUPPORTED:
    		return "kSupported"
        case VmBackupStatus_KUNSUPPORTEDCONFIG:
    		return "kUnsupportedConfig"
        case VmBackupStatus_KMISSING:
    		return "kMissing"
        default:
        	return "kSupported"
    }
}

/**
 * Converts VmBackupStatusEnum Array to its string Array representation
*/
func VmBackupStatusEnumArrayToValue(vmBackupStatusEnum []VmBackupStatusEnum) []string {
    convArray := make([]string,len( vmBackupStatusEnum))
    for i:=0; i<len(vmBackupStatusEnum);i++ {
        convArray[i] = VmBackupStatusEnumToValue(vmBackupStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func VmBackupStatusEnumFromValue(value string) VmBackupStatusEnum {
    switch value {
        case "kSupported":
            return VmBackupStatus_KSUPPORTED
        case "kUnsupportedConfig":
            return VmBackupStatus_KUNSUPPORTEDCONFIG
        case "kMissing":
            return VmBackupStatus_KMISSING
        default:
            return VmBackupStatus_KSUPPORTED
    }
}
