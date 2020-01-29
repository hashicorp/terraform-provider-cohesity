// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SnapshotTaskStatusEnum enum
 */
type SnapshotTaskStatusEnum int

/**
 * Value collection for SnapshotTaskStatusEnum enum
 */
const (
    SnapshotTaskStatus_KJOBRUNNING SnapshotTaskStatusEnum = 1 + iota
    SnapshotTaskStatus_KJOBFINISHED
    SnapshotTaskStatus_KJOBFAILED
    SnapshotTaskStatus_KJOBCANCELED
    SnapshotTaskStatus_KJOBPAUSED
)

func (r SnapshotTaskStatusEnum) MarshalJSON() ([]byte, error) {
    s := SnapshotTaskStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *SnapshotTaskStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SnapshotTaskStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SnapshotTaskStatusEnum to its string representation
 */
func SnapshotTaskStatusEnumToValue(snapshotTaskStatusEnum SnapshotTaskStatusEnum) string {
    switch snapshotTaskStatusEnum {
        case SnapshotTaskStatus_KJOBRUNNING:
    		return "kJobRunning"
        case SnapshotTaskStatus_KJOBFINISHED:
    		return "kJobFinished"
        case SnapshotTaskStatus_KJOBFAILED:
    		return "kJobFailed"
        case SnapshotTaskStatus_KJOBCANCELED:
    		return "kJobCanceled"
        case SnapshotTaskStatus_KJOBPAUSED:
    		return "kJobPaused"
        default:
        	return "kJobRunning"
    }
}

/**
 * Converts SnapshotTaskStatusEnum Array to its string Array representation
*/
func SnapshotTaskStatusEnumArrayToValue(snapshotTaskStatusEnum []SnapshotTaskStatusEnum) []string {
    convArray := make([]string,len( snapshotTaskStatusEnum))
    for i:=0; i<len(snapshotTaskStatusEnum);i++ {
        convArray[i] = SnapshotTaskStatusEnumToValue(snapshotTaskStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SnapshotTaskStatusEnumFromValue(value string) SnapshotTaskStatusEnum {
    switch value {
        case "kJobRunning":
            return SnapshotTaskStatus_KJOBRUNNING
        case "kJobFinished":
            return SnapshotTaskStatus_KJOBFINISHED
        case "kJobFailed":
            return SnapshotTaskStatus_KJOBFAILED
        case "kJobCanceled":
            return SnapshotTaskStatus_KJOBCANCELED
        case "kJobPaused":
            return SnapshotTaskStatus_KJOBPAUSED
        default:
            return SnapshotTaskStatus_KJOBRUNNING
    }
}
