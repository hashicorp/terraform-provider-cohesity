// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for GrootIOPreferentialTierEnum enum
 */
type GrootIOPreferentialTierEnum int

/**
 * Value collection for GrootIOPreferentialTierEnum enum
 */
const (
    GrootIOPreferentialTier_KPCIESSD GrootIOPreferentialTierEnum = 1 + iota
    GrootIOPreferentialTier_KSATASSD
    GrootIOPreferentialTier_KSATAHDD
    GrootIOPreferentialTier_KCLOUD
)

func (r GrootIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := GrootIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *GrootIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  GrootIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts GrootIOPreferentialTierEnum to its string representation
 */
func GrootIOPreferentialTierEnumToValue(grootIOPreferentialTierEnum GrootIOPreferentialTierEnum) string {
    switch grootIOPreferentialTierEnum {
        case GrootIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case GrootIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case GrootIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case GrootIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts GrootIOPreferentialTierEnum Array to its string Array representation
*/
func GrootIOPreferentialTierEnumArrayToValue(grootIOPreferentialTierEnum []GrootIOPreferentialTierEnum) []string {
    convArray := make([]string,len( grootIOPreferentialTierEnum))
    for i:=0; i<len(grootIOPreferentialTierEnum);i++ {
        convArray[i] = GrootIOPreferentialTierEnumToValue(grootIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func GrootIOPreferentialTierEnumFromValue(value string) GrootIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return GrootIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return GrootIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return GrootIOPreferentialTier_KSATAHDD
        case "kCloud":
            return GrootIOPreferentialTier_KCLOUD
        default:
            return GrootIOPreferentialTier_KPCIESSD
    }
}
