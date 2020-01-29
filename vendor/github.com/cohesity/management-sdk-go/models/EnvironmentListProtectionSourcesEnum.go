// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentListProtectionSourcesEnum enum
 */
type EnvironmentListProtectionSourcesEnum int

/**
 * Value collection for EnvironmentListProtectionSourcesEnum enum
 */
const (
    EnvironmentListProtectionSources_KVMWARE EnvironmentListProtectionSourcesEnum = 1 + iota
    EnvironmentListProtectionSources_KSQL
    EnvironmentListProtectionSources_KVIEW
    EnvironmentListProtectionSources_KPUPPETEER
    EnvironmentListProtectionSources_KPHYSICAL
    EnvironmentListProtectionSources_KPURE
    EnvironmentListProtectionSources_KNETAPP
    EnvironmentListProtectionSources_KGENERICNAS
    EnvironmentListProtectionSources_KHYPERV
    EnvironmentListProtectionSources_KACROPOLIS
    EnvironmentListProtectionSources_KAZURE
)

func (r EnvironmentListProtectionSourcesEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentListProtectionSourcesEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentListProtectionSourcesEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentListProtectionSourcesEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentListProtectionSourcesEnum to its string representation
 */
func EnvironmentListProtectionSourcesEnumToValue(environmentListProtectionSourcesEnum EnvironmentListProtectionSourcesEnum) string {
    switch environmentListProtectionSourcesEnum {
        case EnvironmentListProtectionSources_KVMWARE:
    		return "kVMware"
        case EnvironmentListProtectionSources_KSQL:
    		return "kSQL"
        case EnvironmentListProtectionSources_KVIEW:
    		return "kView"
        case EnvironmentListProtectionSources_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentListProtectionSources_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentListProtectionSources_KPURE:
    		return "kPure"
        case EnvironmentListProtectionSources_KNETAPP:
    		return "kNetapp"
        case EnvironmentListProtectionSources_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentListProtectionSources_KHYPERV:
    		return "kHyperV"
        case EnvironmentListProtectionSources_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentListProtectionSources_KAZURE:
    		return "kAzure"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentListProtectionSourcesEnum Array to its string Array representation
*/
func EnvironmentListProtectionSourcesEnumArrayToValue(environmentListProtectionSourcesEnum []EnvironmentListProtectionSourcesEnum) []string {
    convArray := make([]string,len( environmentListProtectionSourcesEnum))
    for i:=0; i<len(environmentListProtectionSourcesEnum);i++ {
        convArray[i] = EnvironmentListProtectionSourcesEnumToValue(environmentListProtectionSourcesEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentListProtectionSourcesEnumFromValue(value string) EnvironmentListProtectionSourcesEnum {
    switch value {
        case "kVMware":
            return EnvironmentListProtectionSources_KVMWARE
        case "kSQL":
            return EnvironmentListProtectionSources_KSQL
        case "kView":
            return EnvironmentListProtectionSources_KVIEW
        case "kPuppeteer":
            return EnvironmentListProtectionSources_KPUPPETEER
        case "kPhysical":
            return EnvironmentListProtectionSources_KPHYSICAL
        case "kPure":
            return EnvironmentListProtectionSources_KPURE
        case "kNetapp":
            return EnvironmentListProtectionSources_KNETAPP
        case "kGenericNas":
            return EnvironmentListProtectionSources_KGENERICNAS
        case "kHyperV":
            return EnvironmentListProtectionSources_KHYPERV
        case "kAcropolis":
            return EnvironmentListProtectionSources_KACROPOLIS
        case "kAzure":
            return EnvironmentListProtectionSources_KAZURE
        default:
            return EnvironmentListProtectionSources_KVMWARE
    }
}
