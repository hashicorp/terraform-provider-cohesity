// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for WormRetentionTypeEnum enum
 */
type WormRetentionTypeEnum int

/**
 * Value collection for WormRetentionTypeEnum enum
 */
const (
    WormRetentionType_KNONE WormRetentionTypeEnum = 1 + iota
    WormRetentionType_KCOMPLIANCE
    WormRetentionType_KADMINISTRATIVE
)

func (r WormRetentionTypeEnum) MarshalJSON() ([]byte, error) {
    s := WormRetentionTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *WormRetentionTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  WormRetentionTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts WormRetentionTypeEnum to its string representation
 */
func WormRetentionTypeEnumToValue(wormRetentionTypeEnum WormRetentionTypeEnum) string {
    switch wormRetentionTypeEnum {
        case WormRetentionType_KNONE:
    		return "kNone"
        case WormRetentionType_KCOMPLIANCE:
    		return "kCompliance"
        case WormRetentionType_KADMINISTRATIVE:
    		return "kAdministrative"
        default:
        	return "kNone"
    }
}

/**
 * Converts WormRetentionTypeEnum Array to its string Array representation
*/
func WormRetentionTypeEnumArrayToValue(wormRetentionTypeEnum []WormRetentionTypeEnum) []string {
    convArray := make([]string,len( wormRetentionTypeEnum))
    for i:=0; i<len(wormRetentionTypeEnum);i++ {
        convArray[i] = WormRetentionTypeEnumToValue(wormRetentionTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func WormRetentionTypeEnumFromValue(value string) WormRetentionTypeEnum {
    switch value {
        case "kNone":
            return WormRetentionType_KNONE
        case "kCompliance":
            return WormRetentionType_KCOMPLIANCE
        case "kAdministrative":
            return WormRetentionType_KADMINISTRATIVE
        default:
            return WormRetentionType_KNONE
    }
}
