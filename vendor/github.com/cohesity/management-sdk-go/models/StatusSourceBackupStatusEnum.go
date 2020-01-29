// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusSourceBackupStatusEnum enum
 */
type StatusSourceBackupStatusEnum int

/**
 * Value collection for StatusSourceBackupStatusEnum enum
 */
const (
    StatusSourceBackupStatus_KACCEPTED StatusSourceBackupStatusEnum = 1 + iota
    StatusSourceBackupStatus_KRUNNING
    StatusSourceBackupStatus_KCANCELING
    StatusSourceBackupStatus_KCANCELED
    StatusSourceBackupStatus_KSUCCESS
    StatusSourceBackupStatus_KFAILURE
)

func (r StatusSourceBackupStatusEnum) MarshalJSON() ([]byte, error) {
    s := StatusSourceBackupStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusSourceBackupStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusSourceBackupStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusSourceBackupStatusEnum to its string representation
 */
func StatusSourceBackupStatusEnumToValue(statusSourceBackupStatusEnum StatusSourceBackupStatusEnum) string {
    switch statusSourceBackupStatusEnum {
        case StatusSourceBackupStatus_KACCEPTED:
    		return "kAccepted"
        case StatusSourceBackupStatus_KRUNNING:
    		return "kRunning"
        case StatusSourceBackupStatus_KCANCELING:
    		return "kCanceling"
        case StatusSourceBackupStatus_KCANCELED:
    		return "kCanceled"
        case StatusSourceBackupStatus_KSUCCESS:
    		return "kSuccess"
        case StatusSourceBackupStatus_KFAILURE:
    		return "kFailure"
        default:
        	return "kAccepted"
    }
}

/**
 * Converts StatusSourceBackupStatusEnum Array to its string Array representation
*/
func StatusSourceBackupStatusEnumArrayToValue(statusSourceBackupStatusEnum []StatusSourceBackupStatusEnum) []string {
    convArray := make([]string,len( statusSourceBackupStatusEnum))
    for i:=0; i<len(statusSourceBackupStatusEnum);i++ {
        convArray[i] = StatusSourceBackupStatusEnumToValue(statusSourceBackupStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusSourceBackupStatusEnumFromValue(value string) StatusSourceBackupStatusEnum {
    switch value {
        case "kAccepted":
            return StatusSourceBackupStatus_KACCEPTED
        case "kRunning":
            return StatusSourceBackupStatus_KRUNNING
        case "kCanceling":
            return StatusSourceBackupStatus_KCANCELING
        case "kCanceled":
            return StatusSourceBackupStatus_KCANCELED
        case "kSuccess":
            return StatusSourceBackupStatus_KSUCCESS
        case "kFailure":
            return StatusSourceBackupStatus_KFAILURE
        default:
            return StatusSourceBackupStatus_KACCEPTED
    }
}
