// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ToolsRunningStatusEnum enum
 */
type ToolsRunningStatusEnum int

/**
 * Value collection for ToolsRunningStatusEnum enum
 */
const (
    ToolsRunningStatus_KUNKNOWN ToolsRunningStatusEnum = 1 + iota
    ToolsRunningStatus_KGUESTTOOLSEXECUTINGSCRIPTS
    ToolsRunningStatus_KGUESTTOOLSNOTRUNNING
    ToolsRunningStatus_KGUESTTOOLSRUNNING
)

func (r ToolsRunningStatusEnum) MarshalJSON() ([]byte, error) {
    s := ToolsRunningStatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *ToolsRunningStatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ToolsRunningStatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ToolsRunningStatusEnum to its string representation
 */
func ToolsRunningStatusEnumToValue(toolsRunningStatusEnum ToolsRunningStatusEnum) string {
    switch toolsRunningStatusEnum {
        case ToolsRunningStatus_KUNKNOWN:
    		return "kUnknown"
        case ToolsRunningStatus_KGUESTTOOLSEXECUTINGSCRIPTS:
    		return "kGuestToolsExecutingScripts"
        case ToolsRunningStatus_KGUESTTOOLSNOTRUNNING:
    		return "kGuestToolsNotRunning"
        case ToolsRunningStatus_KGUESTTOOLSRUNNING:
    		return "kGuestToolsRunning"
        default:
        	return "kUnknown"
    }
}

/**
 * Converts ToolsRunningStatusEnum Array to its string Array representation
*/
func ToolsRunningStatusEnumArrayToValue(toolsRunningStatusEnum []ToolsRunningStatusEnum) []string {
    convArray := make([]string,len( toolsRunningStatusEnum))
    for i:=0; i<len(toolsRunningStatusEnum);i++ {
        convArray[i] = ToolsRunningStatusEnumToValue(toolsRunningStatusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ToolsRunningStatusEnumFromValue(value string) ToolsRunningStatusEnum {
    switch value {
        case "kUnknown":
            return ToolsRunningStatus_KUNKNOWN
        case "kGuestToolsExecutingScripts":
            return ToolsRunningStatus_KGUESTTOOLSEXECUTINGSCRIPTS
        case "kGuestToolsNotRunning":
            return ToolsRunningStatus_KGUESTTOOLSNOTRUNNING
        case "kGuestToolsRunning":
            return ToolsRunningStatus_KGUESTTOOLSRUNNING
        default:
            return ToolsRunningStatus_KUNKNOWN
    }
}
