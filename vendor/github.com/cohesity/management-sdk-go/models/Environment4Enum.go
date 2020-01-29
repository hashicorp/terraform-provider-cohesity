package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment4Enum enum
 */
type Environment4Enum int

/**
 * Value collection for Environment4Enum enum
 */
const (
    Environment4_KVMWARE Environment4Enum = 1 + iota
    Environment4_KHYPERV
    Environment4_KSQL
    Environment4_KVIEW
    Environment4_KPUPPETEER
    Environment4_KPHYSICAL
    Environment4_KPURE
    Environment4_KAZURE
    Environment4_KNETAPP
    Environment4_KAGENT
    Environment4_KGENERICNAS
    Environment4_KACROPOLIS
    Environment4_KPHYSICALFILES
    Environment4_KISILON
    Environment4_KKVM
    Environment4_KAWS
    Environment4_KEXCHANGE
    Environment4_KHYPERVVSS
    Environment4_KORACLE
    Environment4_KGCP
    Environment4_KFLASHBLADE
    Environment4_KAWSNATIVE
    Environment4_KVCD
    Environment4_KO365
    Environment4_KO365OUTLOOK
    Environment4_KHYPERFLEX
    Environment4_KGCPNATIVE
    Environment4_KAZURENATIVE
)

func (r Environment4Enum) MarshalJSON() ([]byte, error) { 
    s := Environment4EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment4Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment4EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment4Enum to its string representation
 */
func Environment4EnumToValue(environment4Enum Environment4Enum) string {
    switch environment4Enum {
        case Environment4_KVMWARE:
    		return "kVMware"		
        case Environment4_KHYPERV:
    		return "kHyperV"		
        case Environment4_KSQL:
    		return "kSQL"		
        case Environment4_KVIEW:
    		return "kView"		
        case Environment4_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment4_KPHYSICAL:
    		return "kPhysical"		
        case Environment4_KPURE:
    		return "kPure"		
        case Environment4_KAZURE:
    		return "kAzure"		
        case Environment4_KNETAPP:
    		return "kNetapp"		
        case Environment4_KAGENT:
    		return "kAgent"		
        case Environment4_KGENERICNAS:
    		return "kGenericNas"		
        case Environment4_KACROPOLIS:
    		return "kAcropolis"		
        case Environment4_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment4_KISILON:
    		return "kIsilon"		
        case Environment4_KKVM:
    		return "kKVM"		
        case Environment4_KAWS:
    		return "kAWS"		
        case Environment4_KEXCHANGE:
    		return "kExchange"		
        case Environment4_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment4_KORACLE:
    		return "kOracle"		
        case Environment4_KGCP:
    		return "kGCP"		
        case Environment4_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment4_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment4_KVCD:
    		return "kVCD"		
        case Environment4_KO365:
    		return "kO365"		
        case Environment4_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment4_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment4_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment4_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment4Enum Array to its string Array representation
*/
func Environment4EnumArrayToValue(environment4Enum []Environment4Enum) []string {
    convArray := make([]string,len( environment4Enum))
    for i:=0; i<len(environment4Enum);i++ {
        convArray[i] = Environment4EnumToValue(environment4Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment4EnumFromValue(value string) Environment4Enum {
    switch value {
        case "kVMware":
            return Environment4_KVMWARE
        case "kHyperV":
            return Environment4_KHYPERV
        case "kSQL":
            return Environment4_KSQL
        case "kView":
            return Environment4_KVIEW
        case "kPuppeteer":
            return Environment4_KPUPPETEER
        case "kPhysical":
            return Environment4_KPHYSICAL
        case "kPure":
            return Environment4_KPURE
        case "kAzure":
            return Environment4_KAZURE
        case "kNetapp":
            return Environment4_KNETAPP
        case "kAgent":
            return Environment4_KAGENT
        case "kGenericNas":
            return Environment4_KGENERICNAS
        case "kAcropolis":
            return Environment4_KACROPOLIS
        case "kPhysicalFiles":
            return Environment4_KPHYSICALFILES
        case "kIsilon":
            return Environment4_KISILON
        case "kKVM":
            return Environment4_KKVM
        case "kAWS":
            return Environment4_KAWS
        case "kExchange":
            return Environment4_KEXCHANGE
        case "kHyperVVSS":
            return Environment4_KHYPERVVSS
        case "kOracle":
            return Environment4_KORACLE
        case "kGCP":
            return Environment4_KGCP
        case "kFlashBlade":
            return Environment4_KFLASHBLADE
        case "kAWSNative":
            return Environment4_KAWSNATIVE
        case "kVCD":
            return Environment4_KVCD
        case "kO365":
            return Environment4_KO365
        case "kO365Outlook":
            return Environment4_KO365OUTLOOK
        case "kHyperFlex":
            return Environment4_KHYPERFLEX
        case "kGCPNative":
            return Environment4_KGCPNATIVE
        case "kAzureNative":
            return Environment4_KAZURENATIVE
        default:
            return Environment4_KVMWARE
    }
}
