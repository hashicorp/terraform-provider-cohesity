// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TaskStateEnum enum
 */
type TaskStateEnum int

/**
 * Value collection for TaskStateEnum enum
 */
const (
    TaskState_KDETACHDISKSDONE TaskStateEnum = 1 + iota
    TaskState_KSETUPDISKSDONE
    TaskState_KMIGRATEDISKSSTARTED
    TaskState_KMIGRATEDISKSDONE
    TaskState_KUNMOUNTDATASTOREDONE
)

func (r TaskStateEnum) MarshalJSON() ([]byte, error) {
    s := TaskStateEnumToValue(r)
    return json.Marshal(s)
}

func (r *TaskStateEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TaskStateEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TaskStateEnum to its string representation
 */
func TaskStateEnumToValue(taskStateEnum TaskStateEnum) string {
    switch taskStateEnum {
        case TaskState_KDETACHDISKSDONE:
    		return "kDetachDisksDone"
        case TaskState_KSETUPDISKSDONE:
    		return "kSetupDisksDone"
        case TaskState_KMIGRATEDISKSSTARTED:
    		return "kMigrateDisksStarted"
        case TaskState_KMIGRATEDISKSDONE:
    		return "kMigrateDisksDone"
        case TaskState_KUNMOUNTDATASTOREDONE:
    		return "kUnMountDatastoreDone"
        default:
        	return "kDetachDisksDone"
    }
}

/**
 * Converts TaskStateEnum Array to its string Array representation
*/
func TaskStateEnumArrayToValue(taskStateEnum []TaskStateEnum) []string {
    convArray := make([]string,len( taskStateEnum))
    for i:=0; i<len(taskStateEnum);i++ {
        convArray[i] = TaskStateEnumToValue(taskStateEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TaskStateEnumFromValue(value string) TaskStateEnum {
    switch value {
        case "kDetachDisksDone":
            return TaskState_KDETACHDISKSDONE
        case "kSetupDisksDone":
            return TaskState_KSETUPDISKSDONE
        case "kMigrateDisksStarted":
            return TaskState_KMIGRATEDISKSSTARTED
        case "kMigrateDisksDone":
            return TaskState_KMIGRATEDISKSDONE
        case "kUnMountDatastoreDone":
            return TaskState_KUNMOUNTDATASTOREDONE
        default:
            return TaskState_KDETACHDISKSDONE
    }
}
