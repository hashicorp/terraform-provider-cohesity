// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ActionUpdateProtectionJobsStateParamsEnum enum
 */
type ActionUpdateProtectionJobsStateParamsEnum int

/**
 * Value collection for ActionUpdateProtectionJobsStateParamsEnum enum
 */
const (
    ActionUpdateProtectionJobsStateParams_KACTIVATE ActionUpdateProtectionJobsStateParamsEnum = 1 + iota
    ActionUpdateProtectionJobsStateParams_KDEACTIVATE
    ActionUpdateProtectionJobsStateParams_KPAUSE
    ActionUpdateProtectionJobsStateParams_KRESUME
)

func (r ActionUpdateProtectionJobsStateParamsEnum) MarshalJSON() ([]byte, error) {
    s := ActionUpdateProtectionJobsStateParamsEnumToValue(r)
    return json.Marshal(s)
}

func (r *ActionUpdateProtectionJobsStateParamsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ActionUpdateProtectionJobsStateParamsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ActionUpdateProtectionJobsStateParamsEnum to its string representation
 */
func ActionUpdateProtectionJobsStateParamsEnumToValue(actionUpdateProtectionJobsStateParamsEnum ActionUpdateProtectionJobsStateParamsEnum) string {
    switch actionUpdateProtectionJobsStateParamsEnum {
        case ActionUpdateProtectionJobsStateParams_KACTIVATE:
    		return "kActivate"
        case ActionUpdateProtectionJobsStateParams_KDEACTIVATE:
    		return "kDeactivate"
        case ActionUpdateProtectionJobsStateParams_KPAUSE:
    		return "kPause"
        case ActionUpdateProtectionJobsStateParams_KRESUME:
    		return "kResume"
        default:
        	return "kActivate"
    }
}

/**
 * Converts ActionUpdateProtectionJobsStateParamsEnum Array to its string Array representation
*/
func ActionUpdateProtectionJobsStateParamsEnumArrayToValue(actionUpdateProtectionJobsStateParamsEnum []ActionUpdateProtectionJobsStateParamsEnum) []string {
    convArray := make([]string,len( actionUpdateProtectionJobsStateParamsEnum))
    for i:=0; i<len(actionUpdateProtectionJobsStateParamsEnum);i++ {
        convArray[i] = ActionUpdateProtectionJobsStateParamsEnumToValue(actionUpdateProtectionJobsStateParamsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ActionUpdateProtectionJobsStateParamsEnumFromValue(value string) ActionUpdateProtectionJobsStateParamsEnum {
    switch value {
        case "kActivate":
            return ActionUpdateProtectionJobsStateParams_KACTIVATE
        case "kDeactivate":
            return ActionUpdateProtectionJobsStateParams_KDEACTIVATE
        case "kPause":
            return ActionUpdateProtectionJobsStateParams_KPAUSE
        case "kResume":
            return ActionUpdateProtectionJobsStateParams_KRESUME
        default:
            return ActionUpdateProtectionJobsStateParams_KACTIVATE
    }
}
