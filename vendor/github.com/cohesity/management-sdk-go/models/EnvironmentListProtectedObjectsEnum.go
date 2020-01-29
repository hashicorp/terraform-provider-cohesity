// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentListProtectedObjectsEnum enum
 */
type EnvironmentListProtectedObjectsEnum int

/**
 * Value collection for EnvironmentListProtectedObjectsEnum enum
 */
const (
    EnvironmentListProtectedObjects_KVMWARE EnvironmentListProtectedObjectsEnum = 1 + iota
    EnvironmentListProtectedObjects_KHYPERV
    EnvironmentListProtectedObjects_KSQL
    EnvironmentListProtectedObjects_KVIEW
    EnvironmentListProtectedObjects_KPUPPETEER
    EnvironmentListProtectedObjects_KPHYSICAL
    EnvironmentListProtectedObjects_KPURE
    EnvironmentListProtectedObjects_KAZURE
    EnvironmentListProtectedObjects_KNETAPP
    EnvironmentListProtectedObjects_KAGENT
    EnvironmentListProtectedObjects_KGENERICNAS
    EnvironmentListProtectedObjects_KACROPOLIS
    EnvironmentListProtectedObjects_KPHYSICALFILES
    EnvironmentListProtectedObjects_KISILON
    EnvironmentListProtectedObjects_KGPFS
    EnvironmentListProtectedObjects_KKVM
    EnvironmentListProtectedObjects_KAWS
    EnvironmentListProtectedObjects_KEXCHANGE
    EnvironmentListProtectedObjects_KHYPERVVSS
    EnvironmentListProtectedObjects_KORACLE
    EnvironmentListProtectedObjects_KGCP
    EnvironmentListProtectedObjects_KFLASHBLADE
    EnvironmentListProtectedObjects_KAWSNATIVE
    EnvironmentListProtectedObjects_KVCD
    EnvironmentListProtectedObjects_KO365
    EnvironmentListProtectedObjects_KO365OUTLOOK
    EnvironmentListProtectedObjects_KHYPERFLEX
    EnvironmentListProtectedObjects_KGCPNATIVE
    EnvironmentListProtectedObjects_KAZURENATIVE
    EnvironmentListProtectedObjects_KKUBERNETES
)

func (r EnvironmentListProtectedObjectsEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentListProtectedObjectsEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentListProtectedObjectsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentListProtectedObjectsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentListProtectedObjectsEnum to its string representation
 */
func EnvironmentListProtectedObjectsEnumToValue(environmentListProtectedObjectsEnum EnvironmentListProtectedObjectsEnum) string {
    switch environmentListProtectedObjectsEnum {
        case EnvironmentListProtectedObjects_KVMWARE:
    		return "kVMware"
        case EnvironmentListProtectedObjects_KHYPERV:
    		return "kHyperV"
        case EnvironmentListProtectedObjects_KSQL:
    		return "kSQL"
        case EnvironmentListProtectedObjects_KVIEW:
    		return "kView"
        case EnvironmentListProtectedObjects_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentListProtectedObjects_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentListProtectedObjects_KPURE:
    		return "kPure"
        case EnvironmentListProtectedObjects_KAZURE:
    		return "kAzure"
        case EnvironmentListProtectedObjects_KNETAPP:
    		return "kNetapp"
        case EnvironmentListProtectedObjects_KAGENT:
    		return "kAgent"
        case EnvironmentListProtectedObjects_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentListProtectedObjects_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentListProtectedObjects_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentListProtectedObjects_KISILON:
    		return "kIsilon"
        case EnvironmentListProtectedObjects_KGPFS:
    		return "kGPFS"
        case EnvironmentListProtectedObjects_KKVM:
    		return "kKVM"
        case EnvironmentListProtectedObjects_KAWS:
    		return "kAWS"
        case EnvironmentListProtectedObjects_KEXCHANGE:
    		return "kExchange"
        case EnvironmentListProtectedObjects_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentListProtectedObjects_KORACLE:
    		return "kOracle"
        case EnvironmentListProtectedObjects_KGCP:
    		return "kGCP"
        case EnvironmentListProtectedObjects_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentListProtectedObjects_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentListProtectedObjects_KVCD:
    		return "kVCD"
        case EnvironmentListProtectedObjects_KO365:
    		return "kO365"
        case EnvironmentListProtectedObjects_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentListProtectedObjects_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentListProtectedObjects_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentListProtectedObjects_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentListProtectedObjects_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentListProtectedObjectsEnum Array to its string Array representation
*/
func EnvironmentListProtectedObjectsEnumArrayToValue(environmentListProtectedObjectsEnum []EnvironmentListProtectedObjectsEnum) []string {
    convArray := make([]string,len( environmentListProtectedObjectsEnum))
    for i:=0; i<len(environmentListProtectedObjectsEnum);i++ {
        convArray[i] = EnvironmentListProtectedObjectsEnumToValue(environmentListProtectedObjectsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentListProtectedObjectsEnumFromValue(value string) EnvironmentListProtectedObjectsEnum {
    switch value {
        case "kVMware":
            return EnvironmentListProtectedObjects_KVMWARE
        case "kHyperV":
            return EnvironmentListProtectedObjects_KHYPERV
        case "kSQL":
            return EnvironmentListProtectedObjects_KSQL
        case "kView":
            return EnvironmentListProtectedObjects_KVIEW
        case "kPuppeteer":
            return EnvironmentListProtectedObjects_KPUPPETEER
        case "kPhysical":
            return EnvironmentListProtectedObjects_KPHYSICAL
        case "kPure":
            return EnvironmentListProtectedObjects_KPURE
        case "kAzure":
            return EnvironmentListProtectedObjects_KAZURE
        case "kNetapp":
            return EnvironmentListProtectedObjects_KNETAPP
        case "kAgent":
            return EnvironmentListProtectedObjects_KAGENT
        case "kGenericNas":
            return EnvironmentListProtectedObjects_KGENERICNAS
        case "kAcropolis":
            return EnvironmentListProtectedObjects_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentListProtectedObjects_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentListProtectedObjects_KISILON
        case "kGPFS":
            return EnvironmentListProtectedObjects_KGPFS
        case "kKVM":
            return EnvironmentListProtectedObjects_KKVM
        case "kAWS":
            return EnvironmentListProtectedObjects_KAWS
        case "kExchange":
            return EnvironmentListProtectedObjects_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentListProtectedObjects_KHYPERVVSS
        case "kOracle":
            return EnvironmentListProtectedObjects_KORACLE
        case "kGCP":
            return EnvironmentListProtectedObjects_KGCP
        case "kFlashBlade":
            return EnvironmentListProtectedObjects_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentListProtectedObjects_KAWSNATIVE
        case "kVCD":
            return EnvironmentListProtectedObjects_KVCD
        case "kO365":
            return EnvironmentListProtectedObjects_KO365
        case "kO365Outlook":
            return EnvironmentListProtectedObjects_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentListProtectedObjects_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentListProtectedObjects_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentListProtectedObjects_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentListProtectedObjects_KKUBERNETES
        default:
            return EnvironmentListProtectedObjects_KVMWARE
    }
}
