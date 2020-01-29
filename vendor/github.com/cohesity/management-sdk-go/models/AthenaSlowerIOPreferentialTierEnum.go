// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AthenaSlowerIOPreferentialTierEnum enum
 */
type AthenaSlowerIOPreferentialTierEnum int

/**
 * Value collection for AthenaSlowerIOPreferentialTierEnum enum
 */
const (
    AthenaSlowerIOPreferentialTier_KPCIESSD AthenaSlowerIOPreferentialTierEnum = 1 + iota
    AthenaSlowerIOPreferentialTier_KSATASSD
    AthenaSlowerIOPreferentialTier_KSATAHDD
    AthenaSlowerIOPreferentialTier_KCLOUD
)

func (r AthenaSlowerIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := AthenaSlowerIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *AthenaSlowerIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AthenaSlowerIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AthenaSlowerIOPreferentialTierEnum to its string representation
 */
func AthenaSlowerIOPreferentialTierEnumToValue(athenaSlowerIOPreferentialTierEnum AthenaSlowerIOPreferentialTierEnum) string {
    switch athenaSlowerIOPreferentialTierEnum {
        case AthenaSlowerIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case AthenaSlowerIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case AthenaSlowerIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case AthenaSlowerIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts AthenaSlowerIOPreferentialTierEnum Array to its string Array representation
*/
func AthenaSlowerIOPreferentialTierEnumArrayToValue(athenaSlowerIOPreferentialTierEnum []AthenaSlowerIOPreferentialTierEnum) []string {
    convArray := make([]string,len( athenaSlowerIOPreferentialTierEnum))
    for i:=0; i<len(athenaSlowerIOPreferentialTierEnum);i++ {
        convArray[i] = AthenaSlowerIOPreferentialTierEnumToValue(athenaSlowerIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AthenaSlowerIOPreferentialTierEnumFromValue(value string) AthenaSlowerIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return AthenaSlowerIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return AthenaSlowerIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return AthenaSlowerIOPreferentialTier_KSATAHDD
        case "kCloud":
            return AthenaSlowerIOPreferentialTier_KCLOUD
        default:
            return AthenaSlowerIOPreferentialTier_KPCIESSD
    }
}
