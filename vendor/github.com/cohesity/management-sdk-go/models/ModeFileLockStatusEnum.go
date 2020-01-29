// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ModeFileLockStatusEnum enum
 */
type ModeFileLockStatusEnum int

/**
 * Value collection for ModeFileLockStatusEnum enum
 */
const (
    ModeFileLockStatus_KCOMPLIANCE ModeFileLockStatusEnum = 1 + iota
    ModeFileLockStatus_KENTERPRISE
)

func (r ModeFileLockStatusEnum) MarshalJSON() ([]byte, error) {
    s := ModeFileLockStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *ModeFileLockStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ModeFileLockStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ModeFileLockStatusEnum to its string representation
 */
func ModeFileLockStatusEnumToValue(modeFileLockStatusEnum ModeFileLockStatusEnum) string {
    switch modeFileLockStatusEnum {
        case ModeFileLockStatus_KCOMPLIANCE:
    		return "kCompliance"
        case ModeFileLockStatus_KENTERPRISE:
    		return "kEnterprise"
        default:
        	return "kCompliance"
    }
}

/**
 * Converts ModeFileLockStatusEnum Array to its string Array representation
*/
func ModeFileLockStatusEnumArrayToValue(modeFileLockStatusEnum []ModeFileLockStatusEnum) []string {
    convArray := make([]string,len( modeFileLockStatusEnum))
    for i:=0; i<len(modeFileLockStatusEnum);i++ {
        convArray[i] = ModeFileLockStatusEnumToValue(modeFileLockStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ModeFileLockStatusEnumFromValue(value string) ModeFileLockStatusEnum {
    switch value {
        case "kCompliance":
            return ModeFileLockStatus_KCOMPLIANCE
        case "kEnterprise":
            return ModeFileLockStatus_KENTERPRISE
        default:
            return ModeFileLockStatus_KCOMPLIANCE
    }
}
