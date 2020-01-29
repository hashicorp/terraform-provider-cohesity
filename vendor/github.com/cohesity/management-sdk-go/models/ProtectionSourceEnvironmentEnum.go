// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ProtectionSourceEnvironmentEnum enum
 */
type ProtectionSourceEnvironmentEnum int

/**
 * Value collection for ProtectionSourceEnvironmentEnum enum
 */
const (
    ProtectionSourceEnvironment_KVMWARE ProtectionSourceEnvironmentEnum = 1 + iota
    ProtectionSourceEnvironment_KHYPERV
    ProtectionSourceEnvironment_KSQL
    ProtectionSourceEnvironment_KVIEW
    ProtectionSourceEnvironment_KPUPPETEER
    ProtectionSourceEnvironment_KPHYSICAL
    ProtectionSourceEnvironment_KPURE
    ProtectionSourceEnvironment_KAZURE
    ProtectionSourceEnvironment_KNETAPP
    ProtectionSourceEnvironment_KAGENT
    ProtectionSourceEnvironment_KGENERICNAS
    ProtectionSourceEnvironment_KACROPOLIS
    ProtectionSourceEnvironment_KPHYSICALFILES
    ProtectionSourceEnvironment_KISILON
    ProtectionSourceEnvironment_KGPFS
    ProtectionSourceEnvironment_KKVM
    ProtectionSourceEnvironment_KAWS
    ProtectionSourceEnvironment_KEXCHANGE
    ProtectionSourceEnvironment_KHYPERVVSS
    ProtectionSourceEnvironment_KORACLE
    ProtectionSourceEnvironment_KGCP
    ProtectionSourceEnvironment_KFLASHBLADE
    ProtectionSourceEnvironment_KAWSNATIVE
    ProtectionSourceEnvironment_KVCD
    ProtectionSourceEnvironment_KO365
    ProtectionSourceEnvironment_KO365OUTLOOK
    ProtectionSourceEnvironment_KHYPERFLEX
    ProtectionSourceEnvironment_KGCPNATIVE
    ProtectionSourceEnvironment_KAZURENATIVE
    ProtectionSourceEnvironment_KKUBERNETES
)

func (r ProtectionSourceEnvironmentEnum) MarshalJSON() ([]byte, error) {
    s := ProtectionSourceEnvironmentEnumToValue(r)
    return json.Marshal(s)
}

func (r *ProtectionSourceEnvironmentEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ProtectionSourceEnvironmentEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ProtectionSourceEnvironmentEnum to its string representation
 */
func ProtectionSourceEnvironmentEnumToValue(protectionSourceEnvironmentEnum ProtectionSourceEnvironmentEnum) string {
    switch protectionSourceEnvironmentEnum {
        case ProtectionSourceEnvironment_KVMWARE:
    		return "kVMware"
        case ProtectionSourceEnvironment_KHYPERV:
    		return "kHyperV"
        case ProtectionSourceEnvironment_KSQL:
    		return "kSQL"
        case ProtectionSourceEnvironment_KVIEW:
    		return "kView"
        case ProtectionSourceEnvironment_KPUPPETEER:
    		return "kPuppeteer"
        case ProtectionSourceEnvironment_KPHYSICAL:
    		return "kPhysical"
        case ProtectionSourceEnvironment_KPURE:
    		return "kPure"
        case ProtectionSourceEnvironment_KAZURE:
    		return "kAzure"
        case ProtectionSourceEnvironment_KNETAPP:
    		return "kNetapp"
        case ProtectionSourceEnvironment_KAGENT:
    		return "kAgent"
        case ProtectionSourceEnvironment_KGENERICNAS:
    		return "kGenericNas"
        case ProtectionSourceEnvironment_KACROPOLIS:
    		return "kAcropolis"
        case ProtectionSourceEnvironment_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case ProtectionSourceEnvironment_KISILON:
    		return "kIsilon"
        case ProtectionSourceEnvironment_KGPFS:
    		return "kGPFS"
        case ProtectionSourceEnvironment_KKVM:
    		return "kKVM"
        case ProtectionSourceEnvironment_KAWS:
    		return "kAWS"
        case ProtectionSourceEnvironment_KEXCHANGE:
    		return "kExchange"
        case ProtectionSourceEnvironment_KHYPERVVSS:
    		return "kHyperVVSS"
        case ProtectionSourceEnvironment_KORACLE:
    		return "kOracle"
        case ProtectionSourceEnvironment_KGCP:
    		return "kGCP"
        case ProtectionSourceEnvironment_KFLASHBLADE:
    		return "kFlashBlade"
        case ProtectionSourceEnvironment_KAWSNATIVE:
    		return "kAWSNative"
        case ProtectionSourceEnvironment_KVCD:
    		return "kVCD"
        case ProtectionSourceEnvironment_KO365:
    		return "kO365"
        case ProtectionSourceEnvironment_KO365OUTLOOK:
    		return "kO365Outlook"
        case ProtectionSourceEnvironment_KHYPERFLEX:
    		return "kHyperFlex"
        case ProtectionSourceEnvironment_KGCPNATIVE:
    		return "kGCPNative"
        case ProtectionSourceEnvironment_KAZURENATIVE:
    		return "kAzureNative"
        case ProtectionSourceEnvironment_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts ProtectionSourceEnvironmentEnum Array to its string Array representation
*/
func ProtectionSourceEnvironmentEnumArrayToValue(protectionSourceEnvironmentEnum []ProtectionSourceEnvironmentEnum) []string {
    convArray := make([]string,len( protectionSourceEnvironmentEnum))
    for i:=0; i<len(protectionSourceEnvironmentEnum);i++ {
        convArray[i] = ProtectionSourceEnvironmentEnumToValue(protectionSourceEnvironmentEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ProtectionSourceEnvironmentEnumFromValue(value string) ProtectionSourceEnvironmentEnum {
    switch value {
        case "kVMware":
            return ProtectionSourceEnvironment_KVMWARE
        case "kHyperV":
            return ProtectionSourceEnvironment_KHYPERV
        case "kSQL":
            return ProtectionSourceEnvironment_KSQL
        case "kView":
            return ProtectionSourceEnvironment_KVIEW
        case "kPuppeteer":
            return ProtectionSourceEnvironment_KPUPPETEER
        case "kPhysical":
            return ProtectionSourceEnvironment_KPHYSICAL
        case "kPure":
            return ProtectionSourceEnvironment_KPURE
        case "kAzure":
            return ProtectionSourceEnvironment_KAZURE
        case "kNetapp":
            return ProtectionSourceEnvironment_KNETAPP
        case "kAgent":
            return ProtectionSourceEnvironment_KAGENT
        case "kGenericNas":
            return ProtectionSourceEnvironment_KGENERICNAS
        case "kAcropolis":
            return ProtectionSourceEnvironment_KACROPOLIS
        case "kPhysicalFiles":
            return ProtectionSourceEnvironment_KPHYSICALFILES
        case "kIsilon":
            return ProtectionSourceEnvironment_KISILON
        case "kGPFS":
            return ProtectionSourceEnvironment_KGPFS
        case "kKVM":
            return ProtectionSourceEnvironment_KKVM
        case "kAWS":
            return ProtectionSourceEnvironment_KAWS
        case "kExchange":
            return ProtectionSourceEnvironment_KEXCHANGE
        case "kHyperVVSS":
            return ProtectionSourceEnvironment_KHYPERVVSS
        case "kOracle":
            return ProtectionSourceEnvironment_KORACLE
        case "kGCP":
            return ProtectionSourceEnvironment_KGCP
        case "kFlashBlade":
            return ProtectionSourceEnvironment_KFLASHBLADE
        case "kAWSNative":
            return ProtectionSourceEnvironment_KAWSNATIVE
        case "kVCD":
            return ProtectionSourceEnvironment_KVCD
        case "kO365":
            return ProtectionSourceEnvironment_KO365
        case "kO365Outlook":
            return ProtectionSourceEnvironment_KO365OUTLOOK
        case "kHyperFlex":
            return ProtectionSourceEnvironment_KHYPERFLEX
        case "kGCPNative":
            return ProtectionSourceEnvironment_KGCPNATIVE
        case "kAzureNative":
            return ProtectionSourceEnvironment_KAZURENATIVE
        case "kKubernetes":
            return ProtectionSourceEnvironment_KKUBERNETES
        default:
            return ProtectionSourceEnvironment_KVMWARE
    }
}
