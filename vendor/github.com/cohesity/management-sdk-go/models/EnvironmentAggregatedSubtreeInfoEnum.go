// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentAggregatedSubtreeInfoEnum enum
 */
type EnvironmentAggregatedSubtreeInfoEnum int

/**
 * Value collection for EnvironmentAggregatedSubtreeInfoEnum enum
 */
const (
    EnvironmentAggregatedSubtreeInfo_KVMWARE EnvironmentAggregatedSubtreeInfoEnum = 1 + iota
    EnvironmentAggregatedSubtreeInfo_KHYPERV
    EnvironmentAggregatedSubtreeInfo_KSQL
    EnvironmentAggregatedSubtreeInfo_KVIEW
    EnvironmentAggregatedSubtreeInfo_KPUPPETEER
    EnvironmentAggregatedSubtreeInfo_KPHYSICAL
    EnvironmentAggregatedSubtreeInfo_KPURE
    EnvironmentAggregatedSubtreeInfo_KAZURE
    EnvironmentAggregatedSubtreeInfo_KNETAPP
    EnvironmentAggregatedSubtreeInfo_KAGENT
    EnvironmentAggregatedSubtreeInfo_KGENERICNAS
    EnvironmentAggregatedSubtreeInfo_KACROPOLIS
    EnvironmentAggregatedSubtreeInfo_KPHYSICALFILES
    EnvironmentAggregatedSubtreeInfo_KISILON
    EnvironmentAggregatedSubtreeInfo_KGPFS
    EnvironmentAggregatedSubtreeInfo_KKVM
    EnvironmentAggregatedSubtreeInfo_KAWS
    EnvironmentAggregatedSubtreeInfo_KEXCHANGE
    EnvironmentAggregatedSubtreeInfo_KHYPERVVSS
    EnvironmentAggregatedSubtreeInfo_KORACLE
    EnvironmentAggregatedSubtreeInfo_KGCP
    EnvironmentAggregatedSubtreeInfo_KFLASHBLADE
    EnvironmentAggregatedSubtreeInfo_KAWSNATIVE
    EnvironmentAggregatedSubtreeInfo_KVCD
    EnvironmentAggregatedSubtreeInfo_KO365
    EnvironmentAggregatedSubtreeInfo_KO365OUTLOOK
    EnvironmentAggregatedSubtreeInfo_KHYPERFLEX
    EnvironmentAggregatedSubtreeInfo_KGCPNATIVE
    EnvironmentAggregatedSubtreeInfo_KAZURENATIVE
    EnvironmentAggregatedSubtreeInfo_KKUBERNETES
)

func (r EnvironmentAggregatedSubtreeInfoEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentAggregatedSubtreeInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentAggregatedSubtreeInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentAggregatedSubtreeInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentAggregatedSubtreeInfoEnum to its string representation
 */
func EnvironmentAggregatedSubtreeInfoEnumToValue(environmentAggregatedSubtreeInfoEnum EnvironmentAggregatedSubtreeInfoEnum) string {
    switch environmentAggregatedSubtreeInfoEnum {
        case EnvironmentAggregatedSubtreeInfo_KVMWARE:
    		return "kVMware"
        case EnvironmentAggregatedSubtreeInfo_KHYPERV:
    		return "kHyperV"
        case EnvironmentAggregatedSubtreeInfo_KSQL:
    		return "kSQL"
        case EnvironmentAggregatedSubtreeInfo_KVIEW:
    		return "kView"
        case EnvironmentAggregatedSubtreeInfo_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentAggregatedSubtreeInfo_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentAggregatedSubtreeInfo_KPURE:
    		return "kPure"
        case EnvironmentAggregatedSubtreeInfo_KAZURE:
    		return "kAzure"
        case EnvironmentAggregatedSubtreeInfo_KNETAPP:
    		return "kNetapp"
        case EnvironmentAggregatedSubtreeInfo_KAGENT:
    		return "kAgent"
        case EnvironmentAggregatedSubtreeInfo_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentAggregatedSubtreeInfo_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentAggregatedSubtreeInfo_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentAggregatedSubtreeInfo_KISILON:
    		return "kIsilon"
        case EnvironmentAggregatedSubtreeInfo_KGPFS:
    		return "kGPFS"
        case EnvironmentAggregatedSubtreeInfo_KKVM:
    		return "kKVM"
        case EnvironmentAggregatedSubtreeInfo_KAWS:
    		return "kAWS"
        case EnvironmentAggregatedSubtreeInfo_KEXCHANGE:
    		return "kExchange"
        case EnvironmentAggregatedSubtreeInfo_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentAggregatedSubtreeInfo_KORACLE:
    		return "kOracle"
        case EnvironmentAggregatedSubtreeInfo_KGCP:
    		return "kGCP"
        case EnvironmentAggregatedSubtreeInfo_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentAggregatedSubtreeInfo_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentAggregatedSubtreeInfo_KVCD:
    		return "kVCD"
        case EnvironmentAggregatedSubtreeInfo_KO365:
    		return "kO365"
        case EnvironmentAggregatedSubtreeInfo_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentAggregatedSubtreeInfo_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentAggregatedSubtreeInfo_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentAggregatedSubtreeInfo_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentAggregatedSubtreeInfo_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentAggregatedSubtreeInfoEnum Array to its string Array representation
*/
func EnvironmentAggregatedSubtreeInfoEnumArrayToValue(environmentAggregatedSubtreeInfoEnum []EnvironmentAggregatedSubtreeInfoEnum) []string {
    convArray := make([]string,len( environmentAggregatedSubtreeInfoEnum))
    for i:=0; i<len(environmentAggregatedSubtreeInfoEnum);i++ {
        convArray[i] = EnvironmentAggregatedSubtreeInfoEnumToValue(environmentAggregatedSubtreeInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentAggregatedSubtreeInfoEnumFromValue(value string) EnvironmentAggregatedSubtreeInfoEnum {
    switch value {
        case "kVMware":
            return EnvironmentAggregatedSubtreeInfo_KVMWARE
        case "kHyperV":
            return EnvironmentAggregatedSubtreeInfo_KHYPERV
        case "kSQL":
            return EnvironmentAggregatedSubtreeInfo_KSQL
        case "kView":
            return EnvironmentAggregatedSubtreeInfo_KVIEW
        case "kPuppeteer":
            return EnvironmentAggregatedSubtreeInfo_KPUPPETEER
        case "kPhysical":
            return EnvironmentAggregatedSubtreeInfo_KPHYSICAL
        case "kPure":
            return EnvironmentAggregatedSubtreeInfo_KPURE
        case "kAzure":
            return EnvironmentAggregatedSubtreeInfo_KAZURE
        case "kNetapp":
            return EnvironmentAggregatedSubtreeInfo_KNETAPP
        case "kAgent":
            return EnvironmentAggregatedSubtreeInfo_KAGENT
        case "kGenericNas":
            return EnvironmentAggregatedSubtreeInfo_KGENERICNAS
        case "kAcropolis":
            return EnvironmentAggregatedSubtreeInfo_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentAggregatedSubtreeInfo_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentAggregatedSubtreeInfo_KISILON
        case "kGPFS":
            return EnvironmentAggregatedSubtreeInfo_KGPFS
        case "kKVM":
            return EnvironmentAggregatedSubtreeInfo_KKVM
        case "kAWS":
            return EnvironmentAggregatedSubtreeInfo_KAWS
        case "kExchange":
            return EnvironmentAggregatedSubtreeInfo_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentAggregatedSubtreeInfo_KHYPERVVSS
        case "kOracle":
            return EnvironmentAggregatedSubtreeInfo_KORACLE
        case "kGCP":
            return EnvironmentAggregatedSubtreeInfo_KGCP
        case "kFlashBlade":
            return EnvironmentAggregatedSubtreeInfo_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentAggregatedSubtreeInfo_KAWSNATIVE
        case "kVCD":
            return EnvironmentAggregatedSubtreeInfo_KVCD
        case "kO365":
            return EnvironmentAggregatedSubtreeInfo_KO365
        case "kO365Outlook":
            return EnvironmentAggregatedSubtreeInfo_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentAggregatedSubtreeInfo_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentAggregatedSubtreeInfo_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentAggregatedSubtreeInfo_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentAggregatedSubtreeInfo_KKUBERNETES
        default:
            return EnvironmentAggregatedSubtreeInfo_KVMWARE
    }
}
