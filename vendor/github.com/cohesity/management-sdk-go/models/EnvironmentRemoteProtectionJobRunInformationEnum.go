// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentRemoteProtectionJobRunInformationEnum enum
 */
type EnvironmentRemoteProtectionJobRunInformationEnum int

/**
 * Value collection for EnvironmentRemoteProtectionJobRunInformationEnum enum
 */
const (
    EnvironmentRemoteProtectionJobRunInformation_KVMWARE EnvironmentRemoteProtectionJobRunInformationEnum = 1 + iota
    EnvironmentRemoteProtectionJobRunInformation_KHYPERV
    EnvironmentRemoteProtectionJobRunInformation_KSQL
    EnvironmentRemoteProtectionJobRunInformation_KVIEW
    EnvironmentRemoteProtectionJobRunInformation_KPUPPETEER
    EnvironmentRemoteProtectionJobRunInformation_KPHYSICAL
    EnvironmentRemoteProtectionJobRunInformation_KPURE
    EnvironmentRemoteProtectionJobRunInformation_KAZURE
    EnvironmentRemoteProtectionJobRunInformation_KNETAPP
    EnvironmentRemoteProtectionJobRunInformation_KAGENT
    EnvironmentRemoteProtectionJobRunInformation_KGENERICNAS
    EnvironmentRemoteProtectionJobRunInformation_KACROPOLIS
    EnvironmentRemoteProtectionJobRunInformation_KPHYSICALFILES
    EnvironmentRemoteProtectionJobRunInformation_KISILON
    EnvironmentRemoteProtectionJobRunInformation_KGPFS
    EnvironmentRemoteProtectionJobRunInformation_KKVM
    EnvironmentRemoteProtectionJobRunInformation_KAWS
    EnvironmentRemoteProtectionJobRunInformation_KEXCHANGE
    EnvironmentRemoteProtectionJobRunInformation_KHYPERVVSS
    EnvironmentRemoteProtectionJobRunInformation_KORACLE
    EnvironmentRemoteProtectionJobRunInformation_KGCP
    EnvironmentRemoteProtectionJobRunInformation_KFLASHBLADE
    EnvironmentRemoteProtectionJobRunInformation_KAWSNATIVE
    EnvironmentRemoteProtectionJobRunInformation_KVCD
    EnvironmentRemoteProtectionJobRunInformation_KO365
    EnvironmentRemoteProtectionJobRunInformation_KO365OUTLOOK
    EnvironmentRemoteProtectionJobRunInformation_KHYPERFLEX
    EnvironmentRemoteProtectionJobRunInformation_KGCPNATIVE
    EnvironmentRemoteProtectionJobRunInformation_KAZURENATIVE
    EnvironmentRemoteProtectionJobRunInformation_KKUBERNETES
)

func (r EnvironmentRemoteProtectionJobRunInformationEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentRemoteProtectionJobRunInformationEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentRemoteProtectionJobRunInformationEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentRemoteProtectionJobRunInformationEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentRemoteProtectionJobRunInformationEnum to its string representation
 */
func EnvironmentRemoteProtectionJobRunInformationEnumToValue(environmentRemoteProtectionJobRunInformationEnum EnvironmentRemoteProtectionJobRunInformationEnum) string {
    switch environmentRemoteProtectionJobRunInformationEnum {
        case EnvironmentRemoteProtectionJobRunInformation_KVMWARE:
    		return "kVMware"
        case EnvironmentRemoteProtectionJobRunInformation_KHYPERV:
    		return "kHyperV"
        case EnvironmentRemoteProtectionJobRunInformation_KSQL:
    		return "kSQL"
        case EnvironmentRemoteProtectionJobRunInformation_KVIEW:
    		return "kView"
        case EnvironmentRemoteProtectionJobRunInformation_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentRemoteProtectionJobRunInformation_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentRemoteProtectionJobRunInformation_KPURE:
    		return "kPure"
        case EnvironmentRemoteProtectionJobRunInformation_KAZURE:
    		return "kAzure"
        case EnvironmentRemoteProtectionJobRunInformation_KNETAPP:
    		return "kNetapp"
        case EnvironmentRemoteProtectionJobRunInformation_KAGENT:
    		return "kAgent"
        case EnvironmentRemoteProtectionJobRunInformation_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentRemoteProtectionJobRunInformation_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentRemoteProtectionJobRunInformation_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentRemoteProtectionJobRunInformation_KISILON:
    		return "kIsilon"
        case EnvironmentRemoteProtectionJobRunInformation_KGPFS:
    		return "kGPFS"
        case EnvironmentRemoteProtectionJobRunInformation_KKVM:
    		return "kKVM"
        case EnvironmentRemoteProtectionJobRunInformation_KAWS:
    		return "kAWS"
        case EnvironmentRemoteProtectionJobRunInformation_KEXCHANGE:
    		return "kExchange"
        case EnvironmentRemoteProtectionJobRunInformation_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentRemoteProtectionJobRunInformation_KORACLE:
    		return "kOracle"
        case EnvironmentRemoteProtectionJobRunInformation_KGCP:
    		return "kGCP"
        case EnvironmentRemoteProtectionJobRunInformation_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentRemoteProtectionJobRunInformation_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentRemoteProtectionJobRunInformation_KVCD:
    		return "kVCD"
        case EnvironmentRemoteProtectionJobRunInformation_KO365:
    		return "kO365"
        case EnvironmentRemoteProtectionJobRunInformation_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentRemoteProtectionJobRunInformation_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentRemoteProtectionJobRunInformation_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentRemoteProtectionJobRunInformation_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentRemoteProtectionJobRunInformation_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentRemoteProtectionJobRunInformationEnum Array to its string Array representation
*/
func EnvironmentRemoteProtectionJobRunInformationEnumArrayToValue(environmentRemoteProtectionJobRunInformationEnum []EnvironmentRemoteProtectionJobRunInformationEnum) []string {
    convArray := make([]string,len( environmentRemoteProtectionJobRunInformationEnum))
    for i:=0; i<len(environmentRemoteProtectionJobRunInformationEnum);i++ {
        convArray[i] = EnvironmentRemoteProtectionJobRunInformationEnumToValue(environmentRemoteProtectionJobRunInformationEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentRemoteProtectionJobRunInformationEnumFromValue(value string) EnvironmentRemoteProtectionJobRunInformationEnum {
    switch value {
        case "kVMware":
            return EnvironmentRemoteProtectionJobRunInformation_KVMWARE
        case "kHyperV":
            return EnvironmentRemoteProtectionJobRunInformation_KHYPERV
        case "kSQL":
            return EnvironmentRemoteProtectionJobRunInformation_KSQL
        case "kView":
            return EnvironmentRemoteProtectionJobRunInformation_KVIEW
        case "kPuppeteer":
            return EnvironmentRemoteProtectionJobRunInformation_KPUPPETEER
        case "kPhysical":
            return EnvironmentRemoteProtectionJobRunInformation_KPHYSICAL
        case "kPure":
            return EnvironmentRemoteProtectionJobRunInformation_KPURE
        case "kAzure":
            return EnvironmentRemoteProtectionJobRunInformation_KAZURE
        case "kNetapp":
            return EnvironmentRemoteProtectionJobRunInformation_KNETAPP
        case "kAgent":
            return EnvironmentRemoteProtectionJobRunInformation_KAGENT
        case "kGenericNas":
            return EnvironmentRemoteProtectionJobRunInformation_KGENERICNAS
        case "kAcropolis":
            return EnvironmentRemoteProtectionJobRunInformation_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentRemoteProtectionJobRunInformation_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentRemoteProtectionJobRunInformation_KISILON
        case "kGPFS":
            return EnvironmentRemoteProtectionJobRunInformation_KGPFS
        case "kKVM":
            return EnvironmentRemoteProtectionJobRunInformation_KKVM
        case "kAWS":
            return EnvironmentRemoteProtectionJobRunInformation_KAWS
        case "kExchange":
            return EnvironmentRemoteProtectionJobRunInformation_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentRemoteProtectionJobRunInformation_KHYPERVVSS
        case "kOracle":
            return EnvironmentRemoteProtectionJobRunInformation_KORACLE
        case "kGCP":
            return EnvironmentRemoteProtectionJobRunInformation_KGCP
        case "kFlashBlade":
            return EnvironmentRemoteProtectionJobRunInformation_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentRemoteProtectionJobRunInformation_KAWSNATIVE
        case "kVCD":
            return EnvironmentRemoteProtectionJobRunInformation_KVCD
        case "kO365":
            return EnvironmentRemoteProtectionJobRunInformation_KO365
        case "kO365Outlook":
            return EnvironmentRemoteProtectionJobRunInformation_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentRemoteProtectionJobRunInformation_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentRemoteProtectionJobRunInformation_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentRemoteProtectionJobRunInformation_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentRemoteProtectionJobRunInformation_KKUBERNETES
        default:
            return EnvironmentRemoteProtectionJobRunInformation_KVMWARE
    }
}
