// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AlgorithmEnum enum
 */
type AlgorithmEnum int

/**
 * Value collection for AlgorithmEnum enum
 */
const (
    Algorithm_REED_SOLOMON AlgorithmEnum = 1 + iota
    Algorithm_LRC
)

func (r AlgorithmEnum) MarshalJSON() ([]byte, error) {
    s := AlgorithmEnumToValue(r)
    return json.Marshal(s)
}

func (r *AlgorithmEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AlgorithmEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AlgorithmEnum to its string representation
 */
func AlgorithmEnumToValue(algorithmEnum AlgorithmEnum) string {
    switch algorithmEnum {
        case Algorithm_REED_SOLOMON:
    		return "REED_SOLOMON"
        case Algorithm_LRC:
    		return "LRC"
        default:
        	return "REED_SOLOMON"
    }
}

/**
 * Converts AlgorithmEnum Array to its string Array representation
*/
func AlgorithmEnumArrayToValue(algorithmEnum []AlgorithmEnum) []string {
    convArray := make([]string,len( algorithmEnum))
    for i:=0; i<len(algorithmEnum);i++ {
        convArray[i] = AlgorithmEnumToValue(algorithmEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlgorithmEnumFromValue(value string) AlgorithmEnum {
    switch value {
        case "REED_SOLOMON":
            return Algorithm_REED_SOLOMON
        case "LRC":
            return Algorithm_LRC
        default:
            return Algorithm_REED_SOLOMON
    }
}
