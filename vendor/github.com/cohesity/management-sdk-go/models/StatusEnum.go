// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusEnum enum
 */
type StatusEnum int

/**
 * Value collection for StatusEnum enum
 */
const (
    Status_KUNKNOWN StatusEnum = 1 + iota
    Status_KUNREACHABLE
    Status_KHEALTHY
    Status_KDEGRADED
)

func (r StatusEnum) MarshalJSON() ([]byte, error) {
    s := StatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusEnum to its string representation
 */
func StatusEnumToValue(statusEnum StatusEnum) string {
    switch statusEnum {
        case Status_KUNKNOWN:
    		return "kUnknown"
        case Status_KUNREACHABLE:
    		return "kUnreachable"
        case Status_KHEALTHY:
    		return "kHealthy"
        case Status_KDEGRADED:
    		return "kDegraded"
        default:
        	return "kUnknown"
    }
}

/**
 * Converts StatusEnum Array to its string Array representation
*/
func StatusEnumArrayToValue(statusEnum []StatusEnum) []string {
    convArray := make([]string,len( statusEnum))
    for i:=0; i<len(statusEnum);i++ {
        convArray[i] = StatusEnumToValue(statusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusEnumFromValue(value string) StatusEnum {
    switch value {
        case "kUnknown":
            return Status_KUNKNOWN
        case "kUnreachable":
            return Status_KUNREACHABLE
        case "kHealthy":
            return Status_KHEALTHY
        case "kDegraded":
            return Status_KDEGRADED
        default:
            return Status_KUNKNOWN
    }
}
