// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ApplicationEnum enum
 */
type ApplicationEnum int

/**
 * Value collection for ApplicationEnum enum
 */
const (
    Application_KVMWARE ApplicationEnum = 1 + iota
    Application_KHYPERV
    Application_KSQL
    Application_KVIEW
    Application_KPUPPETEER
    Application_KPHYSICAL
    Application_KPURE
    Application_KAZURE
    Application_KNETAPP
    Application_KAGENT
    Application_KGENERICNAS
    Application_KACROPOLIS
    Application_KPHYSICALFILES
    Application_KISILON
    Application_KGPFS
    Application_KKVM
    Application_KAWS
    Application_KEXCHANGE
    Application_KHYPERVVSS
    Application_KORACLE
    Application_KGCP
    Application_KFLASHBLADE
    Application_KAWSNATIVE
    Application_KVCD
    Application_KO365
    Application_KO365OUTLOOK
    Application_KHYPERFLEX
    Application_KGCPNATIVE
    Application_KAZURENATIVE
    Application_KKUBERNETES
)

func (r ApplicationEnum) MarshalJSON() ([]byte, error) {
    s := ApplicationEnumToValue(r)
    return json.Marshal(s)
}

func (r *ApplicationEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ApplicationEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ApplicationEnum to its string representation
 */
func ApplicationEnumToValue(applicationEnum ApplicationEnum) string {
    switch applicationEnum {
        case Application_KVMWARE:
    		return "kVMware"
        case Application_KHYPERV:
    		return "kHyperV"
        case Application_KSQL:
    		return "kSQL"
        case Application_KVIEW:
    		return "kView"
        case Application_KPUPPETEER:
    		return "kPuppeteer"
        case Application_KPHYSICAL:
    		return "kPhysical"
        case Application_KPURE:
    		return "kPure"
        case Application_KAZURE:
    		return "kAzure"
        case Application_KNETAPP:
    		return "kNetapp"
        case Application_KAGENT:
    		return "kAgent"
        case Application_KGENERICNAS:
    		return "kGenericNas"
        case Application_KACROPOLIS:
    		return "kAcropolis"
        case Application_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case Application_KISILON:
    		return "kIsilon"
        case Application_KGPFS:
    		return "kGPFS"
        case Application_KKVM:
    		return "kKVM"
        case Application_KAWS:
    		return "kAWS"
        case Application_KEXCHANGE:
    		return "kExchange"
        case Application_KHYPERVVSS:
    		return "kHyperVVSS"
        case Application_KORACLE:
    		return "kOracle"
        case Application_KGCP:
    		return "kGCP"
        case Application_KFLASHBLADE:
    		return "kFlashBlade"
        case Application_KAWSNATIVE:
    		return "kAWSNative"
        case Application_KVCD:
    		return "kVCD"
        case Application_KO365:
    		return "kO365"
        case Application_KO365OUTLOOK:
    		return "kO365Outlook"
        case Application_KHYPERFLEX:
    		return "kHyperFlex"
        case Application_KGCPNATIVE:
    		return "kGCPNative"
        case Application_KAZURENATIVE:
    		return "kAzureNative"
        case Application_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts ApplicationEnum Array to its string Array representation
*/
func ApplicationEnumArrayToValue(applicationEnum []ApplicationEnum) []string {
    convArray := make([]string,len( applicationEnum))
    for i:=0; i<len(applicationEnum);i++ {
        convArray[i] = ApplicationEnumToValue(applicationEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ApplicationEnumFromValue(value string) ApplicationEnum {
    switch value {
        case "kVMware":
            return Application_KVMWARE
        case "kHyperV":
            return Application_KHYPERV
        case "kSQL":
            return Application_KSQL
        case "kView":
            return Application_KVIEW
        case "kPuppeteer":
            return Application_KPUPPETEER
        case "kPhysical":
            return Application_KPHYSICAL
        case "kPure":
            return Application_KPURE
        case "kAzure":
            return Application_KAZURE
        case "kNetapp":
            return Application_KNETAPP
        case "kAgent":
            return Application_KAGENT
        case "kGenericNas":
            return Application_KGENERICNAS
        case "kAcropolis":
            return Application_KACROPOLIS
        case "kPhysicalFiles":
            return Application_KPHYSICALFILES
        case "kIsilon":
            return Application_KISILON
        case "kGPFS":
            return Application_KGPFS
        case "kKVM":
            return Application_KKVM
        case "kAWS":
            return Application_KAWS
        case "kExchange":
            return Application_KEXCHANGE
        case "kHyperVVSS":
            return Application_KHYPERVVSS
        case "kOracle":
            return Application_KORACLE
        case "kGCP":
            return Application_KGCP
        case "kFlashBlade":
            return Application_KFLASHBLADE
        case "kAWSNative":
            return Application_KAWSNATIVE
        case "kVCD":
            return Application_KVCD
        case "kO365":
            return Application_KO365
        case "kO365Outlook":
            return Application_KO365OUTLOOK
        case "kHyperFlex":
            return Application_KHYPERFLEX
        case "kGCPNative":
            return Application_KGCPNATIVE
        case "kAzureNative":
            return Application_KAZURENATIVE
        case "kKubernetes":
            return Application_KKUBERNETES
        default:
            return Application_KVMWARE
    }
}
