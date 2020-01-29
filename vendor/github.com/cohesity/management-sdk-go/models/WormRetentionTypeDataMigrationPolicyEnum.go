// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for WormRetentionTypeDataMigrationPolicyEnum enum
 */
type WormRetentionTypeDataMigrationPolicyEnum int

/**
 * Value collection for WormRetentionTypeDataMigrationPolicyEnum enum
 */
const (
    WormRetentionTypeDataMigrationPolicy_KNONE WormRetentionTypeDataMigrationPolicyEnum = 1 + iota
    WormRetentionTypeDataMigrationPolicy_KCOMPLIANCE
    WormRetentionTypeDataMigrationPolicy_KADMINISTRATIVE
)

func (r WormRetentionTypeDataMigrationPolicyEnum) MarshalJSON() ([]byte, error) {
    s := WormRetentionTypeDataMigrationPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *WormRetentionTypeDataMigrationPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  WormRetentionTypeDataMigrationPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts WormRetentionTypeDataMigrationPolicyEnum to its string representation
 */
func WormRetentionTypeDataMigrationPolicyEnumToValue(wormRetentionTypeDataMigrationPolicyEnum WormRetentionTypeDataMigrationPolicyEnum) string {
    switch wormRetentionTypeDataMigrationPolicyEnum {
        case WormRetentionTypeDataMigrationPolicy_KNONE:
    		return "kNone"
        case WormRetentionTypeDataMigrationPolicy_KCOMPLIANCE:
    		return "kCompliance"
        case WormRetentionTypeDataMigrationPolicy_KADMINISTRATIVE:
    		return "kAdministrative"
        default:
        	return "kNone"
    }
}

/**
 * Converts WormRetentionTypeDataMigrationPolicyEnum Array to its string Array representation
*/
func WormRetentionTypeDataMigrationPolicyEnumArrayToValue(wormRetentionTypeDataMigrationPolicyEnum []WormRetentionTypeDataMigrationPolicyEnum) []string {
    convArray := make([]string,len( wormRetentionTypeDataMigrationPolicyEnum))
    for i:=0; i<len(wormRetentionTypeDataMigrationPolicyEnum);i++ {
        convArray[i] = WormRetentionTypeDataMigrationPolicyEnumToValue(wormRetentionTypeDataMigrationPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func WormRetentionTypeDataMigrationPolicyEnumFromValue(value string) WormRetentionTypeDataMigrationPolicyEnum {
    switch value {
        case "kNone":
            return WormRetentionTypeDataMigrationPolicy_KNONE
        case "kCompliance":
            return WormRetentionTypeDataMigrationPolicy_KCOMPLIANCE
        case "kAdministrative":
            return WormRetentionTypeDataMigrationPolicy_KADMINISTRATIVE
        default:
            return WormRetentionTypeDataMigrationPolicy_KNONE
    }
}
