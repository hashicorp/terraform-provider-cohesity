// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SearchJobStatusRemoteVaultSearchJobResultsEnum enum
 */
type SearchJobStatusRemoteVaultSearchJobResultsEnum int

/**
 * Value collection for SearchJobStatusRemoteVaultSearchJobResultsEnum enum
 */
const (
    SearchJobStatusRemoteVaultSearchJobResults_KJOBRUNNING SearchJobStatusRemoteVaultSearchJobResultsEnum = 1 + iota
    SearchJobStatusRemoteVaultSearchJobResults_KJOBFINISHED
    SearchJobStatusRemoteVaultSearchJobResults_KJOBFAILED
    SearchJobStatusRemoteVaultSearchJobResults_KJOBCANCELED
    SearchJobStatusRemoteVaultSearchJobResults_KJOBPAUSED
)

func (r SearchJobStatusRemoteVaultSearchJobResultsEnum) MarshalJSON() ([]byte, error) {
    s := SearchJobStatusRemoteVaultSearchJobResultsEnumToValue(r)
    return json.Marshal(s)
}

func (r *SearchJobStatusRemoteVaultSearchJobResultsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SearchJobStatusRemoteVaultSearchJobResultsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SearchJobStatusRemoteVaultSearchJobResultsEnum to its string representation
 */
func SearchJobStatusRemoteVaultSearchJobResultsEnumToValue(searchJobStatusRemoteVaultSearchJobResultsEnum SearchJobStatusRemoteVaultSearchJobResultsEnum) string {
    switch searchJobStatusRemoteVaultSearchJobResultsEnum {
        case SearchJobStatusRemoteVaultSearchJobResults_KJOBRUNNING:
    		return "kJobRunning"
        case SearchJobStatusRemoteVaultSearchJobResults_KJOBFINISHED:
    		return "kJobFinished"
        case SearchJobStatusRemoteVaultSearchJobResults_KJOBFAILED:
    		return "kJobFailed"
        case SearchJobStatusRemoteVaultSearchJobResults_KJOBCANCELED:
    		return "kJobCanceled"
        case SearchJobStatusRemoteVaultSearchJobResults_KJOBPAUSED:
    		return "kJobPaused"
        default:
        	return "kJobRunning"
    }
}

/**
 * Converts SearchJobStatusRemoteVaultSearchJobResultsEnum Array to its string Array representation
*/
func SearchJobStatusRemoteVaultSearchJobResultsEnumArrayToValue(searchJobStatusRemoteVaultSearchJobResultsEnum []SearchJobStatusRemoteVaultSearchJobResultsEnum) []string {
    convArray := make([]string,len( searchJobStatusRemoteVaultSearchJobResultsEnum))
    for i:=0; i<len(searchJobStatusRemoteVaultSearchJobResultsEnum);i++ {
        convArray[i] = SearchJobStatusRemoteVaultSearchJobResultsEnumToValue(searchJobStatusRemoteVaultSearchJobResultsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SearchJobStatusRemoteVaultSearchJobResultsEnumFromValue(value string) SearchJobStatusRemoteVaultSearchJobResultsEnum {
    switch value {
        case "kJobRunning":
            return SearchJobStatusRemoteVaultSearchJobResults_KJOBRUNNING
        case "kJobFinished":
            return SearchJobStatusRemoteVaultSearchJobResults_KJOBFINISHED
        case "kJobFailed":
            return SearchJobStatusRemoteVaultSearchJobResults_KJOBFAILED
        case "kJobCanceled":
            return SearchJobStatusRemoteVaultSearchJobResults_KJOBCANCELED
        case "kJobPaused":
            return SearchJobStatusRemoteVaultSearchJobResults_KJOBPAUSED
        default:
            return SearchJobStatusRemoteVaultSearchJobResults_KJOBRUNNING
    }
}
