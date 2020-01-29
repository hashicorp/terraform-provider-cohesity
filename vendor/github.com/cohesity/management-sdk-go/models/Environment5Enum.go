package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment5Enum enum
 */
type Environment5Enum int

/**
 * Value collection for Environment5Enum enum
 */
const (
    Environment5_KVMWARE Environment5Enum = 1 + iota
    Environment5_KHYPERV
    Environment5_KSQL
    Environment5_KVIEW
    Environment5_KPUPPETEER
    Environment5_KPHYSICAL
    Environment5_KPURE
    Environment5_KAZURE
    Environment5_KNETAPP
    Environment5_KAGENT
    Environment5_KGENERICNAS
    Environment5_KACROPOLIS
    Environment5_KPHYSICALFILES
    Environment5_KISILON
    Environment5_KKVM
    Environment5_KAWS
    Environment5_KEXCHANGE
    Environment5_KHYPERVVSS
    Environment5_KORACLE
    Environment5_KGCP
    Environment5_KFLASHBLADE
    Environment5_KAWSNATIVE
    Environment5_KVCD
    Environment5_KO365
    Environment5_KO365OUTLOOK
    Environment5_KHYPERFLEX
    Environment5_KGCPNATIVE
    Environment5_KAZURENATIVE
)

func (r Environment5Enum) MarshalJSON() ([]byte, error) { 
    s := Environment5EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment5Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment5EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment5Enum to its string representation
 */
func Environment5EnumToValue(environment5Enum Environment5Enum) string {
    switch environment5Enum {
        case Environment5_KVMWARE:
    		return "kVMware"		
        case Environment5_KHYPERV:
    		return "kHyperV"		
        case Environment5_KSQL:
    		return "kSQL"		
        case Environment5_KVIEW:
    		return "kView"		
        case Environment5_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment5_KPHYSICAL:
    		return "kPhysical"		
        case Environment5_KPURE:
    		return "kPure"		
        case Environment5_KAZURE:
    		return "kAzure"		
        case Environment5_KNETAPP:
    		return "kNetapp"		
        case Environment5_KAGENT:
    		return "kAgent"		
        case Environment5_KGENERICNAS:
    		return "kGenericNas"		
        case Environment5_KACROPOLIS:
    		return "kAcropolis"		
        case Environment5_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment5_KISILON:
    		return "kIsilon"		
        case Environment5_KKVM:
    		return "kKVM"		
        case Environment5_KAWS:
    		return "kAWS"		
        case Environment5_KEXCHANGE:
    		return "kExchange"		
        case Environment5_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment5_KORACLE:
    		return "kOracle"		
        case Environment5_KGCP:
    		return "kGCP"		
        case Environment5_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment5_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment5_KVCD:
    		return "kVCD"		
        case Environment5_KO365:
    		return "kO365"		
        case Environment5_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment5_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment5_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment5_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment5Enum Array to its string Array representation
*/
func Environment5EnumArrayToValue(environment5Enum []Environment5Enum) []string {
    convArray := make([]string,len( environment5Enum))
    for i:=0; i<len(environment5Enum);i++ {
        convArray[i] = Environment5EnumToValue(environment5Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment5EnumFromValue(value string) Environment5Enum {
    switch value {
        case "kVMware":
            return Environment5_KVMWARE
        case "kHyperV":
            return Environment5_KHYPERV
        case "kSQL":
            return Environment5_KSQL
        case "kView":
            return Environment5_KVIEW
        case "kPuppeteer":
            return Environment5_KPUPPETEER
        case "kPhysical":
            return Environment5_KPHYSICAL
        case "kPure":
            return Environment5_KPURE
        case "kAzure":
            return Environment5_KAZURE
        case "kNetapp":
            return Environment5_KNETAPP
        case "kAgent":
            return Environment5_KAGENT
        case "kGenericNas":
            return Environment5_KGENERICNAS
        case "kAcropolis":
            return Environment5_KACROPOLIS
        case "kPhysicalFiles":
            return Environment5_KPHYSICALFILES
        case "kIsilon":
            return Environment5_KISILON
        case "kKVM":
            return Environment5_KKVM
        case "kAWS":
            return Environment5_KAWS
        case "kExchange":
            return Environment5_KEXCHANGE
        case "kHyperVVSS":
            return Environment5_KHYPERVVSS
        case "kOracle":
            return Environment5_KORACLE
        case "kGCP":
            return Environment5_KGCP
        case "kFlashBlade":
            return Environment5_KFLASHBLADE
        case "kAWSNative":
            return Environment5_KAWSNATIVE
        case "kVCD":
            return Environment5_KVCD
        case "kO365":
            return Environment5_KO365
        case "kO365Outlook":
            return Environment5_KO365OUTLOOK
        case "kHyperFlex":
            return Environment5_KHYPERFLEX
        case "kGCPNative":
            return Environment5_KGCPNATIVE
        case "kAzureNative":
            return Environment5_KAZURENATIVE
        default:
            return Environment5_KVMWARE
    }
}
