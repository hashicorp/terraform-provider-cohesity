// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HealthStatusEnum enum
 */
type HealthStatusEnum int

/**
 * Value collection for HealthStatusEnum enum
 */
const (
    HealthStatus_KUNKNOWN HealthStatusEnum = 1 + iota
    HealthStatus_KUNREACHABLE
    HealthStatus_KHEALTHY
    HealthStatus_KDEGRADED
)

func (r HealthStatusEnum) MarshalJSON() ([]byte, error) {
    s := HealthStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *HealthStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HealthStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HealthStatusEnum to its string representation
 */
func HealthStatusEnumToValue(healthStatusEnum HealthStatusEnum) string {
    switch healthStatusEnum {
        case HealthStatus_KUNKNOWN:
    		return "kUnknown"
        case HealthStatus_KUNREACHABLE:
    		return "kUnreachable"
        case HealthStatus_KHEALTHY:
    		return "kHealthy"
        case HealthStatus_KDEGRADED:
    		return "kDegraded"
        default:
        	return "kUnknown"
    }
}

/**
 * Converts HealthStatusEnum Array to its string Array representation
*/
func HealthStatusEnumArrayToValue(healthStatusEnum []HealthStatusEnum) []string {
    convArray := make([]string,len( healthStatusEnum))
    for i:=0; i<len(healthStatusEnum);i++ {
        convArray[i] = HealthStatusEnumToValue(healthStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HealthStatusEnumFromValue(value string) HealthStatusEnum {
    switch value {
        case "kUnknown":
            return HealthStatus_KUNKNOWN
        case "kUnreachable":
            return HealthStatus_KUNREACHABLE
        case "kHealthy":
            return HealthStatus_KHEALTHY
        case "kDegraded":
            return HealthStatus_KDEGRADED
        default:
            return HealthStatus_KUNKNOWN
    }
}
