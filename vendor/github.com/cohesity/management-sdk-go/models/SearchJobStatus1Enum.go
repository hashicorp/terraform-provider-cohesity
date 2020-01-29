package models

import(
    "encoding/json"
)

/**
 * Type definition for SearchJobStatus1Enum enum
 */
type SearchJobStatus1Enum int

/**
 * Value collection for SearchJobStatus1Enum enum
 */
const (
    SearchJobStatus1_KJOBRUNNING SearchJobStatus1Enum = 1 + iota
    SearchJobStatus1_KJOBFINISHED
    SearchJobStatus1_KJOBFAILED
    SearchJobStatus1_KJOBCANCELED
    SearchJobStatus1_KJOBPAUSED
)

func (r SearchJobStatus1Enum) MarshalJSON() ([]byte, error) { 
    s := SearchJobStatus1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *SearchJobStatus1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  SearchJobStatus1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts SearchJobStatus1Enum to its string representation
 */
func SearchJobStatus1EnumToValue(searchJobStatus1Enum SearchJobStatus1Enum) string {
    switch searchJobStatus1Enum {
        case SearchJobStatus1_KJOBRUNNING:
    		return "kJobRunning"		
        case SearchJobStatus1_KJOBFINISHED:
    		return "kJobFinished"		
        case SearchJobStatus1_KJOBFAILED:
    		return "kJobFailed"		
        case SearchJobStatus1_KJOBCANCELED:
    		return "kJobCanceled"		
        case SearchJobStatus1_KJOBPAUSED:
    		return "kJobPaused"		
        default:
        	return "kJobRunning"
    }
}

/**
 * Converts SearchJobStatus1Enum Array to its string Array representation
*/
func SearchJobStatus1EnumArrayToValue(searchJobStatus1Enum []SearchJobStatus1Enum) []string {
    convArray := make([]string,len( searchJobStatus1Enum))
    for i:=0; i<len(searchJobStatus1Enum);i++ {
        convArray[i] = SearchJobStatus1EnumToValue(searchJobStatus1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SearchJobStatus1EnumFromValue(value string) SearchJobStatus1Enum {
    switch value {
        case "kJobRunning":
            return SearchJobStatus1_KJOBRUNNING
        case "kJobFinished":
            return SearchJobStatus1_KJOBFINISHED
        case "kJobFailed":
            return SearchJobStatus1_KJOBFAILED
        case "kJobCanceled":
            return SearchJobStatus1_KJOBCANCELED
        case "kJobPaused":
            return SearchJobStatus1_KJOBPAUSED
        default:
            return SearchJobStatus1_KJOBRUNNING
    }
}
