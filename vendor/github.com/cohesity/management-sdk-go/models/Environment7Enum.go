package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment7Enum enum
 */
type Environment7Enum int

/**
 * Value collection for Environment7Enum enum
 */
const (
    Environment7_KVMWARE Environment7Enum = 1 + iota
    Environment7_KHYPERV
    Environment7_KSQL
    Environment7_KVIEW
    Environment7_KPUPPETEER
    Environment7_KPHYSICAL
    Environment7_KPURE
    Environment7_KAZURE
    Environment7_KNETAPP
    Environment7_KAGENT
    Environment7_KGENERICNAS
    Environment7_KACROPOLIS
    Environment7_KPHYSICALFILES
    Environment7_KISILON
    Environment7_KKVM
    Environment7_KAWS
    Environment7_KEXCHANGE
    Environment7_KHYPERVVSS
    Environment7_KORACLE
    Environment7_KGCP
    Environment7_KFLASHBLADE
    Environment7_KAWSNATIVE
    Environment7_KVCD
    Environment7_KO365
    Environment7_KO365OUTLOOK
    Environment7_KHYPERFLEX
    Environment7_KGCPNATIVE
    Environment7_KAZURENATIVE
)

func (r Environment7Enum) MarshalJSON() ([]byte, error) { 
    s := Environment7EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment7Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment7EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment7Enum to its string representation
 */
func Environment7EnumToValue(environment7Enum Environment7Enum) string {
    switch environment7Enum {
        case Environment7_KVMWARE:
    		return "kVMware"		
        case Environment7_KHYPERV:
    		return "kHyperV"		
        case Environment7_KSQL:
    		return "kSQL"		
        case Environment7_KVIEW:
    		return "kView"		
        case Environment7_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment7_KPHYSICAL:
    		return "kPhysical"		
        case Environment7_KPURE:
    		return "kPure"		
        case Environment7_KAZURE:
    		return "kAzure"		
        case Environment7_KNETAPP:
    		return "kNetapp"		
        case Environment7_KAGENT:
    		return "kAgent"		
        case Environment7_KGENERICNAS:
    		return "kGenericNas"		
        case Environment7_KACROPOLIS:
    		return "kAcropolis"		
        case Environment7_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment7_KISILON:
    		return "kIsilon"		
        case Environment7_KKVM:
    		return "kKVM"		
        case Environment7_KAWS:
    		return "kAWS"		
        case Environment7_KEXCHANGE:
    		return "kExchange"		
        case Environment7_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment7_KORACLE:
    		return "kOracle"		
        case Environment7_KGCP:
    		return "kGCP"		
        case Environment7_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment7_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment7_KVCD:
    		return "kVCD"		
        case Environment7_KO365:
    		return "kO365"		
        case Environment7_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment7_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment7_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment7_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment7Enum Array to its string Array representation
*/
func Environment7EnumArrayToValue(environment7Enum []Environment7Enum) []string {
    convArray := make([]string,len( environment7Enum))
    for i:=0; i<len(environment7Enum);i++ {
        convArray[i] = Environment7EnumToValue(environment7Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment7EnumFromValue(value string) Environment7Enum {
    switch value {
        case "kVMware":
            return Environment7_KVMWARE
        case "kHyperV":
            return Environment7_KHYPERV
        case "kSQL":
            return Environment7_KSQL
        case "kView":
            return Environment7_KVIEW
        case "kPuppeteer":
            return Environment7_KPUPPETEER
        case "kPhysical":
            return Environment7_KPHYSICAL
        case "kPure":
            return Environment7_KPURE
        case "kAzure":
            return Environment7_KAZURE
        case "kNetapp":
            return Environment7_KNETAPP
        case "kAgent":
            return Environment7_KAGENT
        case "kGenericNas":
            return Environment7_KGENERICNAS
        case "kAcropolis":
            return Environment7_KACROPOLIS
        case "kPhysicalFiles":
            return Environment7_KPHYSICALFILES
        case "kIsilon":
            return Environment7_KISILON
        case "kKVM":
            return Environment7_KKVM
        case "kAWS":
            return Environment7_KAWS
        case "kExchange":
            return Environment7_KEXCHANGE
        case "kHyperVVSS":
            return Environment7_KHYPERVVSS
        case "kOracle":
            return Environment7_KORACLE
        case "kGCP":
            return Environment7_KGCP
        case "kFlashBlade":
            return Environment7_KFLASHBLADE
        case "kAWSNative":
            return Environment7_KAWSNATIVE
        case "kVCD":
            return Environment7_KVCD
        case "kO365":
            return Environment7_KO365
        case "kO365Outlook":
            return Environment7_KO365OUTLOOK
        case "kHyperFlex":
            return Environment7_KHYPERFLEX
        case "kGCPNative":
            return Environment7_KGCPNATIVE
        case "kAzureNative":
            return Environment7_KAZURENATIVE
        default:
            return Environment7_KVMWARE
    }
}
