// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentApplicationInfoEnum enum
 */
type EnvironmentApplicationInfoEnum int

/**
 * Value collection for EnvironmentApplicationInfoEnum enum
 */
const (
    EnvironmentApplicationInfo_KVMWARE EnvironmentApplicationInfoEnum = 1 + iota
    EnvironmentApplicationInfo_KHYPERV
    EnvironmentApplicationInfo_KSQL
    EnvironmentApplicationInfo_KVIEW
    EnvironmentApplicationInfo_KPUPPETEER
    EnvironmentApplicationInfo_KPHYSICAL
    EnvironmentApplicationInfo_KPURE
    EnvironmentApplicationInfo_KAZURE
    EnvironmentApplicationInfo_KNETAPP
    EnvironmentApplicationInfo_KAGENT
    EnvironmentApplicationInfo_KGENERICNAS
    EnvironmentApplicationInfo_KACROPOLIS
    EnvironmentApplicationInfo_KPHYSICALFILES
    EnvironmentApplicationInfo_KISILON
    EnvironmentApplicationInfo_KGPFS
    EnvironmentApplicationInfo_KKVM
    EnvironmentApplicationInfo_KAWS
    EnvironmentApplicationInfo_KEXCHANGE
    EnvironmentApplicationInfo_KHYPERVVSS
    EnvironmentApplicationInfo_KORACLE
    EnvironmentApplicationInfo_KGCP
    EnvironmentApplicationInfo_KFLASHBLADE
    EnvironmentApplicationInfo_KAWSNATIVE
    EnvironmentApplicationInfo_KVCD
    EnvironmentApplicationInfo_KO365
    EnvironmentApplicationInfo_KO365OUTLOOK
    EnvironmentApplicationInfo_KHYPERFLEX
    EnvironmentApplicationInfo_KGCPNATIVE
    EnvironmentApplicationInfo_KAZURENATIVE
    EnvironmentApplicationInfo_KKUBERNETES
)

func (r EnvironmentApplicationInfoEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentApplicationInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentApplicationInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentApplicationInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentApplicationInfoEnum to its string representation
 */
func EnvironmentApplicationInfoEnumToValue(environmentApplicationInfoEnum EnvironmentApplicationInfoEnum) string {
    switch environmentApplicationInfoEnum {
        case EnvironmentApplicationInfo_KVMWARE:
    		return "kVMware"
        case EnvironmentApplicationInfo_KHYPERV:
    		return "kHyperV"
        case EnvironmentApplicationInfo_KSQL:
    		return "kSQL"
        case EnvironmentApplicationInfo_KVIEW:
    		return "kView"
        case EnvironmentApplicationInfo_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentApplicationInfo_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentApplicationInfo_KPURE:
    		return "kPure"
        case EnvironmentApplicationInfo_KAZURE:
    		return "kAzure"
        case EnvironmentApplicationInfo_KNETAPP:
    		return "kNetapp"
        case EnvironmentApplicationInfo_KAGENT:
    		return "kAgent"
        case EnvironmentApplicationInfo_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentApplicationInfo_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentApplicationInfo_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentApplicationInfo_KISILON:
    		return "kIsilon"
        case EnvironmentApplicationInfo_KGPFS:
    		return "kGPFS"
        case EnvironmentApplicationInfo_KKVM:
    		return "kKVM"
        case EnvironmentApplicationInfo_KAWS:
    		return "kAWS"
        case EnvironmentApplicationInfo_KEXCHANGE:
    		return "kExchange"
        case EnvironmentApplicationInfo_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentApplicationInfo_KORACLE:
    		return "kOracle"
        case EnvironmentApplicationInfo_KGCP:
    		return "kGCP"
        case EnvironmentApplicationInfo_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentApplicationInfo_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentApplicationInfo_KVCD:
    		return "kVCD"
        case EnvironmentApplicationInfo_KO365:
    		return "kO365"
        case EnvironmentApplicationInfo_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentApplicationInfo_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentApplicationInfo_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentApplicationInfo_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentApplicationInfo_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentApplicationInfoEnum Array to its string Array representation
*/
func EnvironmentApplicationInfoEnumArrayToValue(environmentApplicationInfoEnum []EnvironmentApplicationInfoEnum) []string {
    convArray := make([]string,len( environmentApplicationInfoEnum))
    for i:=0; i<len(environmentApplicationInfoEnum);i++ {
        convArray[i] = EnvironmentApplicationInfoEnumToValue(environmentApplicationInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentApplicationInfoEnumFromValue(value string) EnvironmentApplicationInfoEnum {
    switch value {
        case "kVMware":
            return EnvironmentApplicationInfo_KVMWARE
        case "kHyperV":
            return EnvironmentApplicationInfo_KHYPERV
        case "kSQL":
            return EnvironmentApplicationInfo_KSQL
        case "kView":
            return EnvironmentApplicationInfo_KVIEW
        case "kPuppeteer":
            return EnvironmentApplicationInfo_KPUPPETEER
        case "kPhysical":
            return EnvironmentApplicationInfo_KPHYSICAL
        case "kPure":
            return EnvironmentApplicationInfo_KPURE
        case "kAzure":
            return EnvironmentApplicationInfo_KAZURE
        case "kNetapp":
            return EnvironmentApplicationInfo_KNETAPP
        case "kAgent":
            return EnvironmentApplicationInfo_KAGENT
        case "kGenericNas":
            return EnvironmentApplicationInfo_KGENERICNAS
        case "kAcropolis":
            return EnvironmentApplicationInfo_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentApplicationInfo_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentApplicationInfo_KISILON
        case "kGPFS":
            return EnvironmentApplicationInfo_KGPFS
        case "kKVM":
            return EnvironmentApplicationInfo_KKVM
        case "kAWS":
            return EnvironmentApplicationInfo_KAWS
        case "kExchange":
            return EnvironmentApplicationInfo_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentApplicationInfo_KHYPERVVSS
        case "kOracle":
            return EnvironmentApplicationInfo_KORACLE
        case "kGCP":
            return EnvironmentApplicationInfo_KGCP
        case "kFlashBlade":
            return EnvironmentApplicationInfo_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentApplicationInfo_KAWSNATIVE
        case "kVCD":
            return EnvironmentApplicationInfo_KVCD
        case "kO365":
            return EnvironmentApplicationInfo_KO365
        case "kO365Outlook":
            return EnvironmentApplicationInfo_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentApplicationInfo_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentApplicationInfo_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentApplicationInfo_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentApplicationInfo_KKUBERNETES
        default:
            return EnvironmentApplicationInfo_KVMWARE
    }
}
