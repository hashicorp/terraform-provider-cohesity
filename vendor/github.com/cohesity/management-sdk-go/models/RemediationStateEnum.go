// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RemediationStateEnum enum
 */
type RemediationStateEnum int

/**
 * Value collection for RemediationStateEnum enum
 */
const (
    RemediationState_KQUARANTINE RemediationStateEnum = 1 + iota
    RemediationState_KUNQUARANTINE
)

func (r RemediationStateEnum) MarshalJSON() ([]byte, error) {
    s := RemediationStateEnumToValue(r)
    return json.Marshal(s)
}

func (r *RemediationStateEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RemediationStateEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RemediationStateEnum to its string representation
 */
func RemediationStateEnumToValue(remediationStateEnum RemediationStateEnum) string {
    switch remediationStateEnum {
        case RemediationState_KQUARANTINE:
    		return "kQuarantine"
        case RemediationState_KUNQUARANTINE:
    		return "kUnquarantine"
        default:
        	return "kQuarantine"
    }
}

/**
 * Converts RemediationStateEnum Array to its string Array representation
*/
func RemediationStateEnumArrayToValue(remediationStateEnum []RemediationStateEnum) []string {
    convArray := make([]string,len( remediationStateEnum))
    for i:=0; i<len(remediationStateEnum);i++ {
        convArray[i] = RemediationStateEnumToValue(remediationStateEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RemediationStateEnumFromValue(value string) RemediationStateEnum {
    switch value {
        case "kQuarantine":
            return RemediationState_KQUARANTINE
        case "kUnquarantine":
            return RemediationState_KUNQUARANTINE
        default:
            return RemediationState_KQUARANTINE
    }
}
