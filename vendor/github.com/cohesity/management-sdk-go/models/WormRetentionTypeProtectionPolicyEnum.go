// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for WormRetentionTypeProtectionPolicyEnum enum
 */
type WormRetentionTypeProtectionPolicyEnum int

/**
 * Value collection for WormRetentionTypeProtectionPolicyEnum enum
 */
const (
    WormRetentionTypeProtectionPolicy_KNONE WormRetentionTypeProtectionPolicyEnum = 1 + iota
    WormRetentionTypeProtectionPolicy_KCOMPLIANCE
    WormRetentionTypeProtectionPolicy_KADMINISTRATIVE
)

func (r WormRetentionTypeProtectionPolicyEnum) MarshalJSON() ([]byte, error) {
    s := WormRetentionTypeProtectionPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *WormRetentionTypeProtectionPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  WormRetentionTypeProtectionPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts WormRetentionTypeProtectionPolicyEnum to its string representation
 */
func WormRetentionTypeProtectionPolicyEnumToValue(wormRetentionTypeProtectionPolicyEnum WormRetentionTypeProtectionPolicyEnum) string {
    switch wormRetentionTypeProtectionPolicyEnum {
        case WormRetentionTypeProtectionPolicy_KNONE:
    		return "kNone"
        case WormRetentionTypeProtectionPolicy_KCOMPLIANCE:
    		return "kCompliance"
        case WormRetentionTypeProtectionPolicy_KADMINISTRATIVE:
    		return "kAdministrative"
        default:
        	return "kNone"
    }
}

/**
 * Converts WormRetentionTypeProtectionPolicyEnum Array to its string Array representation
*/
func WormRetentionTypeProtectionPolicyEnumArrayToValue(wormRetentionTypeProtectionPolicyEnum []WormRetentionTypeProtectionPolicyEnum) []string {
    convArray := make([]string,len( wormRetentionTypeProtectionPolicyEnum))
    for i:=0; i<len(wormRetentionTypeProtectionPolicyEnum);i++ {
        convArray[i] = WormRetentionTypeProtectionPolicyEnumToValue(wormRetentionTypeProtectionPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func WormRetentionTypeProtectionPolicyEnumFromValue(value string) WormRetentionTypeProtectionPolicyEnum {
    switch value {
        case "kNone":
            return WormRetentionTypeProtectionPolicy_KNONE
        case "kCompliance":
            return WormRetentionTypeProtectionPolicy_KCOMPLIANCE
        case "kAdministrative":
            return WormRetentionTypeProtectionPolicy_KADMINISTRATIVE
        default:
            return WormRetentionTypeProtectionPolicy_KNONE
    }
}
