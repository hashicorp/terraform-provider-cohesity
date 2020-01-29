// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentSearchRestoredFilesEnum enum
 */
type EnvironmentSearchRestoredFilesEnum int

/**
 * Value collection for EnvironmentSearchRestoredFilesEnum enum
 */
const (
    EnvironmentSearchRestoredFiles_KVMWARE EnvironmentSearchRestoredFilesEnum = 1 + iota
    EnvironmentSearchRestoredFiles_KHYPERV
    EnvironmentSearchRestoredFiles_KSQL
    EnvironmentSearchRestoredFiles_KVIEW
    EnvironmentSearchRestoredFiles_KPUPPETEER
    EnvironmentSearchRestoredFiles_KPHYSICAL
    EnvironmentSearchRestoredFiles_KPURE
    EnvironmentSearchRestoredFiles_KAZURE
    EnvironmentSearchRestoredFiles_KNETAPP
    EnvironmentSearchRestoredFiles_KAGENT
    EnvironmentSearchRestoredFiles_KGENERICNAS
    EnvironmentSearchRestoredFiles_KACROPOLIS
    EnvironmentSearchRestoredFiles_KPHYSICALFILES
    EnvironmentSearchRestoredFiles_KISILON
    EnvironmentSearchRestoredFiles_KGPFS
    EnvironmentSearchRestoredFiles_KKVM
    EnvironmentSearchRestoredFiles_KAWS
    EnvironmentSearchRestoredFiles_KEXCHANGE
    EnvironmentSearchRestoredFiles_KHYPERVVSS
    EnvironmentSearchRestoredFiles_KORACLE
    EnvironmentSearchRestoredFiles_KGCP
    EnvironmentSearchRestoredFiles_KFLASHBLADE
    EnvironmentSearchRestoredFiles_KAWSNATIVE
    EnvironmentSearchRestoredFiles_KVCD
    EnvironmentSearchRestoredFiles_KO365
    EnvironmentSearchRestoredFiles_KO365OUTLOOK
    EnvironmentSearchRestoredFiles_KHYPERFLEX
    EnvironmentSearchRestoredFiles_KGCPNATIVE
    EnvironmentSearchRestoredFiles_KKUBERNETES
)

func (r EnvironmentSearchRestoredFilesEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentSearchRestoredFilesEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentSearchRestoredFilesEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentSearchRestoredFilesEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentSearchRestoredFilesEnum to its string representation
 */
func EnvironmentSearchRestoredFilesEnumToValue(environmentSearchRestoredFilesEnum EnvironmentSearchRestoredFilesEnum) string {
    switch environmentSearchRestoredFilesEnum {
        case EnvironmentSearchRestoredFiles_KVMWARE:
    		return "kVMware"
        case EnvironmentSearchRestoredFiles_KHYPERV:
    		return "kHyperV"
        case EnvironmentSearchRestoredFiles_KSQL:
    		return "kSQL"
        case EnvironmentSearchRestoredFiles_KVIEW:
    		return "kView"
        case EnvironmentSearchRestoredFiles_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentSearchRestoredFiles_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentSearchRestoredFiles_KPURE:
    		return "kPure"
        case EnvironmentSearchRestoredFiles_KAZURE:
    		return "kAzure"
        case EnvironmentSearchRestoredFiles_KNETAPP:
    		return "kNetapp"
        case EnvironmentSearchRestoredFiles_KAGENT:
    		return "kAgent"
        case EnvironmentSearchRestoredFiles_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentSearchRestoredFiles_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentSearchRestoredFiles_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentSearchRestoredFiles_KISILON:
    		return "kIsilon"
        case EnvironmentSearchRestoredFiles_KGPFS:
    		return "kGPFS"
        case EnvironmentSearchRestoredFiles_KKVM:
    		return "kKVM"
        case EnvironmentSearchRestoredFiles_KAWS:
    		return "kAWS"
        case EnvironmentSearchRestoredFiles_KEXCHANGE:
    		return "kExchange"
        case EnvironmentSearchRestoredFiles_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentSearchRestoredFiles_KORACLE:
    		return "kOracle"
        case EnvironmentSearchRestoredFiles_KGCP:
    		return "kGCP"
        case EnvironmentSearchRestoredFiles_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentSearchRestoredFiles_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentSearchRestoredFiles_KVCD:
    		return "kVCD"
        case EnvironmentSearchRestoredFiles_KO365:
    		return "kO365"
        case EnvironmentSearchRestoredFiles_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentSearchRestoredFiles_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentSearchRestoredFiles_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentSearchRestoredFiles_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentSearchRestoredFilesEnum Array to its string Array representation
*/
func EnvironmentSearchRestoredFilesEnumArrayToValue(environmentSearchRestoredFilesEnum []EnvironmentSearchRestoredFilesEnum) []string {
    convArray := make([]string,len( environmentSearchRestoredFilesEnum))
    for i:=0; i<len(environmentSearchRestoredFilesEnum);i++ {
        convArray[i] = EnvironmentSearchRestoredFilesEnumToValue(environmentSearchRestoredFilesEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentSearchRestoredFilesEnumFromValue(value string) EnvironmentSearchRestoredFilesEnum {
    switch value {
        case "kVMware":
            return EnvironmentSearchRestoredFiles_KVMWARE
        case "kHyperV":
            return EnvironmentSearchRestoredFiles_KHYPERV
        case "kSQL":
            return EnvironmentSearchRestoredFiles_KSQL
        case "kView":
            return EnvironmentSearchRestoredFiles_KVIEW
        case "kPuppeteer":
            return EnvironmentSearchRestoredFiles_KPUPPETEER
        case "kPhysical":
            return EnvironmentSearchRestoredFiles_KPHYSICAL
        case "kPure":
            return EnvironmentSearchRestoredFiles_KPURE
        case "kAzure":
            return EnvironmentSearchRestoredFiles_KAZURE
        case "kNetapp":
            return EnvironmentSearchRestoredFiles_KNETAPP
        case "kAgent":
            return EnvironmentSearchRestoredFiles_KAGENT
        case "kGenericNas":
            return EnvironmentSearchRestoredFiles_KGENERICNAS
        case "kAcropolis":
            return EnvironmentSearchRestoredFiles_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentSearchRestoredFiles_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentSearchRestoredFiles_KISILON
        case "kGPFS":
            return EnvironmentSearchRestoredFiles_KGPFS
        case "kKVM":
            return EnvironmentSearchRestoredFiles_KKVM
        case "kAWS":
            return EnvironmentSearchRestoredFiles_KAWS
        case "kExchange":
            return EnvironmentSearchRestoredFiles_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentSearchRestoredFiles_KHYPERVVSS
        case "kOracle":
            return EnvironmentSearchRestoredFiles_KORACLE
        case "kGCP":
            return EnvironmentSearchRestoredFiles_KGCP
        case "kFlashBlade":
            return EnvironmentSearchRestoredFiles_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentSearchRestoredFiles_KAWSNATIVE
        case "kVCD":
            return EnvironmentSearchRestoredFiles_KVCD
        case "kO365":
            return EnvironmentSearchRestoredFiles_KO365
        case "kO365Outlook":
            return EnvironmentSearchRestoredFiles_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentSearchRestoredFiles_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentSearchRestoredFiles_KGCPNATIVE
        case "kKubernetes":
            return EnvironmentSearchRestoredFiles_KKUBERNETES
        default:
            return EnvironmentSearchRestoredFiles_KVMWARE
    }
}
