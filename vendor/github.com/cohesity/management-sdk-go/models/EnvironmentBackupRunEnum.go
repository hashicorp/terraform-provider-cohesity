// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentBackupRunEnum enum
 */
type EnvironmentBackupRunEnum int

/**
 * Value collection for EnvironmentBackupRunEnum enum
 */
const (
    EnvironmentBackupRun_KVMWARE EnvironmentBackupRunEnum = 1 + iota
    EnvironmentBackupRun_KHYPERV
    EnvironmentBackupRun_KSQL
    EnvironmentBackupRun_KVIEW
    EnvironmentBackupRun_KPUPPETEER
    EnvironmentBackupRun_KPHYSICAL
    EnvironmentBackupRun_KPURE
    EnvironmentBackupRun_KAZURE
    EnvironmentBackupRun_KNETAPP
    EnvironmentBackupRun_KAGENT
    EnvironmentBackupRun_KGENERICNAS
    EnvironmentBackupRun_KACROPOLIS
    EnvironmentBackupRun_KPHYSICALFILES
    EnvironmentBackupRun_KISILON
    EnvironmentBackupRun_KGPFS
    EnvironmentBackupRun_KKVM
    EnvironmentBackupRun_KAWS
    EnvironmentBackupRun_KEXCHANGE
    EnvironmentBackupRun_KHYPERVVSS
    EnvironmentBackupRun_KORACLE
    EnvironmentBackupRun_KGCP
    EnvironmentBackupRun_KFLASHBLADE
    EnvironmentBackupRun_KAWSNATIVE
    EnvironmentBackupRun_KVCD
    EnvironmentBackupRun_KO365
    EnvironmentBackupRun_KO365OUTLOOK
    EnvironmentBackupRun_KHYPERFLEX
    EnvironmentBackupRun_KGCPNATIVE
    EnvironmentBackupRun_KAZURENATIVE
    EnvironmentBackupRun_KKUBERNETES
)

func (r EnvironmentBackupRunEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentBackupRunEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentBackupRunEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentBackupRunEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentBackupRunEnum to its string representation
 */
func EnvironmentBackupRunEnumToValue(environmentBackupRunEnum EnvironmentBackupRunEnum) string {
    switch environmentBackupRunEnum {
        case EnvironmentBackupRun_KVMWARE:
    		return "kVMware"
        case EnvironmentBackupRun_KHYPERV:
    		return "kHyperV"
        case EnvironmentBackupRun_KSQL:
    		return "kSQL"
        case EnvironmentBackupRun_KVIEW:
    		return "kView"
        case EnvironmentBackupRun_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentBackupRun_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentBackupRun_KPURE:
    		return "kPure"
        case EnvironmentBackupRun_KAZURE:
    		return "kAzure"
        case EnvironmentBackupRun_KNETAPP:
    		return "kNetapp"
        case EnvironmentBackupRun_KAGENT:
    		return "kAgent"
        case EnvironmentBackupRun_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentBackupRun_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentBackupRun_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentBackupRun_KISILON:
    		return "kIsilon"
        case EnvironmentBackupRun_KGPFS:
    		return "kGPFS"
        case EnvironmentBackupRun_KKVM:
    		return "kKVM"
        case EnvironmentBackupRun_KAWS:
    		return "kAWS"
        case EnvironmentBackupRun_KEXCHANGE:
    		return "kExchange"
        case EnvironmentBackupRun_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentBackupRun_KORACLE:
    		return "kOracle"
        case EnvironmentBackupRun_KGCP:
    		return "kGCP"
        case EnvironmentBackupRun_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentBackupRun_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentBackupRun_KVCD:
    		return "kVCD"
        case EnvironmentBackupRun_KO365:
    		return "kO365"
        case EnvironmentBackupRun_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentBackupRun_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentBackupRun_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentBackupRun_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentBackupRun_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentBackupRunEnum Array to its string Array representation
*/
func EnvironmentBackupRunEnumArrayToValue(environmentBackupRunEnum []EnvironmentBackupRunEnum) []string {
    convArray := make([]string,len( environmentBackupRunEnum))
    for i:=0; i<len(environmentBackupRunEnum);i++ {
        convArray[i] = EnvironmentBackupRunEnumToValue(environmentBackupRunEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentBackupRunEnumFromValue(value string) EnvironmentBackupRunEnum {
    switch value {
        case "kVMware":
            return EnvironmentBackupRun_KVMWARE
        case "kHyperV":
            return EnvironmentBackupRun_KHYPERV
        case "kSQL":
            return EnvironmentBackupRun_KSQL
        case "kView":
            return EnvironmentBackupRun_KVIEW
        case "kPuppeteer":
            return EnvironmentBackupRun_KPUPPETEER
        case "kPhysical":
            return EnvironmentBackupRun_KPHYSICAL
        case "kPure":
            return EnvironmentBackupRun_KPURE
        case "kAzure":
            return EnvironmentBackupRun_KAZURE
        case "kNetapp":
            return EnvironmentBackupRun_KNETAPP
        case "kAgent":
            return EnvironmentBackupRun_KAGENT
        case "kGenericNas":
            return EnvironmentBackupRun_KGENERICNAS
        case "kAcropolis":
            return EnvironmentBackupRun_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentBackupRun_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentBackupRun_KISILON
        case "kGPFS":
            return EnvironmentBackupRun_KGPFS
        case "kKVM":
            return EnvironmentBackupRun_KKVM
        case "kAWS":
            return EnvironmentBackupRun_KAWS
        case "kExchange":
            return EnvironmentBackupRun_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentBackupRun_KHYPERVVSS
        case "kOracle":
            return EnvironmentBackupRun_KORACLE
        case "kGCP":
            return EnvironmentBackupRun_KGCP
        case "kFlashBlade":
            return EnvironmentBackupRun_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentBackupRun_KAWSNATIVE
        case "kVCD":
            return EnvironmentBackupRun_KVCD
        case "kO365":
            return EnvironmentBackupRun_KO365
        case "kO365Outlook":
            return EnvironmentBackupRun_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentBackupRun_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentBackupRun_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentBackupRun_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentBackupRun_KKUBERNETES
        default:
            return EnvironmentBackupRun_KVMWARE
    }
}
