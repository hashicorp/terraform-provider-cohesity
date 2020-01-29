// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HostTypeDownloadPhysicalAgentEnum enum
 */
type HostTypeDownloadPhysicalAgentEnum int

/**
 * Value collection for HostTypeDownloadPhysicalAgentEnum enum
 */
const (
    HostTypeDownloadPhysicalAgent_KLINUX HostTypeDownloadPhysicalAgentEnum = 1 + iota
    HostTypeDownloadPhysicalAgent_KWINDOWS
    HostTypeDownloadPhysicalAgent_KAIX
    HostTypeDownloadPhysicalAgent_KSOLARIS
    HostTypeDownloadPhysicalAgent_KSAPHANA
    HostTypeDownloadPhysicalAgent_KOTHER
)

func (r HostTypeDownloadPhysicalAgentEnum) MarshalJSON() ([]byte, error) {
    s := HostTypeDownloadPhysicalAgentEnumToValue(r)
    return json.Marshal(s)
}

func (r *HostTypeDownloadPhysicalAgentEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HostTypeDownloadPhysicalAgentEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HostTypeDownloadPhysicalAgentEnum to its string representation
 */
func HostTypeDownloadPhysicalAgentEnumToValue(hostTypeDownloadPhysicalAgentEnum HostTypeDownloadPhysicalAgentEnum) string {
    switch hostTypeDownloadPhysicalAgentEnum {
        case HostTypeDownloadPhysicalAgent_KLINUX:
    		return "kLinux"
        case HostTypeDownloadPhysicalAgent_KWINDOWS:
    		return "kWindows"
        case HostTypeDownloadPhysicalAgent_KAIX:
    		return "kAix"
        case HostTypeDownloadPhysicalAgent_KSOLARIS:
    		return "kSolaris"
        case HostTypeDownloadPhysicalAgent_KSAPHANA:
    		return "kSapHana"
        case HostTypeDownloadPhysicalAgent_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostTypeDownloadPhysicalAgentEnum Array to its string Array representation
*/
func HostTypeDownloadPhysicalAgentEnumArrayToValue(hostTypeDownloadPhysicalAgentEnum []HostTypeDownloadPhysicalAgentEnum) []string {
    convArray := make([]string,len( hostTypeDownloadPhysicalAgentEnum))
    for i:=0; i<len(hostTypeDownloadPhysicalAgentEnum);i++ {
        convArray[i] = HostTypeDownloadPhysicalAgentEnumToValue(hostTypeDownloadPhysicalAgentEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostTypeDownloadPhysicalAgentEnumFromValue(value string) HostTypeDownloadPhysicalAgentEnum {
    switch value {
        case "kLinux":
            return HostTypeDownloadPhysicalAgent_KLINUX
        case "kWindows":
            return HostTypeDownloadPhysicalAgent_KWINDOWS
        case "kAix":
            return HostTypeDownloadPhysicalAgent_KAIX
        case "kSolaris":
            return HostTypeDownloadPhysicalAgent_KSOLARIS
        case "kSapHana":
            return HostTypeDownloadPhysicalAgent_KSAPHANA
        case "kOther":
            return HostTypeDownloadPhysicalAgent_KOTHER
        default:
            return HostTypeDownloadPhysicalAgent_KLINUX
    }
}
