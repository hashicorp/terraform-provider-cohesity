// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ApplicationEnvironmentApplicationsRestoreTaskRequestEnum enum
 */
type ApplicationEnvironmentApplicationsRestoreTaskRequestEnum int

/**
 * Value collection for ApplicationEnvironmentApplicationsRestoreTaskRequestEnum enum
 */
const (
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KVMWARE ApplicationEnvironmentApplicationsRestoreTaskRequestEnum = 1 + iota
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KHYPERV
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KSQL
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KVIEW
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KPUPPETEER
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KPHYSICAL
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KPURE
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KAZURE
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KNETAPP
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KAGENT
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KGENERICNAS
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KACROPOLIS
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KPHYSICALFILES
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KISILON
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KGPFS
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KKVM
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KAWS
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KEXCHANGE
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KHYPERVVSS
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KORACLE
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KGCP
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KFLASHBLADE
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KAWSNATIVE
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KVCD
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KO365
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KO365OUTLOOK
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KHYPERFLEX
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KGCPNATIVE
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KAZURENATIVE
    ApplicationEnvironmentApplicationsRestoreTaskRequest_KKUBERNETES
)

func (r ApplicationEnvironmentApplicationsRestoreTaskRequestEnum) MarshalJSON() ([]byte, error) {
    s := ApplicationEnvironmentApplicationsRestoreTaskRequestEnumToValue(r)
    return json.Marshal(s)
}

func (r *ApplicationEnvironmentApplicationsRestoreTaskRequestEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ApplicationEnvironmentApplicationsRestoreTaskRequestEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ApplicationEnvironmentApplicationsRestoreTaskRequestEnum to its string representation
 */
func ApplicationEnvironmentApplicationsRestoreTaskRequestEnumToValue(applicationEnvironmentApplicationsRestoreTaskRequestEnum ApplicationEnvironmentApplicationsRestoreTaskRequestEnum) string {
    switch applicationEnvironmentApplicationsRestoreTaskRequestEnum {
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KVMWARE:
    		return "kVMware"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KHYPERV:
    		return "kHyperV"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KSQL:
    		return "kSQL"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KVIEW:
    		return "kView"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KPUPPETEER:
    		return "kPuppeteer"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KPHYSICAL:
    		return "kPhysical"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KPURE:
    		return "kPure"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KAZURE:
    		return "kAzure"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KNETAPP:
    		return "kNetapp"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KAGENT:
    		return "kAgent"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KGENERICNAS:
    		return "kGenericNas"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KACROPOLIS:
    		return "kAcropolis"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KISILON:
    		return "kIsilon"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KGPFS:
    		return "kGPFS"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KKVM:
    		return "kKVM"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KAWS:
    		return "kAWS"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KEXCHANGE:
    		return "kExchange"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KHYPERVVSS:
    		return "kHyperVVSS"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KORACLE:
    		return "kOracle"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KGCP:
    		return "kGCP"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KFLASHBLADE:
    		return "kFlashBlade"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KAWSNATIVE:
    		return "kAWSNative"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KVCD:
    		return "kVCD"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KO365:
    		return "kO365"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KO365OUTLOOK:
    		return "kO365Outlook"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KHYPERFLEX:
    		return "kHyperFlex"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KGCPNATIVE:
    		return "kGCPNative"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KAZURENATIVE:
    		return "kAzureNative"
        case ApplicationEnvironmentApplicationsRestoreTaskRequest_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts ApplicationEnvironmentApplicationsRestoreTaskRequestEnum Array to its string Array representation
*/
func ApplicationEnvironmentApplicationsRestoreTaskRequestEnumArrayToValue(applicationEnvironmentApplicationsRestoreTaskRequestEnum []ApplicationEnvironmentApplicationsRestoreTaskRequestEnum) []string {
    convArray := make([]string,len( applicationEnvironmentApplicationsRestoreTaskRequestEnum))
    for i:=0; i<len(applicationEnvironmentApplicationsRestoreTaskRequestEnum);i++ {
        convArray[i] = ApplicationEnvironmentApplicationsRestoreTaskRequestEnumToValue(applicationEnvironmentApplicationsRestoreTaskRequestEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ApplicationEnvironmentApplicationsRestoreTaskRequestEnumFromValue(value string) ApplicationEnvironmentApplicationsRestoreTaskRequestEnum {
    switch value {
        case "kVMware":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KVMWARE
        case "kHyperV":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KHYPERV
        case "kSQL":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KSQL
        case "kView":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KVIEW
        case "kPuppeteer":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KPUPPETEER
        case "kPhysical":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KPHYSICAL
        case "kPure":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KPURE
        case "kAzure":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KAZURE
        case "kNetapp":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KNETAPP
        case "kAgent":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KAGENT
        case "kGenericNas":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KGENERICNAS
        case "kAcropolis":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KACROPOLIS
        case "kPhysicalFiles":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KPHYSICALFILES
        case "kIsilon":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KISILON
        case "kGPFS":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KGPFS
        case "kKVM":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KKVM
        case "kAWS":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KAWS
        case "kExchange":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KEXCHANGE
        case "kHyperVVSS":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KHYPERVVSS
        case "kOracle":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KORACLE
        case "kGCP":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KGCP
        case "kFlashBlade":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KFLASHBLADE
        case "kAWSNative":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KAWSNATIVE
        case "kVCD":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KVCD
        case "kO365":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KO365
        case "kO365Outlook":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KO365OUTLOOK
        case "kHyperFlex":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KHYPERFLEX
        case "kGCPNative":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KGCPNATIVE
        case "kAzureNative":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KAZURENATIVE
        case "kKubernetes":
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KKUBERNETES
        default:
            return ApplicationEnvironmentApplicationsRestoreTaskRequest_KVMWARE
    }
}
