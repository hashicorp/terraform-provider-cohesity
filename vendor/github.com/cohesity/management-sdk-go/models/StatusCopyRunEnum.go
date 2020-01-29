// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusCopyRunEnum enum
 */
type StatusCopyRunEnum int

/**
 * Value collection for StatusCopyRunEnum enum
 */
const (
    StatusCopyRun_KACCEPTED StatusCopyRunEnum = 1 + iota
    StatusCopyRun_KRUNNING
    StatusCopyRun_KCANCELING
    StatusCopyRun_KCANCELED
    StatusCopyRun_KSUCCESS
    StatusCopyRun_KFAILURE
)

func (r StatusCopyRunEnum) MarshalJSON() ([]byte, error) {
    s := StatusCopyRunEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusCopyRunEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusCopyRunEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusCopyRunEnum to its string representation
 */
func StatusCopyRunEnumToValue(statusCopyRunEnum StatusCopyRunEnum) string {
    switch statusCopyRunEnum {
        case StatusCopyRun_KACCEPTED:
    		return "kAccepted"
        case StatusCopyRun_KRUNNING:
    		return "kRunning"
        case StatusCopyRun_KCANCELING:
    		return "kCanceling"
        case StatusCopyRun_KCANCELED:
    		return "kCanceled"
        case StatusCopyRun_KSUCCESS:
    		return "kSuccess"
        case StatusCopyRun_KFAILURE:
    		return "kFailure"
        default:
        	return "kAccepted"
    }
}

/**
 * Converts StatusCopyRunEnum Array to its string Array representation
*/
func StatusCopyRunEnumArrayToValue(statusCopyRunEnum []StatusCopyRunEnum) []string {
    convArray := make([]string,len( statusCopyRunEnum))
    for i:=0; i<len(statusCopyRunEnum);i++ {
        convArray[i] = StatusCopyRunEnumToValue(statusCopyRunEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusCopyRunEnumFromValue(value string) StatusCopyRunEnum {
    switch value {
        case "kAccepted":
            return StatusCopyRun_KACCEPTED
        case "kRunning":
            return StatusCopyRun_KRUNNING
        case "kCanceling":
            return StatusCopyRun_KCANCELING
        case "kCanceled":
            return StatusCopyRun_KCANCELED
        case "kSuccess":
            return StatusCopyRun_KSUCCESS
        case "kFailure":
            return StatusCopyRun_KFAILURE
        default:
            return StatusCopyRun_KACCEPTED
    }
}
