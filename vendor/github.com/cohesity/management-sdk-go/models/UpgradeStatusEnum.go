// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for UpgradeStatusEnum enum
 */
type UpgradeStatusEnum int

/**
 * Value collection for UpgradeStatusEnum enum
 */
const (
    UpgradeStatus_KIDLE UpgradeStatusEnum = 1 + iota
    UpgradeStatus_KACCEPTED
    UpgradeStatus_KSTARTED
    UpgradeStatus_KFINISHED
)

func (r UpgradeStatusEnum) MarshalJSON() ([]byte, error) {
    s := UpgradeStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *UpgradeStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  UpgradeStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts UpgradeStatusEnum to its string representation
 */
func UpgradeStatusEnumToValue(upgradeStatusEnum UpgradeStatusEnum) string {
    switch upgradeStatusEnum {
        case UpgradeStatus_KIDLE:
    		return "kIdle"
        case UpgradeStatus_KACCEPTED:
    		return "kAccepted"
        case UpgradeStatus_KSTARTED:
    		return "kStarted"
        case UpgradeStatus_KFINISHED:
    		return "kFinished"
        default:
        	return "kIdle"
    }
}

/**
 * Converts UpgradeStatusEnum Array to its string Array representation
*/
func UpgradeStatusEnumArrayToValue(upgradeStatusEnum []UpgradeStatusEnum) []string {
    convArray := make([]string,len( upgradeStatusEnum))
    for i:=0; i<len(upgradeStatusEnum);i++ {
        convArray[i] = UpgradeStatusEnumToValue(upgradeStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func UpgradeStatusEnumFromValue(value string) UpgradeStatusEnum {
    switch value {
        case "kIdle":
            return UpgradeStatus_KIDLE
        case "kAccepted":
            return UpgradeStatus_KACCEPTED
        case "kStarted":
            return UpgradeStatus_KSTARTED
        case "kFinished":
            return UpgradeStatus_KFINISHED
        default:
            return UpgradeStatus_KIDLE
    }
}
