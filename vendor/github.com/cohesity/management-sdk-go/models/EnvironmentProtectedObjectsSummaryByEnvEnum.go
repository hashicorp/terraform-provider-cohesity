// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentProtectedObjectsSummaryByEnvEnum enum
 */
type EnvironmentProtectedObjectsSummaryByEnvEnum int

/**
 * Value collection for EnvironmentProtectedObjectsSummaryByEnvEnum enum
 */
const (
    EnvironmentProtectedObjectsSummaryByEnv_KVMWARE EnvironmentProtectedObjectsSummaryByEnvEnum = 1 + iota
    EnvironmentProtectedObjectsSummaryByEnv_KHYPERV
    EnvironmentProtectedObjectsSummaryByEnv_KSQL
    EnvironmentProtectedObjectsSummaryByEnv_KVIEW
    EnvironmentProtectedObjectsSummaryByEnv_KPUPPETEER
    EnvironmentProtectedObjectsSummaryByEnv_KPHYSICAL
    EnvironmentProtectedObjectsSummaryByEnv_KPURE
    EnvironmentProtectedObjectsSummaryByEnv_KAZURE
    EnvironmentProtectedObjectsSummaryByEnv_KNETAPP
    EnvironmentProtectedObjectsSummaryByEnv_KAGENT
    EnvironmentProtectedObjectsSummaryByEnv_KGENERICNAS
    EnvironmentProtectedObjectsSummaryByEnv_KACROPOLIS
    EnvironmentProtectedObjectsSummaryByEnv_KPHYSICALFILES
    EnvironmentProtectedObjectsSummaryByEnv_KISILON
    EnvironmentProtectedObjectsSummaryByEnv_KGPFS
    EnvironmentProtectedObjectsSummaryByEnv_KKVM
    EnvironmentProtectedObjectsSummaryByEnv_KAWS
    EnvironmentProtectedObjectsSummaryByEnv_KEXCHANGE
    EnvironmentProtectedObjectsSummaryByEnv_KHYPERVVSS
    EnvironmentProtectedObjectsSummaryByEnv_KORACLE
    EnvironmentProtectedObjectsSummaryByEnv_KGCP
    EnvironmentProtectedObjectsSummaryByEnv_KFLASHBLADE
    EnvironmentProtectedObjectsSummaryByEnv_KAWSNATIVE
    EnvironmentProtectedObjectsSummaryByEnv_KVCD
    EnvironmentProtectedObjectsSummaryByEnv_KO365
    EnvironmentProtectedObjectsSummaryByEnv_KO365OUTLOOK
    EnvironmentProtectedObjectsSummaryByEnv_KHYPERFLEX
    EnvironmentProtectedObjectsSummaryByEnv_KGCPNATIVE
    EnvironmentProtectedObjectsSummaryByEnv_KAZURENATIVE
    EnvironmentProtectedObjectsSummaryByEnv_KAD
    EnvironmentProtectedObjectsSummaryByEnv_KAWSSNAPSHOTMANAGER
)

func (r EnvironmentProtectedObjectsSummaryByEnvEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentProtectedObjectsSummaryByEnvEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentProtectedObjectsSummaryByEnvEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentProtectedObjectsSummaryByEnvEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentProtectedObjectsSummaryByEnvEnum to its string representation
 */
func EnvironmentProtectedObjectsSummaryByEnvEnumToValue(environmentProtectedObjectsSummaryByEnvEnum EnvironmentProtectedObjectsSummaryByEnvEnum) string {
    switch environmentProtectedObjectsSummaryByEnvEnum {
        case EnvironmentProtectedObjectsSummaryByEnv_KVMWARE:
    		return "kVMware"
        case EnvironmentProtectedObjectsSummaryByEnv_KHYPERV:
    		return "kHyperV"
        case EnvironmentProtectedObjectsSummaryByEnv_KSQL:
    		return "kSQL"
        case EnvironmentProtectedObjectsSummaryByEnv_KVIEW:
    		return "kView"
        case EnvironmentProtectedObjectsSummaryByEnv_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentProtectedObjectsSummaryByEnv_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentProtectedObjectsSummaryByEnv_KPURE:
    		return "kPure"
        case EnvironmentProtectedObjectsSummaryByEnv_KAZURE:
    		return "kAzure"
        case EnvironmentProtectedObjectsSummaryByEnv_KNETAPP:
    		return "kNetapp"
        case EnvironmentProtectedObjectsSummaryByEnv_KAGENT:
    		return "kAgent"
        case EnvironmentProtectedObjectsSummaryByEnv_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentProtectedObjectsSummaryByEnv_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentProtectedObjectsSummaryByEnv_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentProtectedObjectsSummaryByEnv_KISILON:
    		return "kIsilon"
        case EnvironmentProtectedObjectsSummaryByEnv_KGPFS:
    		return "kGPFS"
        case EnvironmentProtectedObjectsSummaryByEnv_KKVM:
    		return "kKVM"
        case EnvironmentProtectedObjectsSummaryByEnv_KAWS:
    		return "kAWS"
        case EnvironmentProtectedObjectsSummaryByEnv_KEXCHANGE:
    		return "kExchange"
        case EnvironmentProtectedObjectsSummaryByEnv_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentProtectedObjectsSummaryByEnv_KORACLE:
    		return "kOracle"
        case EnvironmentProtectedObjectsSummaryByEnv_KGCP:
    		return "kGCP"
        case EnvironmentProtectedObjectsSummaryByEnv_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentProtectedObjectsSummaryByEnv_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentProtectedObjectsSummaryByEnv_KVCD:
    		return "kVCD"
        case EnvironmentProtectedObjectsSummaryByEnv_KO365:
    		return "kO365"
        case EnvironmentProtectedObjectsSummaryByEnv_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentProtectedObjectsSummaryByEnv_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentProtectedObjectsSummaryByEnv_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentProtectedObjectsSummaryByEnv_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentProtectedObjectsSummaryByEnv_KAD:
    		return "kAD"
        case EnvironmentProtectedObjectsSummaryByEnv_KAWSSNAPSHOTMANAGER:
    		return "kAWSSnapshotManager"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentProtectedObjectsSummaryByEnvEnum Array to its string Array representation
*/
func EnvironmentProtectedObjectsSummaryByEnvEnumArrayToValue(environmentProtectedObjectsSummaryByEnvEnum []EnvironmentProtectedObjectsSummaryByEnvEnum) []string {
    convArray := make([]string,len( environmentProtectedObjectsSummaryByEnvEnum))
    for i:=0; i<len(environmentProtectedObjectsSummaryByEnvEnum);i++ {
        convArray[i] = EnvironmentProtectedObjectsSummaryByEnvEnumToValue(environmentProtectedObjectsSummaryByEnvEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentProtectedObjectsSummaryByEnvEnumFromValue(value string) EnvironmentProtectedObjectsSummaryByEnvEnum {
    switch value {
        case "kVMware":
            return EnvironmentProtectedObjectsSummaryByEnv_KVMWARE
        case "kHyperV":
            return EnvironmentProtectedObjectsSummaryByEnv_KHYPERV
        case "kSQL":
            return EnvironmentProtectedObjectsSummaryByEnv_KSQL
        case "kView":
            return EnvironmentProtectedObjectsSummaryByEnv_KVIEW
        case "kPuppeteer":
            return EnvironmentProtectedObjectsSummaryByEnv_KPUPPETEER
        case "kPhysical":
            return EnvironmentProtectedObjectsSummaryByEnv_KPHYSICAL
        case "kPure":
            return EnvironmentProtectedObjectsSummaryByEnv_KPURE
        case "kAzure":
            return EnvironmentProtectedObjectsSummaryByEnv_KAZURE
        case "kNetapp":
            return EnvironmentProtectedObjectsSummaryByEnv_KNETAPP
        case "kAgent":
            return EnvironmentProtectedObjectsSummaryByEnv_KAGENT
        case "kGenericNas":
            return EnvironmentProtectedObjectsSummaryByEnv_KGENERICNAS
        case "kAcropolis":
            return EnvironmentProtectedObjectsSummaryByEnv_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentProtectedObjectsSummaryByEnv_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentProtectedObjectsSummaryByEnv_KISILON
        case "kGPFS":
            return EnvironmentProtectedObjectsSummaryByEnv_KGPFS
        case "kKVM":
            return EnvironmentProtectedObjectsSummaryByEnv_KKVM
        case "kAWS":
            return EnvironmentProtectedObjectsSummaryByEnv_KAWS
        case "kExchange":
            return EnvironmentProtectedObjectsSummaryByEnv_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentProtectedObjectsSummaryByEnv_KHYPERVVSS
        case "kOracle":
            return EnvironmentProtectedObjectsSummaryByEnv_KORACLE
        case "kGCP":
            return EnvironmentProtectedObjectsSummaryByEnv_KGCP
        case "kFlashBlade":
            return EnvironmentProtectedObjectsSummaryByEnv_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentProtectedObjectsSummaryByEnv_KAWSNATIVE
        case "kVCD":
            return EnvironmentProtectedObjectsSummaryByEnv_KVCD
        case "kO365":
            return EnvironmentProtectedObjectsSummaryByEnv_KO365
        case "kO365Outlook":
            return EnvironmentProtectedObjectsSummaryByEnv_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentProtectedObjectsSummaryByEnv_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentProtectedObjectsSummaryByEnv_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentProtectedObjectsSummaryByEnv_KAZURENATIVE
        case "kAD":
            return EnvironmentProtectedObjectsSummaryByEnv_KAD
        case "kAWSSnapshotManager":
            return EnvironmentProtectedObjectsSummaryByEnv_KAWSSNAPSHOTMANAGER
        default:
            return EnvironmentProtectedObjectsSummaryByEnv_KVMWARE
    }
}
