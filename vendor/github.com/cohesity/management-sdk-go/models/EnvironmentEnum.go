// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentEnum enum
 */
type EnvironmentEnum int

/**
 * Value collection for EnvironmentEnum enum
 */
const (
    Environment_KVMWARE EnvironmentEnum = 1 + iota
    Environment_KHYPERV
    Environment_KSQL
    Environment_KVIEW
    Environment_KPUPPETEER
    Environment_KPHYSICAL
    Environment_KPURE
    Environment_KAZURE
    Environment_KNETAPP
    Environment_KAGENT
    Environment_KGENERICNAS
    Environment_KACROPOLIS
    Environment_KPHYSICALFILES
    Environment_KISILON
    Environment_KGPFS
    Environment_KKVM
    Environment_KAWS
    Environment_KEXCHANGE
    Environment_KHYPERVVSS
    Environment_KORACLE
    Environment_KGCP
    Environment_KFLASHBLADE
    Environment_KAWSNATIVE
    Environment_KVCD
    Environment_KO365
    Environment_KO365OUTLOOK
    Environment_KHYPERFLEX
    Environment_KGCPNATIVE
    Environment_KAZURENATIVE
    Environment_KKUBERNETES
)

func (r EnvironmentEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentEnum to its string representation
 */
func EnvironmentEnumToValue(environmentEnum EnvironmentEnum) string {
    switch environmentEnum {
        case Environment_KVMWARE:
    		return "kVMware"
        case Environment_KHYPERV:
    		return "kHyperV"
        case Environment_KSQL:
    		return "kSQL"
        case Environment_KVIEW:
    		return "kView"
        case Environment_KPUPPETEER:
    		return "kPuppeteer"
        case Environment_KPHYSICAL:
    		return "kPhysical"
        case Environment_KPURE:
    		return "kPure"
        case Environment_KAZURE:
    		return "kAzure"
        case Environment_KNETAPP:
    		return "kNetapp"
        case Environment_KAGENT:
    		return "kAgent"
        case Environment_KGENERICNAS:
    		return "kGenericNas"
        case Environment_KACROPOLIS:
    		return "kAcropolis"
        case Environment_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case Environment_KISILON:
    		return "kIsilon"
        case Environment_KGPFS:
    		return "kGPFS"
        case Environment_KKVM:
    		return "kKVM"
        case Environment_KAWS:
    		return "kAWS"
        case Environment_KEXCHANGE:
    		return "kExchange"
        case Environment_KHYPERVVSS:
    		return "kHyperVVSS"
        case Environment_KORACLE:
    		return "kOracle"
        case Environment_KGCP:
    		return "kGCP"
        case Environment_KFLASHBLADE:
    		return "kFlashBlade"
        case Environment_KAWSNATIVE:
    		return "kAWSNative"
        case Environment_KVCD:
    		return "kVCD"
        case Environment_KO365:
    		return "kO365"
        case Environment_KO365OUTLOOK:
    		return "kO365Outlook"
        case Environment_KHYPERFLEX:
    		return "kHyperFlex"
        case Environment_KGCPNATIVE:
    		return "kGCPNative"
        case Environment_KAZURENATIVE:
    		return "kAzureNative"
        case Environment_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentEnum Array to its string Array representation
*/
func EnvironmentEnumArrayToValue(environmentEnum []EnvironmentEnum) []string {
    convArray := make([]string,len( environmentEnum))
    for i:=0; i<len(environmentEnum);i++ {
        convArray[i] = EnvironmentEnumToValue(environmentEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentEnumFromValue(value string) EnvironmentEnum {
    switch value {
        case "kVMware":
            return Environment_KVMWARE
        case "kHyperV":
            return Environment_KHYPERV
        case "kSQL":
            return Environment_KSQL
        case "kView":
            return Environment_KVIEW
        case "kPuppeteer":
            return Environment_KPUPPETEER
        case "kPhysical":
            return Environment_KPHYSICAL
        case "kPure":
            return Environment_KPURE
        case "kAzure":
            return Environment_KAZURE
        case "kNetapp":
            return Environment_KNETAPP
        case "kAgent":
            return Environment_KAGENT
        case "kGenericNas":
            return Environment_KGENERICNAS
        case "kAcropolis":
            return Environment_KACROPOLIS
        case "kPhysicalFiles":
            return Environment_KPHYSICALFILES
        case "kIsilon":
            return Environment_KISILON
        case "kGPFS":
            return Environment_KGPFS
        case "kKVM":
            return Environment_KKVM
        case "kAWS":
            return Environment_KAWS
        case "kExchange":
            return Environment_KEXCHANGE
        case "kHyperVVSS":
            return Environment_KHYPERVVSS
        case "kOracle":
            return Environment_KORACLE
        case "kGCP":
            return Environment_KGCP
        case "kFlashBlade":
            return Environment_KFLASHBLADE
        case "kAWSNative":
            return Environment_KAWSNATIVE
        case "kVCD":
            return Environment_KVCD
        case "kO365":
            return Environment_KO365
        case "kO365Outlook":
            return Environment_KO365OUTLOOK
        case "kHyperFlex":
            return Environment_KHYPERFLEX
        case "kGCPNative":
            return Environment_KGCPNATIVE
        case "kAzureNative":
            return Environment_KAZURENATIVE
        case "kKubernetes":
            return Environment_KKUBERNETES
        default:
            return Environment_KVMWARE
    }
}
