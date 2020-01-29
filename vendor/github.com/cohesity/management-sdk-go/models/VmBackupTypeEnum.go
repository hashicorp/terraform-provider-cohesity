// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for VmBackupTypeEnum enum
 */
type VmBackupTypeEnum int

/**
 * Value collection for VmBackupTypeEnum enum
 */
const (
    VmBackupType_KRCTBACKUP VmBackupTypeEnum = 1 + iota
    VmBackupType_KVSSBACKUP
)

func (r VmBackupTypeEnum) MarshalJSON() ([]byte, error) {
    s := VmBackupTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *VmBackupTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  VmBackupTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts VmBackupTypeEnum to its string representation
 */
func VmBackupTypeEnumToValue(vmBackupTypeEnum VmBackupTypeEnum) string {
    switch vmBackupTypeEnum {
        case VmBackupType_KRCTBACKUP:
    		return "kRctBackup"
        case VmBackupType_KVSSBACKUP:
    		return "kVssBackup"
        default:
        	return "kRctBackup"
    }
}

/**
 * Converts VmBackupTypeEnum Array to its string Array representation
*/
func VmBackupTypeEnumArrayToValue(vmBackupTypeEnum []VmBackupTypeEnum) []string {
    convArray := make([]string,len( vmBackupTypeEnum))
    for i:=0; i<len(vmBackupTypeEnum);i++ {
        convArray[i] = VmBackupTypeEnumToValue(vmBackupTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func VmBackupTypeEnumFromValue(value string) VmBackupTypeEnum {
    switch value {
        case "kRctBackup":
            return VmBackupType_KRCTBACKUP
        case "kVssBackup":
            return VmBackupType_KVSSBACKUP
        default:
            return VmBackupType_KRCTBACKUP
    }
}
