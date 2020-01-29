// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for IndexingTaskStatusEnum enum
 */
type IndexingTaskStatusEnum int

/**
 * Value collection for IndexingTaskStatusEnum enum
 */
const (
    IndexingTaskStatus_KJOBRUNNING IndexingTaskStatusEnum = 1 + iota
    IndexingTaskStatus_KJOBFINISHED
    IndexingTaskStatus_KJOBFAILED
    IndexingTaskStatus_KJOBCANCELED
    IndexingTaskStatus_KJOBPAUSED
)

func (r IndexingTaskStatusEnum) MarshalJSON() ([]byte, error) {
    s := IndexingTaskStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *IndexingTaskStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  IndexingTaskStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts IndexingTaskStatusEnum to its string representation
 */
func IndexingTaskStatusEnumToValue(indexingTaskStatusEnum IndexingTaskStatusEnum) string {
    switch indexingTaskStatusEnum {
        case IndexingTaskStatus_KJOBRUNNING:
    		return "kJobRunning"
        case IndexingTaskStatus_KJOBFINISHED:
    		return "kJobFinished"
        case IndexingTaskStatus_KJOBFAILED:
    		return "kJobFailed"
        case IndexingTaskStatus_KJOBCANCELED:
    		return "kJobCanceled"
        case IndexingTaskStatus_KJOBPAUSED:
    		return "kJobPaused"
        default:
        	return "kJobRunning"
    }
}

/**
 * Converts IndexingTaskStatusEnum Array to its string Array representation
*/
func IndexingTaskStatusEnumArrayToValue(indexingTaskStatusEnum []IndexingTaskStatusEnum) []string {
    convArray := make([]string,len( indexingTaskStatusEnum))
    for i:=0; i<len(indexingTaskStatusEnum);i++ {
        convArray[i] = IndexingTaskStatusEnumToValue(indexingTaskStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func IndexingTaskStatusEnumFromValue(value string) IndexingTaskStatusEnum {
    switch value {
        case "kJobRunning":
            return IndexingTaskStatus_KJOBRUNNING
        case "kJobFinished":
            return IndexingTaskStatus_KJOBFINISHED
        case "kJobFailed":
            return IndexingTaskStatus_KJOBFAILED
        case "kJobCanceled":
            return IndexingTaskStatus_KJOBCANCELED
        case "kJobPaused":
            return IndexingTaskStatus_KJOBPAUSED
        default:
            return IndexingTaskStatus_KJOBRUNNING
    }
}
