package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment9Enum enum
 */
type Environment9Enum int

/**
 * Value collection for Environment9Enum enum
 */
const (
    Environment9_KVMWARE Environment9Enum = 1 + iota
    Environment9_KHYPERV
    Environment9_KSQL
    Environment9_KVIEW
    Environment9_KPUPPETEER
    Environment9_KPHYSICAL
    Environment9_KPURE
    Environment9_KAZURE
    Environment9_KNETAPP
    Environment9_KAGENT
    Environment9_KGENERICNAS
    Environment9_KACROPOLIS
    Environment9_KPHYSICALFILES
    Environment9_KISILON
    Environment9_KKVM
    Environment9_KAWS
    Environment9_KEXCHANGE
    Environment9_KHYPERVVSS
    Environment9_KORACLE
    Environment9_KGCP
    Environment9_KFLASHBLADE
    Environment9_KAWSNATIVE
    Environment9_KVCD
    Environment9_KO365
    Environment9_KO365OUTLOOK
    Environment9_KHYPERFLEX
    Environment9_KGCPNATIVE
    Environment9_KAZURENATIVE
)

func (r Environment9Enum) MarshalJSON() ([]byte, error) { 
    s := Environment9EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment9Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment9EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment9Enum to its string representation
 */
func Environment9EnumToValue(environment9Enum Environment9Enum) string {
    switch environment9Enum {
        case Environment9_KVMWARE:
    		return "kVMware"		
        case Environment9_KHYPERV:
    		return "kHyperV"		
        case Environment9_KSQL:
    		return "kSQL"		
        case Environment9_KVIEW:
    		return "kView"		
        case Environment9_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment9_KPHYSICAL:
    		return "kPhysical"		
        case Environment9_KPURE:
    		return "kPure"		
        case Environment9_KAZURE:
    		return "kAzure"		
        case Environment9_KNETAPP:
    		return "kNetapp"		
        case Environment9_KAGENT:
    		return "kAgent"		
        case Environment9_KGENERICNAS:
    		return "kGenericNas"		
        case Environment9_KACROPOLIS:
    		return "kAcropolis"		
        case Environment9_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment9_KISILON:
    		return "kIsilon"		
        case Environment9_KKVM:
    		return "kKVM"		
        case Environment9_KAWS:
    		return "kAWS"		
        case Environment9_KEXCHANGE:
    		return "kExchange"		
        case Environment9_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment9_KORACLE:
    		return "kOracle"		
        case Environment9_KGCP:
    		return "kGCP"		
        case Environment9_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment9_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment9_KVCD:
    		return "kVCD"		
        case Environment9_KO365:
    		return "kO365"		
        case Environment9_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment9_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment9_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment9_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment9Enum Array to its string Array representation
*/
func Environment9EnumArrayToValue(environment9Enum []Environment9Enum) []string {
    convArray := make([]string,len( environment9Enum))
    for i:=0; i<len(environment9Enum);i++ {
        convArray[i] = Environment9EnumToValue(environment9Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment9EnumFromValue(value string) Environment9Enum {
    switch value {
        case "kVMware":
            return Environment9_KVMWARE
        case "kHyperV":
            return Environment9_KHYPERV
        case "kSQL":
            return Environment9_KSQL
        case "kView":
            return Environment9_KVIEW
        case "kPuppeteer":
            return Environment9_KPUPPETEER
        case "kPhysical":
            return Environment9_KPHYSICAL
        case "kPure":
            return Environment9_KPURE
        case "kAzure":
            return Environment9_KAZURE
        case "kNetapp":
            return Environment9_KNETAPP
        case "kAgent":
            return Environment9_KAGENT
        case "kGenericNas":
            return Environment9_KGENERICNAS
        case "kAcropolis":
            return Environment9_KACROPOLIS
        case "kPhysicalFiles":
            return Environment9_KPHYSICALFILES
        case "kIsilon":
            return Environment9_KISILON
        case "kKVM":
            return Environment9_KKVM
        case "kAWS":
            return Environment9_KAWS
        case "kExchange":
            return Environment9_KEXCHANGE
        case "kHyperVVSS":
            return Environment9_KHYPERVVSS
        case "kOracle":
            return Environment9_KORACLE
        case "kGCP":
            return Environment9_KGCP
        case "kFlashBlade":
            return Environment9_KFLASHBLADE
        case "kAWSNative":
            return Environment9_KAWSNATIVE
        case "kVCD":
            return Environment9_KVCD
        case "kO365":
            return Environment9_KO365
        case "kO365Outlook":
            return Environment9_KO365OUTLOOK
        case "kHyperFlex":
            return Environment9_KHYPERFLEX
        case "kGCPNative":
            return Environment9_KGCPNATIVE
        case "kAzureNative":
            return Environment9_KAZURENATIVE
        default:
            return Environment9_KVMWARE
    }
}
