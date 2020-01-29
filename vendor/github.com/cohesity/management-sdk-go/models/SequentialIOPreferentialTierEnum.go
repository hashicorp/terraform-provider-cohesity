// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SequentialIOPreferentialTierEnum enum
 */
type SequentialIOPreferentialTierEnum int

/**
 * Value collection for SequentialIOPreferentialTierEnum enum
 */
const (
    SequentialIOPreferentialTier_KPCIESSD SequentialIOPreferentialTierEnum = 1 + iota
    SequentialIOPreferentialTier_KSATASSD
    SequentialIOPreferentialTier_KSATAHDD
    SequentialIOPreferentialTier_KCLOUD
)

func (r SequentialIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := SequentialIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *SequentialIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SequentialIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SequentialIOPreferentialTierEnum to its string representation
 */
func SequentialIOPreferentialTierEnumToValue(sequentialIOPreferentialTierEnum SequentialIOPreferentialTierEnum) string {
    switch sequentialIOPreferentialTierEnum {
        case SequentialIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case SequentialIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case SequentialIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case SequentialIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts SequentialIOPreferentialTierEnum Array to its string Array representation
*/
func SequentialIOPreferentialTierEnumArrayToValue(sequentialIOPreferentialTierEnum []SequentialIOPreferentialTierEnum) []string {
    convArray := make([]string,len( sequentialIOPreferentialTierEnum))
    for i:=0; i<len(sequentialIOPreferentialTierEnum);i++ {
        convArray[i] = SequentialIOPreferentialTierEnumToValue(sequentialIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SequentialIOPreferentialTierEnumFromValue(value string) SequentialIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return SequentialIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return SequentialIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return SequentialIOPreferentialTier_KSATAHDD
        case "kCloud":
            return SequentialIOPreferentialTier_KCLOUD
        default:
            return SequentialIOPreferentialTier_KPCIESSD
    }
}
