// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentListProtectionSourcesRegistrationInfoEnum enum
 */
type EnvironmentListProtectionSourcesRegistrationInfoEnum int

/**
 * Value collection for EnvironmentListProtectionSourcesRegistrationInfoEnum enum
 */
const (
    EnvironmentListProtectionSourcesRegistrationInfo_KVMWARE EnvironmentListProtectionSourcesRegistrationInfoEnum = 1 + iota
    EnvironmentListProtectionSourcesRegistrationInfo_KSQL
    EnvironmentListProtectionSourcesRegistrationInfo_KVIEW
    EnvironmentListProtectionSourcesRegistrationInfo_KPUPPETEER
    EnvironmentListProtectionSourcesRegistrationInfo_KPHYSICAL
    EnvironmentListProtectionSourcesRegistrationInfo_KPURE
    EnvironmentListProtectionSourcesRegistrationInfo_KNETAPP
    EnvironmentListProtectionSourcesRegistrationInfo_KGENERICNAS
    EnvironmentListProtectionSourcesRegistrationInfo_KHYPERV
    EnvironmentListProtectionSourcesRegistrationInfo_KACROPOLIS
    EnvironmentListProtectionSourcesRegistrationInfo_KAZURE
)

func (r EnvironmentListProtectionSourcesRegistrationInfoEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentListProtectionSourcesRegistrationInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentListProtectionSourcesRegistrationInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentListProtectionSourcesRegistrationInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentListProtectionSourcesRegistrationInfoEnum to its string representation
 */
func EnvironmentListProtectionSourcesRegistrationInfoEnumToValue(environmentListProtectionSourcesRegistrationInfoEnum EnvironmentListProtectionSourcesRegistrationInfoEnum) string {
    switch environmentListProtectionSourcesRegistrationInfoEnum {
        case EnvironmentListProtectionSourcesRegistrationInfo_KVMWARE:
    		return "kVMware"
        case EnvironmentListProtectionSourcesRegistrationInfo_KSQL:
    		return "kSQL"
        case EnvironmentListProtectionSourcesRegistrationInfo_KVIEW:
    		return "kView"
        case EnvironmentListProtectionSourcesRegistrationInfo_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentListProtectionSourcesRegistrationInfo_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentListProtectionSourcesRegistrationInfo_KPURE:
    		return "kPure"
        case EnvironmentListProtectionSourcesRegistrationInfo_KNETAPP:
    		return "kNetapp"
        case EnvironmentListProtectionSourcesRegistrationInfo_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentListProtectionSourcesRegistrationInfo_KHYPERV:
    		return "kHyperV"
        case EnvironmentListProtectionSourcesRegistrationInfo_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentListProtectionSourcesRegistrationInfo_KAZURE:
    		return "kAzure"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentListProtectionSourcesRegistrationInfoEnum Array to its string Array representation
*/
func EnvironmentListProtectionSourcesRegistrationInfoEnumArrayToValue(environmentListProtectionSourcesRegistrationInfoEnum []EnvironmentListProtectionSourcesRegistrationInfoEnum) []string {
    convArray := make([]string,len( environmentListProtectionSourcesRegistrationInfoEnum))
    for i:=0; i<len(environmentListProtectionSourcesRegistrationInfoEnum);i++ {
        convArray[i] = EnvironmentListProtectionSourcesRegistrationInfoEnumToValue(environmentListProtectionSourcesRegistrationInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentListProtectionSourcesRegistrationInfoEnumFromValue(value string) EnvironmentListProtectionSourcesRegistrationInfoEnum {
    switch value {
        case "kVMware":
            return EnvironmentListProtectionSourcesRegistrationInfo_KVMWARE
        case "kSQL":
            return EnvironmentListProtectionSourcesRegistrationInfo_KSQL
        case "kView":
            return EnvironmentListProtectionSourcesRegistrationInfo_KVIEW
        case "kPuppeteer":
            return EnvironmentListProtectionSourcesRegistrationInfo_KPUPPETEER
        case "kPhysical":
            return EnvironmentListProtectionSourcesRegistrationInfo_KPHYSICAL
        case "kPure":
            return EnvironmentListProtectionSourcesRegistrationInfo_KPURE
        case "kNetapp":
            return EnvironmentListProtectionSourcesRegistrationInfo_KNETAPP
        case "kGenericNas":
            return EnvironmentListProtectionSourcesRegistrationInfo_KGENERICNAS
        case "kHyperV":
            return EnvironmentListProtectionSourcesRegistrationInfo_KHYPERV
        case "kAcropolis":
            return EnvironmentListProtectionSourcesRegistrationInfo_KACROPOLIS
        case "kAzure":
            return EnvironmentListProtectionSourcesRegistrationInfo_KAZURE
        default:
            return EnvironmentListProtectionSourcesRegistrationInfo_KVMWARE
    }
}
