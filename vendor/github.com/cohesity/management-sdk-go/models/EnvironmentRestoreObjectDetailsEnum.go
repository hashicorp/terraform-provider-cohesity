// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentRestoreObjectDetailsEnum enum
 */
type EnvironmentRestoreObjectDetailsEnum int

/**
 * Value collection for EnvironmentRestoreObjectDetailsEnum enum
 */
const (
    EnvironmentRestoreObjectDetails_KVMWARE EnvironmentRestoreObjectDetailsEnum = 1 + iota
    EnvironmentRestoreObjectDetails_KHYPERV
    EnvironmentRestoreObjectDetails_KSQL
    EnvironmentRestoreObjectDetails_KVIEW
    EnvironmentRestoreObjectDetails_KPUPPETEER
    EnvironmentRestoreObjectDetails_KPHYSICAL
    EnvironmentRestoreObjectDetails_KPURE
    EnvironmentRestoreObjectDetails_KAZURE
    EnvironmentRestoreObjectDetails_KNETAPP
    EnvironmentRestoreObjectDetails_KAGENT
    EnvironmentRestoreObjectDetails_KGENERICNAS
    EnvironmentRestoreObjectDetails_KACROPOLIS
    EnvironmentRestoreObjectDetails_KPHYSICALFILES
    EnvironmentRestoreObjectDetails_KISILON
    EnvironmentRestoreObjectDetails_KGPFS
    EnvironmentRestoreObjectDetails_KKVM
    EnvironmentRestoreObjectDetails_KAWS
    EnvironmentRestoreObjectDetails_KEXCHANGE
    EnvironmentRestoreObjectDetails_KHYPERVVSS
    EnvironmentRestoreObjectDetails_KORACLE
    EnvironmentRestoreObjectDetails_KGCP
    EnvironmentRestoreObjectDetails_KFLASHBLADE
    EnvironmentRestoreObjectDetails_KAWSNATIVE
    EnvironmentRestoreObjectDetails_KVCD
    EnvironmentRestoreObjectDetails_KO365
    EnvironmentRestoreObjectDetails_KO365OUTLOOK
    EnvironmentRestoreObjectDetails_KHYPERFLEX
    EnvironmentRestoreObjectDetails_KGCPNATIVE
    EnvironmentRestoreObjectDetails_KAZURENATIVE
    EnvironmentRestoreObjectDetails_KKUBERNETES
)

func (r EnvironmentRestoreObjectDetailsEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentRestoreObjectDetailsEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentRestoreObjectDetailsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentRestoreObjectDetailsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentRestoreObjectDetailsEnum to its string representation
 */
func EnvironmentRestoreObjectDetailsEnumToValue(environmentRestoreObjectDetailsEnum EnvironmentRestoreObjectDetailsEnum) string {
    switch environmentRestoreObjectDetailsEnum {
        case EnvironmentRestoreObjectDetails_KVMWARE:
    		return "kVMware"
        case EnvironmentRestoreObjectDetails_KHYPERV:
    		return "kHyperV"
        case EnvironmentRestoreObjectDetails_KSQL:
    		return "kSQL"
        case EnvironmentRestoreObjectDetails_KVIEW:
    		return "kView"
        case EnvironmentRestoreObjectDetails_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentRestoreObjectDetails_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentRestoreObjectDetails_KPURE:
    		return "kPure"
        case EnvironmentRestoreObjectDetails_KAZURE:
    		return "kAzure"
        case EnvironmentRestoreObjectDetails_KNETAPP:
    		return "kNetapp"
        case EnvironmentRestoreObjectDetails_KAGENT:
    		return "kAgent"
        case EnvironmentRestoreObjectDetails_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentRestoreObjectDetails_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentRestoreObjectDetails_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentRestoreObjectDetails_KISILON:
    		return "kIsilon"
        case EnvironmentRestoreObjectDetails_KGPFS:
    		return "kGPFS"
        case EnvironmentRestoreObjectDetails_KKVM:
    		return "kKVM"
        case EnvironmentRestoreObjectDetails_KAWS:
    		return "kAWS"
        case EnvironmentRestoreObjectDetails_KEXCHANGE:
    		return "kExchange"
        case EnvironmentRestoreObjectDetails_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentRestoreObjectDetails_KORACLE:
    		return "kOracle"
        case EnvironmentRestoreObjectDetails_KGCP:
    		return "kGCP"
        case EnvironmentRestoreObjectDetails_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentRestoreObjectDetails_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentRestoreObjectDetails_KVCD:
    		return "kVCD"
        case EnvironmentRestoreObjectDetails_KO365:
    		return "kO365"
        case EnvironmentRestoreObjectDetails_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentRestoreObjectDetails_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentRestoreObjectDetails_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentRestoreObjectDetails_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentRestoreObjectDetails_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentRestoreObjectDetailsEnum Array to its string Array representation
*/
func EnvironmentRestoreObjectDetailsEnumArrayToValue(environmentRestoreObjectDetailsEnum []EnvironmentRestoreObjectDetailsEnum) []string {
    convArray := make([]string,len( environmentRestoreObjectDetailsEnum))
    for i:=0; i<len(environmentRestoreObjectDetailsEnum);i++ {
        convArray[i] = EnvironmentRestoreObjectDetailsEnumToValue(environmentRestoreObjectDetailsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentRestoreObjectDetailsEnumFromValue(value string) EnvironmentRestoreObjectDetailsEnum {
    switch value {
        case "kVMware":
            return EnvironmentRestoreObjectDetails_KVMWARE
        case "kHyperV":
            return EnvironmentRestoreObjectDetails_KHYPERV
        case "kSQL":
            return EnvironmentRestoreObjectDetails_KSQL
        case "kView":
            return EnvironmentRestoreObjectDetails_KVIEW
        case "kPuppeteer":
            return EnvironmentRestoreObjectDetails_KPUPPETEER
        case "kPhysical":
            return EnvironmentRestoreObjectDetails_KPHYSICAL
        case "kPure":
            return EnvironmentRestoreObjectDetails_KPURE
        case "kAzure":
            return EnvironmentRestoreObjectDetails_KAZURE
        case "kNetapp":
            return EnvironmentRestoreObjectDetails_KNETAPP
        case "kAgent":
            return EnvironmentRestoreObjectDetails_KAGENT
        case "kGenericNas":
            return EnvironmentRestoreObjectDetails_KGENERICNAS
        case "kAcropolis":
            return EnvironmentRestoreObjectDetails_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentRestoreObjectDetails_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentRestoreObjectDetails_KISILON
        case "kGPFS":
            return EnvironmentRestoreObjectDetails_KGPFS
        case "kKVM":
            return EnvironmentRestoreObjectDetails_KKVM
        case "kAWS":
            return EnvironmentRestoreObjectDetails_KAWS
        case "kExchange":
            return EnvironmentRestoreObjectDetails_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentRestoreObjectDetails_KHYPERVVSS
        case "kOracle":
            return EnvironmentRestoreObjectDetails_KORACLE
        case "kGCP":
            return EnvironmentRestoreObjectDetails_KGCP
        case "kFlashBlade":
            return EnvironmentRestoreObjectDetails_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentRestoreObjectDetails_KAWSNATIVE
        case "kVCD":
            return EnvironmentRestoreObjectDetails_KVCD
        case "kO365":
            return EnvironmentRestoreObjectDetails_KO365
        case "kO365Outlook":
            return EnvironmentRestoreObjectDetails_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentRestoreObjectDetails_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentRestoreObjectDetails_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentRestoreObjectDetails_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentRestoreObjectDetails_KKUBERNETES
        default:
            return EnvironmentRestoreObjectDetails_KVMWARE
    }
}
