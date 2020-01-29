// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ApplicationEnvironmentEnum enum
 */
type ApplicationEnvironmentEnum int

/**
 * Value collection for ApplicationEnvironmentEnum enum
 */
const (
    ApplicationEnvironment_KVMWARE ApplicationEnvironmentEnum = 1 + iota
    ApplicationEnvironment_KHYPERV
    ApplicationEnvironment_KSQL
    ApplicationEnvironment_KVIEW
    ApplicationEnvironment_KPUPPETEER
    ApplicationEnvironment_KPHYSICAL
    ApplicationEnvironment_KPURE
    ApplicationEnvironment_KAZURE
    ApplicationEnvironment_KNETAPP
    ApplicationEnvironment_KAGENT
    ApplicationEnvironment_KGENERICNAS
    ApplicationEnvironment_KACROPOLIS
    ApplicationEnvironment_KPHYSICALFILES
    ApplicationEnvironment_KISILON
    ApplicationEnvironment_KGPFS
    ApplicationEnvironment_KKVM
    ApplicationEnvironment_KAWS
    ApplicationEnvironment_KEXCHANGE
    ApplicationEnvironment_KHYPERVVSS
    ApplicationEnvironment_KORACLE
    ApplicationEnvironment_KGCP
    ApplicationEnvironment_KFLASHBLADE
    ApplicationEnvironment_KAWSNATIVE
    ApplicationEnvironment_KVCD
    ApplicationEnvironment_KO365
    ApplicationEnvironment_KO365OUTLOOK
    ApplicationEnvironment_KHYPERFLEX
    ApplicationEnvironment_KGCPNATIVE
    ApplicationEnvironment_KAZURENATIVE
    ApplicationEnvironment_KKUBERNETES
)

func (r ApplicationEnvironmentEnum) MarshalJSON() ([]byte, error) {
    s := ApplicationEnvironmentEnumToValue(r)
    return json.Marshal(s)
}

func (r *ApplicationEnvironmentEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ApplicationEnvironmentEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ApplicationEnvironmentEnum to its string representation
 */
func ApplicationEnvironmentEnumToValue(applicationEnvironmentEnum ApplicationEnvironmentEnum) string {
    switch applicationEnvironmentEnum {
        case ApplicationEnvironment_KVMWARE:
    		return "kVMware"
        case ApplicationEnvironment_KHYPERV:
    		return "kHyperV"
        case ApplicationEnvironment_KSQL:
    		return "kSQL"
        case ApplicationEnvironment_KVIEW:
    		return "kView"
        case ApplicationEnvironment_KPUPPETEER:
    		return "kPuppeteer"
        case ApplicationEnvironment_KPHYSICAL:
    		return "kPhysical"
        case ApplicationEnvironment_KPURE:
    		return "kPure"
        case ApplicationEnvironment_KAZURE:
    		return "kAzure"
        case ApplicationEnvironment_KNETAPP:
    		return "kNetapp"
        case ApplicationEnvironment_KAGENT:
    		return "kAgent"
        case ApplicationEnvironment_KGENERICNAS:
    		return "kGenericNas"
        case ApplicationEnvironment_KACROPOLIS:
    		return "kAcropolis"
        case ApplicationEnvironment_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case ApplicationEnvironment_KISILON:
    		return "kIsilon"
        case ApplicationEnvironment_KGPFS:
    		return "kGPFS"
        case ApplicationEnvironment_KKVM:
    		return "kKVM"
        case ApplicationEnvironment_KAWS:
    		return "kAWS"
        case ApplicationEnvironment_KEXCHANGE:
    		return "kExchange"
        case ApplicationEnvironment_KHYPERVVSS:
    		return "kHyperVVSS"
        case ApplicationEnvironment_KORACLE:
    		return "kOracle"
        case ApplicationEnvironment_KGCP:
    		return "kGCP"
        case ApplicationEnvironment_KFLASHBLADE:
    		return "kFlashBlade"
        case ApplicationEnvironment_KAWSNATIVE:
    		return "kAWSNative"
        case ApplicationEnvironment_KVCD:
    		return "kVCD"
        case ApplicationEnvironment_KO365:
    		return "kO365"
        case ApplicationEnvironment_KO365OUTLOOK:
    		return "kO365Outlook"
        case ApplicationEnvironment_KHYPERFLEX:
    		return "kHyperFlex"
        case ApplicationEnvironment_KGCPNATIVE:
    		return "kGCPNative"
        case ApplicationEnvironment_KAZURENATIVE:
    		return "kAzureNative"
        case ApplicationEnvironment_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts ApplicationEnvironmentEnum Array to its string Array representation
*/
func ApplicationEnvironmentEnumArrayToValue(applicationEnvironmentEnum []ApplicationEnvironmentEnum) []string {
    convArray := make([]string,len( applicationEnvironmentEnum))
    for i:=0; i<len(applicationEnvironmentEnum);i++ {
        convArray[i] = ApplicationEnvironmentEnumToValue(applicationEnvironmentEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ApplicationEnvironmentEnumFromValue(value string) ApplicationEnvironmentEnum {
    switch value {
        case "kVMware":
            return ApplicationEnvironment_KVMWARE
        case "kHyperV":
            return ApplicationEnvironment_KHYPERV
        case "kSQL":
            return ApplicationEnvironment_KSQL
        case "kView":
            return ApplicationEnvironment_KVIEW
        case "kPuppeteer":
            return ApplicationEnvironment_KPUPPETEER
        case "kPhysical":
            return ApplicationEnvironment_KPHYSICAL
        case "kPure":
            return ApplicationEnvironment_KPURE
        case "kAzure":
            return ApplicationEnvironment_KAZURE
        case "kNetapp":
            return ApplicationEnvironment_KNETAPP
        case "kAgent":
            return ApplicationEnvironment_KAGENT
        case "kGenericNas":
            return ApplicationEnvironment_KGENERICNAS
        case "kAcropolis":
            return ApplicationEnvironment_KACROPOLIS
        case "kPhysicalFiles":
            return ApplicationEnvironment_KPHYSICALFILES
        case "kIsilon":
            return ApplicationEnvironment_KISILON
        case "kGPFS":
            return ApplicationEnvironment_KGPFS
        case "kKVM":
            return ApplicationEnvironment_KKVM
        case "kAWS":
            return ApplicationEnvironment_KAWS
        case "kExchange":
            return ApplicationEnvironment_KEXCHANGE
        case "kHyperVVSS":
            return ApplicationEnvironment_KHYPERVVSS
        case "kOracle":
            return ApplicationEnvironment_KORACLE
        case "kGCP":
            return ApplicationEnvironment_KGCP
        case "kFlashBlade":
            return ApplicationEnvironment_KFLASHBLADE
        case "kAWSNative":
            return ApplicationEnvironment_KAWSNATIVE
        case "kVCD":
            return ApplicationEnvironment_KVCD
        case "kO365":
            return ApplicationEnvironment_KO365
        case "kO365Outlook":
            return ApplicationEnvironment_KO365OUTLOOK
        case "kHyperFlex":
            return ApplicationEnvironment_KHYPERFLEX
        case "kGCPNative":
            return ApplicationEnvironment_KGCPNATIVE
        case "kAzureNative":
            return ApplicationEnvironment_KAZURENATIVE
        case "kKubernetes":
            return ApplicationEnvironment_KKUBERNETES
        default:
            return ApplicationEnvironment_KVMWARE
    }
}
