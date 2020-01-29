// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusRestoreTaskEnum enum
 */
type StatusRestoreTaskEnum int

/**
 * Value collection for StatusRestoreTaskEnum enum
 */
const (
    StatusRestoreTask_KREADYTOSCHEDULE StatusRestoreTaskEnum = 1 + iota
    StatusRestoreTask_KPROGRESSMONITORCREATED
    StatusRestoreTask_KRETRIEVEDFROMARCHIVE
    StatusRestoreTask_KADMITTED
    StatusRestoreTask_KINPROGRESS
    StatusRestoreTask_KFINISHINGPROGRESSMONITOR
    StatusRestoreTask_KFINISHED
    StatusRestoreTask_KINTERNALVIEWCREATED
    StatusRestoreTask_KZIPFILEREQUESTED
    StatusRestoreTask_KCANCELLED
)

func (r StatusRestoreTaskEnum) MarshalJSON() ([]byte, error) {
    s := StatusRestoreTaskEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusRestoreTaskEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusRestoreTaskEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusRestoreTaskEnum to its string representation
 */
func StatusRestoreTaskEnumToValue(statusRestoreTaskEnum StatusRestoreTaskEnum) string {
    switch statusRestoreTaskEnum {
        case StatusRestoreTask_KREADYTOSCHEDULE:
    		return "kReadyToSchedule"
        case StatusRestoreTask_KPROGRESSMONITORCREATED:
    		return "kProgressMonitorCreated"
        case StatusRestoreTask_KRETRIEVEDFROMARCHIVE:
    		return "kRetrievedFromArchive"
        case StatusRestoreTask_KADMITTED:
    		return "kAdmitted"
        case StatusRestoreTask_KINPROGRESS:
    		return "kInProgress"
        case StatusRestoreTask_KFINISHINGPROGRESSMONITOR:
    		return "kFinishingProgressMonitor"
        case StatusRestoreTask_KFINISHED:
    		return "kFinished"
        case StatusRestoreTask_KINTERNALVIEWCREATED:
    		return "kInternalViewCreated"
        case StatusRestoreTask_KZIPFILEREQUESTED:
    		return "kZipFileRequested"
        case StatusRestoreTask_KCANCELLED:
    		return "kCancelled"
        default:
        	return "kReadyToSchedule"
    }
}

/**
 * Converts StatusRestoreTaskEnum Array to its string Array representation
*/
func StatusRestoreTaskEnumArrayToValue(statusRestoreTaskEnum []StatusRestoreTaskEnum) []string {
    convArray := make([]string,len( statusRestoreTaskEnum))
    for i:=0; i<len(statusRestoreTaskEnum);i++ {
        convArray[i] = StatusRestoreTaskEnumToValue(statusRestoreTaskEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusRestoreTaskEnumFromValue(value string) StatusRestoreTaskEnum {
    switch value {
        case "kReadyToSchedule":
            return StatusRestoreTask_KREADYTOSCHEDULE
        case "kProgressMonitorCreated":
            return StatusRestoreTask_KPROGRESSMONITORCREATED
        case "kRetrievedFromArchive":
            return StatusRestoreTask_KRETRIEVEDFROMARCHIVE
        case "kAdmitted":
            return StatusRestoreTask_KADMITTED
        case "kInProgress":
            return StatusRestoreTask_KINPROGRESS
        case "kFinishingProgressMonitor":
            return StatusRestoreTask_KFINISHINGPROGRESSMONITOR
        case "kFinished":
            return StatusRestoreTask_KFINISHED
        case "kInternalViewCreated":
            return StatusRestoreTask_KINTERNALVIEWCREATED
        case "kZipFileRequested":
            return StatusRestoreTask_KZIPFILEREQUESTED
        case "kCancelled":
            return StatusRestoreTask_KCANCELLED
        default:
            return StatusRestoreTask_KREADYTOSCHEDULE
    }
}
