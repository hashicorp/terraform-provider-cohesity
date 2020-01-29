// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeProtectionJobInfoEnum enum
 */
type TypeProtectionJobInfoEnum int

/**
 * Value collection for TypeProtectionJobInfoEnum enum
 */
const (
    TypeProtectionJobInfo_KVMWARE TypeProtectionJobInfoEnum = 1 + iota
    TypeProtectionJobInfo_KHYPERV
    TypeProtectionJobInfo_KSQL
    TypeProtectionJobInfo_KVIEW
    TypeProtectionJobInfo_KPUPPETEER
    TypeProtectionJobInfo_KPHYSICAL
    TypeProtectionJobInfo_KPURE
    TypeProtectionJobInfo_KAZURE
    TypeProtectionJobInfo_KNETAPP
    TypeProtectionJobInfo_KAGENT
    TypeProtectionJobInfo_KGENERICNAS
    TypeProtectionJobInfo_KACROPOLIS
    TypeProtectionJobInfo_KPHYSICALFILES
    TypeProtectionJobInfo_KISILON
    TypeProtectionJobInfo_KGPFS
    TypeProtectionJobInfo_KKVM
    TypeProtectionJobInfo_KAWS
    TypeProtectionJobInfo_KEXCHANGE
    TypeProtectionJobInfo_KHYPERVVSS
    TypeProtectionJobInfo_KORACLE
    TypeProtectionJobInfo_KGCP
    TypeProtectionJobInfo_KFLASHBLADE
    TypeProtectionJobInfo_KAWSNATIVE
    TypeProtectionJobInfo_KVCD
    TypeProtectionJobInfo_KO365
    TypeProtectionJobInfo_KO365OUTLOOK
    TypeProtectionJobInfo_KHYPERFLEX
    TypeProtectionJobInfo_KGCPNATIVE
    TypeProtectionJobInfo_KAZURENATIVE
    TypeProtectionJobInfo_KKUBERNETES
)

func (r TypeProtectionJobInfoEnum) MarshalJSON() ([]byte, error) {
    s := TypeProtectionJobInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeProtectionJobInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeProtectionJobInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeProtectionJobInfoEnum to its string representation
 */
func TypeProtectionJobInfoEnumToValue(typeProtectionJobInfoEnum TypeProtectionJobInfoEnum) string {
    switch typeProtectionJobInfoEnum {
        case TypeProtectionJobInfo_KVMWARE:
    		return "kVMware"
        case TypeProtectionJobInfo_KHYPERV:
    		return "kHyperV"
        case TypeProtectionJobInfo_KSQL:
    		return "kSQL"
        case TypeProtectionJobInfo_KVIEW:
    		return "kView"
        case TypeProtectionJobInfo_KPUPPETEER:
    		return "kPuppeteer"
        case TypeProtectionJobInfo_KPHYSICAL:
    		return "kPhysical"
        case TypeProtectionJobInfo_KPURE:
    		return "kPure"
        case TypeProtectionJobInfo_KAZURE:
    		return "kAzure"
        case TypeProtectionJobInfo_KNETAPP:
    		return "kNetapp"
        case TypeProtectionJobInfo_KAGENT:
    		return "kAgent"
        case TypeProtectionJobInfo_KGENERICNAS:
    		return "kGenericNas"
        case TypeProtectionJobInfo_KACROPOLIS:
    		return "kAcropolis"
        case TypeProtectionJobInfo_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case TypeProtectionJobInfo_KISILON:
    		return "kIsilon"
        case TypeProtectionJobInfo_KGPFS:
    		return "kGPFS"
        case TypeProtectionJobInfo_KKVM:
    		return "kKVM"
        case TypeProtectionJobInfo_KAWS:
    		return "kAWS"
        case TypeProtectionJobInfo_KEXCHANGE:
    		return "kExchange"
        case TypeProtectionJobInfo_KHYPERVVSS:
    		return "kHyperVVSS"
        case TypeProtectionJobInfo_KORACLE:
    		return "kOracle"
        case TypeProtectionJobInfo_KGCP:
    		return "kGCP"
        case TypeProtectionJobInfo_KFLASHBLADE:
    		return "kFlashBlade"
        case TypeProtectionJobInfo_KAWSNATIVE:
    		return "kAWSNative"
        case TypeProtectionJobInfo_KVCD:
    		return "kVCD"
        case TypeProtectionJobInfo_KO365:
    		return "kO365"
        case TypeProtectionJobInfo_KO365OUTLOOK:
    		return "kO365Outlook"
        case TypeProtectionJobInfo_KHYPERFLEX:
    		return "kHyperFlex"
        case TypeProtectionJobInfo_KGCPNATIVE:
    		return "kGCPNative"
        case TypeProtectionJobInfo_KAZURENATIVE:
    		return "kAzureNative"
        case TypeProtectionJobInfo_KKUBERNETES:
    		return "kKubernetes"
        default:
        	return "kVMware"
    }
}

/**
 * Converts TypeProtectionJobInfoEnum Array to its string Array representation
*/
func TypeProtectionJobInfoEnumArrayToValue(typeProtectionJobInfoEnum []TypeProtectionJobInfoEnum) []string {
    convArray := make([]string,len( typeProtectionJobInfoEnum))
    for i:=0; i<len(typeProtectionJobInfoEnum);i++ {
        convArray[i] = TypeProtectionJobInfoEnumToValue(typeProtectionJobInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeProtectionJobInfoEnumFromValue(value string) TypeProtectionJobInfoEnum {
    switch value {
        case "kVMware":
            return TypeProtectionJobInfo_KVMWARE
        case "kHyperV":
            return TypeProtectionJobInfo_KHYPERV
        case "kSQL":
            return TypeProtectionJobInfo_KSQL
        case "kView":
            return TypeProtectionJobInfo_KVIEW
        case "kPuppeteer":
            return TypeProtectionJobInfo_KPUPPETEER
        case "kPhysical":
            return TypeProtectionJobInfo_KPHYSICAL
        case "kPure":
            return TypeProtectionJobInfo_KPURE
        case "kAzure":
            return TypeProtectionJobInfo_KAZURE
        case "kNetapp":
            return TypeProtectionJobInfo_KNETAPP
        case "kAgent":
            return TypeProtectionJobInfo_KAGENT
        case "kGenericNas":
            return TypeProtectionJobInfo_KGENERICNAS
        case "kAcropolis":
            return TypeProtectionJobInfo_KACROPOLIS
        case "kPhysicalFiles":
            return TypeProtectionJobInfo_KPHYSICALFILES
        case "kIsilon":
            return TypeProtectionJobInfo_KISILON
        case "kGPFS":
            return TypeProtectionJobInfo_KGPFS
        case "kKVM":
            return TypeProtectionJobInfo_KKVM
        case "kAWS":
            return TypeProtectionJobInfo_KAWS
        case "kExchange":
            return TypeProtectionJobInfo_KEXCHANGE
        case "kHyperVVSS":
            return TypeProtectionJobInfo_KHYPERVVSS
        case "kOracle":
            return TypeProtectionJobInfo_KORACLE
        case "kGCP":
            return TypeProtectionJobInfo_KGCP
        case "kFlashBlade":
            return TypeProtectionJobInfo_KFLASHBLADE
        case "kAWSNative":
            return TypeProtectionJobInfo_KAWSNATIVE
        case "kVCD":
            return TypeProtectionJobInfo_KVCD
        case "kO365":
            return TypeProtectionJobInfo_KO365
        case "kO365Outlook":
            return TypeProtectionJobInfo_KO365OUTLOOK
        case "kHyperFlex":
            return TypeProtectionJobInfo_KHYPERFLEX
        case "kGCPNative":
            return TypeProtectionJobInfo_KGCPNATIVE
        case "kAzureNative":
            return TypeProtectionJobInfo_KAZURENATIVE
        case "kKubernetes":
            return TypeProtectionJobInfo_KKUBERNETES
        default:
            return TypeProtectionJobInfo_KVMWARE
    }
}
