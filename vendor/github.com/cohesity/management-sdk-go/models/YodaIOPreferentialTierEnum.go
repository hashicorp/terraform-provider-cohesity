// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for YodaIOPreferentialTierEnum enum
 */
type YodaIOPreferentialTierEnum int

/**
 * Value collection for YodaIOPreferentialTierEnum enum
 */
const (
    YodaIOPreferentialTier_KPCIESSD YodaIOPreferentialTierEnum = 1 + iota
    YodaIOPreferentialTier_KSATASSD
    YodaIOPreferentialTier_KSATAHDD
    YodaIOPreferentialTier_KCLOUD
)

func (r YodaIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := YodaIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *YodaIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  YodaIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts YodaIOPreferentialTierEnum to its string representation
 */
func YodaIOPreferentialTierEnumToValue(yodaIOPreferentialTierEnum YodaIOPreferentialTierEnum) string {
    switch yodaIOPreferentialTierEnum {
        case YodaIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case YodaIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case YodaIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case YodaIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts YodaIOPreferentialTierEnum Array to its string Array representation
*/
func YodaIOPreferentialTierEnumArrayToValue(yodaIOPreferentialTierEnum []YodaIOPreferentialTierEnum) []string {
    convArray := make([]string,len( yodaIOPreferentialTierEnum))
    for i:=0; i<len(yodaIOPreferentialTierEnum);i++ {
        convArray[i] = YodaIOPreferentialTierEnumToValue(yodaIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func YodaIOPreferentialTierEnumFromValue(value string) YodaIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return YodaIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return YodaIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return YodaIOPreferentialTier_KSATAHDD
        case "kCloud":
            return YodaIOPreferentialTier_KCLOUD
        default:
            return YodaIOPreferentialTier_KPCIESSD
    }
}
