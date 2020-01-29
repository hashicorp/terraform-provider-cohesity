// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentProtectionSummaryByEnvEnum enum
 */
type EnvironmentProtectionSummaryByEnvEnum int

/**
 * Value collection for EnvironmentProtectionSummaryByEnvEnum enum
 */
const (
    EnvironmentProtectionSummaryByEnv_KVMWARE EnvironmentProtectionSummaryByEnvEnum = 1 + iota
    EnvironmentProtectionSummaryByEnv_KHYPERV
    EnvironmentProtectionSummaryByEnv_KSQL
    EnvironmentProtectionSummaryByEnv_KVIEW
    EnvironmentProtectionSummaryByEnv_KPUPPETEER
    EnvironmentProtectionSummaryByEnv_KPHYSICAL
    EnvironmentProtectionSummaryByEnv_KPURE
    EnvironmentProtectionSummaryByEnv_KAZURE
    EnvironmentProtectionSummaryByEnv_KNETAPP
    EnvironmentProtectionSummaryByEnv_KAGENT
    EnvironmentProtectionSummaryByEnv_KGENERICNAS
    EnvironmentProtectionSummaryByEnv_KACROPOLIS
    EnvironmentProtectionSummaryByEnv_KPHYSICALFILES
    EnvironmentProtectionSummaryByEnv_KISILON
    EnvironmentProtectionSummaryByEnv_KGPFS
    EnvironmentProtectionSummaryByEnv_KKVM
    EnvironmentProtectionSummaryByEnv_KAWS
    EnvironmentProtectionSummaryByEnv_KEXCHANGE
    EnvironmentProtectionSummaryByEnv_KHYPERVVSS
    EnvironmentProtectionSummaryByEnv_KORACLE
    EnvironmentProtectionSummaryByEnv_KGCP
    EnvironmentProtectionSummaryByEnv_KFLASHBLADE
    EnvironmentProtectionSummaryByEnv_KAWSNATIVE
    EnvironmentProtectionSummaryByEnv_KVCD
    EnvironmentProtectionSummaryByEnv_KO365
    EnvironmentProtectionSummaryByEnv_KO365OUTLOOK
    EnvironmentProtectionSummaryByEnv_KHYPERFLEX
    EnvironmentProtectionSummaryByEnv_KGCPNATIVE
    EnvironmentProtectionSummaryByEnv_KAZURENATIVE
    EnvironmentProtectionSummaryByEnv_KKUBERNETES
)

func (r EnvironmentProtectionSummaryByEnvEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentProtectionSummaryByEnvEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentProtectionSummaryByEnvEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentProtectionSummaryByEnvEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentProtectionSummaryByEnvEnum to its string representation
 */
func EnvironmentProtectionSummaryByEnvEnumToValue(environmentProtectionSummaryByEnvEnum EnvironmentProtectionSummaryByEnvEnum) string {
    switch environmentProtectionSummaryByEnvEnum {
        case EnvironmentProtectionSummaryByEnv_KVMWARE:
    		return "kVMware"
        case EnvironmentProtectionSummaryByEnv_KHYPERV:
    		return "kHyperV"
        case EnvironmentProtectionSummaryByEnv_KSQL:
    		return "kSQL"
        case EnvironmentProtectionSummaryByEnv_KVIEW:
    		return "kView"
        case EnvironmentProtectionSummaryByEnv_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentProtectionSummaryByEnv_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentProtectionSummaryByEnv_KPURE:
    		return "kPure"
        case EnvironmentProtectionSummaryByEnv_KAZURE:
    		return "kAzure"
        case EnvironmentProtectionSummaryByEnv_KNETAPP:
    		return "kNetapp"
        case EnvironmentProtectionSummaryByEnv_KAGENT:
    		return "kAgent"
        case EnvironmentProtectionSummaryByEnv_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentProtectionSummaryByEnv_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentProtectionSummaryByEnv_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentProtectionSummaryByEnv_KISILON:
    		return "kIsilon"
        case EnvironmentProtectionSummaryByEnv_KGPFS:
    		return "kGPFS"
        case EnvironmentProtectionSummaryByEnv_KKVM:
    		return "kKVM"
        case EnvironmentProtectionSummaryByEnv_KAWS:
    		return "kAWS"
        case EnvironmentProtectionSummaryByEnv_KEXCHANGE:
    		return "kExchange"
        case EnvironmentProtectionSummaryByEnv_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentProtectionSummaryByEnv_KORACLE:
    		return "kOracle"
        case EnvironmentProtectionSummaryByEnv_KGCP:
    		return "kGCP"
        case EnvironmentProtectionSummaryByEnv_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentProtectionSummaryByEnv_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentProtectionSummaryByEnv_KVCD:
    		return "kVCD"
        case EnvironmentProtectionSummaryByEnv_KO365:
    		return "kO365"
        case EnvironmentProtectionSummaryByEnv_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentProtectionSummaryByEnv_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentProtectionSummaryByEnv_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentProtectionSummaryByEnv_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentProtectionSummaryByEnv_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentProtectionSummaryByEnvEnum Array to its string Array representation
*/
func EnvironmentProtectionSummaryByEnvEnumArrayToValue(environmentProtectionSummaryByEnvEnum []EnvironmentProtectionSummaryByEnvEnum) []string {
    convArray := make([]string,len( environmentProtectionSummaryByEnvEnum))
    for i:=0; i<len(environmentProtectionSummaryByEnvEnum);i++ {
        convArray[i] = EnvironmentProtectionSummaryByEnvEnumToValue(environmentProtectionSummaryByEnvEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentProtectionSummaryByEnvEnumFromValue(value string) EnvironmentProtectionSummaryByEnvEnum {
    switch value {
        case "kVMware":
            return EnvironmentProtectionSummaryByEnv_KVMWARE
        case "kHyperV":
            return EnvironmentProtectionSummaryByEnv_KHYPERV
        case "kSQL":
            return EnvironmentProtectionSummaryByEnv_KSQL
        case "kView":
            return EnvironmentProtectionSummaryByEnv_KVIEW
        case "kPuppeteer":
            return EnvironmentProtectionSummaryByEnv_KPUPPETEER
        case "kPhysical":
            return EnvironmentProtectionSummaryByEnv_KPHYSICAL
        case "kPure":
            return EnvironmentProtectionSummaryByEnv_KPURE
        case "kAzure":
            return EnvironmentProtectionSummaryByEnv_KAZURE
        case "kNetapp":
            return EnvironmentProtectionSummaryByEnv_KNETAPP
        case "kAgent":
            return EnvironmentProtectionSummaryByEnv_KAGENT
        case "kGenericNas":
            return EnvironmentProtectionSummaryByEnv_KGENERICNAS
        case "kAcropolis":
            return EnvironmentProtectionSummaryByEnv_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentProtectionSummaryByEnv_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentProtectionSummaryByEnv_KISILON
        case "kGPFS":
            return EnvironmentProtectionSummaryByEnv_KGPFS
        case "kKVM":
            return EnvironmentProtectionSummaryByEnv_KKVM
        case "kAWS":
            return EnvironmentProtectionSummaryByEnv_KAWS
        case "kExchange":
            return EnvironmentProtectionSummaryByEnv_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentProtectionSummaryByEnv_KHYPERVVSS
        case "kOracle":
            return EnvironmentProtectionSummaryByEnv_KORACLE
        case "kGCP":
            return EnvironmentProtectionSummaryByEnv_KGCP
        case "kFlashBlade":
            return EnvironmentProtectionSummaryByEnv_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentProtectionSummaryByEnv_KAWSNATIVE
        case "kVCD":
            return EnvironmentProtectionSummaryByEnv_KVCD
        case "kO365":
            return EnvironmentProtectionSummaryByEnv_KO365
        case "kO365Outlook":
            return EnvironmentProtectionSummaryByEnv_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentProtectionSummaryByEnv_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentProtectionSummaryByEnv_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentProtectionSummaryByEnv_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentProtectionSummaryByEnv_KKUBERNETES
        default:
            return EnvironmentProtectionSummaryByEnv_KVMWARE
    }
}
