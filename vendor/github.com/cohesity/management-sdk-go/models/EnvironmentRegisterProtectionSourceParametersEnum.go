// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentRegisterProtectionSourceParametersEnum enum
 */
type EnvironmentRegisterProtectionSourceParametersEnum int

/**
 * Value collection for EnvironmentRegisterProtectionSourceParametersEnum enum
 */
const (
    EnvironmentRegisterProtectionSourceParameters_KVMWARE EnvironmentRegisterProtectionSourceParametersEnum = 1 + iota
    EnvironmentRegisterProtectionSourceParameters_KHYPERV
    EnvironmentRegisterProtectionSourceParameters_KSQL
    EnvironmentRegisterProtectionSourceParameters_KVIEW
    EnvironmentRegisterProtectionSourceParameters_KPUPPETEER
    EnvironmentRegisterProtectionSourceParameters_KPHYSICAL
    EnvironmentRegisterProtectionSourceParameters_KPURE
    EnvironmentRegisterProtectionSourceParameters_KAZURE
    EnvironmentRegisterProtectionSourceParameters_KNETAPP
    EnvironmentRegisterProtectionSourceParameters_KAGENT
    EnvironmentRegisterProtectionSourceParameters_KGENERICNAS
    EnvironmentRegisterProtectionSourceParameters_KACROPOLIS
    EnvironmentRegisterProtectionSourceParameters_KPHYSICALFILES
    EnvironmentRegisterProtectionSourceParameters_KISILON
    EnvironmentRegisterProtectionSourceParameters_KGPFS
    EnvironmentRegisterProtectionSourceParameters_KKVM
    EnvironmentRegisterProtectionSourceParameters_KAWS
    EnvironmentRegisterProtectionSourceParameters_KEXCHANGE
    EnvironmentRegisterProtectionSourceParameters_KHYPERVVSS
    EnvironmentRegisterProtectionSourceParameters_KORACLE
    EnvironmentRegisterProtectionSourceParameters_KGCP
    EnvironmentRegisterProtectionSourceParameters_KFLASHBLADE
    EnvironmentRegisterProtectionSourceParameters_KAWSNATIVE
    EnvironmentRegisterProtectionSourceParameters_KVCD
    EnvironmentRegisterProtectionSourceParameters_KO365
    EnvironmentRegisterProtectionSourceParameters_KO365OUTLOOK
    EnvironmentRegisterProtectionSourceParameters_KHYPERFLEX
    EnvironmentRegisterProtectionSourceParameters_KGCPNATIVE
    EnvironmentRegisterProtectionSourceParameters_KAZURENATIVE
    EnvironmentRegisterProtectionSourceParameters_KKUBERNETES
)

func (r EnvironmentRegisterProtectionSourceParametersEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentRegisterProtectionSourceParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentRegisterProtectionSourceParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentRegisterProtectionSourceParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentRegisterProtectionSourceParametersEnum to its string representation
 */
func EnvironmentRegisterProtectionSourceParametersEnumToValue(environmentRegisterProtectionSourceParametersEnum EnvironmentRegisterProtectionSourceParametersEnum) string {
    switch environmentRegisterProtectionSourceParametersEnum {
        case EnvironmentRegisterProtectionSourceParameters_KVMWARE:
    		return "kVMware"
        case EnvironmentRegisterProtectionSourceParameters_KHYPERV:
    		return "kHyperV"
        case EnvironmentRegisterProtectionSourceParameters_KSQL:
    		return "kSQL"
        case EnvironmentRegisterProtectionSourceParameters_KVIEW:
    		return "kView"
        case EnvironmentRegisterProtectionSourceParameters_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentRegisterProtectionSourceParameters_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentRegisterProtectionSourceParameters_KPURE:
    		return "kPure"
        case EnvironmentRegisterProtectionSourceParameters_KAZURE:
    		return "kAzure"
        case EnvironmentRegisterProtectionSourceParameters_KNETAPP:
    		return "kNetapp"
        case EnvironmentRegisterProtectionSourceParameters_KAGENT:
    		return "kAgent"
        case EnvironmentRegisterProtectionSourceParameters_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentRegisterProtectionSourceParameters_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentRegisterProtectionSourceParameters_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentRegisterProtectionSourceParameters_KISILON:
    		return "kIsilon"
        case EnvironmentRegisterProtectionSourceParameters_KGPFS:
    		return "kGPFS"
        case EnvironmentRegisterProtectionSourceParameters_KKVM:
    		return "kKVM"
        case EnvironmentRegisterProtectionSourceParameters_KAWS:
    		return "kAWS"
        case EnvironmentRegisterProtectionSourceParameters_KEXCHANGE:
    		return "kExchange"
        case EnvironmentRegisterProtectionSourceParameters_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentRegisterProtectionSourceParameters_KORACLE:
    		return "kOracle"
        case EnvironmentRegisterProtectionSourceParameters_KGCP:
    		return "kGCP"
        case EnvironmentRegisterProtectionSourceParameters_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentRegisterProtectionSourceParameters_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentRegisterProtectionSourceParameters_KVCD:
    		return "kVCD"
        case EnvironmentRegisterProtectionSourceParameters_KO365:
    		return "kO365"
        case EnvironmentRegisterProtectionSourceParameters_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentRegisterProtectionSourceParameters_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentRegisterProtectionSourceParameters_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentRegisterProtectionSourceParameters_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentRegisterProtectionSourceParameters_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentRegisterProtectionSourceParametersEnum Array to its string Array representation
*/
func EnvironmentRegisterProtectionSourceParametersEnumArrayToValue(environmentRegisterProtectionSourceParametersEnum []EnvironmentRegisterProtectionSourceParametersEnum) []string {
    convArray := make([]string,len( environmentRegisterProtectionSourceParametersEnum))
    for i:=0; i<len(environmentRegisterProtectionSourceParametersEnum);i++ {
        convArray[i] = EnvironmentRegisterProtectionSourceParametersEnumToValue(environmentRegisterProtectionSourceParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentRegisterProtectionSourceParametersEnumFromValue(value string) EnvironmentRegisterProtectionSourceParametersEnum {
    switch value {
        case "kVMware":
            return EnvironmentRegisterProtectionSourceParameters_KVMWARE
        case "kHyperV":
            return EnvironmentRegisterProtectionSourceParameters_KHYPERV
        case "kSQL":
            return EnvironmentRegisterProtectionSourceParameters_KSQL
        case "kView":
            return EnvironmentRegisterProtectionSourceParameters_KVIEW
        case "kPuppeteer":
            return EnvironmentRegisterProtectionSourceParameters_KPUPPETEER
        case "kPhysical":
            return EnvironmentRegisterProtectionSourceParameters_KPHYSICAL
        case "kPure":
            return EnvironmentRegisterProtectionSourceParameters_KPURE
        case "kAzure":
            return EnvironmentRegisterProtectionSourceParameters_KAZURE
        case "kNetapp":
            return EnvironmentRegisterProtectionSourceParameters_KNETAPP
        case "kAgent":
            return EnvironmentRegisterProtectionSourceParameters_KAGENT
        case "kGenericNas":
            return EnvironmentRegisterProtectionSourceParameters_KGENERICNAS
        case "kAcropolis":
            return EnvironmentRegisterProtectionSourceParameters_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentRegisterProtectionSourceParameters_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentRegisterProtectionSourceParameters_KISILON
        case "kGPFS":
            return EnvironmentRegisterProtectionSourceParameters_KGPFS
        case "kKVM":
            return EnvironmentRegisterProtectionSourceParameters_KKVM
        case "kAWS":
            return EnvironmentRegisterProtectionSourceParameters_KAWS
        case "kExchange":
            return EnvironmentRegisterProtectionSourceParameters_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentRegisterProtectionSourceParameters_KHYPERVVSS
        case "kOracle":
            return EnvironmentRegisterProtectionSourceParameters_KORACLE
        case "kGCP":
            return EnvironmentRegisterProtectionSourceParameters_KGCP
        case "kFlashBlade":
            return EnvironmentRegisterProtectionSourceParameters_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentRegisterProtectionSourceParameters_KAWSNATIVE
        case "kVCD":
            return EnvironmentRegisterProtectionSourceParameters_KVCD
        case "kO365":
            return EnvironmentRegisterProtectionSourceParameters_KO365
        case "kO365Outlook":
            return EnvironmentRegisterProtectionSourceParameters_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentRegisterProtectionSourceParameters_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentRegisterProtectionSourceParameters_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentRegisterProtectionSourceParameters_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentRegisterProtectionSourceParameters_KKUBERNETES
        default:
            return EnvironmentRegisterProtectionSourceParameters_KVMWARE
    }
}
