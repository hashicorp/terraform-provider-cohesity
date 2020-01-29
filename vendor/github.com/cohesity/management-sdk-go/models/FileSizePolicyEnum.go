// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for FileSizePolicyEnum enum
 */
type FileSizePolicyEnum int

/**
 * Value collection for FileSizePolicyEnum enum
 */
const (
    FileSizePolicy_KGREATERTHAN FileSizePolicyEnum = 1 + iota
    FileSizePolicy_KSMALLERTHAN
)

func (r FileSizePolicyEnum) MarshalJSON() ([]byte, error) {
    s := FileSizePolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *FileSizePolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  FileSizePolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts FileSizePolicyEnum to its string representation
 */
func FileSizePolicyEnumToValue(fileSizePolicyEnum FileSizePolicyEnum) string {
    switch fileSizePolicyEnum {
        case FileSizePolicy_KGREATERTHAN:
    		return "kGreaterThan"
        case FileSizePolicy_KSMALLERTHAN:
    		return "kSmallerThan"
        default:
        	return "kGreaterThan"
    }
}

/**
 * Converts FileSizePolicyEnum Array to its string Array representation
*/
func FileSizePolicyEnumArrayToValue(fileSizePolicyEnum []FileSizePolicyEnum) []string {
    convArray := make([]string,len( fileSizePolicyEnum))
    for i:=0; i<len(fileSizePolicyEnum);i++ {
        convArray[i] = FileSizePolicyEnumToValue(fileSizePolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func FileSizePolicyEnumFromValue(value string) FileSizePolicyEnum {
    switch value {
        case "kGreaterThan":
            return FileSizePolicy_KGREATERTHAN
        case "kSmallerThan":
            return FileSizePolicy_KSMALLERTHAN
        default:
            return FileSizePolicy_KGREATERTHAN
    }
}
