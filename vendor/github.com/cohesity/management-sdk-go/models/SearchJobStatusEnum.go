// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SearchJobStatusEnum enum
 */
type SearchJobStatusEnum int

/**
 * Value collection for SearchJobStatusEnum enum
 */
const (
    SearchJobStatus_KJOBRUNNING SearchJobStatusEnum = 1 + iota
    SearchJobStatus_KJOBFINISHED
    SearchJobStatus_KJOBFAILED
    SearchJobStatus_KJOBCANCELED
    SearchJobStatus_KJOBPAUSED
)

func (r SearchJobStatusEnum) MarshalJSON() ([]byte, error) {
    s := SearchJobStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *SearchJobStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SearchJobStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SearchJobStatusEnum to its string representation
 */
func SearchJobStatusEnumToValue(searchJobStatusEnum SearchJobStatusEnum) string {
    switch searchJobStatusEnum {
        case SearchJobStatus_KJOBRUNNING:
    		return "kJobRunning"
        case SearchJobStatus_KJOBFINISHED:
    		return "kJobFinished"
        case SearchJobStatus_KJOBFAILED:
    		return "kJobFailed"
        case SearchJobStatus_KJOBCANCELED:
    		return "kJobCanceled"
        case SearchJobStatus_KJOBPAUSED:
    		return "kJobPaused"
        default:
        	return "kJobRunning"
    }
}

/**
 * Converts SearchJobStatusEnum Array to its string Array representation
*/
func SearchJobStatusEnumArrayToValue(searchJobStatusEnum []SearchJobStatusEnum) []string {
    convArray := make([]string,len( searchJobStatusEnum))
    for i:=0; i<len(searchJobStatusEnum);i++ {
        convArray[i] = SearchJobStatusEnumToValue(searchJobStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SearchJobStatusEnumFromValue(value string) SearchJobStatusEnum {
    switch value {
        case "kJobRunning":
            return SearchJobStatus_KJOBRUNNING
        case "kJobFinished":
            return SearchJobStatus_KJOBFINISHED
        case "kJobFailed":
            return SearchJobStatus_KJOBFAILED
        case "kJobCanceled":
            return SearchJobStatus_KJOBCANCELED
        case "kJobPaused":
            return SearchJobStatus_KJOBPAUSED
        default:
            return SearchJobStatus_KJOBRUNNING
    }
}
