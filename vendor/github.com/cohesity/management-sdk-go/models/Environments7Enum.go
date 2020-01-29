package models

import(
    "encoding/json"
)

/**
 * Type definition for Environments7Enum enum
 */
type Environments7Enum int

/**
 * Value collection for Environments7Enum enum
 */
const (
    Environments7_KVMWARE Environments7Enum = 1 + iota
    Environments7_KHYPERV
    Environments7_KSQL
    Environments7_KVIEW
    Environments7_KPUPPETEER
    Environments7_KPHYSICAL
    Environments7_KPURE
    Environments7_KAZURE
    Environments7_KNETAPP
    Environments7_KAGENT
    Environments7_KGENERICNAS
    Environments7_KACROPOLIS
    Environments7_KPHYSICALFILES
    Environments7_KISILON
    Environments7_KKVM
    Environments7_KAWS
    Environments7_KEXCHANGE
    Environments7_KHYPERVVSS
    Environments7_KORACLE
    Environments7_KGCP
    Environments7_KFLASHBLADE
    Environments7_KAWSNATIVE
    Environments7_KVCD
    Environments7_KO365
    Environments7_KO365OUTLOOK
    Environments7_KHYPERFLEX
    Environments7_KGCPNATIVE
    Environments7_KAZURENATIVE
)

func (r Environments7Enum) MarshalJSON() ([]byte, error) { 
    s := Environments7EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environments7Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environments7EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environments7Enum to its string representation
 */
func Environments7EnumToValue(environments7Enum Environments7Enum) string {
    switch environments7Enum {
        case Environments7_KVMWARE:
    		return "kVMware"		
        case Environments7_KHYPERV:
    		return "kHyperV"		
        case Environments7_KSQL:
    		return "kSQL"		
        case Environments7_KVIEW:
    		return "kView"		
        case Environments7_KPUPPETEER:
    		return "kPuppeteer"		
        case Environments7_KPHYSICAL:
    		return "kPhysical"		
        case Environments7_KPURE:
    		return "kPure"		
        case Environments7_KAZURE:
    		return "kAzure"		
        case Environments7_KNETAPP:
    		return "kNetapp"		
        case Environments7_KAGENT:
    		return "kAgent"		
        case Environments7_KGENERICNAS:
    		return "kGenericNas"		
        case Environments7_KACROPOLIS:
    		return "kAcropolis"		
        case Environments7_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environments7_KISILON:
    		return "kIsilon"		
        case Environments7_KKVM:
    		return "kKVM"		
        case Environments7_KAWS:
    		return "kAWS"		
        case Environments7_KEXCHANGE:
    		return "kExchange"		
        case Environments7_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environments7_KORACLE:
    		return "kOracle"		
        case Environments7_KGCP:
    		return "kGCP"		
        case Environments7_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environments7_KAWSNATIVE:
    		return "kAWSNative"		
        case Environments7_KVCD:
    		return "kVCD"		
        case Environments7_KO365:
    		return "kO365"		
        case Environments7_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environments7_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environments7_KGCPNATIVE:
    		return "kGCPNative"		
        case Environments7_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environments7Enum Array to its string Array representation
*/
func Environments7EnumArrayToValue(environments7Enum []Environments7Enum) []string {
    convArray := make([]string,len( environments7Enum))
    for i:=0; i<len(environments7Enum);i++ {
        convArray[i] = Environments7EnumToValue(environments7Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environments7EnumFromValue(value string) Environments7Enum {
    switch value {
        case "kVMware":
            return Environments7_KVMWARE
        case "kHyperV":
            return Environments7_KHYPERV
        case "kSQL":
            return Environments7_KSQL
        case "kView":
            return Environments7_KVIEW
        case "kPuppeteer":
            return Environments7_KPUPPETEER
        case "kPhysical":
            return Environments7_KPHYSICAL
        case "kPure":
            return Environments7_KPURE
        case "kAzure":
            return Environments7_KAZURE
        case "kNetapp":
            return Environments7_KNETAPP
        case "kAgent":
            return Environments7_KAGENT
        case "kGenericNas":
            return Environments7_KGENERICNAS
        case "kAcropolis":
            return Environments7_KACROPOLIS
        case "kPhysicalFiles":
            return Environments7_KPHYSICALFILES
        case "kIsilon":
            return Environments7_KISILON
        case "kKVM":
            return Environments7_KKVM
        case "kAWS":
            return Environments7_KAWS
        case "kExchange":
            return Environments7_KEXCHANGE
        case "kHyperVVSS":
            return Environments7_KHYPERVVSS
        case "kOracle":
            return Environments7_KORACLE
        case "kGCP":
            return Environments7_KGCP
        case "kFlashBlade":
            return Environments7_KFLASHBLADE
        case "kAWSNative":
            return Environments7_KAWSNATIVE
        case "kVCD":
            return Environments7_KVCD
        case "kO365":
            return Environments7_KO365
        case "kO365Outlook":
            return Environments7_KO365OUTLOOK
        case "kHyperFlex":
            return Environments7_KHYPERFLEX
        case "kGCPNative":
            return Environments7_KGCPNATIVE
        case "kAzureNative":
            return Environments7_KAZURENATIVE
        default:
            return Environments7_KVMWARE
    }
}
