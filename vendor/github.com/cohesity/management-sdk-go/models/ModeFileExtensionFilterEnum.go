// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ModeFileExtensionFilterEnum enum
 */
type ModeFileExtensionFilterEnum int

/**
 * Value collection for ModeFileExtensionFilterEnum enum
 */
const (
    ModeFileExtensionFilter_KWHITELIST ModeFileExtensionFilterEnum = 1 + iota
    ModeFileExtensionFilter_KBLACKLIST
)

func (r ModeFileExtensionFilterEnum) MarshalJSON() ([]byte, error) {
    s := ModeFileExtensionFilterEnumToValue(r)
    return json.Marshal(s)
}

func (r *ModeFileExtensionFilterEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ModeFileExtensionFilterEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ModeFileExtensionFilterEnum to its string representation
 */
func ModeFileExtensionFilterEnumToValue(modeFileExtensionFilterEnum ModeFileExtensionFilterEnum) string {
    switch modeFileExtensionFilterEnum {
        case ModeFileExtensionFilter_KWHITELIST:
    		return "kWhitelist"
        case ModeFileExtensionFilter_KBLACKLIST:
    		return "kBlacklist"
        default:
        	return "kWhitelist"
    }
}

/**
 * Converts ModeFileExtensionFilterEnum Array to its string Array representation
*/
func ModeFileExtensionFilterEnumArrayToValue(modeFileExtensionFilterEnum []ModeFileExtensionFilterEnum) []string {
    convArray := make([]string,len( modeFileExtensionFilterEnum))
    for i:=0; i<len(modeFileExtensionFilterEnum);i++ {
        convArray[i] = ModeFileExtensionFilterEnumToValue(modeFileExtensionFilterEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ModeFileExtensionFilterEnumFromValue(value string) ModeFileExtensionFilterEnum {
    switch value {
        case "kWhitelist":
            return ModeFileExtensionFilter_KWHITELIST
        case "kBlacklist":
            return ModeFileExtensionFilter_KBLACKLIST
        default:
            return ModeFileExtensionFilter_KWHITELIST
    }
}
