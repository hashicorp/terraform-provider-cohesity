// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentSearchObjectsEnum enum
 */
type EnvironmentSearchObjectsEnum int

/**
 * Value collection for EnvironmentSearchObjectsEnum enum
 */
const (
    EnvironmentSearchObjects_KVMWARE EnvironmentSearchObjectsEnum = 1 + iota
    EnvironmentSearchObjects_KHYPERV
    EnvironmentSearchObjects_KSQL
    EnvironmentSearchObjects_KVIEW
    EnvironmentSearchObjects_KPUPPETEER
    EnvironmentSearchObjects_KPHYSICAL
    EnvironmentSearchObjects_KPURE
    EnvironmentSearchObjects_KAZURE
    EnvironmentSearchObjects_KNETAPP
    EnvironmentSearchObjects_KAGENT
    EnvironmentSearchObjects_KGENERICNAS
    EnvironmentSearchObjects_KACROPOLIS
    EnvironmentSearchObjects_KPHYSICALFILES
    EnvironmentSearchObjects_KISILON
    EnvironmentSearchObjects_KGPFS
    EnvironmentSearchObjects_KKVM
    EnvironmentSearchObjects_KAWS
    EnvironmentSearchObjects_KEXCHANGE
    EnvironmentSearchObjects_KHYPERVVSS
    EnvironmentSearchObjects_KORACLE
    EnvironmentSearchObjects_KGCP
    EnvironmentSearchObjects_KFLASHBLADE
    EnvironmentSearchObjects_KAWSNATIVE
    EnvironmentSearchObjects_KVCD
    EnvironmentSearchObjects_KO365
    EnvironmentSearchObjects_KO365OUTLOOK
    EnvironmentSearchObjects_KHYPERFLEX
    EnvironmentSearchObjects_KGCPNATIVE
    EnvironmentSearchObjects_KKUBERNETES
)

func (r EnvironmentSearchObjectsEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentSearchObjectsEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentSearchObjectsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentSearchObjectsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentSearchObjectsEnum to its string representation
 */
func EnvironmentSearchObjectsEnumToValue(environmentSearchObjectsEnum EnvironmentSearchObjectsEnum) string {
    switch environmentSearchObjectsEnum {
        case EnvironmentSearchObjects_KVMWARE:
    		return "kVMware"
        case EnvironmentSearchObjects_KHYPERV:
    		return "kHyperV"
        case EnvironmentSearchObjects_KSQL:
    		return "kSQL"
        case EnvironmentSearchObjects_KVIEW:
    		return "kView"
        case EnvironmentSearchObjects_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentSearchObjects_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentSearchObjects_KPURE:
    		return "kPure"
        case EnvironmentSearchObjects_KAZURE:
    		return "kAzure"
        case EnvironmentSearchObjects_KNETAPP:
    		return "kNetapp"
        case EnvironmentSearchObjects_KAGENT:
    		return "kAgent"
        case EnvironmentSearchObjects_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentSearchObjects_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentSearchObjects_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentSearchObjects_KISILON:
    		return "kIsilon"
        case EnvironmentSearchObjects_KGPFS:
    		return "kGPFS"
        case EnvironmentSearchObjects_KKVM:
    		return "kKVM"
        case EnvironmentSearchObjects_KAWS:
    		return "kAWS"
        case EnvironmentSearchObjects_KEXCHANGE:
    		return "kExchange"
        case EnvironmentSearchObjects_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentSearchObjects_KORACLE:
    		return "kOracle"
        case EnvironmentSearchObjects_KGCP:
    		return "kGCP"
        case EnvironmentSearchObjects_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentSearchObjects_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentSearchObjects_KVCD:
    		return "kVCD"
        case EnvironmentSearchObjects_KO365:
    		return "kO365"
        case EnvironmentSearchObjects_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentSearchObjects_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentSearchObjects_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentSearchObjects_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentSearchObjectsEnum Array to its string Array representation
*/
func EnvironmentSearchObjectsEnumArrayToValue(environmentSearchObjectsEnum []EnvironmentSearchObjectsEnum) []string {
    convArray := make([]string,len( environmentSearchObjectsEnum))
    for i:=0; i<len(environmentSearchObjectsEnum);i++ {
        convArray[i] = EnvironmentSearchObjectsEnumToValue(environmentSearchObjectsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentSearchObjectsEnumFromValue(value string) EnvironmentSearchObjectsEnum {
    switch value {
        case "kVMware":
            return EnvironmentSearchObjects_KVMWARE
        case "kHyperV":
            return EnvironmentSearchObjects_KHYPERV
        case "kSQL":
            return EnvironmentSearchObjects_KSQL
        case "kView":
            return EnvironmentSearchObjects_KVIEW
        case "kPuppeteer":
            return EnvironmentSearchObjects_KPUPPETEER
        case "kPhysical":
            return EnvironmentSearchObjects_KPHYSICAL
        case "kPure":
            return EnvironmentSearchObjects_KPURE
        case "kAzure":
            return EnvironmentSearchObjects_KAZURE
        case "kNetapp":
            return EnvironmentSearchObjects_KNETAPP
        case "kAgent":
            return EnvironmentSearchObjects_KAGENT
        case "kGenericNas":
            return EnvironmentSearchObjects_KGENERICNAS
        case "kAcropolis":
            return EnvironmentSearchObjects_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentSearchObjects_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentSearchObjects_KISILON
        case "kGPFS":
            return EnvironmentSearchObjects_KGPFS
        case "kKVM":
            return EnvironmentSearchObjects_KKVM
        case "kAWS":
            return EnvironmentSearchObjects_KAWS
        case "kExchange":
            return EnvironmentSearchObjects_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentSearchObjects_KHYPERVVSS
        case "kOracle":
            return EnvironmentSearchObjects_KORACLE
        case "kGCP":
            return EnvironmentSearchObjects_KGCP
        case "kFlashBlade":
            return EnvironmentSearchObjects_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentSearchObjects_KAWSNATIVE
        case "kVCD":
            return EnvironmentSearchObjects_KVCD
        case "kO365":
            return EnvironmentSearchObjects_KO365
        case "kO365Outlook":
            return EnvironmentSearchObjects_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentSearchObjects_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentSearchObjects_KGCPNATIVE
        case "kKubernetes":
            return EnvironmentSearchObjects_KKUBERNETES
        default:
            return EnvironmentSearchObjects_KVMWARE
    }
}
