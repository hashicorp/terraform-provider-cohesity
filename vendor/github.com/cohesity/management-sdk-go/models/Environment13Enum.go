package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment13Enum enum
 */
type Environment13Enum int

/**
 * Value collection for Environment13Enum enum
 */
const (
    Environment13_KVMWARE Environment13Enum = 1 + iota
    Environment13_KHYPERV
    Environment13_KSQL
    Environment13_KVIEW
    Environment13_KPUPPETEER
    Environment13_KPHYSICAL
    Environment13_KPURE
    Environment13_KAZURE
    Environment13_KNETAPP
    Environment13_KAGENT
    Environment13_KGENERICNAS
    Environment13_KACROPOLIS
    Environment13_KPHYSICALFILES
    Environment13_KISILON
    Environment13_KKVM
    Environment13_KAWS
    Environment13_KEXCHANGE
    Environment13_KHYPERVVSS
    Environment13_KORACLE
    Environment13_KGCP
    Environment13_KFLASHBLADE
    Environment13_KAWSNATIVE
    Environment13_KVCD
    Environment13_KO365
    Environment13_KO365OUTLOOK
    Environment13_KHYPERFLEX
    Environment13_KGCPNATIVE
    Environment13_KAZURENATIVE
)

func (r Environment13Enum) MarshalJSON() ([]byte, error) { 
    s := Environment13EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment13Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment13EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment13Enum to its string representation
 */
func Environment13EnumToValue(environment13Enum Environment13Enum) string {
    switch environment13Enum {
        case Environment13_KVMWARE:
    		return "kVMware"		
        case Environment13_KHYPERV:
    		return "kHyperV"		
        case Environment13_KSQL:
    		return "kSQL"		
        case Environment13_KVIEW:
    		return "kView"		
        case Environment13_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment13_KPHYSICAL:
    		return "kPhysical"		
        case Environment13_KPURE:
    		return "kPure"		
        case Environment13_KAZURE:
    		return "kAzure"		
        case Environment13_KNETAPP:
    		return "kNetapp"		
        case Environment13_KAGENT:
    		return "kAgent"		
        case Environment13_KGENERICNAS:
    		return "kGenericNas"		
        case Environment13_KACROPOLIS:
    		return "kAcropolis"		
        case Environment13_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment13_KISILON:
    		return "kIsilon"		
        case Environment13_KKVM:
    		return "kKVM"		
        case Environment13_KAWS:
    		return "kAWS"		
        case Environment13_KEXCHANGE:
    		return "kExchange"		
        case Environment13_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment13_KORACLE:
    		return "kOracle"		
        case Environment13_KGCP:
    		return "kGCP"		
        case Environment13_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment13_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment13_KVCD:
    		return "kVCD"		
        case Environment13_KO365:
    		return "kO365"		
        case Environment13_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment13_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment13_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment13_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment13Enum Array to its string Array representation
*/
func Environment13EnumArrayToValue(environment13Enum []Environment13Enum) []string {
    convArray := make([]string,len( environment13Enum))
    for i:=0; i<len(environment13Enum);i++ {
        convArray[i] = Environment13EnumToValue(environment13Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment13EnumFromValue(value string) Environment13Enum {
    switch value {
        case "kVMware":
            return Environment13_KVMWARE
        case "kHyperV":
            return Environment13_KHYPERV
        case "kSQL":
            return Environment13_KSQL
        case "kView":
            return Environment13_KVIEW
        case "kPuppeteer":
            return Environment13_KPUPPETEER
        case "kPhysical":
            return Environment13_KPHYSICAL
        case "kPure":
            return Environment13_KPURE
        case "kAzure":
            return Environment13_KAZURE
        case "kNetapp":
            return Environment13_KNETAPP
        case "kAgent":
            return Environment13_KAGENT
        case "kGenericNas":
            return Environment13_KGENERICNAS
        case "kAcropolis":
            return Environment13_KACROPOLIS
        case "kPhysicalFiles":
            return Environment13_KPHYSICALFILES
        case "kIsilon":
            return Environment13_KISILON
        case "kKVM":
            return Environment13_KKVM
        case "kAWS":
            return Environment13_KAWS
        case "kExchange":
            return Environment13_KEXCHANGE
        case "kHyperVVSS":
            return Environment13_KHYPERVVSS
        case "kOracle":
            return Environment13_KORACLE
        case "kGCP":
            return Environment13_KGCP
        case "kFlashBlade":
            return Environment13_KFLASHBLADE
        case "kAWSNative":
            return Environment13_KAWSNATIVE
        case "kVCD":
            return Environment13_KVCD
        case "kO365":
            return Environment13_KO365
        case "kO365Outlook":
            return Environment13_KO365OUTLOOK
        case "kHyperFlex":
            return Environment13_KHYPERFLEX
        case "kGCPNative":
            return Environment13_KGCPNATIVE
        case "kAzureNative":
            return Environment13_KAZURENATIVE
        default:
            return Environment13_KVMWARE
    }
}
