// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for UpgradabilityAgentDeploymentStatusResponseEnum enum
 */
type UpgradabilityAgentDeploymentStatusResponseEnum int

/**
 * Value collection for UpgradabilityAgentDeploymentStatusResponseEnum enum
 */
const (
    UpgradabilityAgentDeploymentStatusResponse_KUPGRADABLE UpgradabilityAgentDeploymentStatusResponseEnum = 1 + iota
    UpgradabilityAgentDeploymentStatusResponse_KCURRENT
    UpgradabilityAgentDeploymentStatusResponse_KUNKNOWN
    UpgradabilityAgentDeploymentStatusResponse_KNONUPGRADABLEINVALIDVERSION
    UpgradabilityAgentDeploymentStatusResponse_KNONUPGRADABLEAGENTISNEWER
    UpgradabilityAgentDeploymentStatusResponse_KNONUPGRADABLEAGENTISOLD
)

func (r UpgradabilityAgentDeploymentStatusResponseEnum) MarshalJSON() ([]byte, error) {
    s := UpgradabilityAgentDeploymentStatusResponseEnumToValue(r)
    return json.Marshal(s)
}

func (r *UpgradabilityAgentDeploymentStatusResponseEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  UpgradabilityAgentDeploymentStatusResponseEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts UpgradabilityAgentDeploymentStatusResponseEnum to its string representation
 */
func UpgradabilityAgentDeploymentStatusResponseEnumToValue(upgradabilityAgentDeploymentStatusResponseEnum UpgradabilityAgentDeploymentStatusResponseEnum) string {
    switch upgradabilityAgentDeploymentStatusResponseEnum {
        case UpgradabilityAgentDeploymentStatusResponse_KUPGRADABLE:
    		return "kUpgradable"
        case UpgradabilityAgentDeploymentStatusResponse_KCURRENT:
    		return "kCurrent"
        case UpgradabilityAgentDeploymentStatusResponse_KUNKNOWN:
    		return "kUnknown"
        case UpgradabilityAgentDeploymentStatusResponse_KNONUPGRADABLEINVALIDVERSION:
    		return "kNonUpgradableInvalidVersion"
        case UpgradabilityAgentDeploymentStatusResponse_KNONUPGRADABLEAGENTISNEWER:
    		return "kNonUpgradableAgentIsNewer"
        case UpgradabilityAgentDeploymentStatusResponse_KNONUPGRADABLEAGENTISOLD:
    		return "kNonUpgradableAgentIsOld"
        default:
        	return "kUpgradable"
    }
}

/**
 * Converts UpgradabilityAgentDeploymentStatusResponseEnum Array to its string Array representation
*/
func UpgradabilityAgentDeploymentStatusResponseEnumArrayToValue(upgradabilityAgentDeploymentStatusResponseEnum []UpgradabilityAgentDeploymentStatusResponseEnum) []string {
    convArray := make([]string,len( upgradabilityAgentDeploymentStatusResponseEnum))
    for i:=0; i<len(upgradabilityAgentDeploymentStatusResponseEnum);i++ {
        convArray[i] = UpgradabilityAgentDeploymentStatusResponseEnumToValue(upgradabilityAgentDeploymentStatusResponseEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func UpgradabilityAgentDeploymentStatusResponseEnumFromValue(value string) UpgradabilityAgentDeploymentStatusResponseEnum {
    switch value {
        case "kUpgradable":
            return UpgradabilityAgentDeploymentStatusResponse_KUPGRADABLE
        case "kCurrent":
            return UpgradabilityAgentDeploymentStatusResponse_KCURRENT
        case "kUnknown":
            return UpgradabilityAgentDeploymentStatusResponse_KUNKNOWN
        case "kNonUpgradableInvalidVersion":
            return UpgradabilityAgentDeploymentStatusResponse_KNONUPGRADABLEINVALIDVERSION
        case "kNonUpgradableAgentIsNewer":
            return UpgradabilityAgentDeploymentStatusResponse_KNONUPGRADABLEAGENTISNEWER
        case "kNonUpgradableAgentIsOld":
            return UpgradabilityAgentDeploymentStatusResponse_KNONUPGRADABLEAGENTISOLD
        default:
            return UpgradabilityAgentDeploymentStatusResponse_KUPGRADABLE
    }
}
