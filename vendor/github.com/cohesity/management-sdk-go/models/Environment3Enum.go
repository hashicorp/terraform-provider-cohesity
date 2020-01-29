package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment3Enum enum
 */
type Environment3Enum int

/**
 * Value collection for Environment3Enum enum
 */
const (
    Environment3_KVMWARE Environment3Enum = 1 + iota
    Environment3_KHYPERV
    Environment3_KSQL
    Environment3_KVIEW
    Environment3_KPUPPETEER
    Environment3_KPHYSICAL
    Environment3_KPURE
    Environment3_KAZURE
    Environment3_KNETAPP
    Environment3_KAGENT
    Environment3_KGENERICNAS
    Environment3_KACROPOLIS
    Environment3_KPHYSICALFILES
    Environment3_KISILON
    Environment3_KKVM
    Environment3_KAWS
    Environment3_KEXCHANGE
    Environment3_KHYPERVVSS
    Environment3_KORACLE
    Environment3_KGCP
    Environment3_KFLASHBLADE
    Environment3_KAWSNATIVE
    Environment3_KVCD
    Environment3_KO365
    Environment3_KO365OUTLOOK
    Environment3_KHYPERFLEX
    Environment3_KGCPNATIVE
    Environment3_KAZURENATIVE
)

func (r Environment3Enum) MarshalJSON() ([]byte, error) { 
    s := Environment3EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment3Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment3EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment3Enum to its string representation
 */
func Environment3EnumToValue(environment3Enum Environment3Enum) string {
    switch environment3Enum {
        case Environment3_KVMWARE:
    		return "kVMware"		
        case Environment3_KHYPERV:
    		return "kHyperV"		
        case Environment3_KSQL:
    		return "kSQL"		
        case Environment3_KVIEW:
    		return "kView"		
        case Environment3_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment3_KPHYSICAL:
    		return "kPhysical"		
        case Environment3_KPURE:
    		return "kPure"		
        case Environment3_KAZURE:
    		return "kAzure"		
        case Environment3_KNETAPP:
    		return "kNetapp"		
        case Environment3_KAGENT:
    		return "kAgent"		
        case Environment3_KGENERICNAS:
    		return "kGenericNas"		
        case Environment3_KACROPOLIS:
    		return "kAcropolis"		
        case Environment3_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment3_KISILON:
    		return "kIsilon"		
        case Environment3_KKVM:
    		return "kKVM"		
        case Environment3_KAWS:
    		return "kAWS"		
        case Environment3_KEXCHANGE:
    		return "kExchange"		
        case Environment3_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment3_KORACLE:
    		return "kOracle"		
        case Environment3_KGCP:
    		return "kGCP"		
        case Environment3_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment3_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment3_KVCD:
    		return "kVCD"		
        case Environment3_KO365:
    		return "kO365"		
        case Environment3_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment3_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment3_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment3_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment3Enum Array to its string Array representation
*/
func Environment3EnumArrayToValue(environment3Enum []Environment3Enum) []string {
    convArray := make([]string,len( environment3Enum))
    for i:=0; i<len(environment3Enum);i++ {
        convArray[i] = Environment3EnumToValue(environment3Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment3EnumFromValue(value string) Environment3Enum {
    switch value {
        case "kVMware":
            return Environment3_KVMWARE
        case "kHyperV":
            return Environment3_KHYPERV
        case "kSQL":
            return Environment3_KSQL
        case "kView":
            return Environment3_KVIEW
        case "kPuppeteer":
            return Environment3_KPUPPETEER
        case "kPhysical":
            return Environment3_KPHYSICAL
        case "kPure":
            return Environment3_KPURE
        case "kAzure":
            return Environment3_KAZURE
        case "kNetapp":
            return Environment3_KNETAPP
        case "kAgent":
            return Environment3_KAGENT
        case "kGenericNas":
            return Environment3_KGENERICNAS
        case "kAcropolis":
            return Environment3_KACROPOLIS
        case "kPhysicalFiles":
            return Environment3_KPHYSICALFILES
        case "kIsilon":
            return Environment3_KISILON
        case "kKVM":
            return Environment3_KKVM
        case "kAWS":
            return Environment3_KAWS
        case "kExchange":
            return Environment3_KEXCHANGE
        case "kHyperVVSS":
            return Environment3_KHYPERVVSS
        case "kOracle":
            return Environment3_KORACLE
        case "kGCP":
            return Environment3_KGCP
        case "kFlashBlade":
            return Environment3_KFLASHBLADE
        case "kAWSNative":
            return Environment3_KAWSNATIVE
        case "kVCD":
            return Environment3_KVCD
        case "kO365":
            return Environment3_KO365
        case "kO365Outlook":
            return Environment3_KO365OUTLOOK
        case "kHyperFlex":
            return Environment3_KHYPERFLEX
        case "kGCPNative":
            return Environment3_KGCPNATIVE
        case "kAzureNative":
            return Environment3_KAZURENATIVE
        default:
            return Environment3_KVMWARE
    }
}
