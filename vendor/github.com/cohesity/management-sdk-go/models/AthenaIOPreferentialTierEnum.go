// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AthenaIOPreferentialTierEnum enum
 */
type AthenaIOPreferentialTierEnum int

/**
 * Value collection for AthenaIOPreferentialTierEnum enum
 */
const (
    AthenaIOPreferentialTier_KPCIESSD AthenaIOPreferentialTierEnum = 1 + iota
    AthenaIOPreferentialTier_KSATASSD
    AthenaIOPreferentialTier_KSATAHDD
    AthenaIOPreferentialTier_KCLOUD
)

func (r AthenaIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := AthenaIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *AthenaIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AthenaIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AthenaIOPreferentialTierEnum to its string representation
 */
func AthenaIOPreferentialTierEnumToValue(athenaIOPreferentialTierEnum AthenaIOPreferentialTierEnum) string {
    switch athenaIOPreferentialTierEnum {
        case AthenaIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case AthenaIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case AthenaIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case AthenaIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts AthenaIOPreferentialTierEnum Array to its string Array representation
*/
func AthenaIOPreferentialTierEnumArrayToValue(athenaIOPreferentialTierEnum []AthenaIOPreferentialTierEnum) []string {
    convArray := make([]string,len( athenaIOPreferentialTierEnum))
    for i:=0; i<len(athenaIOPreferentialTierEnum);i++ {
        convArray[i] = AthenaIOPreferentialTierEnumToValue(athenaIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AthenaIOPreferentialTierEnumFromValue(value string) AthenaIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return AthenaIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return AthenaIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return AthenaIOPreferentialTier_KSATAHDD
        case "kCloud":
            return AthenaIOPreferentialTier_KCLOUD
        default:
            return AthenaIOPreferentialTier_KPCIESSD
    }
}
