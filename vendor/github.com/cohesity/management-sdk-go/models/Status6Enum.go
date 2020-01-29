package models

import(
    "encoding/json"
)

/**
 * Type definition for Status6Enum enum
 */
type Status6Enum int

/**
 * Value collection for Status6Enum enum
 */
const (
    Status6_KREADYTOSCHEDULE Status6Enum = 1 + iota
    Status6_KPROGRESSMONITORCREATED
    Status6_KRETRIEVEDFROMARCHIVE
    Status6_KADMITTED
    Status6_KINPROGRESS
    Status6_KFINISHINGPROGRESSMONITOR
    Status6_KFINISHED
    Status6_KINTERNALVIEWCREATED
    Status6_KZIPFILEREQUESTED
    Status6_KCANCELLED
)

func (r Status6Enum) MarshalJSON() ([]byte, error) { 
    s := Status6EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Status6Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Status6EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Status6Enum to its string representation
 */
func Status6EnumToValue(status6Enum Status6Enum) string {
    switch status6Enum {
        case Status6_KREADYTOSCHEDULE:
    		return "kReadyToSchedule"		
        case Status6_KPROGRESSMONITORCREATED:
    		return "kProgressMonitorCreated"		
        case Status6_KRETRIEVEDFROMARCHIVE:
    		return "kRetrievedFromArchive"		
        case Status6_KADMITTED:
    		return "kAdmitted"		
        case Status6_KINPROGRESS:
    		return "kInProgress"		
        case Status6_KFINISHINGPROGRESSMONITOR:
    		return "kFinishingProgressMonitor"		
        case Status6_KFINISHED:
    		return "kFinished"		
        case Status6_KINTERNALVIEWCREATED:
    		return "kInternalViewCreated"		
        case Status6_KZIPFILEREQUESTED:
    		return "kZipFileRequested"		
        case Status6_KCANCELLED:
    		return "kCancelled"		
        default:
        	return "kReadyToSchedule"
    }
}

/**
 * Converts Status6Enum Array to its string Array representation
*/
func Status6EnumArrayToValue(status6Enum []Status6Enum) []string {
    convArray := make([]string,len( status6Enum))
    for i:=0; i<len(status6Enum);i++ {
        convArray[i] = Status6EnumToValue(status6Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Status6EnumFromValue(value string) Status6Enum {
    switch value {
        case "kReadyToSchedule":
            return Status6_KREADYTOSCHEDULE
        case "kProgressMonitorCreated":
            return Status6_KPROGRESSMONITORCREATED
        case "kRetrievedFromArchive":
            return Status6_KRETRIEVEDFROMARCHIVE
        case "kAdmitted":
            return Status6_KADMITTED
        case "kInProgress":
            return Status6_KINPROGRESS
        case "kFinishingProgressMonitor":
            return Status6_KFINISHINGPROGRESSMONITOR
        case "kFinished":
            return Status6_KFINISHED
        case "kInternalViewCreated":
            return Status6_KINTERNALVIEWCREATED
        case "kZipFileRequested":
            return Status6_KZIPFILEREQUESTED
        case "kCancelled":
            return Status6_KCANCELLED
        default:
            return Status6_KREADYTOSCHEDULE
    }
}
