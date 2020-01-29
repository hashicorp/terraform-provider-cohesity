package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment1Enum enum
 */
type Environment1Enum int

/**
 * Value collection for Environment1Enum enum
 */
const (
    Environment1_KVMWARE Environment1Enum = 1 + iota
    Environment1_KHYPERV
    Environment1_KSQL
    Environment1_KVIEW
    Environment1_KPUPPETEER
    Environment1_KPHYSICAL
    Environment1_KPURE
    Environment1_KAZURE
    Environment1_KNETAPP
    Environment1_KAGENT
    Environment1_KGENERICNAS
    Environment1_KACROPOLIS
    Environment1_KPHYSICALFILES
    Environment1_KISILON
    Environment1_KKVM
    Environment1_KAWS
    Environment1_KEXCHANGE
    Environment1_KHYPERVVSS
    Environment1_KORACLE
    Environment1_KGCP
    Environment1_KFLASHBLADE
    Environment1_KAWSNATIVE
    Environment1_KVCD
    Environment1_KO365
    Environment1_KO365OUTLOOK
    Environment1_KHYPERFLEX
    Environment1_KGCPNATIVE
    Environment1_KAZURENATIVE
)

func (r Environment1Enum) MarshalJSON() ([]byte, error) { 
    s := Environment1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment1Enum to its string representation
 */
func Environment1EnumToValue(environment1Enum Environment1Enum) string {
    switch environment1Enum {
        case Environment1_KVMWARE:
    		return "kVMware"		
        case Environment1_KHYPERV:
    		return "kHyperV"		
        case Environment1_KSQL:
    		return "kSQL"		
        case Environment1_KVIEW:
    		return "kView"		
        case Environment1_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment1_KPHYSICAL:
    		return "kPhysical"		
        case Environment1_KPURE:
    		return "kPure"		
        case Environment1_KAZURE:
    		return "kAzure"		
        case Environment1_KNETAPP:
    		return "kNetapp"		
        case Environment1_KAGENT:
    		return "kAgent"		
        case Environment1_KGENERICNAS:
    		return "kGenericNas"		
        case Environment1_KACROPOLIS:
    		return "kAcropolis"		
        case Environment1_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment1_KISILON:
    		return "kIsilon"		
        case Environment1_KKVM:
    		return "kKVM"		
        case Environment1_KAWS:
    		return "kAWS"		
        case Environment1_KEXCHANGE:
    		return "kExchange"		
        case Environment1_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment1_KORACLE:
    		return "kOracle"		
        case Environment1_KGCP:
    		return "kGCP"		
        case Environment1_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment1_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment1_KVCD:
    		return "kVCD"		
        case Environment1_KO365:
    		return "kO365"		
        case Environment1_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment1_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment1_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment1_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment1Enum Array to its string Array representation
*/
func Environment1EnumArrayToValue(environment1Enum []Environment1Enum) []string {
    convArray := make([]string,len( environment1Enum))
    for i:=0; i<len(environment1Enum);i++ {
        convArray[i] = Environment1EnumToValue(environment1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment1EnumFromValue(value string) Environment1Enum {
    switch value {
        case "kVMware":
            return Environment1_KVMWARE
        case "kHyperV":
            return Environment1_KHYPERV
        case "kSQL":
            return Environment1_KSQL
        case "kView":
            return Environment1_KVIEW
        case "kPuppeteer":
            return Environment1_KPUPPETEER
        case "kPhysical":
            return Environment1_KPHYSICAL
        case "kPure":
            return Environment1_KPURE
        case "kAzure":
            return Environment1_KAZURE
        case "kNetapp":
            return Environment1_KNETAPP
        case "kAgent":
            return Environment1_KAGENT
        case "kGenericNas":
            return Environment1_KGENERICNAS
        case "kAcropolis":
            return Environment1_KACROPOLIS
        case "kPhysicalFiles":
            return Environment1_KPHYSICALFILES
        case "kIsilon":
            return Environment1_KISILON
        case "kKVM":
            return Environment1_KKVM
        case "kAWS":
            return Environment1_KAWS
        case "kExchange":
            return Environment1_KEXCHANGE
        case "kHyperVVSS":
            return Environment1_KHYPERVVSS
        case "kOracle":
            return Environment1_KORACLE
        case "kGCP":
            return Environment1_KGCP
        case "kFlashBlade":
            return Environment1_KFLASHBLADE
        case "kAWSNative":
            return Environment1_KAWSNATIVE
        case "kVCD":
            return Environment1_KVCD
        case "kO365":
            return Environment1_KO365
        case "kO365Outlook":
            return Environment1_KO365OUTLOOK
        case "kHyperFlex":
            return Environment1_KHYPERFLEX
        case "kGCPNative":
            return Environment1_KGCPNATIVE
        case "kAzureNative":
            return Environment1_KAZURENATIVE
        default:
            return Environment1_KVMWARE
    }
}
