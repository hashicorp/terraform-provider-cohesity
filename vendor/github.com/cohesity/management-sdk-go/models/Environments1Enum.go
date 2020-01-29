package models

import(
    "encoding/json"
)

/**
 * Type definition for Environments1Enum enum
 */
type Environments1Enum int

/**
 * Value collection for Environments1Enum enum
 */
const (
    Environments1_KVMWARE Environments1Enum = 1 + iota
    Environments1_KHYPERV
    Environments1_KSQL
    Environments1_KVIEW
    Environments1_KPUPPETEER
    Environments1_KPHYSICAL
    Environments1_KPURE
    Environments1_KAZURE
    Environments1_KNETAPP
    Environments1_KAGENT
    Environments1_KGENERICNAS
    Environments1_KACROPOLIS
    Environments1_KPHYSICALFILES
    Environments1_KISILON
    Environments1_KKVM
    Environments1_KAWS
    Environments1_KEXCHANGE
    Environments1_KHYPERVVSS
    Environments1_KORACLE
    Environments1_KGCP
    Environments1_KFLASHBLADE
    Environments1_KAWSNATIVE
    Environments1_KVCD
    Environments1_KO365
    Environments1_KO365OUTLOOK
    Environments1_KHYPERFLEX
    Environments1_KGCPNATIVE
)

func (r Environments1Enum) MarshalJSON() ([]byte, error) { 
    s := Environments1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environments1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environments1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environments1Enum to its string representation
 */
func Environments1EnumToValue(environments1Enum Environments1Enum) string {
    switch environments1Enum {
        case Environments1_KVMWARE:
    		return "kVMware"		
        case Environments1_KHYPERV:
    		return "kHyperV"		
        case Environments1_KSQL:
    		return "kSQL"		
        case Environments1_KVIEW:
    		return "kView"		
        case Environments1_KPUPPETEER:
    		return "kPuppeteer"		
        case Environments1_KPHYSICAL:
    		return "kPhysical"		
        case Environments1_KPURE:
    		return "kPure"		
        case Environments1_KAZURE:
    		return "kAzure"		
        case Environments1_KNETAPP:
    		return "kNetapp"		
        case Environments1_KAGENT:
    		return "kAgent"		
        case Environments1_KGENERICNAS:
    		return "kGenericNas"		
        case Environments1_KACROPOLIS:
    		return "kAcropolis"		
        case Environments1_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environments1_KISILON:
    		return "kIsilon"		
        case Environments1_KKVM:
    		return "kKVM"		
        case Environments1_KAWS:
    		return "kAWS"		
        case Environments1_KEXCHANGE:
    		return "kExchange"		
        case Environments1_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environments1_KORACLE:
    		return "kOracle"		
        case Environments1_KGCP:
    		return "kGCP"		
        case Environments1_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environments1_KAWSNATIVE:
    		return "kAWSNative"		
        case Environments1_KVCD:
    		return "kVCD"		
        case Environments1_KO365:
    		return "kO365"		
        case Environments1_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environments1_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environments1_KGCPNATIVE:
    		return "kGCPNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environments1Enum Array to its string Array representation
*/
func Environments1EnumArrayToValue(environments1Enum []Environments1Enum) []string {
    convArray := make([]string,len( environments1Enum))
    for i:=0; i<len(environments1Enum);i++ {
        convArray[i] = Environments1EnumToValue(environments1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environments1EnumFromValue(value string) Environments1Enum {
    switch value {
        case "kVMware":
            return Environments1_KVMWARE
        case "kHyperV":
            return Environments1_KHYPERV
        case "kSQL":
            return Environments1_KSQL
        case "kView":
            return Environments1_KVIEW
        case "kPuppeteer":
            return Environments1_KPUPPETEER
        case "kPhysical":
            return Environments1_KPHYSICAL
        case "kPure":
            return Environments1_KPURE
        case "kAzure":
            return Environments1_KAZURE
        case "kNetapp":
            return Environments1_KNETAPP
        case "kAgent":
            return Environments1_KAGENT
        case "kGenericNas":
            return Environments1_KGENERICNAS
        case "kAcropolis":
            return Environments1_KACROPOLIS
        case "kPhysicalFiles":
            return Environments1_KPHYSICALFILES
        case "kIsilon":
            return Environments1_KISILON
        case "kKVM":
            return Environments1_KKVM
        case "kAWS":
            return Environments1_KAWS
        case "kExchange":
            return Environments1_KEXCHANGE
        case "kHyperVVSS":
            return Environments1_KHYPERVVSS
        case "kOracle":
            return Environments1_KORACLE
        case "kGCP":
            return Environments1_KGCP
        case "kFlashBlade":
            return Environments1_KFLASHBLADE
        case "kAWSNative":
            return Environments1_KAWSNATIVE
        case "kVCD":
            return Environments1_KVCD
        case "kO365":
            return Environments1_KO365
        case "kO365Outlook":
            return Environments1_KO365OUTLOOK
        case "kHyperFlex":
            return Environments1_KHYPERFLEX
        case "kGCPNative":
            return Environments1_KGCPNATIVE
        default:
            return Environments1_KVMWARE
    }
}
