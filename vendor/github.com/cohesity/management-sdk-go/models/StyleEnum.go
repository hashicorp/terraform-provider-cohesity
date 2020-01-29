// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StyleEnum enum
 */
type StyleEnum int

/**
 * Value collection for StyleEnum enum
 */
const (
    Style_KUNIX StyleEnum = 1 + iota
    Style_KNTFS
    Style_KMIXED
    Style_KUNIFIED
)

func (r StyleEnum) MarshalJSON() ([]byte, error) {
    s := StyleEnumToValue(r)
    return json.Marshal(s)
}

func (r *StyleEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StyleEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StyleEnum to its string representation
 */
func StyleEnumToValue(styleEnum StyleEnum) string {
    switch styleEnum {
        case Style_KUNIX:
    		return "kUnix"
        case Style_KNTFS:
    		return "kNtfs"
        case Style_KMIXED:
    		return "kMixed"
        case Style_KUNIFIED:
    		return "kUnified"
        default:
        	return "kUnix"
    }
}

/**
 * Converts StyleEnum Array to its string Array representation
*/
func StyleEnumArrayToValue(styleEnum []StyleEnum) []string {
    convArray := make([]string,len( styleEnum))
    for i:=0; i<len(styleEnum);i++ {
        convArray[i] = StyleEnumToValue(styleEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StyleEnumFromValue(value string) StyleEnum {
    switch value {
        case "kUnix":
            return Style_KUNIX
        case "kNtfs":
            return Style_KNTFS
        case "kMixed":
            return Style_KMIXED
        case "kUnified":
            return Style_KUNIFIED
        default:
            return Style_KUNIX
    }
}
