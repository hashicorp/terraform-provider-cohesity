// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ApolloIOPreferentialTierEnum enum
 */
type ApolloIOPreferentialTierEnum int

/**
 * Value collection for ApolloIOPreferentialTierEnum enum
 */
const (
    ApolloIOPreferentialTier_KPCIESSD ApolloIOPreferentialTierEnum = 1 + iota
    ApolloIOPreferentialTier_KSATASSD
    ApolloIOPreferentialTier_KSATAHDD
    ApolloIOPreferentialTier_KCLOUD
)

func (r ApolloIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := ApolloIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *ApolloIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ApolloIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ApolloIOPreferentialTierEnum to its string representation
 */
func ApolloIOPreferentialTierEnumToValue(apolloIOPreferentialTierEnum ApolloIOPreferentialTierEnum) string {
    switch apolloIOPreferentialTierEnum {
        case ApolloIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case ApolloIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case ApolloIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case ApolloIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts ApolloIOPreferentialTierEnum Array to its string Array representation
*/
func ApolloIOPreferentialTierEnumArrayToValue(apolloIOPreferentialTierEnum []ApolloIOPreferentialTierEnum) []string {
    convArray := make([]string,len( apolloIOPreferentialTierEnum))
    for i:=0; i<len(apolloIOPreferentialTierEnum);i++ {
        convArray[i] = ApolloIOPreferentialTierEnumToValue(apolloIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ApolloIOPreferentialTierEnumFromValue(value string) ApolloIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return ApolloIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return ApolloIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return ApolloIOPreferentialTier_KSATAHDD
        case "kCloud":
            return ApolloIOPreferentialTier_KCLOUD
        default:
            return ApolloIOPreferentialTier_KPCIESSD
    }
}
