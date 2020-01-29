// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AdAttributeFlagEnum enum
 */
type AdAttributeFlagEnum int

/**
 * Value collection for AdAttributeFlagEnum enum
 */
const (
    AdAttributeFlag_KEQUAL AdAttributeFlagEnum = 1 + iota
    AdAttributeFlag_KNOTEQUAL
    AdAttributeFlag_KNOTFOUND
    AdAttributeFlag_KSYSTEM
    AdAttributeFlag_KMULTIVALUE
)

func (r AdAttributeFlagEnum) MarshalJSON() ([]byte, error) {
    s := AdAttributeFlagEnumToValue(r)
    return json.Marshal(s)
}

func (r *AdAttributeFlagEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AdAttributeFlagEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AdAttributeFlagEnum to its string representation
 */
func AdAttributeFlagEnumToValue(adAttributeFlagEnum AdAttributeFlagEnum) string {
    switch adAttributeFlagEnum {
        case AdAttributeFlag_KEQUAL:
    		return "kEqual"
        case AdAttributeFlag_KNOTEQUAL:
    		return "kNotEqual"
        case AdAttributeFlag_KNOTFOUND:
    		return "kNotFound"
        case AdAttributeFlag_KSYSTEM:
    		return "kSystem"
        case AdAttributeFlag_KMULTIVALUE:
    		return "kMultiValue"
        default:
        	return "kEqual"
    }
}

/**
 * Converts AdAttributeFlagEnum Array to its string Array representation
*/
func AdAttributeFlagEnumArrayToValue(adAttributeFlagEnum []AdAttributeFlagEnum) []string {
    convArray := make([]string,len( adAttributeFlagEnum))
    for i:=0; i<len(adAttributeFlagEnum);i++ {
        convArray[i] = AdAttributeFlagEnumToValue(adAttributeFlagEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AdAttributeFlagEnumFromValue(value string) AdAttributeFlagEnum {
    switch value {
        case "kEqual":
            return AdAttributeFlag_KEQUAL
        case "kNotEqual":
            return AdAttributeFlag_KNOTEQUAL
        case "kNotFound":
            return AdAttributeFlag_KNOTFOUND
        case "kSystem":
            return AdAttributeFlag_KSYSTEM
        case "kMultiValue":
            return AdAttributeFlag_KMULTIVALUE
        default:
            return AdAttributeFlag_KEQUAL
    }
}
