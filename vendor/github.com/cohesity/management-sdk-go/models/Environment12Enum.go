package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment12Enum enum
 */
type Environment12Enum int

/**
 * Value collection for Environment12Enum enum
 */
const (
    Environment12_KVMWARE Environment12Enum = 1 + iota
    Environment12_KHYPERV
    Environment12_KSQL
    Environment12_KVIEW
    Environment12_KPUPPETEER
    Environment12_KPHYSICAL
    Environment12_KPURE
    Environment12_KAZURE
    Environment12_KNETAPP
    Environment12_KAGENT
    Environment12_KGENERICNAS
    Environment12_KACROPOLIS
    Environment12_KPHYSICALFILES
    Environment12_KISILON
    Environment12_KKVM
    Environment12_KAWS
    Environment12_KEXCHANGE
    Environment12_KHYPERVVSS
    Environment12_KORACLE
    Environment12_KGCP
    Environment12_KFLASHBLADE
    Environment12_KAWSNATIVE
    Environment12_KVCD
    Environment12_KO365
    Environment12_KO365OUTLOOK
    Environment12_KHYPERFLEX
    Environment12_KGCPNATIVE
    Environment12_KAZURENATIVE
)

func (r Environment12Enum) MarshalJSON() ([]byte, error) { 
    s := Environment12EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment12Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment12EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment12Enum to its string representation
 */
func Environment12EnumToValue(environment12Enum Environment12Enum) string {
    switch environment12Enum {
        case Environment12_KVMWARE:
    		return "kVMware"		
        case Environment12_KHYPERV:
    		return "kHyperV"		
        case Environment12_KSQL:
    		return "kSQL"		
        case Environment12_KVIEW:
    		return "kView"		
        case Environment12_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment12_KPHYSICAL:
    		return "kPhysical"		
        case Environment12_KPURE:
    		return "kPure"		
        case Environment12_KAZURE:
    		return "kAzure"		
        case Environment12_KNETAPP:
    		return "kNetapp"		
        case Environment12_KAGENT:
    		return "kAgent"		
        case Environment12_KGENERICNAS:
    		return "kGenericNas"		
        case Environment12_KACROPOLIS:
    		return "kAcropolis"		
        case Environment12_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment12_KISILON:
    		return "kIsilon"		
        case Environment12_KKVM:
    		return "kKVM"		
        case Environment12_KAWS:
    		return "kAWS"		
        case Environment12_KEXCHANGE:
    		return "kExchange"		
        case Environment12_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment12_KORACLE:
    		return "kOracle"		
        case Environment12_KGCP:
    		return "kGCP"		
        case Environment12_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment12_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment12_KVCD:
    		return "kVCD"		
        case Environment12_KO365:
    		return "kO365"		
        case Environment12_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment12_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment12_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment12_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment12Enum Array to its string Array representation
*/
func Environment12EnumArrayToValue(environment12Enum []Environment12Enum) []string {
    convArray := make([]string,len( environment12Enum))
    for i:=0; i<len(environment12Enum);i++ {
        convArray[i] = Environment12EnumToValue(environment12Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment12EnumFromValue(value string) Environment12Enum {
    switch value {
        case "kVMware":
            return Environment12_KVMWARE
        case "kHyperV":
            return Environment12_KHYPERV
        case "kSQL":
            return Environment12_KSQL
        case "kView":
            return Environment12_KVIEW
        case "kPuppeteer":
            return Environment12_KPUPPETEER
        case "kPhysical":
            return Environment12_KPHYSICAL
        case "kPure":
            return Environment12_KPURE
        case "kAzure":
            return Environment12_KAZURE
        case "kNetapp":
            return Environment12_KNETAPP
        case "kAgent":
            return Environment12_KAGENT
        case "kGenericNas":
            return Environment12_KGENERICNAS
        case "kAcropolis":
            return Environment12_KACROPOLIS
        case "kPhysicalFiles":
            return Environment12_KPHYSICALFILES
        case "kIsilon":
            return Environment12_KISILON
        case "kKVM":
            return Environment12_KKVM
        case "kAWS":
            return Environment12_KAWS
        case "kExchange":
            return Environment12_KEXCHANGE
        case "kHyperVVSS":
            return Environment12_KHYPERVVSS
        case "kOracle":
            return Environment12_KORACLE
        case "kGCP":
            return Environment12_KGCP
        case "kFlashBlade":
            return Environment12_KFLASHBLADE
        case "kAWSNative":
            return Environment12_KAWSNATIVE
        case "kVCD":
            return Environment12_KVCD
        case "kO365":
            return Environment12_KO365
        case "kO365Outlook":
            return Environment12_KO365OUTLOOK
        case "kHyperFlex":
            return Environment12_KHYPERFLEX
        case "kGCPNative":
            return Environment12_KGCPNATIVE
        case "kAzureNative":
            return Environment12_KAZURENATIVE
        default:
            return Environment12_KVMWARE
    }
}
