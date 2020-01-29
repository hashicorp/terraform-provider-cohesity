package models

import(
    "encoding/json"
)

/**
 * Type definition for EnvironmentsEnum enum
 */
type EnvironmentsEnum int

/**
 * Value collection for EnvironmentsEnum enum
 */
const (
    Environments_KVMWARE EnvironmentsEnum = 1 + iota
    Environments_KHYPERV
    Environments_KSQL
    Environments_KVIEW
    Environments_KPUPPETEER
    Environments_KPHYSICAL
    Environments_KPURE
    Environments_KAZURE
    Environments_KNETAPP
    Environments_KGENERICNAS
    Environments_KACROPOLIS
    Environments_KPHYSICALFILES
    Environments_KISILON
    Environments_KKVM
    Environments_KAWS
    Environments_KEXCHANGE
    Environments_KHYPERVVSS
    Environments_KORACLE
    Environments_KGCP
    Environments_KFLASHBLADE
    Environments_KAWSNATIVE
    Environments_KVCD
    Environments_KO365
    Environments_KO365OUTLOOK
    Environments_KHYPERFLEX
    Environments_KGCPNATIVE
)

func (r EnvironmentsEnum) MarshalJSON() ([]byte, error) { 
    s := EnvironmentsEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *EnvironmentsEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  EnvironmentsEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts EnvironmentsEnum to its string representation
 */
func EnvironmentsEnumToValue(environmentsEnum EnvironmentsEnum) string {
    switch environmentsEnum {
        case Environments_KVMWARE:
    		return "kVMware"		
        case Environments_KHYPERV:
    		return "kHyperV"		
        case Environments_KSQL:
    		return "kSQL"		
        case Environments_KVIEW:
    		return "kView"		
        case Environments_KPUPPETEER:
    		return "kPuppeteer"		
        case Environments_KPHYSICAL:
    		return "kPhysical"		
        case Environments_KPURE:
    		return "kPure"		
        case Environments_KAZURE:
    		return "kAzure"		
        case Environments_KNETAPP:
    		return "kNetapp"		
        case Environments_KGENERICNAS:
    		return "kGenericNas"		
        case Environments_KACROPOLIS:
    		return "kAcropolis"		
        case Environments_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environments_KISILON:
    		return "kIsilon"		
        case Environments_KKVM:
    		return "kKVM"		
        case Environments_KAWS:
    		return "kAWS"		
        case Environments_KEXCHANGE:
    		return "kExchange"		
        case Environments_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environments_KORACLE:
    		return "kOracle"		
        case Environments_KGCP:
    		return "kGCP"		
        case Environments_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environments_KAWSNATIVE:
    		return "kAWSNative"		
        case Environments_KVCD:
    		return "kVCD"		
        case Environments_KO365:
    		return "kO365"		
        case Environments_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environments_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environments_KGCPNATIVE:
    		return "kGCPNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts EnvironmentsEnum Array to its string Array representation
*/
func EnvironmentsEnumArrayToValue(environmentsEnum []EnvironmentsEnum) []string {
    convArray := make([]string,len( environmentsEnum))
    for i:=0; i<len(environmentsEnum);i++ {
        convArray[i] = EnvironmentsEnumToValue(environmentsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EnvironmentsEnumFromValue(value string) EnvironmentsEnum {
    switch value {
        case "kVMware":
            return Environments_KVMWARE
        case "kHyperV":
            return Environments_KHYPERV
        case "kSQL":
            return Environments_KSQL
        case "kView":
            return Environments_KVIEW
        case "kPuppeteer":
            return Environments_KPUPPETEER
        case "kPhysical":
            return Environments_KPHYSICAL
        case "kPure":
            return Environments_KPURE
        case "kAzure":
            return Environments_KAZURE
        case "kNetapp":
            return Environments_KNETAPP
        case "kGenericNas":
            return Environments_KGENERICNAS
        case "kAcropolis":
            return Environments_KACROPOLIS
        case "kPhysicalFiles":
            return Environments_KPHYSICALFILES
        case "kIsilon":
            return Environments_KISILON
        case "kKVM":
            return Environments_KKVM
        case "kAWS":
            return Environments_KAWS
        case "kExchange":
            return Environments_KEXCHANGE
        case "kHyperVVSS":
            return Environments_KHYPERVVSS
        case "kOracle":
            return Environments_KORACLE
        case "kGCP":
            return Environments_KGCP
        case "kFlashBlade":
            return Environments_KFLASHBLADE
        case "kAWSNative":
            return Environments_KAWSNATIVE
        case "kVCD":
            return Environments_KVCD
        case "kO365":
            return Environments_KO365
        case "kO365Outlook":
            return Environments_KO365OUTLOOK
        case "kHyperFlex":
            return Environments_KHYPERFLEX
        case "kGCPNative":
            return Environments_KGCPNATIVE
        default:
            return Environments_KVMWARE
    }
}
