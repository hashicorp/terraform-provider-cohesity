// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusBackupRunEnum enum
 */
type StatusBackupRunEnum int

/**
 * Value collection for StatusBackupRunEnum enum
 */
const (
    StatusBackupRun_KACCEPTED StatusBackupRunEnum = 1 + iota
    StatusBackupRun_KRUNNING
    StatusBackupRun_KCANCELING
    StatusBackupRun_KCANCELED
    StatusBackupRun_KSUCCESS
    StatusBackupRun_KFAILURE
)

func (r StatusBackupRunEnum) MarshalJSON() ([]byte, error) {
    s := StatusBackupRunEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusBackupRunEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusBackupRunEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusBackupRunEnum to its string representation
 */
func StatusBackupRunEnumToValue(statusBackupRunEnum StatusBackupRunEnum) string {
    switch statusBackupRunEnum {
        case StatusBackupRun_KACCEPTED:
    		return "kAccepted"
        case StatusBackupRun_KRUNNING:
    		return "kRunning"
        case StatusBackupRun_KCANCELING:
    		return "kCanceling"
        case StatusBackupRun_KCANCELED:
    		return "kCanceled"
        case StatusBackupRun_KSUCCESS:
    		return "kSuccess"
        case StatusBackupRun_KFAILURE:
    		return "kFailure"
        default:
        	return "kAccepted"
    }
}

/**
 * Converts StatusBackupRunEnum Array to its string Array representation
*/
func StatusBackupRunEnumArrayToValue(statusBackupRunEnum []StatusBackupRunEnum) []string {
    convArray := make([]string,len( statusBackupRunEnum))
    for i:=0; i<len(statusBackupRunEnum);i++ {
        convArray[i] = StatusBackupRunEnumToValue(statusBackupRunEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusBackupRunEnumFromValue(value string) StatusBackupRunEnum {
    switch value {
        case "kAccepted":
            return StatusBackupRun_KACCEPTED
        case "kRunning":
            return StatusBackupRun_KRUNNING
        case "kCanceling":
            return StatusBackupRun_KCANCELING
        case "kCanceled":
            return StatusBackupRun_KCANCELED
        case "kSuccess":
            return StatusBackupRun_KSUCCESS
        case "kFailure":
            return StatusBackupRun_KFAILURE
        default:
            return StatusBackupRun_KACCEPTED
    }
}
