// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StateLicenseStateEnum enum
 */
type StateLicenseStateEnum int

/**
 * Value collection for StateLicenseStateEnum enum
 */
const (
    StateLicenseState_KINPROGRESSNEWCLUSTER StateLicenseStateEnum = 1 + iota
    StateLicenseState_KINPROGRESSOLDCLUSTER
    StateLicenseState_KCLAIMED
    StateLicenseState_KSKIPPED
    StateLicenseState_KSTARTED
)

func (r StateLicenseStateEnum) MarshalJSON() ([]byte, error) {
    s := StateLicenseStateEnumToValue(r)
    return json.Marshal(s)
}

func (r *StateLicenseStateEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StateLicenseStateEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StateLicenseStateEnum to its string representation
 */
func StateLicenseStateEnumToValue(stateLicenseStateEnum StateLicenseStateEnum) string {
    switch stateLicenseStateEnum {
        case StateLicenseState_KINPROGRESSNEWCLUSTER:
    		return "kInProgressNewCluster"
        case StateLicenseState_KINPROGRESSOLDCLUSTER:
    		return "kInProgressOldCluster"
        case StateLicenseState_KCLAIMED:
    		return "kClaimed"
        case StateLicenseState_KSKIPPED:
    		return "kSkipped"
        case StateLicenseState_KSTARTED:
    		return "kStarted"
        default:
        	return "kInProgressNewCluster"
    }
}

/**
 * Converts StateLicenseStateEnum Array to its string Array representation
*/
func StateLicenseStateEnumArrayToValue(stateLicenseStateEnum []StateLicenseStateEnum) []string {
    convArray := make([]string,len( stateLicenseStateEnum))
    for i:=0; i<len(stateLicenseStateEnum);i++ {
        convArray[i] = StateLicenseStateEnumToValue(stateLicenseStateEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StateLicenseStateEnumFromValue(value string) StateLicenseStateEnum {
    switch value {
        case "kInProgressNewCluster":
            return StateLicenseState_KINPROGRESSNEWCLUSTER
        case "kInProgressOldCluster":
            return StateLicenseState_KINPROGRESSOLDCLUSTER
        case "kClaimed":
            return StateLicenseState_KCLAIMED
        case "kSkipped":
            return StateLicenseState_KSKIPPED
        case "kStarted":
            return StateLicenseState_KSTARTED
        default:
            return StateLicenseState_KINPROGRESSNEWCLUSTER
    }
}
