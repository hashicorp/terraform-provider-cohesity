// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for BackupRunTypeEnum enum
 */
type BackupRunTypeEnum int

/**
 * Value collection for BackupRunTypeEnum enum
 */
const (
    BackupRunType_KREGULAR BackupRunTypeEnum = 1 + iota
    BackupRunType_KFULL
    BackupRunType_KLOG
    BackupRunType_KSYSTEM
)

func (r BackupRunTypeEnum) MarshalJSON() ([]byte, error) {
    s := BackupRunTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *BackupRunTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  BackupRunTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts BackupRunTypeEnum to its string representation
 */
func BackupRunTypeEnumToValue(backupRunTypeEnum BackupRunTypeEnum) string {
    switch backupRunTypeEnum {
        case BackupRunType_KREGULAR:
    		return "kRegular"
        case BackupRunType_KFULL:
    		return "kFull"
        case BackupRunType_KLOG:
    		return "kLog"
        case BackupRunType_KSYSTEM:
    		return "kSystem"
        default:
        	return "kRegular"
    }
}

/**
 * Converts BackupRunTypeEnum Array to its string Array representation
*/
func BackupRunTypeEnumArrayToValue(backupRunTypeEnum []BackupRunTypeEnum) []string {
    convArray := make([]string,len( backupRunTypeEnum))
    for i:=0; i<len(backupRunTypeEnum);i++ {
        convArray[i] = BackupRunTypeEnumToValue(backupRunTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func BackupRunTypeEnumFromValue(value string) BackupRunTypeEnum {
    switch value {
        case "kRegular":
            return BackupRunType_KREGULAR
        case "kFull":
            return BackupRunType_KFULL
        case "kLog":
            return BackupRunType_KLOG
        case "kSystem":
            return BackupRunType_KSYSTEM
        default:
            return BackupRunType_KREGULAR
    }
}
