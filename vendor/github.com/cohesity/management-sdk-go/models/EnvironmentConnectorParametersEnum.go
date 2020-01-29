// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentConnectorParametersEnum enum
 */
type EnvironmentConnectorParametersEnum int

/**
 * Value collection for EnvironmentConnectorParametersEnum enum
 */
const (
    EnvironmentConnectorParameters_KVMWARE EnvironmentConnectorParametersEnum = 1 + iota
    EnvironmentConnectorParameters_KHYPERV
    EnvironmentConnectorParameters_KSQL
    EnvironmentConnectorParameters_KVIEW
    EnvironmentConnectorParameters_KPUPPETEER
    EnvironmentConnectorParameters_KPHYSICAL
    EnvironmentConnectorParameters_KPURE
    EnvironmentConnectorParameters_KAZURE
    EnvironmentConnectorParameters_KNETAPP
    EnvironmentConnectorParameters_KAGENT
    EnvironmentConnectorParameters_KGENERICNAS
    EnvironmentConnectorParameters_KACROPOLIS
    EnvironmentConnectorParameters_KPHYSICALFILES
    EnvironmentConnectorParameters_KISILON
    EnvironmentConnectorParameters_KGPFS
    EnvironmentConnectorParameters_KKVM
    EnvironmentConnectorParameters_KAWS
    EnvironmentConnectorParameters_KEXCHANGE
    EnvironmentConnectorParameters_KHYPERVVSS
    EnvironmentConnectorParameters_KORACLE
    EnvironmentConnectorParameters_KGCP
    EnvironmentConnectorParameters_KFLASHBLADE
    EnvironmentConnectorParameters_KAWSNATIVE
    EnvironmentConnectorParameters_KVCD
    EnvironmentConnectorParameters_KO365
    EnvironmentConnectorParameters_KO365OUTLOOK
    EnvironmentConnectorParameters_KHYPERFLEX
    EnvironmentConnectorParameters_KGCPNATIVE
    EnvironmentConnectorParameters_KAZURENATIVE
    EnvironmentConnectorParameters_KKUBERNETES
)

func (r EnvironmentConnectorParametersEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentConnectorParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentConnectorParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentConnectorParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentConnectorParametersEnum to its string representation
 */
func EnvironmentConnectorParametersEnumToValue(environmentConnectorParametersEnum EnvironmentConnectorParametersEnum) string {
    switch environmentConnectorParametersEnum {
        case EnvironmentConnectorParameters_KVMWARE:
    		return "kVMware"
        case EnvironmentConnectorParameters_KHYPERV:
    		return "kHyperV"
        case EnvironmentConnectorParameters_KSQL:
    		return "kSQL"
        case EnvironmentConnectorParameters_KVIEW:
    		return "kView"
        case EnvironmentConnectorParameters_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentConnectorParameters_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentConnectorParameters_KPURE:
    		return "kPure"
        case EnvironmentConnectorParameters_KAZURE:
    		return "kAzure"
        case EnvironmentConnectorParameters_KNETAPP:
    		return "kNetapp"
        case EnvironmentConnectorParameters_KAGENT:
    		return "kAgent"
        case EnvironmentConnectorParameters_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentConnectorParameters_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentConnectorParameters_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentConnectorParameters_KISILON:
    		return "kIsilon"
        case EnvironmentConnectorParameters_KGPFS:
    		return "kGPFS"
        case EnvironmentConnectorParameters_KKVM:
    		return "kKVM"
        case EnvironmentConnectorParameters_KAWS:
    		return "kAWS"
        case EnvironmentConnectorParameters_KEXCHANGE:
    		return "kExchange"
        case EnvironmentConnectorParameters_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentConnectorParameters_KORACLE:
    		return "kOracle"
        case EnvironmentConnectorParameters_KGCP:
    		return "kGCP"
        case EnvironmentConnectorParameters_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentConnectorParameters_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentConnectorParameters_KVCD:
    		return "kVCD"
        case EnvironmentConnectorParameters_KO365:
    		return "kO365"
        case EnvironmentConnectorParameters_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentConnectorParameters_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentConnectorParameters_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentConnectorParameters_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentConnectorParameters_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentConnectorParametersEnum Array to its string Array representation
*/
func EnvironmentConnectorParametersEnumArrayToValue(environmentConnectorParametersEnum []EnvironmentConnectorParametersEnum) []string {
    convArray := make([]string,len( environmentConnectorParametersEnum))
    for i:=0; i<len(environmentConnectorParametersEnum);i++ {
        convArray[i] = EnvironmentConnectorParametersEnumToValue(environmentConnectorParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentConnectorParametersEnumFromValue(value string) EnvironmentConnectorParametersEnum {
    switch value {
        case "kVMware":
            return EnvironmentConnectorParameters_KVMWARE
        case "kHyperV":
            return EnvironmentConnectorParameters_KHYPERV
        case "kSQL":
            return EnvironmentConnectorParameters_KSQL
        case "kView":
            return EnvironmentConnectorParameters_KVIEW
        case "kPuppeteer":
            return EnvironmentConnectorParameters_KPUPPETEER
        case "kPhysical":
            return EnvironmentConnectorParameters_KPHYSICAL
        case "kPure":
            return EnvironmentConnectorParameters_KPURE
        case "kAzure":
            return EnvironmentConnectorParameters_KAZURE
        case "kNetapp":
            return EnvironmentConnectorParameters_KNETAPP
        case "kAgent":
            return EnvironmentConnectorParameters_KAGENT
        case "kGenericNas":
            return EnvironmentConnectorParameters_KGENERICNAS
        case "kAcropolis":
            return EnvironmentConnectorParameters_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentConnectorParameters_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentConnectorParameters_KISILON
        case "kGPFS":
            return EnvironmentConnectorParameters_KGPFS
        case "kKVM":
            return EnvironmentConnectorParameters_KKVM
        case "kAWS":
            return EnvironmentConnectorParameters_KAWS
        case "kExchange":
            return EnvironmentConnectorParameters_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentConnectorParameters_KHYPERVVSS
        case "kOracle":
            return EnvironmentConnectorParameters_KORACLE
        case "kGCP":
            return EnvironmentConnectorParameters_KGCP
        case "kFlashBlade":
            return EnvironmentConnectorParameters_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentConnectorParameters_KAWSNATIVE
        case "kVCD":
            return EnvironmentConnectorParameters_KVCD
        case "kO365":
            return EnvironmentConnectorParameters_KO365
        case "kO365Outlook":
            return EnvironmentConnectorParameters_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentConnectorParameters_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentConnectorParameters_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentConnectorParameters_KAZURENATIVE
        case "kKubernetes":
            return EnvironmentConnectorParameters_KKUBERNETES
        default:
            return EnvironmentConnectorParameters_KVMWARE
    }
}
