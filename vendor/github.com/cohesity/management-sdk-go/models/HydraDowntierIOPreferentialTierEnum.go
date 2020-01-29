// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HydraDowntierIOPreferentialTierEnum enum
 */
type HydraDowntierIOPreferentialTierEnum int

/**
 * Value collection for HydraDowntierIOPreferentialTierEnum enum
 */
const (
    HydraDowntierIOPreferentialTier_KPCIESSD HydraDowntierIOPreferentialTierEnum = 1 + iota
    HydraDowntierIOPreferentialTier_KSATASSD
    HydraDowntierIOPreferentialTier_KSATAHDD
    HydraDowntierIOPreferentialTier_KCLOUD
)

func (r HydraDowntierIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := HydraDowntierIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *HydraDowntierIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HydraDowntierIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HydraDowntierIOPreferentialTierEnum to its string representation
 */
func HydraDowntierIOPreferentialTierEnumToValue(hydraDowntierIOPreferentialTierEnum HydraDowntierIOPreferentialTierEnum) string {
    switch hydraDowntierIOPreferentialTierEnum {
        case HydraDowntierIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case HydraDowntierIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case HydraDowntierIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case HydraDowntierIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts HydraDowntierIOPreferentialTierEnum Array to its string Array representation
*/
func HydraDowntierIOPreferentialTierEnumArrayToValue(hydraDowntierIOPreferentialTierEnum []HydraDowntierIOPreferentialTierEnum) []string {
    convArray := make([]string,len( hydraDowntierIOPreferentialTierEnum))
    for i:=0; i<len(hydraDowntierIOPreferentialTierEnum);i++ {
        convArray[i] = HydraDowntierIOPreferentialTierEnumToValue(hydraDowntierIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HydraDowntierIOPreferentialTierEnumFromValue(value string) HydraDowntierIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return HydraDowntierIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return HydraDowntierIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return HydraDowntierIOPreferentialTier_KSATAHDD
        case "kCloud":
            return HydraDowntierIOPreferentialTier_KCLOUD
        default:
            return HydraDowntierIOPreferentialTier_KPCIESSD
    }
}
