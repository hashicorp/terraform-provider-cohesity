// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ExcludeOffice365TypeEnum enum
 */
type ExcludeOffice365TypeEnum int

/**
 * Value collection for ExcludeOffice365TypeEnum enum
 */
const (
    ExcludeOffice365Type_KDOMAIN ExcludeOffice365TypeEnum = 1 + iota
    ExcludeOffice365Type_KOUTLOOK
    ExcludeOffice365Type_KMAILBOX
)

func (r ExcludeOffice365TypeEnum) MarshalJSON() ([]byte, error) {
    s := ExcludeOffice365TypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ExcludeOffice365TypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ExcludeOffice365TypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ExcludeOffice365TypeEnum to its string representation
 */
func ExcludeOffice365TypeEnumToValue(excludeOffice365TypeEnum ExcludeOffice365TypeEnum) string {
    switch excludeOffice365TypeEnum {
        case ExcludeOffice365Type_KDOMAIN:
    		return "kDomain"
        case ExcludeOffice365Type_KOUTLOOK:
    		return "kOutlook"
        case ExcludeOffice365Type_KMAILBOX:
    		return "kMailbox"
        default:
        	return "kDomain"
    }
}

/**
 * Converts ExcludeOffice365TypeEnum Array to its string Array representation
*/
func ExcludeOffice365TypeEnumArrayToValue(excludeOffice365TypeEnum []ExcludeOffice365TypeEnum) []string {
    convArray := make([]string,len( excludeOffice365TypeEnum))
    for i:=0; i<len(excludeOffice365TypeEnum);i++ {
        convArray[i] = ExcludeOffice365TypeEnumToValue(excludeOffice365TypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ExcludeOffice365TypeEnumFromValue(value string) ExcludeOffice365TypeEnum {
    switch value {
        case "kDomain":
            return ExcludeOffice365Type_KDOMAIN
        case "kOutlook":
            return ExcludeOffice365Type_KOUTLOOK
        case "kMailbox":
            return ExcludeOffice365Type_KMAILBOX
        default:
            return ExcludeOffice365Type_KDOMAIN
    }
}
