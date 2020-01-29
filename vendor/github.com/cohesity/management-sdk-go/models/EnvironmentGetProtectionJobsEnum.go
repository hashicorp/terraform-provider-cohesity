// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentGetProtectionJobsEnum enum
 */
type EnvironmentGetProtectionJobsEnum int

/**
 * Value collection for EnvironmentGetProtectionJobsEnum enum
 */
const (
    EnvironmentGetProtectionJobs_KVMWARE EnvironmentGetProtectionJobsEnum = 1 + iota
    EnvironmentGetProtectionJobs_KHYPERV
    EnvironmentGetProtectionJobs_KSQL
    EnvironmentGetProtectionJobs_KVIEW
    EnvironmentGetProtectionJobs_KPUPPETEER
    EnvironmentGetProtectionJobs_KPHYSICAL
    EnvironmentGetProtectionJobs_KPURE
    EnvironmentGetProtectionJobs_KAZURE
    EnvironmentGetProtectionJobs_KNETAPP
    EnvironmentGetProtectionJobs_KGENERICNAS
    EnvironmentGetProtectionJobs_KACROPOLIS
    EnvironmentGetProtectionJobs_KPHYSICALFILES
    EnvironmentGetProtectionJobs_KISILON
    EnvironmentGetProtectionJobs_KGPFS
    EnvironmentGetProtectionJobs_KKVM
    EnvironmentGetProtectionJobs_KAWS
    EnvironmentGetProtectionJobs_KEXCHANGE
    EnvironmentGetProtectionJobs_KHYPERVVSS
    EnvironmentGetProtectionJobs_KORACLE
    EnvironmentGetProtectionJobs_KGCP
    EnvironmentGetProtectionJobs_KFLASHBLADE
    EnvironmentGetProtectionJobs_KAWSNATIVE
    EnvironmentGetProtectionJobs_KVCD
    EnvironmentGetProtectionJobs_KO365
    EnvironmentGetProtectionJobs_KO365OUTLOOK
    EnvironmentGetProtectionJobs_KHYPERFLEX
    EnvironmentGetProtectionJobs_KGCPNATIVE
    EnvironmentGetProtectionJobs_KKUBERNETES
)

func (r EnvironmentGetProtectionJobsEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentGetProtectionJobsEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentGetProtectionJobsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentGetProtectionJobsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentGetProtectionJobsEnum to its string representation
 */
func EnvironmentGetProtectionJobsEnumToValue(environmentGetProtectionJobsEnum EnvironmentGetProtectionJobsEnum) string {
    switch environmentGetProtectionJobsEnum {
        case EnvironmentGetProtectionJobs_KVMWARE:
    		return "kVMware"
        case EnvironmentGetProtectionJobs_KHYPERV:
    		return "kHyperV"
        case EnvironmentGetProtectionJobs_KSQL:
    		return "kSQL"
        case EnvironmentGetProtectionJobs_KVIEW:
    		return "kView"
        case EnvironmentGetProtectionJobs_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentGetProtectionJobs_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentGetProtectionJobs_KPURE:
    		return "kPure"
        case EnvironmentGetProtectionJobs_KAZURE:
    		return "kAzure"
        case EnvironmentGetProtectionJobs_KNETAPP:
    		return "kNetapp"
        case EnvironmentGetProtectionJobs_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentGetProtectionJobs_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentGetProtectionJobs_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentGetProtectionJobs_KISILON:
    		return "kIsilon"
        case EnvironmentGetProtectionJobs_KGPFS:
    		return "kGPFS"
        case EnvironmentGetProtectionJobs_KKVM:
    		return "kKVM"
        case EnvironmentGetProtectionJobs_KAWS:
    		return "kAWS"
        case EnvironmentGetProtectionJobs_KEXCHANGE:
    		return "kExchange"
        case EnvironmentGetProtectionJobs_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentGetProtectionJobs_KORACLE:
    		return "kOracle"
        case EnvironmentGetProtectionJobs_KGCP:
    		return "kGCP"
        case EnvironmentGetProtectionJobs_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentGetProtectionJobs_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentGetProtectionJobs_KVCD:
    		return "kVCD"
        case EnvironmentGetProtectionJobs_KO365:
    		return "kO365"
        case EnvironmentGetProtectionJobs_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentGetProtectionJobs_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentGetProtectionJobs_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentGetProtectionJobs_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentGetProtectionJobsEnum Array to its string Array representation
*/
func EnvironmentGetProtectionJobsEnumArrayToValue(environmentGetProtectionJobsEnum []EnvironmentGetProtectionJobsEnum) []string {
    convArray := make([]string,len( environmentGetProtectionJobsEnum))
    for i:=0; i<len(environmentGetProtectionJobsEnum);i++ {
        convArray[i] = EnvironmentGetProtectionJobsEnumToValue(environmentGetProtectionJobsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentGetProtectionJobsEnumFromValue(value string) EnvironmentGetProtectionJobsEnum {
    switch value {
        case "kVMware":
            return EnvironmentGetProtectionJobs_KVMWARE
        case "kHyperV":
            return EnvironmentGetProtectionJobs_KHYPERV
        case "kSQL":
            return EnvironmentGetProtectionJobs_KSQL
        case "kView":
            return EnvironmentGetProtectionJobs_KVIEW
        case "kPuppeteer":
            return EnvironmentGetProtectionJobs_KPUPPETEER
        case "kPhysical":
            return EnvironmentGetProtectionJobs_KPHYSICAL
        case "kPure":
            return EnvironmentGetProtectionJobs_KPURE
        case "kAzure":
            return EnvironmentGetProtectionJobs_KAZURE
        case "kNetapp":
            return EnvironmentGetProtectionJobs_KNETAPP
        case "kGenericNas":
            return EnvironmentGetProtectionJobs_KGENERICNAS
        case "kAcropolis":
            return EnvironmentGetProtectionJobs_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentGetProtectionJobs_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentGetProtectionJobs_KISILON
        case "kGPFS":
            return EnvironmentGetProtectionJobs_KGPFS
        case "kKVM":
            return EnvironmentGetProtectionJobs_KKVM
        case "kAWS":
            return EnvironmentGetProtectionJobs_KAWS
        case "kExchange":
            return EnvironmentGetProtectionJobs_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentGetProtectionJobs_KHYPERVVSS
        case "kOracle":
            return EnvironmentGetProtectionJobs_KORACLE
        case "kGCP":
            return EnvironmentGetProtectionJobs_KGCP
        case "kFlashBlade":
            return EnvironmentGetProtectionJobs_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentGetProtectionJobs_KAWSNATIVE
        case "kVCD":
            return EnvironmentGetProtectionJobs_KVCD
        case "kO365":
            return EnvironmentGetProtectionJobs_KO365
        case "kO365Outlook":
            return EnvironmentGetProtectionJobs_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentGetProtectionJobs_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentGetProtectionJobs_KGCPNATIVE
        case "kKubernetes":
            return EnvironmentGetProtectionJobs_KKUBERNETES
        default:
            return EnvironmentGetProtectionJobs_KVMWARE
    }
}
