package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment2Enum enum
 */
type Environment2Enum int

/**
 * Value collection for Environment2Enum enum
 */
const (
    Environment2_KVMWARE Environment2Enum = 1 + iota
    Environment2_KHYPERV
    Environment2_KSQL
    Environment2_KVIEW
    Environment2_KPUPPETEER
    Environment2_KPHYSICAL
    Environment2_KPURE
    Environment2_KAZURE
    Environment2_KNETAPP
    Environment2_KAGENT
    Environment2_KGENERICNAS
    Environment2_KACROPOLIS
    Environment2_KPHYSICALFILES
    Environment2_KISILON
    Environment2_KKVM
    Environment2_KAWS
    Environment2_KEXCHANGE
    Environment2_KHYPERVVSS
    Environment2_KORACLE
    Environment2_KGCP
    Environment2_KFLASHBLADE
    Environment2_KAWSNATIVE
    Environment2_KVCD
    Environment2_KO365
    Environment2_KO365OUTLOOK
    Environment2_KHYPERFLEX
    Environment2_KGCPNATIVE
    Environment2_KAZURENATIVE
)

func (r Environment2Enum) MarshalJSON() ([]byte, error) { 
    s := Environment2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment2Enum to its string representation
 */
func Environment2EnumToValue(environment2Enum Environment2Enum) string {
    switch environment2Enum {
        case Environment2_KVMWARE:
    		return "kVMware"		
        case Environment2_KHYPERV:
    		return "kHyperV"		
        case Environment2_KSQL:
    		return "kSQL"		
        case Environment2_KVIEW:
    		return "kView"		
        case Environment2_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment2_KPHYSICAL:
    		return "kPhysical"		
        case Environment2_KPURE:
    		return "kPure"		
        case Environment2_KAZURE:
    		return "kAzure"		
        case Environment2_KNETAPP:
    		return "kNetapp"		
        case Environment2_KAGENT:
    		return "kAgent"		
        case Environment2_KGENERICNAS:
    		return "kGenericNas"		
        case Environment2_KACROPOLIS:
    		return "kAcropolis"		
        case Environment2_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment2_KISILON:
    		return "kIsilon"		
        case Environment2_KKVM:
    		return "kKVM"		
        case Environment2_KAWS:
    		return "kAWS"		
        case Environment2_KEXCHANGE:
    		return "kExchange"		
        case Environment2_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment2_KORACLE:
    		return "kOracle"		
        case Environment2_KGCP:
    		return "kGCP"		
        case Environment2_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment2_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment2_KVCD:
    		return "kVCD"		
        case Environment2_KO365:
    		return "kO365"		
        case Environment2_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment2_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment2_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment2_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment2Enum Array to its string Array representation
*/
func Environment2EnumArrayToValue(environment2Enum []Environment2Enum) []string {
    convArray := make([]string,len( environment2Enum))
    for i:=0; i<len(environment2Enum);i++ {
        convArray[i] = Environment2EnumToValue(environment2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment2EnumFromValue(value string) Environment2Enum {
    switch value {
        case "kVMware":
            return Environment2_KVMWARE
        case "kHyperV":
            return Environment2_KHYPERV
        case "kSQL":
            return Environment2_KSQL
        case "kView":
            return Environment2_KVIEW
        case "kPuppeteer":
            return Environment2_KPUPPETEER
        case "kPhysical":
            return Environment2_KPHYSICAL
        case "kPure":
            return Environment2_KPURE
        case "kAzure":
            return Environment2_KAZURE
        case "kNetapp":
            return Environment2_KNETAPP
        case "kAgent":
            return Environment2_KAGENT
        case "kGenericNas":
            return Environment2_KGENERICNAS
        case "kAcropolis":
            return Environment2_KACROPOLIS
        case "kPhysicalFiles":
            return Environment2_KPHYSICALFILES
        case "kIsilon":
            return Environment2_KISILON
        case "kKVM":
            return Environment2_KKVM
        case "kAWS":
            return Environment2_KAWS
        case "kExchange":
            return Environment2_KEXCHANGE
        case "kHyperVVSS":
            return Environment2_KHYPERVVSS
        case "kOracle":
            return Environment2_KORACLE
        case "kGCP":
            return Environment2_KGCP
        case "kFlashBlade":
            return Environment2_KFLASHBLADE
        case "kAWSNative":
            return Environment2_KAWSNATIVE
        case "kVCD":
            return Environment2_KVCD
        case "kO365":
            return Environment2_KO365
        case "kO365Outlook":
            return Environment2_KO365OUTLOOK
        case "kHyperFlex":
            return Environment2_KHYPERFLEX
        case "kGCPNative":
            return Environment2_KGCPNATIVE
        case "kAzureNative":
            return Environment2_KAZURENATIVE
        default:
            return Environment2_KVMWARE
    }
}
