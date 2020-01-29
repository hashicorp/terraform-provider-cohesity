// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ActionEnum enum
 */
type ActionEnum int

/**
 * Value collection for ActionEnum enum
 */
const (
    Action_KSTOP ActionEnum = 1 + iota
    Action_KSTART
    Action_KRESTART
)

func (r ActionEnum) MarshalJSON() ([]byte, error) {
    s := ActionEnumToValue(r)
    return json.Marshal(s)
}

func (r *ActionEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ActionEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ActionEnum to its string representation
 */
func ActionEnumToValue(actionEnum ActionEnum) string {
    switch actionEnum {
        case Action_KSTOP:
    		return "kStop"
        case Action_KSTART:
    		return "kStart"
        case Action_KRESTART:
    		return "kRestart"
        default:
        	return "kStop"
    }
}

/**
 * Converts ActionEnum Array to its string Array representation
*/
func ActionEnumArrayToValue(actionEnum []ActionEnum) []string {
    convArray := make([]string,len( actionEnum))
    for i:=0; i<len(actionEnum);i++ {
        convArray[i] = ActionEnumToValue(actionEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ActionEnumFromValue(value string) ActionEnum {
    switch value {
        case "kStop":
            return Action_KSTOP
        case "kStart":
            return Action_KSTART
        case "kRestart":
            return Action_KRESTART
        default:
            return Action_KSTOP
    }
}
