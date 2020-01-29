// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentVaultProviderStatsByEnvEnum enum
 */
type EnvironmentVaultProviderStatsByEnvEnum int

/**
 * Value collection for EnvironmentVaultProviderStatsByEnvEnum enum
 */
const (
    EnvironmentVaultProviderStatsByEnv_KVMWARE EnvironmentVaultProviderStatsByEnvEnum = 1 + iota
    EnvironmentVaultProviderStatsByEnv_KHYPERV
    EnvironmentVaultProviderStatsByEnv_KSQL
    EnvironmentVaultProviderStatsByEnv_KVIEW
    EnvironmentVaultProviderStatsByEnv_KPUPPETEER
    EnvironmentVaultProviderStatsByEnv_KPHYSICAL
    EnvironmentVaultProviderStatsByEnv_KPURE
    EnvironmentVaultProviderStatsByEnv_KAZURE
    EnvironmentVaultProviderStatsByEnv_KNETAPP
    EnvironmentVaultProviderStatsByEnv_KAGENT
    EnvironmentVaultProviderStatsByEnv_KGENERICNAS
    EnvironmentVaultProviderStatsByEnv_KACROPOLIS
    EnvironmentVaultProviderStatsByEnv_KPHYSICALFILES
    EnvironmentVaultProviderStatsByEnv_KISILON
    EnvironmentVaultProviderStatsByEnv_KGPFS
    EnvironmentVaultProviderStatsByEnv_KKVM
    EnvironmentVaultProviderStatsByEnv_KAWS
    EnvironmentVaultProviderStatsByEnv_KEXCHANGE
    EnvironmentVaultProviderStatsByEnv_KHYPERVVSS
    EnvironmentVaultProviderStatsByEnv_KORACLE
    EnvironmentVaultProviderStatsByEnv_KGCP
    EnvironmentVaultProviderStatsByEnv_KFLASHBLADE
    EnvironmentVaultProviderStatsByEnv_KAWSNATIVE
    EnvironmentVaultProviderStatsByEnv_KVCD
    EnvironmentVaultProviderStatsByEnv_KO365
    EnvironmentVaultProviderStatsByEnv_KO365OUTLOOK
    EnvironmentVaultProviderStatsByEnv_KHYPERFLEX
    EnvironmentVaultProviderStatsByEnv_KGCPNATIVE
    EnvironmentVaultProviderStatsByEnv_KAZURENATIVE
    EnvironmentVaultProviderStatsByEnv_KAD
    EnvironmentVaultProviderStatsByEnv_KAWSSNAPSHOTMANAGER
)

func (r EnvironmentVaultProviderStatsByEnvEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentVaultProviderStatsByEnvEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentVaultProviderStatsByEnvEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentVaultProviderStatsByEnvEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentVaultProviderStatsByEnvEnum to its string representation
 */
func EnvironmentVaultProviderStatsByEnvEnumToValue(environmentVaultProviderStatsByEnvEnum EnvironmentVaultProviderStatsByEnvEnum) string {
    switch environmentVaultProviderStatsByEnvEnum {
        case EnvironmentVaultProviderStatsByEnv_KVMWARE:
    		return "kVMware"
        case EnvironmentVaultProviderStatsByEnv_KHYPERV:
    		return "kHyperV"
        case EnvironmentVaultProviderStatsByEnv_KSQL:
    		return "kSQL"
        case EnvironmentVaultProviderStatsByEnv_KVIEW:
    		return "kView"
        case EnvironmentVaultProviderStatsByEnv_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentVaultProviderStatsByEnv_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentVaultProviderStatsByEnv_KPURE:
    		return "kPure"
        case EnvironmentVaultProviderStatsByEnv_KAZURE:
    		return "kAzure"
        case EnvironmentVaultProviderStatsByEnv_KNETAPP:
    		return "kNetapp"
        case EnvironmentVaultProviderStatsByEnv_KAGENT:
    		return "kAgent"
        case EnvironmentVaultProviderStatsByEnv_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentVaultProviderStatsByEnv_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentVaultProviderStatsByEnv_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentVaultProviderStatsByEnv_KISILON:
    		return "kIsilon"
        case EnvironmentVaultProviderStatsByEnv_KGPFS:
    		return "kGPFS"
        case EnvironmentVaultProviderStatsByEnv_KKVM:
    		return "kKVM"
        case EnvironmentVaultProviderStatsByEnv_KAWS:
    		return "kAWS"
        case EnvironmentVaultProviderStatsByEnv_KEXCHANGE:
    		return "kExchange"
        case EnvironmentVaultProviderStatsByEnv_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentVaultProviderStatsByEnv_KORACLE:
    		return "kOracle"
        case EnvironmentVaultProviderStatsByEnv_KGCP:
    		return "kGCP"
        case EnvironmentVaultProviderStatsByEnv_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentVaultProviderStatsByEnv_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentVaultProviderStatsByEnv_KVCD:
    		return "kVCD"
        case EnvironmentVaultProviderStatsByEnv_KO365:
    		return "kO365"
        case EnvironmentVaultProviderStatsByEnv_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentVaultProviderStatsByEnv_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentVaultProviderStatsByEnv_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentVaultProviderStatsByEnv_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentVaultProviderStatsByEnv_KAD:
    		return "kAD"
        case EnvironmentVaultProviderStatsByEnv_KAWSSNAPSHOTMANAGER:
    		return "kAWSSnapshotManager"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentVaultProviderStatsByEnvEnum Array to its string Array representation
*/
func EnvironmentVaultProviderStatsByEnvEnumArrayToValue(environmentVaultProviderStatsByEnvEnum []EnvironmentVaultProviderStatsByEnvEnum) []string {
    convArray := make([]string,len( environmentVaultProviderStatsByEnvEnum))
    for i:=0; i<len(environmentVaultProviderStatsByEnvEnum);i++ {
        convArray[i] = EnvironmentVaultProviderStatsByEnvEnumToValue(environmentVaultProviderStatsByEnvEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentVaultProviderStatsByEnvEnumFromValue(value string) EnvironmentVaultProviderStatsByEnvEnum {
    switch value {
        case "kVMware":
            return EnvironmentVaultProviderStatsByEnv_KVMWARE
        case "kHyperV":
            return EnvironmentVaultProviderStatsByEnv_KHYPERV
        case "kSQL":
            return EnvironmentVaultProviderStatsByEnv_KSQL
        case "kView":
            return EnvironmentVaultProviderStatsByEnv_KVIEW
        case "kPuppeteer":
            return EnvironmentVaultProviderStatsByEnv_KPUPPETEER
        case "kPhysical":
            return EnvironmentVaultProviderStatsByEnv_KPHYSICAL
        case "kPure":
            return EnvironmentVaultProviderStatsByEnv_KPURE
        case "kAzure":
            return EnvironmentVaultProviderStatsByEnv_KAZURE
        case "kNetapp":
            return EnvironmentVaultProviderStatsByEnv_KNETAPP
        case "kAgent":
            return EnvironmentVaultProviderStatsByEnv_KAGENT
        case "kGenericNas":
            return EnvironmentVaultProviderStatsByEnv_KGENERICNAS
        case "kAcropolis":
            return EnvironmentVaultProviderStatsByEnv_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentVaultProviderStatsByEnv_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentVaultProviderStatsByEnv_KISILON
        case "kGPFS":
            return EnvironmentVaultProviderStatsByEnv_KGPFS
        case "kKVM":
            return EnvironmentVaultProviderStatsByEnv_KKVM
        case "kAWS":
            return EnvironmentVaultProviderStatsByEnv_KAWS
        case "kExchange":
            return EnvironmentVaultProviderStatsByEnv_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentVaultProviderStatsByEnv_KHYPERVVSS
        case "kOracle":
            return EnvironmentVaultProviderStatsByEnv_KORACLE
        case "kGCP":
            return EnvironmentVaultProviderStatsByEnv_KGCP
        case "kFlashBlade":
            return EnvironmentVaultProviderStatsByEnv_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentVaultProviderStatsByEnv_KAWSNATIVE
        case "kVCD":
            return EnvironmentVaultProviderStatsByEnv_KVCD
        case "kO365":
            return EnvironmentVaultProviderStatsByEnv_KO365
        case "kO365Outlook":
            return EnvironmentVaultProviderStatsByEnv_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentVaultProviderStatsByEnv_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentVaultProviderStatsByEnv_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentVaultProviderStatsByEnv_KAZURENATIVE
        case "kAD":
            return EnvironmentVaultProviderStatsByEnv_KAD
        case "kAWSSnapshotManager":
            return EnvironmentVaultProviderStatsByEnv_KAWSSNAPSHOTMANAGER
        default:
            return EnvironmentVaultProviderStatsByEnv_KVMWARE
    }
}
