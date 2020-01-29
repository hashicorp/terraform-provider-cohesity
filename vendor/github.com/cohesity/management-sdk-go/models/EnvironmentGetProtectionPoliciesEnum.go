// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentGetProtectionPoliciesEnum enum
 */
type EnvironmentGetProtectionPoliciesEnum int

/**
 * Value collection for EnvironmentGetProtectionPoliciesEnum enum
 */
const (
    EnvironmentGetProtectionPolicies_KVMWARE EnvironmentGetProtectionPoliciesEnum = 1 + iota
    EnvironmentGetProtectionPolicies_KHYPERV
    EnvironmentGetProtectionPolicies_KSQL
    EnvironmentGetProtectionPolicies_KVIEW
    EnvironmentGetProtectionPolicies_KPUPPETEER
    EnvironmentGetProtectionPolicies_KPHYSICAL
    EnvironmentGetProtectionPolicies_KPURE
    EnvironmentGetProtectionPolicies_KAZURE
    EnvironmentGetProtectionPolicies_KNETAPP
    EnvironmentGetProtectionPolicies_KAGENT
    EnvironmentGetProtectionPolicies_KGENERICNAS
    EnvironmentGetProtectionPolicies_KACROPOLIS
    EnvironmentGetProtectionPolicies_KPHYSICALFILES
    EnvironmentGetProtectionPolicies_KISILON
    EnvironmentGetProtectionPolicies_KGPFS
    EnvironmentGetProtectionPolicies_KKVM
    EnvironmentGetProtectionPolicies_KAWS
    EnvironmentGetProtectionPolicies_KEXCHANGE
    EnvironmentGetProtectionPolicies_KHYPERVVSS
    EnvironmentGetProtectionPolicies_KORACLE
    EnvironmentGetProtectionPolicies_KGCP
    EnvironmentGetProtectionPolicies_KFLASHBLADE
    EnvironmentGetProtectionPolicies_KAWSNATIVE
    EnvironmentGetProtectionPolicies_KVCD
    EnvironmentGetProtectionPolicies_KO365
    EnvironmentGetProtectionPolicies_KO365OUTLOOK
    EnvironmentGetProtectionPolicies_KHYPERFLEX
    EnvironmentGetProtectionPolicies_KGCPNATIVE
    EnvironmentGetProtectionPolicies_KKUBERNETES
)

func (r EnvironmentGetProtectionPoliciesEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentGetProtectionPoliciesEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentGetProtectionPoliciesEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentGetProtectionPoliciesEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentGetProtectionPoliciesEnum to its string representation
 */
func EnvironmentGetProtectionPoliciesEnumToValue(environmentGetProtectionPoliciesEnum EnvironmentGetProtectionPoliciesEnum) string {
    switch environmentGetProtectionPoliciesEnum {
        case EnvironmentGetProtectionPolicies_KVMWARE:
    		return "kVMware"
        case EnvironmentGetProtectionPolicies_KHYPERV:
    		return "kHyperV"
        case EnvironmentGetProtectionPolicies_KSQL:
    		return "kSQL"
        case EnvironmentGetProtectionPolicies_KVIEW:
    		return "kView"
        case EnvironmentGetProtectionPolicies_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentGetProtectionPolicies_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentGetProtectionPolicies_KPURE:
    		return "kPure"
        case EnvironmentGetProtectionPolicies_KAZURE:
    		return "kAzure"
        case EnvironmentGetProtectionPolicies_KNETAPP:
    		return "kNetapp"
        case EnvironmentGetProtectionPolicies_KAGENT:
    		return "kAgent"
        case EnvironmentGetProtectionPolicies_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentGetProtectionPolicies_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentGetProtectionPolicies_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentGetProtectionPolicies_KISILON:
    		return "kIsilon"
        case EnvironmentGetProtectionPolicies_KGPFS:
    		return "kGPFS"
        case EnvironmentGetProtectionPolicies_KKVM:
    		return "kKVM"
        case EnvironmentGetProtectionPolicies_KAWS:
    		return "kAWS"
        case EnvironmentGetProtectionPolicies_KEXCHANGE:
    		return "kExchange"
        case EnvironmentGetProtectionPolicies_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentGetProtectionPolicies_KORACLE:
    		return "kOracle"
        case EnvironmentGetProtectionPolicies_KGCP:
    		return "kGCP"
        case EnvironmentGetProtectionPolicies_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentGetProtectionPolicies_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentGetProtectionPolicies_KVCD:
    		return "kVCD"
        case EnvironmentGetProtectionPolicies_KO365:
    		return "kO365"
        case EnvironmentGetProtectionPolicies_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentGetProtectionPolicies_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentGetProtectionPolicies_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentGetProtectionPolicies_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentGetProtectionPoliciesEnum Array to its string Array representation
*/
func EnvironmentGetProtectionPoliciesEnumArrayToValue(environmentGetProtectionPoliciesEnum []EnvironmentGetProtectionPoliciesEnum) []string {
    convArray := make([]string,len( environmentGetProtectionPoliciesEnum))
    for i:=0; i<len(environmentGetProtectionPoliciesEnum);i++ {
        convArray[i] = EnvironmentGetProtectionPoliciesEnumToValue(environmentGetProtectionPoliciesEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentGetProtectionPoliciesEnumFromValue(value string) EnvironmentGetProtectionPoliciesEnum {
    switch value {
        case "kVMware":
            return EnvironmentGetProtectionPolicies_KVMWARE
        case "kHyperV":
            return EnvironmentGetProtectionPolicies_KHYPERV
        case "kSQL":
            return EnvironmentGetProtectionPolicies_KSQL
        case "kView":
            return EnvironmentGetProtectionPolicies_KVIEW
        case "kPuppeteer":
            return EnvironmentGetProtectionPolicies_KPUPPETEER
        case "kPhysical":
            return EnvironmentGetProtectionPolicies_KPHYSICAL
        case "kPure":
            return EnvironmentGetProtectionPolicies_KPURE
        case "kAzure":
            return EnvironmentGetProtectionPolicies_KAZURE
        case "kNetapp":
            return EnvironmentGetProtectionPolicies_KNETAPP
        case "kAgent":
            return EnvironmentGetProtectionPolicies_KAGENT
        case "kGenericNas":
            return EnvironmentGetProtectionPolicies_KGENERICNAS
        case "kAcropolis":
            return EnvironmentGetProtectionPolicies_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentGetProtectionPolicies_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentGetProtectionPolicies_KISILON
        case "kGPFS":
            return EnvironmentGetProtectionPolicies_KGPFS
        case "kKVM":
            return EnvironmentGetProtectionPolicies_KKVM
        case "kAWS":
            return EnvironmentGetProtectionPolicies_KAWS
        case "kExchange":
            return EnvironmentGetProtectionPolicies_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentGetProtectionPolicies_KHYPERVVSS
        case "kOracle":
            return EnvironmentGetProtectionPolicies_KORACLE
        case "kGCP":
            return EnvironmentGetProtectionPolicies_KGCP
        case "kFlashBlade":
            return EnvironmentGetProtectionPolicies_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentGetProtectionPolicies_KAWSNATIVE
        case "kVCD":
            return EnvironmentGetProtectionPolicies_KVCD
        case "kO365":
            return EnvironmentGetProtectionPolicies_KO365
        case "kO365Outlook":
            return EnvironmentGetProtectionPolicies_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentGetProtectionPolicies_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentGetProtectionPolicies_KGCPNATIVE
        case "kKubernetes":
            return EnvironmentGetProtectionPolicies_KKUBERNETES
        default:
            return EnvironmentGetProtectionPolicies_KVMWARE
    }
}
