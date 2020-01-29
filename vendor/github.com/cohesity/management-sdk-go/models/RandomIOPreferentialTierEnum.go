// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RandomIOPreferentialTierEnum enum
 */
type RandomIOPreferentialTierEnum int

/**
 * Value collection for RandomIOPreferentialTierEnum enum
 */
const (
    RandomIOPreferentialTier_KPCIESSD RandomIOPreferentialTierEnum = 1 + iota
    RandomIOPreferentialTier_KSATASSD
    RandomIOPreferentialTier_KSATAHDD
    RandomIOPreferentialTier_KCLOUD
)

func (r RandomIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := RandomIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *RandomIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RandomIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RandomIOPreferentialTierEnum to its string representation
 */
func RandomIOPreferentialTierEnumToValue(randomIOPreferentialTierEnum RandomIOPreferentialTierEnum) string {
    switch randomIOPreferentialTierEnum {
        case RandomIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case RandomIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case RandomIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case RandomIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts RandomIOPreferentialTierEnum Array to its string Array representation
*/
func RandomIOPreferentialTierEnumArrayToValue(randomIOPreferentialTierEnum []RandomIOPreferentialTierEnum) []string {
    convArray := make([]string,len( randomIOPreferentialTierEnum))
    for i:=0; i<len(randomIOPreferentialTierEnum);i++ {
        convArray[i] = RandomIOPreferentialTierEnumToValue(randomIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RandomIOPreferentialTierEnumFromValue(value string) RandomIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return RandomIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return RandomIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return RandomIOPreferentialTier_KSATAHDD
        case "kCloud":
            return RandomIOPreferentialTier_KCLOUD
        default:
            return RandomIOPreferentialTier_KPCIESSD
    }
}
