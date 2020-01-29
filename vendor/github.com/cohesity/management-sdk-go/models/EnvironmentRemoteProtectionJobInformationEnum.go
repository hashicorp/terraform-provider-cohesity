// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentRemoteProtectionJobInformationEnum enum
 */
type EnvironmentRemoteProtectionJobInformationEnum int

/**
 * Value collection for EnvironmentRemoteProtectionJobInformationEnum enum
 */
const (
    EnvironmentRemoteProtectionJobInformation_KVMWARE EnvironmentRemoteProtectionJobInformationEnum = 1 + iota
    EnvironmentRemoteProtectionJobInformation_KHYPERV
    EnvironmentRemoteProtectionJobInformation_KSQL
    EnvironmentRemoteProtectionJobInformation_KVIEW
    EnvironmentRemoteProtectionJobInformation_KPUPPETEER
    EnvironmentRemoteProtectionJobInformation_KPHYSICAL
    EnvironmentRemoteProtectionJobInformation_KPURE
    EnvironmentRemoteProtectionJobInformation_KAZURE
    EnvironmentRemoteProtectionJobInformation_KNETAPP
    EnvironmentRemoteProtectionJobInformation_KAGENT
    EnvironmentRemoteProtectionJobInformation_KGENERICNAS
    EnvironmentRemoteProtectionJobInformation_KACROPOLIS
    EnvironmentRemoteProtectionJobInformation_KPHYSICALFILES
    EnvironmentRemoteProtectionJobInformation_KISILON
    EnvironmentRemoteProtectionJobInformation_KGPFS
    EnvironmentRemoteProtectionJobInformation_KKVM
    EnvironmentRemoteProtectionJobInformation_KAWS
    EnvironmentRemoteProtectionJobInformation_KEXCHANGE
    EnvironmentRemoteProtectionJobInformation_KHYPERVVSS
    EnvironmentRemoteProtectionJobInformation_KORACLE
    EnvironmentRemoteProtectionJobInformation_KGCP
    EnvironmentRemoteProtectionJobInformation_KFLASHBLADE
    EnvironmentRemoteProtectionJobInformation_KAWSNATIVE
    EnvironmentRemoteProtectionJobInformation_KVCD
    EnvironmentRemoteProtectionJobInformation_KO365
    EnvironmentRemoteProtectionJobInformation_KO365OUTLOOK
    EnvironmentRemoteProtectionJobInformation_KHYPERFLEX
    EnvironmentRemoteProtectionJobInformation_KGCPNATIVE
    EnvironmentRemoteProtectionJobInformation_KAZURENATIVE
    EnvironmentRemoteProtectionJobInformation_KKUBERNETES
)

func (r EnvironmentRemoteProtectionJobInformationEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentRemoteProtectionJobInformationEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentRemoteProtectionJobInformationEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentRemoteProtectionJobInformationEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentRemoteProtectionJobInformationEnum to its string representation
 */
func EnvironmentRemoteProtectionJobInformationEnumToValue(environmentRemoteProtectionJobInformationEnum EnvironmentRemoteProtectionJobInformationEnum) string {
    switch environmentRemoteProtectionJobInformationEnum {
        case EnvironmentRemoteProtectionJobInformation_KVMWARE:
    		return "kVMware"
        case EnvironmentRemoteProtectionJobInformation_KHYPERV:
    		return "kHyperV"
        case EnvironmentRemoteProtectionJobInformation_KSQL:
    		return "kSQL"
        case EnvironmentRemoteProtectionJobInformation_KVIEW:
    		return "kView"
        case EnvironmentRemoteProtectionJobInformation_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentRemoteProtectionJobInformation_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentRemoteProtectionJobInformation_KPURE:
    		return "kPure"
        case EnvironmentRemoteProtectionJobInformation_KAZURE:
    		return "kAzure"
        case EnvironmentRemoteProtectionJobInformation_KNETAPP:
    		return "kNetapp"
        case EnvironmentRemoteProtectionJobInformation_KAGENT:
    		return "kAgent"
        case EnvironmentRemoteProtectionJobInformation_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentRemoteProtectionJobInformation_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentRemoteProtectionJobInformation_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentRemoteProtectionJobInformation_KISILON:
    		return "kIsilon"
        case EnvironmentRemoteProtectionJobInformation_KGPFS:
    		return "kGPFS"
        case EnvironmentRemoteProtectionJobInformation_KKVM:
    		return "kKVM"
        case EnvironmentRemoteProtectionJobInformation_KAWS:
    		return "kAWS"
        case EnvironmentRemoteProtectionJobInformation_KEXCHANGE:
    		return "kExchange"
        case EnvironmentRemoteProtectionJobInformation_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentRemoteProtectionJobInformation_KORACLE:
    		return "kOracle"
        case EnvironmentRemoteProtectionJobInformation_KGCP:
    		return "kGCP"
        case EnvironmentRemoteProtectionJobInformation_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentRemoteProtectionJobInformation_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentRemoteProtectionJobInformation_KVCD:
    		return "kVCD"
        case EnvironmentRemoteProtectionJobInformation_KO365:
    		return "kO365"
        case EnvironmentRemoteProtectionJobInformation_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentRemoteProtectionJobInformation_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentRemoteProtectionJobInformation_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentRemoteProtectionJobInformation_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentRemoteProtectionJobInformation_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentRemoteProtectionJobInformationEnum Array to its string Array representation
*/
func EnvironmentRemoteProtectionJobInformationEnumArrayToValue(environmentRemoteProtectionJobInformationEnum []EnvironmentRemoteProtectionJobInformationEnum) []string {
    convArray := make([]string,len( environmentRemoteProtectionJobInformationEnum))
    for i:=0; i<len(environmentRemoteProtectionJobInformationEnum);i++ {
        convArray[i] = EnvironmentRemoteProtectionJobInformationEnumToValue(environmentRemoteProtectionJobInformationEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentRemoteProtectionJobInformationEnumFromValue(value string) EnvironmentRemoteProtectionJobInformationEnum {
    switch value {
        case "kVMware":
            return EnvironmentRemoteProtectionJobInformation_KVMWARE
        case "kHyperV":
            return EnvironmentRemoteProtectionJobInformation_KHYPERV
        case "kSQL":
            return EnvironmentRemoteProtectionJobInformation_KSQL
        case "kView":
            return EnvironmentRemoteProtectionJobInformation_KVIEW
        case "kPuppeteer":
            return EnvironmentRemoteProtectionJobInformation_KPUPPETEER
        case "kPhysical":
            return EnvironmentRemoteProtectionJobInformation_KPHYSICAL
        case "kPure":
            return EnvironmentRemoteProtectionJobInformation_KPURE
        case "kAzure":
            return EnvironmentRemoteProtectionJobInformation_KAZURE
        case "kNetapp":
            return EnvironmentRemoteProtectionJobInformation_KNETAPP
        case "kAgent":
            return EnvironmentRemoteProtectionJobInformation_KAGENT
        case "kGenericNas":
            return EnvironmentRemoteProtectionJobInformation_KGENERICNAS
        case "kAcropolis":
            return EnvironmentRemoteProtectionJobInformation_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentRemoteProtectionJobInformation_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentRemoteProtectionJobInformation_KISILON
        case "kGPFS":
            return EnvironmentRemoteProtectionJobInformation_KGPFS
        case "kKVM":
            return EnvironmentRemoteProtectionJobInformation_KKVM
        case "kAWS":
            return EnvironmentRemoteProtectionJobInformation_KAWS
        case "kExchange":
            return EnvironmentRemoteProtectionJobInformation_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentRemoteProtectionJobInformation_KHYPERVVSS
        case "kOracle":
            return EnvironmentRemoteProtectionJobInformation_KORACLE
        case "kGCP":
            return EnvironmentRemoteProtectionJobInformation_KGCP
        case "kFlashBlade":
            return EnvironmentRemoteProtectionJobInformation_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentRemoteProtectionJobInformation_KAWSNATIVE
        case "kVCD":
            return EnvironmentRemoteProtectionJobInformation_KVCD
        case "kO365":
            return EnvironmentRemoteProtectionJobInformation_KO365
        case "kO365Outlook":
            return EnvironmentRemoteProtectionJobInformation_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentRemoteProtectionJobInformation_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentRemoteProtectionJobInformation_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentRemoteProtectionJobInformation_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentRemoteProtectionJobInformation_KKUBERNETES
        default:
            return EnvironmentRemoteProtectionJobInformation_KVMWARE
    }
}
