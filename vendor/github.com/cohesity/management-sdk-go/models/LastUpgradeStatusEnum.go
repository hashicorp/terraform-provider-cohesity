// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for LastUpgradeStatusEnum enum
 */
type LastUpgradeStatusEnum int

/**
 * Value collection for LastUpgradeStatusEnum enum
 */
const (
    LastUpgradeStatus_KIDLE LastUpgradeStatusEnum = 1 + iota
    LastUpgradeStatus_KACCEPTED
    LastUpgradeStatus_KSTARTED
    LastUpgradeStatus_KFINISHED
)

func (r LastUpgradeStatusEnum) MarshalJSON() ([]byte, error) {
    s := LastUpgradeStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *LastUpgradeStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  LastUpgradeStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts LastUpgradeStatusEnum to its string representation
 */
func LastUpgradeStatusEnumToValue(lastUpgradeStatusEnum LastUpgradeStatusEnum) string {
    switch lastUpgradeStatusEnum {
        case LastUpgradeStatus_KIDLE:
    		return "kIdle"
        case LastUpgradeStatus_KACCEPTED:
    		return "kAccepted"
        case LastUpgradeStatus_KSTARTED:
    		return "kStarted"
        case LastUpgradeStatus_KFINISHED:
    		return "kFinished"
        default:
        	return "kIdle"
    }
}

/**
 * Converts LastUpgradeStatusEnum Array to its string Array representation
*/
func LastUpgradeStatusEnumArrayToValue(lastUpgradeStatusEnum []LastUpgradeStatusEnum) []string {
    convArray := make([]string,len( lastUpgradeStatusEnum))
    for i:=0; i<len(lastUpgradeStatusEnum);i++ {
        convArray[i] = LastUpgradeStatusEnumToValue(lastUpgradeStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func LastUpgradeStatusEnumFromValue(value string) LastUpgradeStatusEnum {
    switch value {
        case "kIdle":
            return LastUpgradeStatus_KIDLE
        case "kAccepted":
            return LastUpgradeStatus_KACCEPTED
        case "kStarted":
            return LastUpgradeStatus_KSTARTED
        case "kFinished":
            return LastUpgradeStatus_KFINISHED
        default:
            return LastUpgradeStatus_KIDLE
    }
}
