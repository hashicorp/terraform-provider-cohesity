// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RemediationStateUpdateInfectedFileParamsEnum enum
 */
type RemediationStateUpdateInfectedFileParamsEnum int

/**
 * Value collection for RemediationStateUpdateInfectedFileParamsEnum enum
 */
const (
    RemediationStateUpdateInfectedFileParams_KQUARANTINE RemediationStateUpdateInfectedFileParamsEnum = 1 + iota
    RemediationStateUpdateInfectedFileParams_KUNQUARANTINE
)

func (r RemediationStateUpdateInfectedFileParamsEnum) MarshalJSON() ([]byte, error) {
    s := RemediationStateUpdateInfectedFileParamsEnumToValue(r)
    return json.Marshal(s)
}

func (r *RemediationStateUpdateInfectedFileParamsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RemediationStateUpdateInfectedFileParamsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RemediationStateUpdateInfectedFileParamsEnum to its string representation
 */
func RemediationStateUpdateInfectedFileParamsEnumToValue(remediationStateUpdateInfectedFileParamsEnum RemediationStateUpdateInfectedFileParamsEnum) string {
    switch remediationStateUpdateInfectedFileParamsEnum {
        case RemediationStateUpdateInfectedFileParams_KQUARANTINE:
    		return "kQuarantine"
        case RemediationStateUpdateInfectedFileParams_KUNQUARANTINE:
    		return "kUnquarantine"
        default:
        	return "kQuarantine"
    }
}

/**
 * Converts RemediationStateUpdateInfectedFileParamsEnum Array to its string Array representation
*/
func RemediationStateUpdateInfectedFileParamsEnumArrayToValue(remediationStateUpdateInfectedFileParamsEnum []RemediationStateUpdateInfectedFileParamsEnum) []string {
    convArray := make([]string,len( remediationStateUpdateInfectedFileParamsEnum))
    for i:=0; i<len(remediationStateUpdateInfectedFileParamsEnum);i++ {
        convArray[i] = RemediationStateUpdateInfectedFileParamsEnumToValue(remediationStateUpdateInfectedFileParamsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RemediationStateUpdateInfectedFileParamsEnumFromValue(value string) RemediationStateUpdateInfectedFileParamsEnum {
    switch value {
        case "kQuarantine":
            return RemediationStateUpdateInfectedFileParams_KQUARANTINE
        case "kUnquarantine":
            return RemediationStateUpdateInfectedFileParams_KUNQUARANTINE
        default:
            return RemediationStateUpdateInfectedFileParams_KQUARANTINE
    }
}
