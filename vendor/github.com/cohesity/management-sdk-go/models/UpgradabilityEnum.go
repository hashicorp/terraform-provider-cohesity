// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for UpgradabilityEnum enum
 */
type UpgradabilityEnum int

/**
 * Value collection for UpgradabilityEnum enum
 */
const (
    Upgradability_KUPGRADABLE UpgradabilityEnum = 1 + iota
    Upgradability_KCURRENT
    Upgradability_KUNKNOWN
    Upgradability_KNONUPGRADABLEINVALIDVERSION
    Upgradability_KNONUPGRADABLEAGENTISNEWER
    Upgradability_KNONUPGRADABLEAGENTISOLD
)

func (r UpgradabilityEnum) MarshalJSON() ([]byte, error) {
    s := UpgradabilityEnumToValue(r)
    return json.Marshal(s)
}

func (r *UpgradabilityEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  UpgradabilityEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts UpgradabilityEnum to its string representation
 */
func UpgradabilityEnumToValue(upgradabilityEnum UpgradabilityEnum) string {
    switch upgradabilityEnum {
        case Upgradability_KUPGRADABLE:
    		return "kUpgradable"
        case Upgradability_KCURRENT:
    		return "kCurrent"
        case Upgradability_KUNKNOWN:
    		return "kUnknown"
        case Upgradability_KNONUPGRADABLEINVALIDVERSION:
    		return "kNonUpgradableInvalidVersion"
        case Upgradability_KNONUPGRADABLEAGENTISNEWER:
    		return "kNonUpgradableAgentIsNewer"
        case Upgradability_KNONUPGRADABLEAGENTISOLD:
    		return "kNonUpgradableAgentIsOld"
        default:
        	return "kUpgradable"
    }
}

/**
 * Converts UpgradabilityEnum Array to its string Array representation
*/
func UpgradabilityEnumArrayToValue(upgradabilityEnum []UpgradabilityEnum) []string {
    convArray := make([]string,len( upgradabilityEnum))
    for i:=0; i<len(upgradabilityEnum);i++ {
        convArray[i] = UpgradabilityEnumToValue(upgradabilityEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func UpgradabilityEnumFromValue(value string) UpgradabilityEnum {
    switch value {
        case "kUpgradable":
            return Upgradability_KUPGRADABLE
        case "kCurrent":
            return Upgradability_KCURRENT
        case "kUnknown":
            return Upgradability_KUNKNOWN
        case "kNonUpgradableInvalidVersion":
            return Upgradability_KNONUPGRADABLEINVALIDVERSION
        case "kNonUpgradableAgentIsNewer":
            return Upgradability_KNONUPGRADABLEAGENTISNEWER
        case "kNonUpgradableAgentIsOld":
            return Upgradability_KNONUPGRADABLEAGENTISOLD
        default:
            return Upgradability_KUPGRADABLE
    }
}
