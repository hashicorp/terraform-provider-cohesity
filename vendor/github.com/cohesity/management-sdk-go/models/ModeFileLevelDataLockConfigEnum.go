// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ModeFileLevelDataLockConfigEnum enum
 */
type ModeFileLevelDataLockConfigEnum int

/**
 * Value collection for ModeFileLevelDataLockConfigEnum enum
 */
const (
    ModeFileLevelDataLockConfig_KCOMPLIANCE ModeFileLevelDataLockConfigEnum = 1 + iota
    ModeFileLevelDataLockConfig_KENTERPRISE
)

func (r ModeFileLevelDataLockConfigEnum) MarshalJSON() ([]byte, error) {
    s := ModeFileLevelDataLockConfigEnumToValue(r)
    return json.Marshal(s)
}

func (r *ModeFileLevelDataLockConfigEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ModeFileLevelDataLockConfigEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ModeFileLevelDataLockConfigEnum to its string representation
 */
func ModeFileLevelDataLockConfigEnumToValue(modeFileLevelDataLockConfigEnum ModeFileLevelDataLockConfigEnum) string {
    switch modeFileLevelDataLockConfigEnum {
        case ModeFileLevelDataLockConfig_KCOMPLIANCE:
    		return "kCompliance"
        case ModeFileLevelDataLockConfig_KENTERPRISE:
    		return "kEnterprise"
        default:
        	return "kCompliance"
    }
}

/**
 * Converts ModeFileLevelDataLockConfigEnum Array to its string Array representation
*/
func ModeFileLevelDataLockConfigEnumArrayToValue(modeFileLevelDataLockConfigEnum []ModeFileLevelDataLockConfigEnum) []string {
    convArray := make([]string,len( modeFileLevelDataLockConfigEnum))
    for i:=0; i<len(modeFileLevelDataLockConfigEnum);i++ {
        convArray[i] = ModeFileLevelDataLockConfigEnumToValue(modeFileLevelDataLockConfigEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ModeFileLevelDataLockConfigEnumFromValue(value string) ModeFileLevelDataLockConfigEnum {
    switch value {
        case "kCompliance":
            return ModeFileLevelDataLockConfig_KCOMPLIANCE
        case "kEnterprise":
            return ModeFileLevelDataLockConfig_KENTERPRISE
        default:
            return ModeFileLevelDataLockConfig_KCOMPLIANCE
    }
}
