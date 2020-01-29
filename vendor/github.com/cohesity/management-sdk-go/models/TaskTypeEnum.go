// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TaskTypeEnum enum
 */
type TaskTypeEnum int

/**
 * Value collection for TaskTypeEnum enum
 */
const (
    TaskType_RESTORE TaskTypeEnum = 1 + iota
    TaskType_CLONE
    TaskType_BACKUPNOW
    TaskType_FIELDMESSAGE
)

func (r TaskTypeEnum) MarshalJSON() ([]byte, error) {
    s := TaskTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *TaskTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TaskTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TaskTypeEnum to its string representation
 */
func TaskTypeEnumToValue(taskTypeEnum TaskTypeEnum) string {
    switch taskTypeEnum {
        case TaskType_RESTORE:
    		return "restore"
        case TaskType_CLONE:
    		return "clone"
        case TaskType_BACKUPNOW:
    		return "backupNow"
        case TaskType_FIELDMESSAGE:
    		return "fieldMessage"
        default:
        	return "restore"
    }
}

/**
 * Converts TaskTypeEnum Array to its string Array representation
*/
func TaskTypeEnumArrayToValue(taskTypeEnum []TaskTypeEnum) []string {
    convArray := make([]string,len( taskTypeEnum))
    for i:=0; i<len(taskTypeEnum);i++ {
        convArray[i] = TaskTypeEnumToValue(taskTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TaskTypeEnumFromValue(value string) TaskTypeEnum {
    switch value {
        case "restore":
            return TaskType_RESTORE
        case "clone":
            return TaskType_CLONE
        case "backupNow":
            return TaskType_BACKUPNOW
        case "fieldMessage":
            return TaskType_FIELDMESSAGE
        default:
            return TaskType_RESTORE
    }
}
