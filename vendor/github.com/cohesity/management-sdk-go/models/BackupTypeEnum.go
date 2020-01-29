// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for BackupTypeEnum enum
 */
type BackupTypeEnum int

/**
 * Value collection for BackupTypeEnum enum
 */
const (
    BackupType_KRCTBACKUP BackupTypeEnum = 1 + iota
    BackupType_KVSSBACKUP
)

func (r BackupTypeEnum) MarshalJSON() ([]byte, error) {
    s := BackupTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *BackupTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  BackupTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts BackupTypeEnum to its string representation
 */
func BackupTypeEnumToValue(backupTypeEnum BackupTypeEnum) string {
    switch backupTypeEnum {
        case BackupType_KRCTBACKUP:
    		return "kRctBackup"
        case BackupType_KVSSBACKUP:
    		return "kVssBackup"
        default:
        	return "kRctBackup"
    }
}

/**
 * Converts BackupTypeEnum Array to its string Array representation
*/
func BackupTypeEnumArrayToValue(backupTypeEnum []BackupTypeEnum) []string {
    convArray := make([]string,len( backupTypeEnum))
    for i:=0; i<len(backupTypeEnum);i++ {
        convArray[i] = BackupTypeEnumToValue(backupTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func BackupTypeEnumFromValue(value string) BackupTypeEnum {
    switch value {
        case "kRctBackup":
            return BackupType_KRCTBACKUP
        case "kVssBackup":
            return BackupType_KVSSBACKUP
        default:
            return BackupType_KRCTBACKUP
    }
}
