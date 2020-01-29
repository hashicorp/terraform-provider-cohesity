// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SeverityNotificationRuleEnum enum
 */
type SeverityNotificationRuleEnum int

/**
 * Value collection for SeverityNotificationRuleEnum enum
 */
const (
    SeverityNotificationRule_KCRITICAL SeverityNotificationRuleEnum = 1 + iota
    SeverityNotificationRule_KWARNING
    SeverityNotificationRule_KINFO
)

func (r SeverityNotificationRuleEnum) MarshalJSON() ([]byte, error) {
    s := SeverityNotificationRuleEnumToValue(r)
    return json.Marshal(s)
}

func (r *SeverityNotificationRuleEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SeverityNotificationRuleEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SeverityNotificationRuleEnum to its string representation
 */
func SeverityNotificationRuleEnumToValue(severityNotificationRuleEnum SeverityNotificationRuleEnum) string {
    switch severityNotificationRuleEnum {
        case SeverityNotificationRule_KCRITICAL:
    		return "kCritical"
        case SeverityNotificationRule_KWARNING:
    		return "kWarning"
        case SeverityNotificationRule_KINFO:
    		return "kInfo"
        default:
        	return "kCritical"
    }
}

/**
 * Converts SeverityNotificationRuleEnum Array to its string Array representation
*/
func SeverityNotificationRuleEnumArrayToValue(severityNotificationRuleEnum []SeverityNotificationRuleEnum) []string {
    convArray := make([]string,len( severityNotificationRuleEnum))
    for i:=0; i<len(severityNotificationRuleEnum);i++ {
        convArray[i] = SeverityNotificationRuleEnumToValue(severityNotificationRuleEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SeverityNotificationRuleEnumFromValue(value string) SeverityNotificationRuleEnum {
    switch value {
        case "kCritical":
            return SeverityNotificationRule_KCRITICAL
        case "kWarning":
            return SeverityNotificationRule_KWARNING
        case "kInfo":
            return SeverityNotificationRule_KINFO
        default:
            return SeverityNotificationRule_KCRITICAL
    }
}
