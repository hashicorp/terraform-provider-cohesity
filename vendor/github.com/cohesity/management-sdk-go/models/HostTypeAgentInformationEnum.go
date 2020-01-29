// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for HostTypeAgentInformationEnum enum
 */
type HostTypeAgentInformationEnum int

/**
 * Value collection for HostTypeAgentInformationEnum enum
 */
const (
    HostTypeAgentInformation_KLINUX HostTypeAgentInformationEnum = 1 + iota
    HostTypeAgentInformation_KWINDOWS
    HostTypeAgentInformation_KAIX
    HostTypeAgentInformation_KSOLARIS
    HostTypeAgentInformation_KSAPHANA
    HostTypeAgentInformation_KOTHER
)

func (r HostTypeAgentInformationEnum) MarshalJSON() ([]byte, error) {
    s := HostTypeAgentInformationEnumToValue(r)
    return json.Marshal(s)
}

func (r *HostTypeAgentInformationEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  HostTypeAgentInformationEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts HostTypeAgentInformationEnum to its string representation
 */
func HostTypeAgentInformationEnumToValue(hostTypeAgentInformationEnum HostTypeAgentInformationEnum) string {
    switch hostTypeAgentInformationEnum {
        case HostTypeAgentInformation_KLINUX:
    		return "kLinux"
        case HostTypeAgentInformation_KWINDOWS:
    		return "kWindows"
        case HostTypeAgentInformation_KAIX:
    		return "kAix"
        case HostTypeAgentInformation_KSOLARIS:
    		return "kSolaris"
        case HostTypeAgentInformation_KSAPHANA:
    		return "kSapHana"
        case HostTypeAgentInformation_KOTHER:
    		return "kOther"
        default:
        	return "kLinux"
    }
}

/**
 * Converts HostTypeAgentInformationEnum Array to its string Array representation
*/
func HostTypeAgentInformationEnumArrayToValue(hostTypeAgentInformationEnum []HostTypeAgentInformationEnum) []string {
    convArray := make([]string,len( hostTypeAgentInformationEnum))
    for i:=0; i<len(hostTypeAgentInformationEnum);i++ {
        convArray[i] = HostTypeAgentInformationEnumToValue(hostTypeAgentInformationEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func HostTypeAgentInformationEnumFromValue(value string) HostTypeAgentInformationEnum {
    switch value {
        case "kLinux":
            return HostTypeAgentInformation_KLINUX
        case "kWindows":
            return HostTypeAgentInformation_KWINDOWS
        case "kAix":
            return HostTypeAgentInformation_KAIX
        case "kSolaris":
            return HostTypeAgentInformation_KSOLARIS
        case "kSapHana":
            return HostTypeAgentInformation_KSAPHANA
        case "kOther":
            return HostTypeAgentInformation_KOTHER
        default:
            return HostTypeAgentInformation_KLINUX
    }
}
