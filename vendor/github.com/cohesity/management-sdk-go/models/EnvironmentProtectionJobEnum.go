// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentProtectionJobEnum enum
 */
type EnvironmentProtectionJobEnum int

/**
 * Value collection for EnvironmentProtectionJobEnum enum
 */
const (
    EnvironmentProtectionJob_KVMWARE EnvironmentProtectionJobEnum = 1 + iota
    EnvironmentProtectionJob_KHYPERV
    EnvironmentProtectionJob_KSQL
    EnvironmentProtectionJob_KVIEW
    EnvironmentProtectionJob_KPUPPETEER
    EnvironmentProtectionJob_KPHYSICAL
    EnvironmentProtectionJob_KPURE
    EnvironmentProtectionJob_KAZURE
    EnvironmentProtectionJob_KNETAPP
    EnvironmentProtectionJob_KAGENT
    EnvironmentProtectionJob_KGENERICNAS
    EnvironmentProtectionJob_KACROPOLIS
    EnvironmentProtectionJob_KPHYSICALFILES
    EnvironmentProtectionJob_KISILON
    EnvironmentProtectionJob_KGPFS
    EnvironmentProtectionJob_KKVM
    EnvironmentProtectionJob_KAWS
    EnvironmentProtectionJob_KEXCHANGE
    EnvironmentProtectionJob_KHYPERVVSS
    EnvironmentProtectionJob_KORACLE
    EnvironmentProtectionJob_KGCP
    EnvironmentProtectionJob_KFLASHBLADE
    EnvironmentProtectionJob_KAWSNATIVE
    EnvironmentProtectionJob_KVCD
    EnvironmentProtectionJob_KO365
    EnvironmentProtectionJob_KO365OUTLOOK
    EnvironmentProtectionJob_KHYPERFLEX
    EnvironmentProtectionJob_KGCPNATIVE
    EnvironmentProtectionJob_KAZURENATIVE
    EnvironmentProtectionJob_KKUBERNETES
)

func (r EnvironmentProtectionJobEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentProtectionJobEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentProtectionJobEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentProtectionJobEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentProtectionJobEnum to its string representation
 */
func EnvironmentProtectionJobEnumToValue(environmentProtectionJobEnum EnvironmentProtectionJobEnum) string {
    switch environmentProtectionJobEnum {
        case EnvironmentProtectionJob_KVMWARE:
    		return "kVMware"
        case EnvironmentProtectionJob_KHYPERV:
    		return "kHyperV"
        case EnvironmentProtectionJob_KSQL:
    		return "kSQL"
        case EnvironmentProtectionJob_KVIEW:
    		return "kView"
        case EnvironmentProtectionJob_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentProtectionJob_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentProtectionJob_KPURE:
    		return "kPure"
        case EnvironmentProtectionJob_KAZURE:
    		return "kAzure"
        case EnvironmentProtectionJob_KNETAPP:
    		return "kNetapp"
        case EnvironmentProtectionJob_KAGENT:
    		return "kAgent"
        case EnvironmentProtectionJob_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentProtectionJob_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentProtectionJob_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentProtectionJob_KISILON:
    		return "kIsilon"
        case EnvironmentProtectionJob_KGPFS:
    		return "kGPFS"
        case EnvironmentProtectionJob_KKVM:
    		return "kKVM"
        case EnvironmentProtectionJob_KAWS:
    		return "kAWS"
        case EnvironmentProtectionJob_KEXCHANGE:
    		return "kExchange"
        case EnvironmentProtectionJob_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentProtectionJob_KORACLE:
    		return "kOracle"
        case EnvironmentProtectionJob_KGCP:
    		return "kGCP"
        case EnvironmentProtectionJob_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentProtectionJob_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentProtectionJob_KVCD:
    		return "kVCD"
        case EnvironmentProtectionJob_KO365:
    		return "kO365"
        case EnvironmentProtectionJob_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentProtectionJob_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentProtectionJob_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentProtectionJob_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentProtectionJob_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentProtectionJobEnum Array to its string Array representation
*/
func EnvironmentProtectionJobEnumArrayToValue(environmentProtectionJobEnum []EnvironmentProtectionJobEnum) []string {
    convArray := make([]string,len( environmentProtectionJobEnum))
    for i:=0; i<len(environmentProtectionJobEnum);i++ {
        convArray[i] = EnvironmentProtectionJobEnumToValue(environmentProtectionJobEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentProtectionJobEnumFromValue(value string) EnvironmentProtectionJobEnum {
    switch value {
        case "kVMware":
            return EnvironmentProtectionJob_KVMWARE
        case "kHyperV":
            return EnvironmentProtectionJob_KHYPERV
        case "kSQL":
            return EnvironmentProtectionJob_KSQL
        case "kView":
            return EnvironmentProtectionJob_KVIEW
        case "kPuppeteer":
            return EnvironmentProtectionJob_KPUPPETEER
        case "kPhysical":
            return EnvironmentProtectionJob_KPHYSICAL
        case "kPure":
            return EnvironmentProtectionJob_KPURE
        case "kAzure":
            return EnvironmentProtectionJob_KAZURE
        case "kNetapp":
            return EnvironmentProtectionJob_KNETAPP
        case "kAgent":
            return EnvironmentProtectionJob_KAGENT
        case "kGenericNas":
            return EnvironmentProtectionJob_KGENERICNAS
        case "kAcropolis":
            return EnvironmentProtectionJob_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentProtectionJob_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentProtectionJob_KISILON
        case "kGPFS":
            return EnvironmentProtectionJob_KGPFS
        case "kKVM":
            return EnvironmentProtectionJob_KKVM
        case "kAWS":
            return EnvironmentProtectionJob_KAWS
        case "kExchange":
            return EnvironmentProtectionJob_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentProtectionJob_KHYPERVVSS
        case "kOracle":
            return EnvironmentProtectionJob_KORACLE
        case "kGCP":
            return EnvironmentProtectionJob_KGCP
        case "kFlashBlade":
            return EnvironmentProtectionJob_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentProtectionJob_KAWSNATIVE
        case "kVCD":
            return EnvironmentProtectionJob_KVCD
        case "kO365":
            return EnvironmentProtectionJob_KO365
        case "kO365Outlook":
            return EnvironmentProtectionJob_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentProtectionJob_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentProtectionJob_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentProtectionJob_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentProtectionJob_KKUBERNETES
        default:
            return EnvironmentProtectionJob_KVMWARE
    }
}
