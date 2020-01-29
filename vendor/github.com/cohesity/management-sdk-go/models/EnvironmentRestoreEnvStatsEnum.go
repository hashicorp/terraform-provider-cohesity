// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentRestoreEnvStatsEnum enum
 */
type EnvironmentRestoreEnvStatsEnum int

/**
 * Value collection for EnvironmentRestoreEnvStatsEnum enum
 */
const (
    EnvironmentRestoreEnvStats_KVMWARE EnvironmentRestoreEnvStatsEnum = 1 + iota
    EnvironmentRestoreEnvStats_KHYPERV
    EnvironmentRestoreEnvStats_KSQL
    EnvironmentRestoreEnvStats_KVIEW
    EnvironmentRestoreEnvStats_KPUPPETEER
    EnvironmentRestoreEnvStats_KPHYSICAL
    EnvironmentRestoreEnvStats_KPURE
    EnvironmentRestoreEnvStats_KAZURE
    EnvironmentRestoreEnvStats_KNETAPP
    EnvironmentRestoreEnvStats_KAGENT
    EnvironmentRestoreEnvStats_KGENERICNAS
    EnvironmentRestoreEnvStats_KACROPOLIS
    EnvironmentRestoreEnvStats_KPHYSICALFILES
    EnvironmentRestoreEnvStats_KISILON
    EnvironmentRestoreEnvStats_KGPFS
    EnvironmentRestoreEnvStats_KKVM
    EnvironmentRestoreEnvStats_KAWS
    EnvironmentRestoreEnvStats_KEXCHANGE
    EnvironmentRestoreEnvStats_KHYPERVVSS
    EnvironmentRestoreEnvStats_KORACLE
    EnvironmentRestoreEnvStats_KGCP
    EnvironmentRestoreEnvStats_KFLASHBLADE
    EnvironmentRestoreEnvStats_KAWSNATIVE
    EnvironmentRestoreEnvStats_KVCD
    EnvironmentRestoreEnvStats_KO365
    EnvironmentRestoreEnvStats_KO365OUTLOOK
    EnvironmentRestoreEnvStats_KHYPERFLEX
    EnvironmentRestoreEnvStats_KGCPNATIVE
    EnvironmentRestoreEnvStats_KAZURENATIVE
    EnvironmentRestoreEnvStats_KAD
    EnvironmentRestoreEnvStats_KAWSSNAPSHOTMANAGER
)

func (r EnvironmentRestoreEnvStatsEnum) MarshalJSON() ([]byte, error) {
    s := EnvironmentRestoreEnvStatsEnumToValue(r)
    return json.Marshal(s)
}

func (r *EnvironmentRestoreEnvStatsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EnvironmentRestoreEnvStatsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EnvironmentRestoreEnvStatsEnum to its string representation
 */
func EnvironmentRestoreEnvStatsEnumToValue(environmentRestoreEnvStatsEnum EnvironmentRestoreEnvStatsEnum) string {
    switch environmentRestoreEnvStatsEnum {
        case EnvironmentRestoreEnvStats_KVMWARE:
    		return "kVMware"
        case EnvironmentRestoreEnvStats_KHYPERV:
    		return "kHyperV"
        case EnvironmentRestoreEnvStats_KSQL:
    		return "kSQL"
        case EnvironmentRestoreEnvStats_KVIEW:
    		return "kView"
        case EnvironmentRestoreEnvStats_KPUPPETEER:
    		return "kPuppeteer"
        case EnvironmentRestoreEnvStats_KPHYSICAL:
    		return "kPhysical"
        case EnvironmentRestoreEnvStats_KPURE:
    		return "kPure"
        case EnvironmentRestoreEnvStats_KAZURE:
    		return "kAzure"
        case EnvironmentRestoreEnvStats_KNETAPP:
    		return "kNetapp"
        case EnvironmentRestoreEnvStats_KAGENT:
    		return "kAgent"
        case EnvironmentRestoreEnvStats_KGENERICNAS:
    		return "kGenericNas"
        case EnvironmentRestoreEnvStats_KACROPOLIS:
    		return "kAcropolis"
        case EnvironmentRestoreEnvStats_KPHYSICALFILES:
    		return "kPhysicalFiles"
        case EnvironmentRestoreEnvStats_KISILON:
    		return "kIsilon"
        case EnvironmentRestoreEnvStats_KGPFS:
    		return "kGPFS"
        case EnvironmentRestoreEnvStats_KKVM:
    		return "kKVM"
        case EnvironmentRestoreEnvStats_KAWS:
    		return "kAWS"
        case EnvironmentRestoreEnvStats_KEXCHANGE:
    		return "kExchange"
        case EnvironmentRestoreEnvStats_KHYPERVVSS:
    		return "kHyperVVSS"
        case EnvironmentRestoreEnvStats_KORACLE:
    		return "kOracle"
        case EnvironmentRestoreEnvStats_KGCP:
    		return "kGCP"
        case EnvironmentRestoreEnvStats_KFLASHBLADE:
    		return "kFlashBlade"
        case EnvironmentRestoreEnvStats_KAWSNATIVE:
    		return "kAWSNative"
        case EnvironmentRestoreEnvStats_KVCD:
    		return "kVCD"
        case EnvironmentRestoreEnvStats_KO365:
    		return "kO365"
        case EnvironmentRestoreEnvStats_KO365OUTLOOK:
    		return "kO365Outlook"
        case EnvironmentRestoreEnvStats_KHYPERFLEX:
    		return "kHyperFlex"
        case EnvironmentRestoreEnvStats_KGCPNATIVE:
    		return "kGCPNative"
        case EnvironmentRestoreEnvStats_KAZURENATIVE:
    		return "kAzureNative"
        case EnvironmentRestoreEnvStats_KAD:
    		return "kAD"
        case EnvironmentRestoreEnvStats_KAWSSNAPSHOTMANAGER:
    		return "kAWSSnapshotManager"
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentRestoreEnvStatsEnum Array to its string Array representation
*/
func EnvironmentRestoreEnvStatsEnumArrayToValue(environmentRestoreEnvStatsEnum []EnvironmentRestoreEnvStatsEnum) []string {
    convArray := make([]string,len( environmentRestoreEnvStatsEnum))
    for i:=0; i<len(environmentRestoreEnvStatsEnum);i++ {
        convArray[i] = EnvironmentRestoreEnvStatsEnumToValue(environmentRestoreEnvStatsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentRestoreEnvStatsEnumFromValue(value string) EnvironmentRestoreEnvStatsEnum {
    switch value {
        case "kVMware":
            return EnvironmentRestoreEnvStats_KVMWARE
        case "kHyperV":
            return EnvironmentRestoreEnvStats_KHYPERV
        case "kSQL":
            return EnvironmentRestoreEnvStats_KSQL
        case "kView":
            return EnvironmentRestoreEnvStats_KVIEW
        case "kPuppeteer":
            return EnvironmentRestoreEnvStats_KPUPPETEER
        case "kPhysical":
            return EnvironmentRestoreEnvStats_KPHYSICAL
        case "kPure":
            return EnvironmentRestoreEnvStats_KPURE
        case "kAzure":
            return EnvironmentRestoreEnvStats_KAZURE
        case "kNetapp":
            return EnvironmentRestoreEnvStats_KNETAPP
        case "kAgent":
            return EnvironmentRestoreEnvStats_KAGENT
        case "kGenericNas":
            return EnvironmentRestoreEnvStats_KGENERICNAS
        case "kAcropolis":
            return EnvironmentRestoreEnvStats_KACROPOLIS
        case "kPhysicalFiles":
            return EnvironmentRestoreEnvStats_KPHYSICALFILES
        case "kIsilon":
            return EnvironmentRestoreEnvStats_KISILON
        case "kGPFS":
            return EnvironmentRestoreEnvStats_KGPFS
        case "kKVM":
            return EnvironmentRestoreEnvStats_KKVM
        case "kAWS":
            return EnvironmentRestoreEnvStats_KAWS
        case "kExchange":
            return EnvironmentRestoreEnvStats_KEXCHANGE
        case "kHyperVVSS":
            return EnvironmentRestoreEnvStats_KHYPERVVSS
        case "kOracle":
            return EnvironmentRestoreEnvStats_KORACLE
        case "kGCP":
            return EnvironmentRestoreEnvStats_KGCP
        case "kFlashBlade":
            return EnvironmentRestoreEnvStats_KFLASHBLADE
        case "kAWSNative":
            return EnvironmentRestoreEnvStats_KAWSNATIVE
        case "kVCD":
            return EnvironmentRestoreEnvStats_KVCD
        case "kO365":
            return EnvironmentRestoreEnvStats_KO365
        case "kO365Outlook":
            return EnvironmentRestoreEnvStats_KO365OUTLOOK
        case "kHyperFlex":
            return EnvironmentRestoreEnvStats_KHYPERFLEX
        case "kGCPNative":
            return EnvironmentRestoreEnvStats_KGCPNATIVE
        case "kAzureNative":
            return EnvironmentRestoreEnvStats_KAZURENATIVE
        case "kAD":
            return EnvironmentRestoreEnvStats_KAD
        case "kAWSSnapshotManager":
            return EnvironmentRestoreEnvStats_KAWSSNAPSHOTMANAGER
        default:
            return EnvironmentRestoreEnvStats_KVMWARE
    }
}
