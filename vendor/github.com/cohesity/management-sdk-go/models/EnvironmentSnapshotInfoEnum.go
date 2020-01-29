// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentSnapshotInfoEnum enum
 */
type EnvironmentSnapshotInfoEnum int

/**
 * Value collection for EnvironmentSnapshotInfoEnum enum
 */
const (
    EnvironmentSnapshotInfo_KVMWARE EnvironmentSnapshotInfoEnum = 1 + iota
    EnvironmentSnapshotInfo_KHYPERV
    EnvironmentSnapshotInfo_KSQL
    EnvironmentSnapshotInfo_KVIEW
    EnvironmentSnapshotInfo_KPUPPETEER
    EnvironmentSnapshotInfo_KPHYSICAL
    EnvironmentSnapshotInfo_KPURE
    EnvironmentSnapshotInfo_KAZURE
    EnvironmentSnapshotInfo_KNETAPP
    EnvironmentSnapshotInfo_KAGENT
    EnvironmentSnapshotInfo_KGENERICNAS
    EnvironmentSnapshotInfo_KACROPOLIS
    EnvironmentSnapshotInfo_KPHYSICALFILES
    EnvironmentSnapshotInfo_KISILON
    EnvironmentSnapshotInfo_KGPFS
    EnvironmentSnapshotInfo_KKVM
    EnvironmentSnapshotInfo_KAWS
    EnvironmentSnapshotInfo_KEXCHANGE
    EnvironmentSnapshotInfo_KHYPERVVSS
    EnvironmentSnapshotInfo_KORACLE
    EnvironmentSnapshotInfo_KGCP
    EnvironmentSnapshotInfo_KFLASHBLADE
    EnvironmentSnapshotInfo_KAWSNATIVE
    EnvironmentSnapshotInfo_KVCD
    EnvironmentSnapshotInfo_KO365
    EnvironmentSnapshotInfo_KO365OUTLOOK
    EnvironmentSnapshotInfo_KHYPERFLEX
    EnvironmentSnapshotInfo_KGCPNATIVE
    EnvironmentSnapshotInfo_KAZURENATIVE
    EnvironmentSnapshotInfo_KKUBERNETES
)

func (r EnvironmentSnapshotInfoEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentSnapshotInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentSnapshotInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentSnapshotInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentSnapshotInfoEnum to its string representation
 */
func EnvironmentSnapshotInfoEnumToValue(environmentSnapshotInfoEnum EnvironmentSnapshotInfoEnum) string {
    switch environmentSnapshotInfoEnum {
        case EnvironmentSnapshotInfo_KVMWARE:
    		return "kVMware"
        case EnvironmentSnapshotInfo_KHYPERV:
    		return "kHyperV"
        case EnvironmentSnapshotInfo_KSQL:
    		return "kSQL"
        case EnvironmentSnapshotInfo_KVIEW:
    		return "kView"
        case EnvironmentSnapshotInfo_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentSnapshotInfo_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentSnapshotInfo_KPURE:
    		return "kPure"
        case EnvironmentSnapshotInfo_KAZURE:
    		return "kAzure"
        case EnvironmentSnapshotInfo_KNETAPP:
    		return "kNetapp"
        case EnvironmentSnapshotInfo_KAGENT:
    		return "kAgent"
        case EnvironmentSnapshotInfo_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentSnapshotInfo_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentSnapshotInfo_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentSnapshotInfo_KISILON:
    		return "kIsilon"
        case EnvironmentSnapshotInfo_KGPFS:
    		return "kGPFS"
        case EnvironmentSnapshotInfo_KKVM:
    		return "kKVM"
        case EnvironmentSnapshotInfo_KAWS:
    		return "kAWS"
        case EnvironmentSnapshotInfo_KEXCHANGE:
    		return "kExchange"
        case EnvironmentSnapshotInfo_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentSnapshotInfo_KORACLE:
    		return "kOracle"
        case EnvironmentSnapshotInfo_KGCP:
    		return "kGCP"
        case EnvironmentSnapshotInfo_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentSnapshotInfo_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentSnapshotInfo_KVCD:
    		return "kVCD"
        case EnvironmentSnapshotInfo_KO365:
    		return "kO365"
        case EnvironmentSnapshotInfo_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentSnapshotInfo_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentSnapshotInfo_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentSnapshotInfo_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentSnapshotInfo_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentSnapshotInfoEnum Array to its string Array representation
*/
func EnvironmentSnapshotInfoEnumArrayToValue(environmentSnapshotInfoEnum []EnvironmentSnapshotInfoEnum) []string {
    convArray := make([]string,len( environmentSnapshotInfoEnum))
    for i:=0; i<len(environmentSnapshotInfoEnum);i++ {
        convArray[i] = EnvironmentSnapshotInfoEnumToValue(environmentSnapshotInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentSnapshotInfoEnumFromValue(value string) EnvironmentSnapshotInfoEnum {
    switch value {
        case "kVMware":
            return EnvironmentSnapshotInfo_KVMWARE
        case "kHyperV":
            return EnvironmentSnapshotInfo_KHYPERV
        case "kSQL":
            return EnvironmentSnapshotInfo_KSQL
        case "kView":
            return EnvironmentSnapshotInfo_KVIEW
        case "kPuppeteer":
            return EnvironmentSnapshotInfo_KPUPPETEER
        case "kPhysical":
            return EnvironmentSnapshotInfo_KPHYSICAL
        case "kPure":
            return EnvironmentSnapshotInfo_KPURE
        case "kAzure":
            return EnvironmentSnapshotInfo_KAZURE
        case "kNetapp":
            return EnvironmentSnapshotInfo_KNETAPP
        case "kAgent":
            return EnvironmentSnapshotInfo_KAGENT
        case "kGenericNas":
            return EnvironmentSnapshotInfo_KGENERICNAS
        case "kAcropolis":
            return EnvironmentSnapshotInfo_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentSnapshotInfo_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentSnapshotInfo_KISILON
        case "kGPFS":
            return EnvironmentSnapshotInfo_KGPFS
        case "kKVM":
            return EnvironmentSnapshotInfo_KKVM
        case "kAWS":
            return EnvironmentSnapshotInfo_KAWS
        case "kExchange":
            return EnvironmentSnapshotInfo_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentSnapshotInfo_KHYPERVVSS
        case "kOracle":
            return EnvironmentSnapshotInfo_KORACLE
        case "kGCP":
            return EnvironmentSnapshotInfo_KGCP
        case "kFlashBlade":
            return EnvironmentSnapshotInfo_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentSnapshotInfo_KAWSNATIVE
        case "kVCD":
            return EnvironmentSnapshotInfo_KVCD
        case "kO365":
            return EnvironmentSnapshotInfo_KO365
        case "kO365Outlook":
            return EnvironmentSnapshotInfo_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentSnapshotInfo_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentSnapshotInfo_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentSnapshotInfo_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentSnapshotInfo_KKUBERNETES
        default:
            return EnvironmentSnapshotInfo_KVMWARE
    }
}
