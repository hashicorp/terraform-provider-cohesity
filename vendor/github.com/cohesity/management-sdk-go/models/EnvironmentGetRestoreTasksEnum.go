// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentGetRestoreTasksEnum enum
 */
type EnvironmentGetRestoreTasksEnum int

/**
 * Value collection for EnvironmentGetRestoreTasksEnum enum
 */
const (
    EnvironmentGetRestoreTasks_KVMWARE EnvironmentGetRestoreTasksEnum = 1 + iota
    EnvironmentGetRestoreTasks_KHYPERV
    EnvironmentGetRestoreTasks_KSQL
    EnvironmentGetRestoreTasks_KVIEW
    EnvironmentGetRestoreTasks_KPUPPETEER
    EnvironmentGetRestoreTasks_KPHYSICAL
    EnvironmentGetRestoreTasks_KPURE
    EnvironmentGetRestoreTasks_KAZURE
    EnvironmentGetRestoreTasks_KNETAPP
    EnvironmentGetRestoreTasks_KAGENT
    EnvironmentGetRestoreTasks_KGENERICNAS
    EnvironmentGetRestoreTasks_KACROPOLIS
    EnvironmentGetRestoreTasks_KPHYSICALFILES
    EnvironmentGetRestoreTasks_KISILON
    EnvironmentGetRestoreTasks_KGPFS
    EnvironmentGetRestoreTasks_KKVM
    EnvironmentGetRestoreTasks_KAWS
    EnvironmentGetRestoreTasks_KEXCHANGE
    EnvironmentGetRestoreTasks_KHYPERVVSS
    EnvironmentGetRestoreTasks_KORACLE
    EnvironmentGetRestoreTasks_KGCP
    EnvironmentGetRestoreTasks_KFLASHBLADE
    EnvironmentGetRestoreTasks_KAWSNATIVE
    EnvironmentGetRestoreTasks_KVCD
    EnvironmentGetRestoreTasks_KO365
    EnvironmentGetRestoreTasks_KO365OUTLOOK
    EnvironmentGetRestoreTasks_KHYPERFLEX
    EnvironmentGetRestoreTasks_KGCPNATIVE
    EnvironmentGetRestoreTasks_KAZURENATIVE
    EnvironmentGetRestoreTasks_KKUBERNETES
)

func (r EnvironmentGetRestoreTasksEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentGetRestoreTasksEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentGetRestoreTasksEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentGetRestoreTasksEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentGetRestoreTasksEnum to its string representation
 */
func EnvironmentGetRestoreTasksEnumToValue(environmentGetRestoreTasksEnum EnvironmentGetRestoreTasksEnum) string {
    switch environmentGetRestoreTasksEnum {
        case EnvironmentGetRestoreTasks_KVMWARE:
    		return "kVMware"
        case EnvironmentGetRestoreTasks_KHYPERV:
    		return "kHyperV"
        case EnvironmentGetRestoreTasks_KSQL:
    		return "kSQL"
        case EnvironmentGetRestoreTasks_KVIEW:
    		return "kView"
        case EnvironmentGetRestoreTasks_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentGetRestoreTasks_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentGetRestoreTasks_KPURE:
    		return "kPure"
        case EnvironmentGetRestoreTasks_KAZURE:
    		return "kAzure"
        case EnvironmentGetRestoreTasks_KNETAPP:
    		return "kNetapp"
        case EnvironmentGetRestoreTasks_KAGENT:
    		return "kAgent"
        case EnvironmentGetRestoreTasks_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentGetRestoreTasks_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentGetRestoreTasks_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentGetRestoreTasks_KISILON:
    		return "kIsilon"
        case EnvironmentGetRestoreTasks_KGPFS:
    		return "kGPFS"
        case EnvironmentGetRestoreTasks_KKVM:
    		return "kKVM"
        case EnvironmentGetRestoreTasks_KAWS:
    		return "kAWS"
        case EnvironmentGetRestoreTasks_KEXCHANGE:
    		return "kExchange"
        case EnvironmentGetRestoreTasks_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentGetRestoreTasks_KORACLE:
    		return "kOracle"
        case EnvironmentGetRestoreTasks_KGCP:
    		return "kGCP"
        case EnvironmentGetRestoreTasks_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentGetRestoreTasks_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentGetRestoreTasks_KVCD:
    		return "kVCD"
        case EnvironmentGetRestoreTasks_KO365:
    		return "kO365"
        case EnvironmentGetRestoreTasks_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentGetRestoreTasks_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentGetRestoreTasks_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentGetRestoreTasks_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentGetRestoreTasks_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentGetRestoreTasksEnum Array to its string Array representation
*/
func EnvironmentGetRestoreTasksEnumArrayToValue(environmentGetRestoreTasksEnum []EnvironmentGetRestoreTasksEnum) []string {
    convArray := make([]string,len( environmentGetRestoreTasksEnum))
    for i:=0; i<len(environmentGetRestoreTasksEnum);i++ {
        convArray[i] = EnvironmentGetRestoreTasksEnumToValue(environmentGetRestoreTasksEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentGetRestoreTasksEnumFromValue(value string) EnvironmentGetRestoreTasksEnum {
    switch value {
        case "kVMware":
            return EnvironmentGetRestoreTasks_KVMWARE
        case "kHyperV":
            return EnvironmentGetRestoreTasks_KHYPERV
        case "kSQL":
            return EnvironmentGetRestoreTasks_KSQL
        case "kView":
            return EnvironmentGetRestoreTasks_KVIEW
        case "kPuppeteer":
            return EnvironmentGetRestoreTasks_KPUPPETEER
        case "kPhysical":
            return EnvironmentGetRestoreTasks_KPHYSICAL
        case "kPure":
            return EnvironmentGetRestoreTasks_KPURE
        case "kAzure":
            return EnvironmentGetRestoreTasks_KAZURE
        case "kNetapp":
            return EnvironmentGetRestoreTasks_KNETAPP
        case "kAgent":
            return EnvironmentGetRestoreTasks_KAGENT
        case "kGenericNas":
            return EnvironmentGetRestoreTasks_KGENERICNAS
        case "kAcropolis":
            return EnvironmentGetRestoreTasks_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentGetRestoreTasks_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentGetRestoreTasks_KISILON
        case "kGPFS":
            return EnvironmentGetRestoreTasks_KGPFS
        case "kKVM":
            return EnvironmentGetRestoreTasks_KKVM
        case "kAWS":
            return EnvironmentGetRestoreTasks_KAWS
        case "kExchange":
            return EnvironmentGetRestoreTasks_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentGetRestoreTasks_KHYPERVVSS
        case "kOracle":
            return EnvironmentGetRestoreTasks_KORACLE
        case "kGCP":
            return EnvironmentGetRestoreTasks_KGCP
        case "kFlashBlade":
            return EnvironmentGetRestoreTasks_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentGetRestoreTasks_KAWSNATIVE
        case "kVCD":
            return EnvironmentGetRestoreTasks_KVCD
        case "kO365":
            return EnvironmentGetRestoreTasks_KO365
        case "kO365Outlook":
            return EnvironmentGetRestoreTasks_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentGetRestoreTasks_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentGetRestoreTasks_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentGetRestoreTasks_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentGetRestoreTasks_KKUBERNETES
        default:
            return EnvironmentGetRestoreTasks_KVMWARE
    }
}
