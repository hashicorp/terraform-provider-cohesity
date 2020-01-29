// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentListApplicationServersEnum enum
 */
type EnvironmentListApplicationServersEnum int

/**
 * Value collection for EnvironmentListApplicationServersEnum enum
 */
const (
    EnvironmentListApplicationServers_KVMWARE EnvironmentListApplicationServersEnum = 1 + iota
    EnvironmentListApplicationServers_KHYPERV
    EnvironmentListApplicationServers_KSQL
    EnvironmentListApplicationServers_KVIEW
    EnvironmentListApplicationServers_KPUPPETEER
    EnvironmentListApplicationServers_KPHYSICAL
    EnvironmentListApplicationServers_KPURE
    EnvironmentListApplicationServers_KAZURE
    EnvironmentListApplicationServers_KNETAPP
    EnvironmentListApplicationServers_KAGENT
    EnvironmentListApplicationServers_KGENERICNAS
    EnvironmentListApplicationServers_KACROPOLIS
    EnvironmentListApplicationServers_KPHYSICALFILES
    EnvironmentListApplicationServers_KISILON
    EnvironmentListApplicationServers_KGPFS
    EnvironmentListApplicationServers_KKVM
    EnvironmentListApplicationServers_KAWS
    EnvironmentListApplicationServers_KEXCHANGE
    EnvironmentListApplicationServers_KHYPERVVSS
    EnvironmentListApplicationServers_KORACLE
    EnvironmentListApplicationServers_KGCP
    EnvironmentListApplicationServers_KFLASHBLADE
    EnvironmentListApplicationServers_KAWSNATIVE
    EnvironmentListApplicationServers_KVCD
    EnvironmentListApplicationServers_KO365
    EnvironmentListApplicationServers_KO365OUTLOOK
    EnvironmentListApplicationServers_KHYPERFLEX
    EnvironmentListApplicationServers_KGCPNATIVE
    EnvironmentListApplicationServers_KAZURENATIVE
    EnvironmentListApplicationServers_KKUBERNETES
)

func (r EnvironmentListApplicationServersEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentListApplicationServersEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentListApplicationServersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentListApplicationServersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentListApplicationServersEnum to its string representation
 */
func EnvironmentListApplicationServersEnumToValue(environmentListApplicationServersEnum EnvironmentListApplicationServersEnum) string {
    switch environmentListApplicationServersEnum {
        case EnvironmentListApplicationServers_KVMWARE:
    		return "kVMware"
        case EnvironmentListApplicationServers_KHYPERV:
    		return "kHyperV"
        case EnvironmentListApplicationServers_KSQL:
    		return "kSQL"
        case EnvironmentListApplicationServers_KVIEW:
    		return "kView"
        case EnvironmentListApplicationServers_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentListApplicationServers_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentListApplicationServers_KPURE:
    		return "kPure"
        case EnvironmentListApplicationServers_KAZURE:
    		return "kAzure"
        case EnvironmentListApplicationServers_KNETAPP:
    		return "kNetapp"
        case EnvironmentListApplicationServers_KAGENT:
    		return "kAgent"
        case EnvironmentListApplicationServers_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentListApplicationServers_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentListApplicationServers_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentListApplicationServers_KISILON:
    		return "kIsilon"
        case EnvironmentListApplicationServers_KGPFS:
    		return "kGPFS"
        case EnvironmentListApplicationServers_KKVM:
    		return "kKVM"
        case EnvironmentListApplicationServers_KAWS:
    		return "kAWS"
        case EnvironmentListApplicationServers_KEXCHANGE:
    		return "kExchange"
        case EnvironmentListApplicationServers_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentListApplicationServers_KORACLE:
    		return "kOracle"
        case EnvironmentListApplicationServers_KGCP:
    		return "kGCP"
        case EnvironmentListApplicationServers_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentListApplicationServers_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentListApplicationServers_KVCD:
    		return "kVCD"
        case EnvironmentListApplicationServers_KO365:
    		return "kO365"
        case EnvironmentListApplicationServers_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentListApplicationServers_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentListApplicationServers_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentListApplicationServers_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentListApplicationServers_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentListApplicationServersEnum Array to its string Array representation
*/
func EnvironmentListApplicationServersEnumArrayToValue(environmentListApplicationServersEnum []EnvironmentListApplicationServersEnum) []string {
    convArray := make([]string,len( environmentListApplicationServersEnum))
    for i:=0; i<len(environmentListApplicationServersEnum);i++ {
        convArray[i] = EnvironmentListApplicationServersEnumToValue(environmentListApplicationServersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentListApplicationServersEnumFromValue(value string) EnvironmentListApplicationServersEnum {
    switch value {
        case "kVMware":
            return EnvironmentListApplicationServers_KVMWARE
        case "kHyperV":
            return EnvironmentListApplicationServers_KHYPERV
        case "kSQL":
            return EnvironmentListApplicationServers_KSQL
        case "kView":
            return EnvironmentListApplicationServers_KVIEW
        case "kPuppeteer":
            return EnvironmentListApplicationServers_KPUPPETEER
        case "kPhysical":
            return EnvironmentListApplicationServers_KPHYSICAL
        case "kPure":
            return EnvironmentListApplicationServers_KPURE
        case "kAzure":
            return EnvironmentListApplicationServers_KAZURE
        case "kNetapp":
            return EnvironmentListApplicationServers_KNETAPP
        case "kAgent":
            return EnvironmentListApplicationServers_KAGENT
        case "kGenericNas":
            return EnvironmentListApplicationServers_KGENERICNAS
        case "kAcropolis":
            return EnvironmentListApplicationServers_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentListApplicationServers_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentListApplicationServers_KISILON
        case "kGPFS":
            return EnvironmentListApplicationServers_KGPFS
        case "kKVM":
            return EnvironmentListApplicationServers_KKVM
        case "kAWS":
            return EnvironmentListApplicationServers_KAWS
        case "kExchange":
            return EnvironmentListApplicationServers_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentListApplicationServers_KHYPERVVSS
        case "kOracle":
            return EnvironmentListApplicationServers_KORACLE
        case "kGCP":
            return EnvironmentListApplicationServers_KGCP
        case "kFlashBlade":
            return EnvironmentListApplicationServers_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentListApplicationServers_KAWSNATIVE
        case "kVCD":
            return EnvironmentListApplicationServers_KVCD
        case "kO365":
            return EnvironmentListApplicationServers_KO365
        case "kO365Outlook":
            return EnvironmentListApplicationServers_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentListApplicationServers_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentListApplicationServers_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentListApplicationServers_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentListApplicationServers_KKUBERNETES
        default:
            return EnvironmentListApplicationServers_KVMWARE
    }
}
