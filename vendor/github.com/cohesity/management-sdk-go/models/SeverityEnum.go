// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SeverityEnum enum
 */
type SeverityEnum int

/**
 * Value collection for SeverityEnum enum
 */
const (
    Severity_KCRITICAL SeverityEnum = 1 + iota
    Severity_KWARNING
    Severity_KINFO
)

func (r SeverityEnum) MarshalJSON() ([]byte, error) {
    s := SeverityEnumToValue(r)
    return json.Marshal(s)
}

func (r *SeverityEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SeverityEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SeverityEnum to its string representation
 */
func SeverityEnumToValue(severityEnum SeverityEnum) string {
    switch severityEnum {
        case Severity_KCRITICAL:
    		return "kCritical"
        case Severity_KWARNING:
    		return "kWarning"
        case Severity_KINFO:
    		return "kInfo"
        default:
        	return "kCritical"
    }
}

/**
 * Converts SeverityEnum Array to its string Array representation
*/
func SeverityEnumArrayToValue(severityEnum []SeverityEnum) []string {
    convArray := make([]string,len( severityEnum))
    for i:=0; i<len(severityEnum);i++ {
        convArray[i] = SeverityEnumToValue(severityEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SeverityEnumFromValue(value string) SeverityEnum {
    switch value {
        case "kCritical":
            return Severity_KCRITICAL
        case "kWarning":
            return Severity_KWARNING
        case "kInfo":
            return Severity_KINFO
        default:
            return Severity_KCRITICAL
    }
}
