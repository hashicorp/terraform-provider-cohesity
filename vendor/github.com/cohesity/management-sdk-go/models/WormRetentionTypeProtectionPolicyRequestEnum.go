// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for WormRetentionTypeProtectionPolicyRequestEnum enum
 */
type WormRetentionTypeProtectionPolicyRequestEnum int

/**
 * Value collection for WormRetentionTypeProtectionPolicyRequestEnum enum
 */
const (
    WormRetentionTypeProtectionPolicyRequest_KNONE WormRetentionTypeProtectionPolicyRequestEnum = 1 + iota
    WormRetentionTypeProtectionPolicyRequest_KCOMPLIANCE
    WormRetentionTypeProtectionPolicyRequest_KADMINISTRATIVE
)

func (r WormRetentionTypeProtectionPolicyRequestEnum) MarshalJSON() ([]byte, error) {
    s := WormRetentionTypeProtectionPolicyRequestEnumToValue(r)
    return json.Marshal(s)
}

func (r *WormRetentionTypeProtectionPolicyRequestEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  WormRetentionTypeProtectionPolicyRequestEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts WormRetentionTypeProtectionPolicyRequestEnum to its string representation
 */
func WormRetentionTypeProtectionPolicyRequestEnumToValue(wormRetentionTypeProtectionPolicyRequestEnum WormRetentionTypeProtectionPolicyRequestEnum) string {
    switch wormRetentionTypeProtectionPolicyRequestEnum {
        case WormRetentionTypeProtectionPolicyRequest_KNONE:
    		return "kNone"
        case WormRetentionTypeProtectionPolicyRequest_KCOMPLIANCE:
    		return "kCompliance"
        case WormRetentionTypeProtectionPolicyRequest_KADMINISTRATIVE:
    		return "kAdministrative"
        default:
        	return "kNone"
    }
}

/**
 * Converts WormRetentionTypeProtectionPolicyRequestEnum Array to its string Array representation
*/
func WormRetentionTypeProtectionPolicyRequestEnumArrayToValue(wormRetentionTypeProtectionPolicyRequestEnum []WormRetentionTypeProtectionPolicyRequestEnum) []string {
    convArray := make([]string,len( wormRetentionTypeProtectionPolicyRequestEnum))
    for i:=0; i<len(wormRetentionTypeProtectionPolicyRequestEnum);i++ {
        convArray[i] = WormRetentionTypeProtectionPolicyRequestEnumToValue(wormRetentionTypeProtectionPolicyRequestEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func WormRetentionTypeProtectionPolicyRequestEnumFromValue(value string) WormRetentionTypeProtectionPolicyRequestEnum {
    switch value {
        case "kNone":
            return WormRetentionTypeProtectionPolicyRequest_KNONE
        case "kCompliance":
            return WormRetentionTypeProtectionPolicyRequest_KCOMPLIANCE
        case "kAdministrative":
            return WormRetentionTypeProtectionPolicyRequest_KADMINISTRATIVE
        default:
            return WormRetentionTypeProtectionPolicyRequest_KNONE
    }
}
