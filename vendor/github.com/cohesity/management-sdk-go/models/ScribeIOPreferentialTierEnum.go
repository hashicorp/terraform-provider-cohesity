// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ScribeIOPreferentialTierEnum enum
 */
type ScribeIOPreferentialTierEnum int

/**
 * Value collection for ScribeIOPreferentialTierEnum enum
 */
const (
    ScribeIOPreferentialTier_KPCIESSD ScribeIOPreferentialTierEnum = 1 + iota
    ScribeIOPreferentialTier_KSATASSD
    ScribeIOPreferentialTier_KSATAHDD
    ScribeIOPreferentialTier_KCLOUD
)

func (r ScribeIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := ScribeIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *ScribeIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ScribeIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ScribeIOPreferentialTierEnum to its string representation
 */
func ScribeIOPreferentialTierEnumToValue(scribeIOPreferentialTierEnum ScribeIOPreferentialTierEnum) string {
    switch scribeIOPreferentialTierEnum {
        case ScribeIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case ScribeIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case ScribeIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case ScribeIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts ScribeIOPreferentialTierEnum Array to its string Array representation
*/
func ScribeIOPreferentialTierEnumArrayToValue(scribeIOPreferentialTierEnum []ScribeIOPreferentialTierEnum) []string {
    convArray := make([]string,len( scribeIOPreferentialTierEnum))
    for i:=0; i<len(scribeIOPreferentialTierEnum);i++ {
        convArray[i] = ScribeIOPreferentialTierEnumToValue(scribeIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ScribeIOPreferentialTierEnumFromValue(value string) ScribeIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return ScribeIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return ScribeIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return ScribeIOPreferentialTier_KSATAHDD
        case "kCloud":
            return ScribeIOPreferentialTier_KCLOUD
        default:
            return ScribeIOPreferentialTier_KPCIESSD
    }
}
