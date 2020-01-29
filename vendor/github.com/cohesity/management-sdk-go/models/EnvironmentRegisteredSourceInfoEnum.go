// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentRegisteredSourceInfoEnum enum
 */
type EnvironmentRegisteredSourceInfoEnum int

/**
 * Value collection for EnvironmentRegisteredSourceInfoEnum enum
 */
const (
    EnvironmentRegisteredSourceInfo_KVMWARE EnvironmentRegisteredSourceInfoEnum = 1 + iota
    EnvironmentRegisteredSourceInfo_KHYPERV
    EnvironmentRegisteredSourceInfo_KSQL
    EnvironmentRegisteredSourceInfo_KVIEW
    EnvironmentRegisteredSourceInfo_KPUPPETEER
    EnvironmentRegisteredSourceInfo_KPHYSICAL
    EnvironmentRegisteredSourceInfo_KPURE
    EnvironmentRegisteredSourceInfo_KAZURE
    EnvironmentRegisteredSourceInfo_KNETAPP
    EnvironmentRegisteredSourceInfo_KAGENT
    EnvironmentRegisteredSourceInfo_KGENERICNAS
    EnvironmentRegisteredSourceInfo_KACROPOLIS
    EnvironmentRegisteredSourceInfo_KPHYSICALFILES
    EnvironmentRegisteredSourceInfo_KISILON
    EnvironmentRegisteredSourceInfo_KGPFS
    EnvironmentRegisteredSourceInfo_KKVM
    EnvironmentRegisteredSourceInfo_KAWS
    EnvironmentRegisteredSourceInfo_KEXCHANGE
    EnvironmentRegisteredSourceInfo_KHYPERVVSS
    EnvironmentRegisteredSourceInfo_KORACLE
    EnvironmentRegisteredSourceInfo_KGCP
    EnvironmentRegisteredSourceInfo_KFLASHBLADE
    EnvironmentRegisteredSourceInfo_KAWSNATIVE
    EnvironmentRegisteredSourceInfo_KVCD
    EnvironmentRegisteredSourceInfo_KO365
    EnvironmentRegisteredSourceInfo_KO365OUTLOOK
    EnvironmentRegisteredSourceInfo_KHYPERFLEX
    EnvironmentRegisteredSourceInfo_KGCPNATIVE
    EnvironmentRegisteredSourceInfo_KAZURENATIVE
    EnvironmentRegisteredSourceInfo_KKUBERNETES
)

func (r EnvironmentRegisteredSourceInfoEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentRegisteredSourceInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentRegisteredSourceInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentRegisteredSourceInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentRegisteredSourceInfoEnum to its string representation
 */
func EnvironmentRegisteredSourceInfoEnumToValue(environmentRegisteredSourceInfoEnum EnvironmentRegisteredSourceInfoEnum) string {
    switch environmentRegisteredSourceInfoEnum {
        case EnvironmentRegisteredSourceInfo_KVMWARE:
    		return "kVMware"
        case EnvironmentRegisteredSourceInfo_KHYPERV:
    		return "kHyperV"
        case EnvironmentRegisteredSourceInfo_KSQL:
    		return "kSQL"
        case EnvironmentRegisteredSourceInfo_KVIEW:
    		return "kView"
        case EnvironmentRegisteredSourceInfo_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentRegisteredSourceInfo_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentRegisteredSourceInfo_KPURE:
    		return "kPure"
        case EnvironmentRegisteredSourceInfo_KAZURE:
    		return "kAzure"
        case EnvironmentRegisteredSourceInfo_KNETAPP:
    		return "kNetapp"
        case EnvironmentRegisteredSourceInfo_KAGENT:
    		return "kAgent"
        case EnvironmentRegisteredSourceInfo_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentRegisteredSourceInfo_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentRegisteredSourceInfo_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentRegisteredSourceInfo_KISILON:
    		return "kIsilon"
        case EnvironmentRegisteredSourceInfo_KGPFS:
    		return "kGPFS"
        case EnvironmentRegisteredSourceInfo_KKVM:
    		return "kKVM"
        case EnvironmentRegisteredSourceInfo_KAWS:
    		return "kAWS"
        case EnvironmentRegisteredSourceInfo_KEXCHANGE:
    		return "kExchange"
        case EnvironmentRegisteredSourceInfo_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentRegisteredSourceInfo_KORACLE:
    		return "kOracle"
        case EnvironmentRegisteredSourceInfo_KGCP:
    		return "kGCP"
        case EnvironmentRegisteredSourceInfo_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentRegisteredSourceInfo_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentRegisteredSourceInfo_KVCD:
    		return "kVCD"
        case EnvironmentRegisteredSourceInfo_KO365:
    		return "kO365"
        case EnvironmentRegisteredSourceInfo_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentRegisteredSourceInfo_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentRegisteredSourceInfo_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentRegisteredSourceInfo_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentRegisteredSourceInfo_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentRegisteredSourceInfoEnum Array to its string Array representation
*/
func EnvironmentRegisteredSourceInfoEnumArrayToValue(environmentRegisteredSourceInfoEnum []EnvironmentRegisteredSourceInfoEnum) []string {
    convArray := make([]string,len( environmentRegisteredSourceInfoEnum))
    for i:=0; i<len(environmentRegisteredSourceInfoEnum);i++ {
        convArray[i] = EnvironmentRegisteredSourceInfoEnumToValue(environmentRegisteredSourceInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentRegisteredSourceInfoEnumFromValue(value string) EnvironmentRegisteredSourceInfoEnum {
    switch value {
        case "kVMware":
            return EnvironmentRegisteredSourceInfo_KVMWARE
        case "kHyperV":
            return EnvironmentRegisteredSourceInfo_KHYPERV
        case "kSQL":
            return EnvironmentRegisteredSourceInfo_KSQL
        case "kView":
            return EnvironmentRegisteredSourceInfo_KVIEW
        case "kPuppeteer":
            return EnvironmentRegisteredSourceInfo_KPUPPETEER
        case "kPhysical":
            return EnvironmentRegisteredSourceInfo_KPHYSICAL
        case "kPure":
            return EnvironmentRegisteredSourceInfo_KPURE
        case "kAzure":
            return EnvironmentRegisteredSourceInfo_KAZURE
        case "kNetapp":
            return EnvironmentRegisteredSourceInfo_KNETAPP
        case "kAgent":
            return EnvironmentRegisteredSourceInfo_KAGENT
        case "kGenericNas":
            return EnvironmentRegisteredSourceInfo_KGENERICNAS
        case "kAcropolis":
            return EnvironmentRegisteredSourceInfo_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentRegisteredSourceInfo_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentRegisteredSourceInfo_KISILON
        case "kGPFS":
            return EnvironmentRegisteredSourceInfo_KGPFS
        case "kKVM":
            return EnvironmentRegisteredSourceInfo_KKVM
        case "kAWS":
            return EnvironmentRegisteredSourceInfo_KAWS
        case "kExchange":
            return EnvironmentRegisteredSourceInfo_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentRegisteredSourceInfo_KHYPERVVSS
        case "kOracle":
            return EnvironmentRegisteredSourceInfo_KORACLE
        case "kGCP":
            return EnvironmentRegisteredSourceInfo_KGCP
        case "kFlashBlade":
            return EnvironmentRegisteredSourceInfo_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentRegisteredSourceInfo_KAWSNATIVE
        case "kVCD":
            return EnvironmentRegisteredSourceInfo_KVCD
        case "kO365":
            return EnvironmentRegisteredSourceInfo_KO365
        case "kO365Outlook":
            return EnvironmentRegisteredSourceInfo_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentRegisteredSourceInfo_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentRegisteredSourceInfo_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentRegisteredSourceInfo_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentRegisteredSourceInfo_KKUBERNETES
        default:
            return EnvironmentRegisteredSourceInfo_KVMWARE
    }
}
