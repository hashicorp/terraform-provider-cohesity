// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HydraIOPreferentialTierEnum enum
 */
type HydraIOPreferentialTierEnum int

/**
 * Value collection for HydraIOPreferentialTierEnum enum
 */
const (
    HydraIOPreferentialTier_KPCIESSD HydraIOPreferentialTierEnum = 1 + iota
    HydraIOPreferentialTier_KSATASSD
    HydraIOPreferentialTier_KSATAHDD
    HydraIOPreferentialTier_KCLOUD
)

func (r HydraIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := HydraIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *HydraIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HydraIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HydraIOPreferentialTierEnum to its string representation
 */
func HydraIOPreferentialTierEnumToValue(hydraIOPreferentialTierEnum HydraIOPreferentialTierEnum) string {
    switch hydraIOPreferentialTierEnum {
        case HydraIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case HydraIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case HydraIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case HydraIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts HydraIOPreferentialTierEnum Array to its string Array representation
*/
func HydraIOPreferentialTierEnumArrayToValue(hydraIOPreferentialTierEnum []HydraIOPreferentialTierEnum) []string {
    convArray := make([]string,len( hydraIOPreferentialTierEnum))
    for i:=0; i<len(hydraIOPreferentialTierEnum);i++ {
        convArray[i] = HydraIOPreferentialTierEnumToValue(hydraIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HydraIOPreferentialTierEnumFromValue(value string) HydraIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return HydraIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return HydraIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return HydraIOPreferentialTier_KSATAHDD
        case "kCloud":
            return HydraIOPreferentialTier_KCLOUD
        default:
            return HydraIOPreferentialTier_KPCIESSD
    }
}
