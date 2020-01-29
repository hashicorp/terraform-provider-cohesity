// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusCopySnapshotTaskStatusEnum enum
 */
type StatusCopySnapshotTaskStatusEnum int

/**
 * Value collection for StatusCopySnapshotTaskStatusEnum enum
 */
const (
    StatusCopySnapshotTaskStatus_KACCEPTED StatusCopySnapshotTaskStatusEnum = 1 + iota
    StatusCopySnapshotTaskStatus_KRUNNING
    StatusCopySnapshotTaskStatus_KCANCELING
    StatusCopySnapshotTaskStatus_KCANCELED
    StatusCopySnapshotTaskStatus_KSUCCESS
    StatusCopySnapshotTaskStatus_KFAILURE
)

func (r StatusCopySnapshotTaskStatusEnum) MarshalJSON() ([]byte, error) {
    s := StatusCopySnapshotTaskStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusCopySnapshotTaskStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusCopySnapshotTaskStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusCopySnapshotTaskStatusEnum to its string representation
 */
func StatusCopySnapshotTaskStatusEnumToValue(statusCopySnapshotTaskStatusEnum StatusCopySnapshotTaskStatusEnum) string {
    switch statusCopySnapshotTaskStatusEnum {
        case StatusCopySnapshotTaskStatus_KACCEPTED:
    		return "kAccepted"
        case StatusCopySnapshotTaskStatus_KRUNNING:
    		return "kRunning"
        case StatusCopySnapshotTaskStatus_KCANCELING:
    		return "kCanceling"
        case StatusCopySnapshotTaskStatus_KCANCELED:
    		return "kCanceled"
        case StatusCopySnapshotTaskStatus_KSUCCESS:
    		return "kSuccess"
        case StatusCopySnapshotTaskStatus_KFAILURE:
    		return "kFailure"
        default:
        	return "kAccepted"
    }
}

/**
 * Converts StatusCopySnapshotTaskStatusEnum Array to its string Array representation
*/
func StatusCopySnapshotTaskStatusEnumArrayToValue(statusCopySnapshotTaskStatusEnum []StatusCopySnapshotTaskStatusEnum) []string {
    convArray := make([]string,len( statusCopySnapshotTaskStatusEnum))
    for i:=0; i<len(statusCopySnapshotTaskStatusEnum);i++ {
        convArray[i] = StatusCopySnapshotTaskStatusEnumToValue(statusCopySnapshotTaskStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusCopySnapshotTaskStatusEnumFromValue(value string) StatusCopySnapshotTaskStatusEnum {
    switch value {
        case "kAccepted":
            return StatusCopySnapshotTaskStatus_KACCEPTED
        case "kRunning":
            return StatusCopySnapshotTaskStatus_KRUNNING
        case "kCanceling":
            return StatusCopySnapshotTaskStatus_KCANCELING
        case "kCanceled":
            return StatusCopySnapshotTaskStatus_KCANCELED
        case "kSuccess":
            return StatusCopySnapshotTaskStatus_KSUCCESS
        case "kFailure":
            return StatusCopySnapshotTaskStatus_KFAILURE
        default:
            return StatusCopySnapshotTaskStatus_KACCEPTED
    }
}
