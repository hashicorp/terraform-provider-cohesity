// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ExcludeTypeGetProtectedObjectsSummaryEnum enum
 */
type ExcludeTypeGetProtectedObjectsSummaryEnum int

/**
 * Value collection for ExcludeTypeGetProtectedObjectsSummaryEnum enum
 */
const (
    ExcludeTypeGetProtectedObjectsSummary_KVMWARE ExcludeTypeGetProtectedObjectsSummaryEnum = 1 + iota
    ExcludeTypeGetProtectedObjectsSummary_KHYPERV
    ExcludeTypeGetProtectedObjectsSummary_KSQL
    ExcludeTypeGetProtectedObjectsSummary_KVIEW
    ExcludeTypeGetProtectedObjectsSummary_KPUPPETEER
    ExcludeTypeGetProtectedObjectsSummary_KPHYSICAL
    ExcludeTypeGetProtectedObjectsSummary_KPURE
    ExcludeTypeGetProtectedObjectsSummary_KAZURE
    ExcludeTypeGetProtectedObjectsSummary_KNETAPP
    ExcludeTypeGetProtectedObjectsSummary_KAGENT
    ExcludeTypeGetProtectedObjectsSummary_KGENERICNAS
    ExcludeTypeGetProtectedObjectsSummary_KACROPOLIS
    ExcludeTypeGetProtectedObjectsSummary_KPHYSICALFILES
    ExcludeTypeGetProtectedObjectsSummary_KISILON
    ExcludeTypeGetProtectedObjectsSummary_KGPFS
    ExcludeTypeGetProtectedObjectsSummary_KKVM
    ExcludeTypeGetProtectedObjectsSummary_KAWS
    ExcludeTypeGetProtectedObjectsSummary_KEXCHANGE
    ExcludeTypeGetProtectedObjectsSummary_KHYPERVVSS
    ExcludeTypeGetProtectedObjectsSummary_KORACLE
    ExcludeTypeGetProtectedObjectsSummary_KGCP
    ExcludeTypeGetProtectedObjectsSummary_KFLASHBLADE
    ExcludeTypeGetProtectedObjectsSummary_KAWSNATIVE
    ExcludeTypeGetProtectedObjectsSummary_KVCD
    ExcludeTypeGetProtectedObjectsSummary_KO365
    ExcludeTypeGetProtectedObjectsSummary_KO365OUTLOOK
    ExcludeTypeGetProtectedObjectsSummary_KHYPERFLEX
    ExcludeTypeGetProtectedObjectsSummary_KGCPNATIVE
    ExcludeTypeGetProtectedObjectsSummary_KAZURENATIVE
    ExcludeTypeGetProtectedObjectsSummary_KAD
    ExcludeTypeGetProtectedObjectsSummary_KAWSSNAPSHOTMANAGER
)

func (r ExcludeTypeGetProtectedObjectsSummaryEnum) MarshalJSON() ([]byte, error) {
    s := ExcludeTypeGetProtectedObjectsSummaryEnumToValue(r)
    return json.Marshal(s)
}

func (r *ExcludeTypeGetProtectedObjectsSummaryEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ExcludeTypeGetProtectedObjectsSummaryEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ExcludeTypeGetProtectedObjectsSummaryEnum to its string representation
 */
func ExcludeTypeGetProtectedObjectsSummaryEnumToValue(excludeTypeGetProtectedObjectsSummaryEnum ExcludeTypeGetProtectedObjectsSummaryEnum) string {
    switch excludeTypeGetProtectedObjectsSummaryEnum {
        case ExcludeTypeGetProtectedObjectsSummary_KVMWARE:
    		return "kVMware"
        case ExcludeTypeGetProtectedObjectsSummary_KHYPERV:
    		return "kHyperV"
        case ExcludeTypeGetProtectedObjectsSummary_KSQL:
    		return "kSQL"
        case ExcludeTypeGetProtectedObjectsSummary_KVIEW:
    		return "kView"
        case ExcludeTypeGetProtectedObjectsSummary_KPUPPETEER:
    		return "kPuppeteer"
        case ExcludeTypeGetProtectedObjectsSummary_KPHYSICAL:
    		return "kPhysical"
        case ExcludeTypeGetProtectedObjectsSummary_KPURE:
    		return "kPure"
        case ExcludeTypeGetProtectedObjectsSummary_KAZURE:
    		return "kAzure"
        case ExcludeTypeGetProtectedObjectsSummary_KNETAPP:
    		return "kNetapp"
        case ExcludeTypeGetProtectedObjectsSummary_KAGENT:
    		return "kAgent"
        case ExcludeTypeGetProtectedObjectsSummary_KGENERICNAS:
    		return "kGenericNas"
        case ExcludeTypeGetProtectedObjectsSummary_KACROPOLIS:
    		return "kAcropolis"
        case ExcludeTypeGetProtectedObjectsSummary_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case ExcludeTypeGetProtectedObjectsSummary_KISILON:
    		return "kIsilon"
        case ExcludeTypeGetProtectedObjectsSummary_KGPFS:
    		return "kGPFS"
        case ExcludeTypeGetProtectedObjectsSummary_KKVM:
    		return "kKVM"
        case ExcludeTypeGetProtectedObjectsSummary_KAWS:
    		return "kAWS"
        case ExcludeTypeGetProtectedObjectsSummary_KEXCHANGE:
    		return "kExchange"
        case ExcludeTypeGetProtectedObjectsSummary_KHYPERVVSS:
    		return "kHyperVVSS"
        case ExcludeTypeGetProtectedObjectsSummary_KORACLE:
    		return "kOracle"
        case ExcludeTypeGetProtectedObjectsSummary_KGCP:
    		return "kGCP"
        case ExcludeTypeGetProtectedObjectsSummary_KFLASHBLADE:
    		return "kFlashBlade"
        case ExcludeTypeGetProtectedObjectsSummary_KAWSNATIVE:
    		return "kAWSNative"
        case ExcludeTypeGetProtectedObjectsSummary_KVCD:
    		return "kVCD"
        case ExcludeTypeGetProtectedObjectsSummary_KO365:
    		return "kO365"
        case ExcludeTypeGetProtectedObjectsSummary_KO365OUTLOOK:
    		return "kO365Outlook"
        case ExcludeTypeGetProtectedObjectsSummary_KHYPERFLEX:
    		return "kHyperFlex"
        case ExcludeTypeGetProtectedObjectsSummary_KGCPNATIVE:
    		return "kGCPNative"
        case ExcludeTypeGetProtectedObjectsSummary_KAZURENATIVE:
    		return "kAzureNative"
        case ExcludeTypeGetProtectedObjectsSummary_KAD:
    		return "kAD"
        case ExcludeTypeGetProtectedObjectsSummary_KAWSSNAPSHOTMANAGER:
    		return "kAWSSnapshotManager"
        default:
        	return "kVMware"
    }
}

/**
 * Converts ExcludeTypeGetProtectedObjectsSummaryEnum Array to its string Array representation
*/
func ExcludeTypeGetProtectedObjectsSummaryEnumArrayToValue(excludeTypeGetProtectedObjectsSummaryEnum []ExcludeTypeGetProtectedObjectsSummaryEnum) []string {
    convArray := make([]string,len( excludeTypeGetProtectedObjectsSummaryEnum))
    for i:=0; i<len(excludeTypeGetProtectedObjectsSummaryEnum);i++ {
        convArray[i] = ExcludeTypeGetProtectedObjectsSummaryEnumToValue(excludeTypeGetProtectedObjectsSummaryEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ExcludeTypeGetProtectedObjectsSummaryEnumFromValue(value string) ExcludeTypeGetProtectedObjectsSummaryEnum {
    switch value {
        case "kVMware":
            return ExcludeTypeGetProtectedObjectsSummary_KVMWARE
        case "kHyperV":
            return ExcludeTypeGetProtectedObjectsSummary_KHYPERV
        case "kSQL":
            return ExcludeTypeGetProtectedObjectsSummary_KSQL
        case "kView":
            return ExcludeTypeGetProtectedObjectsSummary_KVIEW
        case "kPuppeteer":
            return ExcludeTypeGetProtectedObjectsSummary_KPUPPETEER
        case "kPhysical":
            return ExcludeTypeGetProtectedObjectsSummary_KPHYSICAL
        case "kPure":
            return ExcludeTypeGetProtectedObjectsSummary_KPURE
        case "kAzure":
            return ExcludeTypeGetProtectedObjectsSummary_KAZURE
        case "kNetapp":
            return ExcludeTypeGetProtectedObjectsSummary_KNETAPP
        case "kAgent":
            return ExcludeTypeGetProtectedObjectsSummary_KAGENT
        case "kGenericNas":
            return ExcludeTypeGetProtectedObjectsSummary_KGENERICNAS
        case "kAcropolis":
            return ExcludeTypeGetProtectedObjectsSummary_KACROPOLIS
        case "kPhysicalFiles":
            return ExcludeTypeGetProtectedObjectsSummary_KPHYSICALFILES
        case "kIsilon":
            return ExcludeTypeGetProtectedObjectsSummary_KISILON
        case "kGPFS":
            return ExcludeTypeGetProtectedObjectsSummary_KGPFS
        case "kKVM":
            return ExcludeTypeGetProtectedObjectsSummary_KKVM
        case "kAWS":
            return ExcludeTypeGetProtectedObjectsSummary_KAWS
        case "kExchange":
            return ExcludeTypeGetProtectedObjectsSummary_KEXCHANGE
        case "kHyperVVSS":
            return ExcludeTypeGetProtectedObjectsSummary_KHYPERVVSS
        case "kOracle":
            return ExcludeTypeGetProtectedObjectsSummary_KORACLE
        case "kGCP":
            return ExcludeTypeGetProtectedObjectsSummary_KGCP
        case "kFlashBlade":
            return ExcludeTypeGetProtectedObjectsSummary_KFLASHBLADE
        case "kAWSNative":
            return ExcludeTypeGetProtectedObjectsSummary_KAWSNATIVE
        case "kVCD":
            return ExcludeTypeGetProtectedObjectsSummary_KVCD
        case "kO365":
            return ExcludeTypeGetProtectedObjectsSummary_KO365
        case "kO365Outlook":
            return ExcludeTypeGetProtectedObjectsSummary_KO365OUTLOOK
        case "kHyperFlex":
            return ExcludeTypeGetProtectedObjectsSummary_KHYPERFLEX
        case "kGCPNative":
            return ExcludeTypeGetProtectedObjectsSummary_KGCPNATIVE
        case "kAzureNative":
            return ExcludeTypeGetProtectedObjectsSummary_KAZURENATIVE
        case "kAD":
            return ExcludeTypeGetProtectedObjectsSummary_KAD
        case "kAWSSnapshotManager":
            return ExcludeTypeGetProtectedObjectsSummary_KAWSSNAPSHOTMANAGER
        default:
            return ExcludeTypeGetProtectedObjectsSummary_KVMWARE
    }
}
