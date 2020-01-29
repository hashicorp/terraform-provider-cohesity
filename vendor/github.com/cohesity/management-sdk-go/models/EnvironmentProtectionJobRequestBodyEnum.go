// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentProtectionJobRequestBodyEnum enum
 */
type EnvironmentProtectionJobRequestBodyEnum int

/**
 * Value collection for EnvironmentProtectionJobRequestBodyEnum enum
 */
const (
    EnvironmentProtectionJobRequestBody_KVMWARE EnvironmentProtectionJobRequestBodyEnum = 1 + iota
    EnvironmentProtectionJobRequestBody_KHYPERV
    EnvironmentProtectionJobRequestBody_KSQL
    EnvironmentProtectionJobRequestBody_KVIEW
    EnvironmentProtectionJobRequestBody_KPUPPETEER
    EnvironmentProtectionJobRequestBody_KPHYSICAL
    EnvironmentProtectionJobRequestBody_KPURE
    EnvironmentProtectionJobRequestBody_KAZURE
    EnvironmentProtectionJobRequestBody_KNETAPP
    EnvironmentProtectionJobRequestBody_KAGENT
    EnvironmentProtectionJobRequestBody_KGENERICNAS
    EnvironmentProtectionJobRequestBody_KACROPOLIS
    EnvironmentProtectionJobRequestBody_KPHYSICALFILES
    EnvironmentProtectionJobRequestBody_KISILON
    EnvironmentProtectionJobRequestBody_KGPFS
    EnvironmentProtectionJobRequestBody_KKVM
    EnvironmentProtectionJobRequestBody_KAWS
    EnvironmentProtectionJobRequestBody_KEXCHANGE
    EnvironmentProtectionJobRequestBody_KHYPERVVSS
    EnvironmentProtectionJobRequestBody_KORACLE
    EnvironmentProtectionJobRequestBody_KGCP
    EnvironmentProtectionJobRequestBody_KFLASHBLADE
    EnvironmentProtectionJobRequestBody_KAWSNATIVE
    EnvironmentProtectionJobRequestBody_KVCD
    EnvironmentProtectionJobRequestBody_KO365
    EnvironmentProtectionJobRequestBody_KO365OUTLOOK
    EnvironmentProtectionJobRequestBody_KHYPERFLEX
    EnvironmentProtectionJobRequestBody_KGCPNATIVE
    EnvironmentProtectionJobRequestBody_KAZURENATIVE
    EnvironmentProtectionJobRequestBody_KKUBERNETES
)

func (r EnvironmentProtectionJobRequestBodyEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentProtectionJobRequestBodyEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentProtectionJobRequestBodyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentProtectionJobRequestBodyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentProtectionJobRequestBodyEnum to its string representation
 */
func EnvironmentProtectionJobRequestBodyEnumToValue(environmentProtectionJobRequestBodyEnum EnvironmentProtectionJobRequestBodyEnum) string {
    switch environmentProtectionJobRequestBodyEnum {
        case EnvironmentProtectionJobRequestBody_KVMWARE:
    		return "kVMware"
        case EnvironmentProtectionJobRequestBody_KHYPERV:
    		return "kHyperV"
        case EnvironmentProtectionJobRequestBody_KSQL:
    		return "kSQL"
        case EnvironmentProtectionJobRequestBody_KVIEW:
    		return "kView"
        case EnvironmentProtectionJobRequestBody_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentProtectionJobRequestBody_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentProtectionJobRequestBody_KPURE:
    		return "kPure"
        case EnvironmentProtectionJobRequestBody_KAZURE:
    		return "kAzure"
        case EnvironmentProtectionJobRequestBody_KNETAPP:
    		return "kNetapp"
        case EnvironmentProtectionJobRequestBody_KAGENT:
    		return "kAgent"
        case EnvironmentProtectionJobRequestBody_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentProtectionJobRequestBody_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentProtectionJobRequestBody_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentProtectionJobRequestBody_KISILON:
    		return "kIsilon"
        case EnvironmentProtectionJobRequestBody_KGPFS:
    		return "kGPFS"
        case EnvironmentProtectionJobRequestBody_KKVM:
    		return "kKVM"
        case EnvironmentProtectionJobRequestBody_KAWS:
    		return "kAWS"
        case EnvironmentProtectionJobRequestBody_KEXCHANGE:
    		return "kExchange"
        case EnvironmentProtectionJobRequestBody_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentProtectionJobRequestBody_KORACLE:
    		return "kOracle"
        case EnvironmentProtectionJobRequestBody_KGCP:
    		return "kGCP"
        case EnvironmentProtectionJobRequestBody_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentProtectionJobRequestBody_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentProtectionJobRequestBody_KVCD:
    		return "kVCD"
        case EnvironmentProtectionJobRequestBody_KO365:
    		return "kO365"
        case EnvironmentProtectionJobRequestBody_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentProtectionJobRequestBody_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentProtectionJobRequestBody_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentProtectionJobRequestBody_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentProtectionJobRequestBody_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentProtectionJobRequestBodyEnum Array to its string Array representation
*/
func EnvironmentProtectionJobRequestBodyEnumArrayToValue(environmentProtectionJobRequestBodyEnum []EnvironmentProtectionJobRequestBodyEnum) []string {
    convArray := make([]string,len( environmentProtectionJobRequestBodyEnum))
    for i:=0; i<len(environmentProtectionJobRequestBodyEnum);i++ {
        convArray[i] = EnvironmentProtectionJobRequestBodyEnumToValue(environmentProtectionJobRequestBodyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentProtectionJobRequestBodyEnumFromValue(value string) EnvironmentProtectionJobRequestBodyEnum {
    switch value {
        case "kVMware":
            return EnvironmentProtectionJobRequestBody_KVMWARE
        case "kHyperV":
            return EnvironmentProtectionJobRequestBody_KHYPERV
        case "kSQL":
            return EnvironmentProtectionJobRequestBody_KSQL
        case "kView":
            return EnvironmentProtectionJobRequestBody_KVIEW
        case "kPuppeteer":
            return EnvironmentProtectionJobRequestBody_KPUPPETEER
        case "kPhysical":
            return EnvironmentProtectionJobRequestBody_KPHYSICAL
        case "kPure":
            return EnvironmentProtectionJobRequestBody_KPURE
        case "kAzure":
            return EnvironmentProtectionJobRequestBody_KAZURE
        case "kNetapp":
            return EnvironmentProtectionJobRequestBody_KNETAPP
        case "kAgent":
            return EnvironmentProtectionJobRequestBody_KAGENT
        case "kGenericNas":
            return EnvironmentProtectionJobRequestBody_KGENERICNAS
        case "kAcropolis":
            return EnvironmentProtectionJobRequestBody_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentProtectionJobRequestBody_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentProtectionJobRequestBody_KISILON
        case "kGPFS":
            return EnvironmentProtectionJobRequestBody_KGPFS
        case "kKVM":
            return EnvironmentProtectionJobRequestBody_KKVM
        case "kAWS":
            return EnvironmentProtectionJobRequestBody_KAWS
        case "kExchange":
            return EnvironmentProtectionJobRequestBody_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentProtectionJobRequestBody_KHYPERVVSS
        case "kOracle":
            return EnvironmentProtectionJobRequestBody_KORACLE
        case "kGCP":
            return EnvironmentProtectionJobRequestBody_KGCP
        case "kFlashBlade":
            return EnvironmentProtectionJobRequestBody_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentProtectionJobRequestBody_KAWSNATIVE
        case "kVCD":
            return EnvironmentProtectionJobRequestBody_KVCD
        case "kO365":
            return EnvironmentProtectionJobRequestBody_KO365
        case "kO365Outlook":
            return EnvironmentProtectionJobRequestBody_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentProtectionJobRequestBody_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentProtectionJobRequestBody_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentProtectionJobRequestBody_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentProtectionJobRequestBody_KKUBERNETES
        default:
            return EnvironmentProtectionJobRequestBody_KVMWARE
    }
}
