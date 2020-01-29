// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentSearchProtectionSourcesEnum enum
 */
type EnvironmentSearchProtectionSourcesEnum int

/**
 * Value collection for EnvironmentSearchProtectionSourcesEnum enum
 */
const (
    EnvironmentSearchProtectionSources_KVMWARE EnvironmentSearchProtectionSourcesEnum = 1 + iota
    EnvironmentSearchProtectionSources_KHYPERV
    EnvironmentSearchProtectionSources_KSQL
    EnvironmentSearchProtectionSources_KVIEW
    EnvironmentSearchProtectionSources_KPUPPETEER
    EnvironmentSearchProtectionSources_KPHYSICAL
    EnvironmentSearchProtectionSources_KPURE
    EnvironmentSearchProtectionSources_KAZURE
    EnvironmentSearchProtectionSources_KNETAPP
    EnvironmentSearchProtectionSources_KAGENT
    EnvironmentSearchProtectionSources_KGENERICNAS
    EnvironmentSearchProtectionSources_KACROPOLIS
    EnvironmentSearchProtectionSources_KPHYSICALFILES
    EnvironmentSearchProtectionSources_KISILON
    EnvironmentSearchProtectionSources_KGPFS
    EnvironmentSearchProtectionSources_KKVM
    EnvironmentSearchProtectionSources_KAWS
    EnvironmentSearchProtectionSources_KEXCHANGE
    EnvironmentSearchProtectionSources_KHYPERVVSS
    EnvironmentSearchProtectionSources_KORACLE
    EnvironmentSearchProtectionSources_KGCP
    EnvironmentSearchProtectionSources_KFLASHBLADE
    EnvironmentSearchProtectionSources_KAWSNATIVE
    EnvironmentSearchProtectionSources_KVCD
    EnvironmentSearchProtectionSources_KO365
    EnvironmentSearchProtectionSources_KO365OUTLOOK
    EnvironmentSearchProtectionSources_KHYPERFLEX
    EnvironmentSearchProtectionSources_KGCPNATIVE
    EnvironmentSearchProtectionSources_KAZURENATIVE
    EnvironmentSearchProtectionSources_KKUBERNETES
)

func (r EnvironmentSearchProtectionSourcesEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentSearchProtectionSourcesEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentSearchProtectionSourcesEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentSearchProtectionSourcesEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentSearchProtectionSourcesEnum to its string representation
 */
func EnvironmentSearchProtectionSourcesEnumToValue(environmentSearchProtectionSourcesEnum EnvironmentSearchProtectionSourcesEnum) string {
    switch environmentSearchProtectionSourcesEnum {
        case EnvironmentSearchProtectionSources_KVMWARE:
    		return "kVMware"
        case EnvironmentSearchProtectionSources_KHYPERV:
    		return "kHyperV"
        case EnvironmentSearchProtectionSources_KSQL:
    		return "kSQL"
        case EnvironmentSearchProtectionSources_KVIEW:
    		return "kView"
        case EnvironmentSearchProtectionSources_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentSearchProtectionSources_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentSearchProtectionSources_KPURE:
    		return "kPure"
        case EnvironmentSearchProtectionSources_KAZURE:
    		return "kAzure"
        case EnvironmentSearchProtectionSources_KNETAPP:
    		return "kNetapp"
        case EnvironmentSearchProtectionSources_KAGENT:
    		return "kAgent"
        case EnvironmentSearchProtectionSources_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentSearchProtectionSources_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentSearchProtectionSources_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentSearchProtectionSources_KISILON:
    		return "kIsilon"
        case EnvironmentSearchProtectionSources_KGPFS:
    		return "kGPFS"
        case EnvironmentSearchProtectionSources_KKVM:
    		return "kKVM"
        case EnvironmentSearchProtectionSources_KAWS:
    		return "kAWS"
        case EnvironmentSearchProtectionSources_KEXCHANGE:
    		return "kExchange"
        case EnvironmentSearchProtectionSources_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentSearchProtectionSources_KORACLE:
    		return "kOracle"
        case EnvironmentSearchProtectionSources_KGCP:
    		return "kGCP"
        case EnvironmentSearchProtectionSources_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentSearchProtectionSources_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentSearchProtectionSources_KVCD:
    		return "kVCD"
        case EnvironmentSearchProtectionSources_KO365:
    		return "kO365"
        case EnvironmentSearchProtectionSources_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentSearchProtectionSources_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentSearchProtectionSources_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentSearchProtectionSources_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentSearchProtectionSources_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentSearchProtectionSourcesEnum Array to its string Array representation
*/
func EnvironmentSearchProtectionSourcesEnumArrayToValue(environmentSearchProtectionSourcesEnum []EnvironmentSearchProtectionSourcesEnum) []string {
    convArray := make([]string,len( environmentSearchProtectionSourcesEnum))
    for i:=0; i<len(environmentSearchProtectionSourcesEnum);i++ {
        convArray[i] = EnvironmentSearchProtectionSourcesEnumToValue(environmentSearchProtectionSourcesEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentSearchProtectionSourcesEnumFromValue(value string) EnvironmentSearchProtectionSourcesEnum {
    switch value {
        case "kVMware":
            return EnvironmentSearchProtectionSources_KVMWARE
        case "kHyperV":
            return EnvironmentSearchProtectionSources_KHYPERV
        case "kSQL":
            return EnvironmentSearchProtectionSources_KSQL
        case "kView":
            return EnvironmentSearchProtectionSources_KVIEW
        case "kPuppeteer":
            return EnvironmentSearchProtectionSources_KPUPPETEER
        case "kPhysical":
            return EnvironmentSearchProtectionSources_KPHYSICAL
        case "kPure":
            return EnvironmentSearchProtectionSources_KPURE
        case "kAzure":
            return EnvironmentSearchProtectionSources_KAZURE
        case "kNetapp":
            return EnvironmentSearchProtectionSources_KNETAPP
        case "kAgent":
            return EnvironmentSearchProtectionSources_KAGENT
        case "kGenericNas":
            return EnvironmentSearchProtectionSources_KGENERICNAS
        case "kAcropolis":
            return EnvironmentSearchProtectionSources_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentSearchProtectionSources_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentSearchProtectionSources_KISILON
        case "kGPFS":
            return EnvironmentSearchProtectionSources_KGPFS
        case "kKVM":
            return EnvironmentSearchProtectionSources_KKVM
        case "kAWS":
            return EnvironmentSearchProtectionSources_KAWS
        case "kExchange":
            return EnvironmentSearchProtectionSources_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentSearchProtectionSources_KHYPERVVSS
        case "kOracle":
            return EnvironmentSearchProtectionSources_KORACLE
        case "kGCP":
            return EnvironmentSearchProtectionSources_KGCP
        case "kFlashBlade":
            return EnvironmentSearchProtectionSources_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentSearchProtectionSources_KAWSNATIVE
        case "kVCD":
            return EnvironmentSearchProtectionSources_KVCD
        case "kO365":
            return EnvironmentSearchProtectionSources_KO365
        case "kO365Outlook":
            return EnvironmentSearchProtectionSources_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentSearchProtectionSources_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentSearchProtectionSources_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentSearchProtectionSources_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentSearchProtectionSources_KKUBERNETES
        default:
            return EnvironmentSearchProtectionSources_KVMWARE
    }
}
