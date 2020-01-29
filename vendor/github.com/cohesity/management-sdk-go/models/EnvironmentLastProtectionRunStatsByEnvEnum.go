// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentLastProtectionRunStatsByEnvEnum enum
 */
type EnvironmentLastProtectionRunStatsByEnvEnum int

/**
 * Value collection for EnvironmentLastProtectionRunStatsByEnvEnum enum
 */
const (
    EnvironmentLastProtectionRunStatsByEnv_KVMWARE EnvironmentLastProtectionRunStatsByEnvEnum = 1 + iota
    EnvironmentLastProtectionRunStatsByEnv_KHYPERV
    EnvironmentLastProtectionRunStatsByEnv_KSQL
    EnvironmentLastProtectionRunStatsByEnv_KVIEW
    EnvironmentLastProtectionRunStatsByEnv_KPUPPETEER
    EnvironmentLastProtectionRunStatsByEnv_KPHYSICAL
    EnvironmentLastProtectionRunStatsByEnv_KPURE
    EnvironmentLastProtectionRunStatsByEnv_KAZURE
    EnvironmentLastProtectionRunStatsByEnv_KNETAPP
    EnvironmentLastProtectionRunStatsByEnv_KAGENT
    EnvironmentLastProtectionRunStatsByEnv_KGENERICNAS
    EnvironmentLastProtectionRunStatsByEnv_KACROPOLIS
    EnvironmentLastProtectionRunStatsByEnv_KPHYSICALFILES
    EnvironmentLastProtectionRunStatsByEnv_KISILON
    EnvironmentLastProtectionRunStatsByEnv_KGPFS
    EnvironmentLastProtectionRunStatsByEnv_KKVM
    EnvironmentLastProtectionRunStatsByEnv_KAWS
    EnvironmentLastProtectionRunStatsByEnv_KEXCHANGE
    EnvironmentLastProtectionRunStatsByEnv_KHYPERVVSS
    EnvironmentLastProtectionRunStatsByEnv_KORACLE
    EnvironmentLastProtectionRunStatsByEnv_KGCP
    EnvironmentLastProtectionRunStatsByEnv_KFLASHBLADE
    EnvironmentLastProtectionRunStatsByEnv_KAWSNATIVE
    EnvironmentLastProtectionRunStatsByEnv_KVCD
    EnvironmentLastProtectionRunStatsByEnv_KO365
    EnvironmentLastProtectionRunStatsByEnv_KO365OUTLOOK
    EnvironmentLastProtectionRunStatsByEnv_KHYPERFLEX
    EnvironmentLastProtectionRunStatsByEnv_KGCPNATIVE
    EnvironmentLastProtectionRunStatsByEnv_KAZURENATIVE
    EnvironmentLastProtectionRunStatsByEnv_KAD
    EnvironmentLastProtectionRunStatsByEnv_KAWSSNAPSHOTMANAGER
)

func (r EnvironmentLastProtectionRunStatsByEnvEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentLastProtectionRunStatsByEnvEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentLastProtectionRunStatsByEnvEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentLastProtectionRunStatsByEnvEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentLastProtectionRunStatsByEnvEnum to its string representation
 */
func EnvironmentLastProtectionRunStatsByEnvEnumToValue(environmentLastProtectionRunStatsByEnvEnum EnvironmentLastProtectionRunStatsByEnvEnum) string {
    switch environmentLastProtectionRunStatsByEnvEnum {
        case EnvironmentLastProtectionRunStatsByEnv_KVMWARE:
    		return "kVMware"
        case EnvironmentLastProtectionRunStatsByEnv_KHYPERV:
    		return "kHyperV"
        case EnvironmentLastProtectionRunStatsByEnv_KSQL:
    		return "kSQL"
        case EnvironmentLastProtectionRunStatsByEnv_KVIEW:
    		return "kView"
        case EnvironmentLastProtectionRunStatsByEnv_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentLastProtectionRunStatsByEnv_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentLastProtectionRunStatsByEnv_KPURE:
    		return "kPure"
        case EnvironmentLastProtectionRunStatsByEnv_KAZURE:
    		return "kAzure"
        case EnvironmentLastProtectionRunStatsByEnv_KNETAPP:
    		return "kNetapp"
        case EnvironmentLastProtectionRunStatsByEnv_KAGENT:
    		return "kAgent"
        case EnvironmentLastProtectionRunStatsByEnv_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentLastProtectionRunStatsByEnv_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentLastProtectionRunStatsByEnv_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentLastProtectionRunStatsByEnv_KISILON:
    		return "kIsilon"
        case EnvironmentLastProtectionRunStatsByEnv_KGPFS:
    		return "kGPFS"
        case EnvironmentLastProtectionRunStatsByEnv_KKVM:
    		return "kKVM"
        case EnvironmentLastProtectionRunStatsByEnv_KAWS:
    		return "kAWS"
        case EnvironmentLastProtectionRunStatsByEnv_KEXCHANGE:
    		return "kExchange"
        case EnvironmentLastProtectionRunStatsByEnv_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentLastProtectionRunStatsByEnv_KORACLE:
    		return "kOracle"
        case EnvironmentLastProtectionRunStatsByEnv_KGCP:
    		return "kGCP"
        case EnvironmentLastProtectionRunStatsByEnv_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentLastProtectionRunStatsByEnv_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentLastProtectionRunStatsByEnv_KVCD:
    		return "kVCD"
        case EnvironmentLastProtectionRunStatsByEnv_KO365:
    		return "kO365"
        case EnvironmentLastProtectionRunStatsByEnv_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentLastProtectionRunStatsByEnv_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentLastProtectionRunStatsByEnv_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentLastProtectionRunStatsByEnv_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentLastProtectionRunStatsByEnv_KAD:
    		return "kAD"
        case EnvironmentLastProtectionRunStatsByEnv_KAWSSNAPSHOTMANAGER:
    		return "kAWSSnapshotManager"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentLastProtectionRunStatsByEnvEnum Array to its string Array representation
*/
func EnvironmentLastProtectionRunStatsByEnvEnumArrayToValue(environmentLastProtectionRunStatsByEnvEnum []EnvironmentLastProtectionRunStatsByEnvEnum) []string {
    convArray := make([]string,len( environmentLastProtectionRunStatsByEnvEnum))
    for i:=0; i<len(environmentLastProtectionRunStatsByEnvEnum);i++ {
        convArray[i] = EnvironmentLastProtectionRunStatsByEnvEnumToValue(environmentLastProtectionRunStatsByEnvEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentLastProtectionRunStatsByEnvEnumFromValue(value string) EnvironmentLastProtectionRunStatsByEnvEnum {
    switch value {
        case "kVMware":
            return EnvironmentLastProtectionRunStatsByEnv_KVMWARE
        case "kHyperV":
            return EnvironmentLastProtectionRunStatsByEnv_KHYPERV
        case "kSQL":
            return EnvironmentLastProtectionRunStatsByEnv_KSQL
        case "kView":
            return EnvironmentLastProtectionRunStatsByEnv_KVIEW
        case "kPuppeteer":
            return EnvironmentLastProtectionRunStatsByEnv_KPUPPETEER
        case "kPhysical":
            return EnvironmentLastProtectionRunStatsByEnv_KPHYSICAL
        case "kPure":
            return EnvironmentLastProtectionRunStatsByEnv_KPURE
        case "kAzure":
            return EnvironmentLastProtectionRunStatsByEnv_KAZURE
        case "kNetapp":
            return EnvironmentLastProtectionRunStatsByEnv_KNETAPP
        case "kAgent":
            return EnvironmentLastProtectionRunStatsByEnv_KAGENT
        case "kGenericNas":
            return EnvironmentLastProtectionRunStatsByEnv_KGENERICNAS
        case "kAcropolis":
            return EnvironmentLastProtectionRunStatsByEnv_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentLastProtectionRunStatsByEnv_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentLastProtectionRunStatsByEnv_KISILON
        case "kGPFS":
            return EnvironmentLastProtectionRunStatsByEnv_KGPFS
        case "kKVM":
            return EnvironmentLastProtectionRunStatsByEnv_KKVM
        case "kAWS":
            return EnvironmentLastProtectionRunStatsByEnv_KAWS
        case "kExchange":
            return EnvironmentLastProtectionRunStatsByEnv_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentLastProtectionRunStatsByEnv_KHYPERVVSS
        case "kOracle":
            return EnvironmentLastProtectionRunStatsByEnv_KORACLE
        case "kGCP":
            return EnvironmentLastProtectionRunStatsByEnv_KGCP
        case "kFlashBlade":
            return EnvironmentLastProtectionRunStatsByEnv_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentLastProtectionRunStatsByEnv_KAWSNATIVE
        case "kVCD":
            return EnvironmentLastProtectionRunStatsByEnv_KVCD
        case "kO365":
            return EnvironmentLastProtectionRunStatsByEnv_KO365
        case "kO365Outlook":
            return EnvironmentLastProtectionRunStatsByEnv_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentLastProtectionRunStatsByEnv_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentLastProtectionRunStatsByEnv_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentLastProtectionRunStatsByEnv_KAZURENATIVE
        case "kAD":
            return EnvironmentLastProtectionRunStatsByEnv_KAD
        case "kAWSSnapshotManager":
            return EnvironmentLastProtectionRunStatsByEnv_KAWSSNAPSHOTMANAGER
        default:
            return EnvironmentLastProtectionRunStatsByEnv_KVMWARE
    }
}
