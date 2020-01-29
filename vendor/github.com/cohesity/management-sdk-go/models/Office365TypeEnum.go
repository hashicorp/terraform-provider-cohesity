// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for Office365TypeEnum enum
 */
type Office365TypeEnum int

/**
 * Value collection for Office365TypeEnum enum
 */
const (
    Office365Type_KDOMAIN Office365TypeEnum = 1 + iota
    Office365Type_KOUTLOOK
    Office365Type_KMAILBOX
)

func (r Office365TypeEnum) MarshalJSON() ([]byte, error) {
    s := Office365TypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *Office365TypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  Office365TypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts Office365TypeEnum to its string representation
 */
func Office365TypeEnumToValue(office365TypeEnum Office365TypeEnum) string {
    switch office365TypeEnum {
        case Office365Type_KDOMAIN:
    		return "kDomain"
        case Office365Type_KOUTLOOK:
    		return "kOutlook"
        case Office365Type_KMAILBOX:
    		return "kMailbox"
        default:
        	return "kDomain"
    }
}

/**
 * Converts Office365TypeEnum Array to its string Array representation
*/
func Office365TypeEnumArrayToValue(office365TypeEnum []Office365TypeEnum) []string {
    convArray := make([]string,len( office365TypeEnum))
    for i:=0; i<len(office365TypeEnum);i++ {
        convArray[i] = Office365TypeEnumToValue(office365TypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Office365TypeEnumFromValue(value string) Office365TypeEnum {
    switch value {
        case "kDomain":
            return Office365Type_KDOMAIN
        case "kOutlook":
            return Office365Type_KOUTLOOK
        case "kMailbox":
            return Office365Type_KMAILBOX
        default:
            return Office365Type_KDOMAIN
    }
}
