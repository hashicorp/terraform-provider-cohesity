// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ApolloWalIOPreferentialTierEnum enum
 */
type ApolloWalIOPreferentialTierEnum int

/**
 * Value collection for ApolloWalIOPreferentialTierEnum enum
 */
const (
    ApolloWalIOPreferentialTier_KPCIESSD ApolloWalIOPreferentialTierEnum = 1 + iota
    ApolloWalIOPreferentialTier_KSATASSD
    ApolloWalIOPreferentialTier_KSATAHDD
    ApolloWalIOPreferentialTier_KCLOUD
)

func (r ApolloWalIOPreferentialTierEnum) MarshalJSON() ([]byte, error) {
    s := ApolloWalIOPreferentialTierEnumToValue(r)
    return json.Marshal(s)
}

func (r *ApolloWalIOPreferentialTierEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ApolloWalIOPreferentialTierEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ApolloWalIOPreferentialTierEnum to its string representation
 */
func ApolloWalIOPreferentialTierEnumToValue(apolloWalIOPreferentialTierEnum ApolloWalIOPreferentialTierEnum) string {
    switch apolloWalIOPreferentialTierEnum {
        case ApolloWalIOPreferentialTier_KPCIESSD:
    		return "kPcieSsd"
        case ApolloWalIOPreferentialTier_KSATASSD:
    		return "kSataSsd"
        case ApolloWalIOPreferentialTier_KSATAHDD:
    		return "kSataHdd"
        case ApolloWalIOPreferentialTier_KCLOUD:
    		return "kCloud"
        default:
        	return "kPcieSsd"
    }
}

/**
 * Converts ApolloWalIOPreferentialTierEnum Array to its string Array representation
*/
func ApolloWalIOPreferentialTierEnumArrayToValue(apolloWalIOPreferentialTierEnum []ApolloWalIOPreferentialTierEnum) []string {
    convArray := make([]string,len( apolloWalIOPreferentialTierEnum))
    for i:=0; i<len(apolloWalIOPreferentialTierEnum);i++ {
        convArray[i] = ApolloWalIOPreferentialTierEnumToValue(apolloWalIOPreferentialTierEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ApolloWalIOPreferentialTierEnumFromValue(value string) ApolloWalIOPreferentialTierEnum {
    switch value {
        case "kPcieSsd":
            return ApolloWalIOPreferentialTier_KPCIESSD
        case "kSataSsd":
            return ApolloWalIOPreferentialTier_KSATASSD
        case "kSataHdd":
            return ApolloWalIOPreferentialTier_KSATAHDD
        case "kCloud":
            return ApolloWalIOPreferentialTier_KCLOUD
        default:
            return ApolloWalIOPreferentialTier_KPCIESSD
    }
}
