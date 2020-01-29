package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment6Enum enum
 */
type Environment6Enum int

/**
 * Value collection for Environment6Enum enum
 */
const (
    Environment6_KVMWARE Environment6Enum = 1 + iota
    Environment6_KHYPERV
    Environment6_KSQL
    Environment6_KVIEW
    Environment6_KPUPPETEER
    Environment6_KPHYSICAL
    Environment6_KPURE
    Environment6_KAZURE
    Environment6_KNETAPP
    Environment6_KAGENT
    Environment6_KGENERICNAS
    Environment6_KACROPOLIS
    Environment6_KPHYSICALFILES
    Environment6_KISILON
    Environment6_KKVM
    Environment6_KAWS
    Environment6_KEXCHANGE
    Environment6_KHYPERVVSS
    Environment6_KORACLE
    Environment6_KGCP
    Environment6_KFLASHBLADE
    Environment6_KAWSNATIVE
    Environment6_KVCD
    Environment6_KO365
    Environment6_KO365OUTLOOK
    Environment6_KHYPERFLEX
    Environment6_KGCPNATIVE
    Environment6_KAZURENATIVE
)

func (r Environment6Enum) MarshalJSON() ([]byte, error) { 
    s := Environment6EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment6Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment6EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment6Enum to its string representation
 */
func Environment6EnumToValue(environment6Enum Environment6Enum) string {
    switch environment6Enum {
        case Environment6_KVMWARE:
    		return "kVMware"		
        case Environment6_KHYPERV:
    		return "kHyperV"		
        case Environment6_KSQL:
    		return "kSQL"		
        case Environment6_KVIEW:
    		return "kView"		
        case Environment6_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment6_KPHYSICAL:
    		return "kPhysical"		
        case Environment6_KPURE:
    		return "kPure"		
        case Environment6_KAZURE:
    		return "kAzure"		
        case Environment6_KNETAPP:
    		return "kNetapp"		
        case Environment6_KAGENT:
    		return "kAgent"		
        case Environment6_KGENERICNAS:
    		return "kGenericNas"		
        case Environment6_KACROPOLIS:
    		return "kAcropolis"		
        case Environment6_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment6_KISILON:
    		return "kIsilon"		
        case Environment6_KKVM:
    		return "kKVM"		
        case Environment6_KAWS:
    		return "kAWS"		
        case Environment6_KEXCHANGE:
    		return "kExchange"		
        case Environment6_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment6_KORACLE:
    		return "kOracle"		
        case Environment6_KGCP:
    		return "kGCP"		
        case Environment6_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment6_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment6_KVCD:
    		return "kVCD"		
        case Environment6_KO365:
    		return "kO365"		
        case Environment6_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment6_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment6_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment6_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment6Enum Array to its string Array representation
*/
func Environment6EnumArrayToValue(environment6Enum []Environment6Enum) []string {
    convArray := make([]string,len( environment6Enum))
    for i:=0; i<len(environment6Enum);i++ {
        convArray[i] = Environment6EnumToValue(environment6Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment6EnumFromValue(value string) Environment6Enum {
    switch value {
        case "kVMware":
            return Environment6_KVMWARE
        case "kHyperV":
            return Environment6_KHYPERV
        case "kSQL":
            return Environment6_KSQL
        case "kView":
            return Environment6_KVIEW
        case "kPuppeteer":
            return Environment6_KPUPPETEER
        case "kPhysical":
            return Environment6_KPHYSICAL
        case "kPure":
            return Environment6_KPURE
        case "kAzure":
            return Environment6_KAZURE
        case "kNetapp":
            return Environment6_KNETAPP
        case "kAgent":
            return Environment6_KAGENT
        case "kGenericNas":
            return Environment6_KGENERICNAS
        case "kAcropolis":
            return Environment6_KACROPOLIS
        case "kPhysicalFiles":
            return Environment6_KPHYSICALFILES
        case "kIsilon":
            return Environment6_KISILON
        case "kKVM":
            return Environment6_KKVM
        case "kAWS":
            return Environment6_KAWS
        case "kExchange":
            return Environment6_KEXCHANGE
        case "kHyperVVSS":
            return Environment6_KHYPERVVSS
        case "kOracle":
            return Environment6_KORACLE
        case "kGCP":
            return Environment6_KGCP
        case "kFlashBlade":
            return Environment6_KFLASHBLADE
        case "kAWSNative":
            return Environment6_KAWSNATIVE
        case "kVCD":
            return Environment6_KVCD
        case "kO365":
            return Environment6_KO365
        case "kO365Outlook":
            return Environment6_KO365OUTLOOK
        case "kHyperFlex":
            return Environment6_KHYPERFLEX
        case "kGCPNative":
            return Environment6_KGCPNATIVE
        case "kAzureNative":
            return Environment6_KAZURENATIVE
        default:
            return Environment6_KVMWARE
    }
}
