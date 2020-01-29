// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentListProtectionSourcesRootNodesEnum enum
 */
type EnvironmentListProtectionSourcesRootNodesEnum int

/**
 * Value collection for EnvironmentListProtectionSourcesRootNodesEnum enum
 */
const (
    EnvironmentListProtectionSourcesRootNodes_KVMWARE EnvironmentListProtectionSourcesRootNodesEnum = 1 + iota
    EnvironmentListProtectionSourcesRootNodes_KSQL
    EnvironmentListProtectionSourcesRootNodes_KVIEW
    EnvironmentListProtectionSourcesRootNodes_KPUPPETEER
    EnvironmentListProtectionSourcesRootNodes_KPHYSICAL
    EnvironmentListProtectionSourcesRootNodes_KPURE
    EnvironmentListProtectionSourcesRootNodes_KNETAPP
    EnvironmentListProtectionSourcesRootNodes_KGENERICNAS
    EnvironmentListProtectionSourcesRootNodes_KHYPERV
    EnvironmentListProtectionSourcesRootNodes_KACROPOLIS
    EnvironmentListProtectionSourcesRootNodes_KAZURE
)

func (r EnvironmentListProtectionSourcesRootNodesEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentListProtectionSourcesRootNodesEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentListProtectionSourcesRootNodesEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentListProtectionSourcesRootNodesEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentListProtectionSourcesRootNodesEnum to its string representation
 */
func EnvironmentListProtectionSourcesRootNodesEnumToValue(environmentListProtectionSourcesRootNodesEnum EnvironmentListProtectionSourcesRootNodesEnum) string {
    switch environmentListProtectionSourcesRootNodesEnum {
        case EnvironmentListProtectionSourcesRootNodes_KVMWARE:
    		return "kVMware"
        case EnvironmentListProtectionSourcesRootNodes_KSQL:
    		return "kSQL"
        case EnvironmentListProtectionSourcesRootNodes_KVIEW:
    		return "kView"
        case EnvironmentListProtectionSourcesRootNodes_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentListProtectionSourcesRootNodes_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentListProtectionSourcesRootNodes_KPURE:
    		return "kPure"
        case EnvironmentListProtectionSourcesRootNodes_KNETAPP:
    		return "kNetapp"
        case EnvironmentListProtectionSourcesRootNodes_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentListProtectionSourcesRootNodes_KHYPERV:
    		return "kHyperV"
        case EnvironmentListProtectionSourcesRootNodes_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentListProtectionSourcesRootNodes_KAZURE:
    		return "kAzure"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentListProtectionSourcesRootNodesEnum Array to its string Array representation
*/
func EnvironmentListProtectionSourcesRootNodesEnumArrayToValue(environmentListProtectionSourcesRootNodesEnum []EnvironmentListProtectionSourcesRootNodesEnum) []string {
    convArray := make([]string,len( environmentListProtectionSourcesRootNodesEnum))
    for i:=0; i<len(environmentListProtectionSourcesRootNodesEnum);i++ {
        convArray[i] = EnvironmentListProtectionSourcesRootNodesEnumToValue(environmentListProtectionSourcesRootNodesEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentListProtectionSourcesRootNodesEnumFromValue(value string) EnvironmentListProtectionSourcesRootNodesEnum {
    switch value {
        case "kVMware":
            return EnvironmentListProtectionSourcesRootNodes_KVMWARE
        case "kSQL":
            return EnvironmentListProtectionSourcesRootNodes_KSQL
        case "kView":
            return EnvironmentListProtectionSourcesRootNodes_KVIEW
        case "kPuppeteer":
            return EnvironmentListProtectionSourcesRootNodes_KPUPPETEER
        case "kPhysical":
            return EnvironmentListProtectionSourcesRootNodes_KPHYSICAL
        case "kPure":
            return EnvironmentListProtectionSourcesRootNodes_KPURE
        case "kNetapp":
            return EnvironmentListProtectionSourcesRootNodes_KNETAPP
        case "kGenericNas":
            return EnvironmentListProtectionSourcesRootNodes_KGENERICNAS
        case "kHyperV":
            return EnvironmentListProtectionSourcesRootNodes_KHYPERV
        case "kAcropolis":
            return EnvironmentListProtectionSourcesRootNodes_KACROPOLIS
        case "kAzure":
            return EnvironmentListProtectionSourcesRootNodes_KAZURE
        default:
            return EnvironmentListProtectionSourcesRootNodes_KVMWARE
    }
}
