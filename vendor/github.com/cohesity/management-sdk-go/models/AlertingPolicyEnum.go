// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AlertingPolicyEnum enum
 */
type AlertingPolicyEnum int

/**
 * Value collection for AlertingPolicyEnum enum
 */
const (
    AlertingPolicy_KSUCCESS AlertingPolicyEnum = 1 + iota
    AlertingPolicy_KFAILURE
    AlertingPolicy_KSLAVIOLATION
)

func (r AlertingPolicyEnum) MarshalJSON() ([]byte, error) {
    s := AlertingPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *AlertingPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AlertingPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AlertingPolicyEnum to its string representation
 */
func AlertingPolicyEnumToValue(alertingPolicyEnum AlertingPolicyEnum) string {
    switch alertingPolicyEnum {
        case AlertingPolicy_KSUCCESS:
    		return "kSuccess"
        case AlertingPolicy_KFAILURE:
    		return "kFailure"
        case AlertingPolicy_KSLAVIOLATION:
    		return "kSlaViolation"
        default:
        	return "kSuccess"
    }
}

/**
 * Converts AlertingPolicyEnum Array to its string Array representation
*/
func AlertingPolicyEnumArrayToValue(alertingPolicyEnum []AlertingPolicyEnum) []string {
    convArray := make([]string,len( alertingPolicyEnum))
    for i:=0; i<len(alertingPolicyEnum);i++ {
        convArray[i] = AlertingPolicyEnumToValue(alertingPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlertingPolicyEnumFromValue(value string) AlertingPolicyEnum {
    switch value {
        case "kSuccess":
            return AlertingPolicy_KSUCCESS
        case "kFailure":
            return AlertingPolicy_KFAILURE
        case "kSlaViolation":
            return AlertingPolicy_KSLAVIOLATION
        default:
            return AlertingPolicy_KSUCCESS
    }
}
