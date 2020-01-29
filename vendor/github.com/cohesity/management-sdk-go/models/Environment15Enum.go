package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment15Enum enum
 */
type Environment15Enum int

/**
 * Value collection for Environment15Enum enum
 */
const (
    Environment15_KVMWARE Environment15Enum = 1 + iota
    Environment15_KHYPERV
    Environment15_KSQL
    Environment15_KVIEW
    Environment15_KPUPPETEER
    Environment15_KPHYSICAL
    Environment15_KPURE
    Environment15_KAZURE
    Environment15_KNETAPP
    Environment15_KAGENT
    Environment15_KGENERICNAS
    Environment15_KACROPOLIS
    Environment15_KPHYSICALFILES
    Environment15_KISILON
    Environment15_KKVM
    Environment15_KAWS
    Environment15_KEXCHANGE
    Environment15_KHYPERVVSS
    Environment15_KORACLE
    Environment15_KGCP
    Environment15_KFLASHBLADE
    Environment15_KAWSNATIVE
    Environment15_KVCD
    Environment15_KO365
    Environment15_KO365OUTLOOK
    Environment15_KHYPERFLEX
    Environment15_KGCPNATIVE
    Environment15_KAZURENATIVE
)

func (r Environment15Enum) MarshalJSON() ([]byte, error) { 
    s := Environment15EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment15Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment15EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment15Enum to its string representation
 */
func Environment15EnumToValue(environment15Enum Environment15Enum) string {
    switch environment15Enum {
        case Environment15_KVMWARE:
    		return "kVMware"		
        case Environment15_KHYPERV:
    		return "kHyperV"		
        case Environment15_KSQL:
    		return "kSQL"		
        case Environment15_KVIEW:
    		return "kView"		
        case Environment15_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment15_KPHYSICAL:
    		return "kPhysical"		
        case Environment15_KPURE:
    		return "kPure"		
        case Environment15_KAZURE:
    		return "kAzure"		
        case Environment15_KNETAPP:
    		return "kNetapp"		
        case Environment15_KAGENT:
    		return "kAgent"		
        case Environment15_KGENERICNAS:
    		return "kGenericNas"		
        case Environment15_KACROPOLIS:
    		return "kAcropolis"		
        case Environment15_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment15_KISILON:
    		return "kIsilon"		
        case Environment15_KKVM:
    		return "kKVM"		
        case Environment15_KAWS:
    		return "kAWS"		
        case Environment15_KEXCHANGE:
    		return "kExchange"		
        case Environment15_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment15_KORACLE:
    		return "kOracle"		
        case Environment15_KGCP:
    		return "kGCP"		
        case Environment15_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment15_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment15_KVCD:
    		return "kVCD"		
        case Environment15_KO365:
    		return "kO365"		
        case Environment15_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment15_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment15_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment15_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment15Enum Array to its string Array representation
*/
func Environment15EnumArrayToValue(environment15Enum []Environment15Enum) []string {
    convArray := make([]string,len( environment15Enum))
    for i:=0; i<len(environment15Enum);i++ {
        convArray[i] = Environment15EnumToValue(environment15Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment15EnumFromValue(value string) Environment15Enum {
    switch value {
        case "kVMware":
            return Environment15_KVMWARE
        case "kHyperV":
            return Environment15_KHYPERV
        case "kSQL":
            return Environment15_KSQL
        case "kView":
            return Environment15_KVIEW
        case "kPuppeteer":
            return Environment15_KPUPPETEER
        case "kPhysical":
            return Environment15_KPHYSICAL
        case "kPure":
            return Environment15_KPURE
        case "kAzure":
            return Environment15_KAZURE
        case "kNetapp":
            return Environment15_KNETAPP
        case "kAgent":
            return Environment15_KAGENT
        case "kGenericNas":
            return Environment15_KGENERICNAS
        case "kAcropolis":
            return Environment15_KACROPOLIS
        case "kPhysicalFiles":
            return Environment15_KPHYSICALFILES
        case "kIsilon":
            return Environment15_KISILON
        case "kKVM":
            return Environment15_KKVM
        case "kAWS":
            return Environment15_KAWS
        case "kExchange":
            return Environment15_KEXCHANGE
        case "kHyperVVSS":
            return Environment15_KHYPERVVSS
        case "kOracle":
            return Environment15_KORACLE
        case "kGCP":
            return Environment15_KGCP
        case "kFlashBlade":
            return Environment15_KFLASHBLADE
        case "kAWSNative":
            return Environment15_KAWSNATIVE
        case "kVCD":
            return Environment15_KVCD
        case "kO365":
            return Environment15_KO365
        case "kO365Outlook":
            return Environment15_KO365OUTLOOK
        case "kHyperFlex":
            return Environment15_KHYPERFLEX
        case "kGCPNative":
            return Environment15_KGCPNATIVE
        case "kAzureNative":
            return Environment15_KAZURENATIVE
        default:
            return Environment15_KVMWARE
    }
}
