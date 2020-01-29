// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentRestorePointsForTimeRangeParamEnum enum
 */
type EnvironmentRestorePointsForTimeRangeParamEnum int

/**
 * Value collection for EnvironmentRestorePointsForTimeRangeParamEnum enum
 */
const (
    EnvironmentRestorePointsForTimeRangeParam_KVMWARE EnvironmentRestorePointsForTimeRangeParamEnum = 1 + iota
    EnvironmentRestorePointsForTimeRangeParam_KHYPERV
    EnvironmentRestorePointsForTimeRangeParam_KSQL
    EnvironmentRestorePointsForTimeRangeParam_KVIEW
    EnvironmentRestorePointsForTimeRangeParam_KPUPPETEER
    EnvironmentRestorePointsForTimeRangeParam_KPHYSICAL
    EnvironmentRestorePointsForTimeRangeParam_KPURE
    EnvironmentRestorePointsForTimeRangeParam_KAZURE
    EnvironmentRestorePointsForTimeRangeParam_KNETAPP
    EnvironmentRestorePointsForTimeRangeParam_KAGENT
    EnvironmentRestorePointsForTimeRangeParam_KGENERICNAS
    EnvironmentRestorePointsForTimeRangeParam_KACROPOLIS
    EnvironmentRestorePointsForTimeRangeParam_KPHYSICALFILES
    EnvironmentRestorePointsForTimeRangeParam_KISILON
    EnvironmentRestorePointsForTimeRangeParam_KGPFS
    EnvironmentRestorePointsForTimeRangeParam_KKVM
    EnvironmentRestorePointsForTimeRangeParam_KAWS
    EnvironmentRestorePointsForTimeRangeParam_KEXCHANGE
    EnvironmentRestorePointsForTimeRangeParam_KHYPERVVSS
    EnvironmentRestorePointsForTimeRangeParam_KORACLE
    EnvironmentRestorePointsForTimeRangeParam_KGCP
    EnvironmentRestorePointsForTimeRangeParam_KFLASHBLADE
    EnvironmentRestorePointsForTimeRangeParam_KAWSNATIVE
    EnvironmentRestorePointsForTimeRangeParam_KVCD
    EnvironmentRestorePointsForTimeRangeParam_KO365
    EnvironmentRestorePointsForTimeRangeParam_KO365OUTLOOK
    EnvironmentRestorePointsForTimeRangeParam_KHYPERFLEX
    EnvironmentRestorePointsForTimeRangeParam_KGCPNATIVE
    EnvironmentRestorePointsForTimeRangeParam_KAZURENATIVE
    EnvironmentRestorePointsForTimeRangeParam_KKUBERNETES
)

func (r EnvironmentRestorePointsForTimeRangeParamEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentRestorePointsForTimeRangeParamEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentRestorePointsForTimeRangeParamEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentRestorePointsForTimeRangeParamEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentRestorePointsForTimeRangeParamEnum to its string representation
 */
func EnvironmentRestorePointsForTimeRangeParamEnumToValue(environmentRestorePointsForTimeRangeParamEnum EnvironmentRestorePointsForTimeRangeParamEnum) string {
    switch environmentRestorePointsForTimeRangeParamEnum {
        case EnvironmentRestorePointsForTimeRangeParam_KVMWARE:
    		return "kVMware"
        case EnvironmentRestorePointsForTimeRangeParam_KHYPERV:
    		return "kHyperV"
        case EnvironmentRestorePointsForTimeRangeParam_KSQL:
    		return "kSQL"
        case EnvironmentRestorePointsForTimeRangeParam_KVIEW:
    		return "kView"
        case EnvironmentRestorePointsForTimeRangeParam_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentRestorePointsForTimeRangeParam_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentRestorePointsForTimeRangeParam_KPURE:
    		return "kPure"
        case EnvironmentRestorePointsForTimeRangeParam_KAZURE:
    		return "kAzure"
        case EnvironmentRestorePointsForTimeRangeParam_KNETAPP:
    		return "kNetapp"
        case EnvironmentRestorePointsForTimeRangeParam_KAGENT:
    		return "kAgent"
        case EnvironmentRestorePointsForTimeRangeParam_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentRestorePointsForTimeRangeParam_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentRestorePointsForTimeRangeParam_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentRestorePointsForTimeRangeParam_KISILON:
    		return "kIsilon"
        case EnvironmentRestorePointsForTimeRangeParam_KGPFS:
    		return "kGPFS"
        case EnvironmentRestorePointsForTimeRangeParam_KKVM:
    		return "kKVM"
        case EnvironmentRestorePointsForTimeRangeParam_KAWS:
    		return "kAWS"
        case EnvironmentRestorePointsForTimeRangeParam_KEXCHANGE:
    		return "kExchange"
        case EnvironmentRestorePointsForTimeRangeParam_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentRestorePointsForTimeRangeParam_KORACLE:
    		return "kOracle"
        case EnvironmentRestorePointsForTimeRangeParam_KGCP:
    		return "kGCP"
        case EnvironmentRestorePointsForTimeRangeParam_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentRestorePointsForTimeRangeParam_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentRestorePointsForTimeRangeParam_KVCD:
    		return "kVCD"
        case EnvironmentRestorePointsForTimeRangeParam_KO365:
    		return "kO365"
        case EnvironmentRestorePointsForTimeRangeParam_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentRestorePointsForTimeRangeParam_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentRestorePointsForTimeRangeParam_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentRestorePointsForTimeRangeParam_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentRestorePointsForTimeRangeParam_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentRestorePointsForTimeRangeParamEnum Array to its string Array representation
*/
func EnvironmentRestorePointsForTimeRangeParamEnumArrayToValue(environmentRestorePointsForTimeRangeParamEnum []EnvironmentRestorePointsForTimeRangeParamEnum) []string {
    convArray := make([]string,len( environmentRestorePointsForTimeRangeParamEnum))
    for i:=0; i<len(environmentRestorePointsForTimeRangeParamEnum);i++ {
        convArray[i] = EnvironmentRestorePointsForTimeRangeParamEnumToValue(environmentRestorePointsForTimeRangeParamEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentRestorePointsForTimeRangeParamEnumFromValue(value string) EnvironmentRestorePointsForTimeRangeParamEnum {
    switch value {
        case "kVMware":
            return EnvironmentRestorePointsForTimeRangeParam_KVMWARE
        case "kHyperV":
            return EnvironmentRestorePointsForTimeRangeParam_KHYPERV
        case "kSQL":
            return EnvironmentRestorePointsForTimeRangeParam_KSQL
        case "kView":
            return EnvironmentRestorePointsForTimeRangeParam_KVIEW
        case "kPuppeteer":
            return EnvironmentRestorePointsForTimeRangeParam_KPUPPETEER
        case "kPhysical":
            return EnvironmentRestorePointsForTimeRangeParam_KPHYSICAL
        case "kPure":
            return EnvironmentRestorePointsForTimeRangeParam_KPURE
        case "kAzure":
            return EnvironmentRestorePointsForTimeRangeParam_KAZURE
        case "kNetapp":
            return EnvironmentRestorePointsForTimeRangeParam_KNETAPP
        case "kAgent":
            return EnvironmentRestorePointsForTimeRangeParam_KAGENT
        case "kGenericNas":
            return EnvironmentRestorePointsForTimeRangeParam_KGENERICNAS
        case "kAcropolis":
            return EnvironmentRestorePointsForTimeRangeParam_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentRestorePointsForTimeRangeParam_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentRestorePointsForTimeRangeParam_KISILON
        case "kGPFS":
            return EnvironmentRestorePointsForTimeRangeParam_KGPFS
        case "kKVM":
            return EnvironmentRestorePointsForTimeRangeParam_KKVM
        case "kAWS":
            return EnvironmentRestorePointsForTimeRangeParam_KAWS
        case "kExchange":
            return EnvironmentRestorePointsForTimeRangeParam_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentRestorePointsForTimeRangeParam_KHYPERVVSS
        case "kOracle":
            return EnvironmentRestorePointsForTimeRangeParam_KORACLE
        case "kGCP":
            return EnvironmentRestorePointsForTimeRangeParam_KGCP
        case "kFlashBlade":
            return EnvironmentRestorePointsForTimeRangeParam_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentRestorePointsForTimeRangeParam_KAWSNATIVE
        case "kVCD":
            return EnvironmentRestorePointsForTimeRangeParam_KVCD
        case "kO365":
            return EnvironmentRestorePointsForTimeRangeParam_KO365
        case "kO365Outlook":
            return EnvironmentRestorePointsForTimeRangeParam_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentRestorePointsForTimeRangeParam_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentRestorePointsForTimeRangeParam_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentRestorePointsForTimeRangeParam_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentRestorePointsForTimeRangeParam_KKUBERNETES
        default:
            return EnvironmentRestorePointsForTimeRangeParam_KVMWARE
    }
}
