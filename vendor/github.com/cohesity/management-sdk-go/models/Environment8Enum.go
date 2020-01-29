package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment8Enum enum
 */
type Environment8Enum int

/**
 * Value collection for Environment8Enum enum
 */
const (
    Environment8_KVMWARE Environment8Enum = 1 + iota
    Environment8_KHYPERV
    Environment8_KSQL
    Environment8_KVIEW
    Environment8_KPUPPETEER
    Environment8_KPHYSICAL
    Environment8_KPURE
    Environment8_KAZURE
    Environment8_KNETAPP
    Environment8_KAGENT
    Environment8_KGENERICNAS
    Environment8_KACROPOLIS
    Environment8_KPHYSICALFILES
    Environment8_KISILON
    Environment8_KKVM
    Environment8_KAWS
    Environment8_KEXCHANGE
    Environment8_KHYPERVVSS
    Environment8_KORACLE
    Environment8_KGCP
    Environment8_KFLASHBLADE
    Environment8_KAWSNATIVE
    Environment8_KVCD
    Environment8_KO365
    Environment8_KO365OUTLOOK
    Environment8_KHYPERFLEX
    Environment8_KGCPNATIVE
    Environment8_KAZURENATIVE
)

func (r Environment8Enum) MarshalJSON() ([]byte, error) { 
    s := Environment8EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment8Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment8EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment8Enum to its string representation
 */
func Environment8EnumToValue(environment8Enum Environment8Enum) string {
    switch environment8Enum {
        case Environment8_KVMWARE:
    		return "kVMware"		
        case Environment8_KHYPERV:
    		return "kHyperV"		
        case Environment8_KSQL:
    		return "kSQL"		
        case Environment8_KVIEW:
    		return "kView"		
        case Environment8_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment8_KPHYSICAL:
    		return "kPhysical"		
        case Environment8_KPURE:
    		return "kPure"		
        case Environment8_KAZURE:
    		return "kAzure"		
        case Environment8_KNETAPP:
    		return "kNetapp"		
        case Environment8_KAGENT:
    		return "kAgent"		
        case Environment8_KGENERICNAS:
    		return "kGenericNas"		
        case Environment8_KACROPOLIS:
    		return "kAcropolis"		
        case Environment8_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment8_KISILON:
    		return "kIsilon"		
        case Environment8_KKVM:
    		return "kKVM"		
        case Environment8_KAWS:
    		return "kAWS"		
        case Environment8_KEXCHANGE:
    		return "kExchange"		
        case Environment8_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment8_KORACLE:
    		return "kOracle"		
        case Environment8_KGCP:
    		return "kGCP"		
        case Environment8_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment8_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment8_KVCD:
    		return "kVCD"		
        case Environment8_KO365:
    		return "kO365"		
        case Environment8_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment8_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment8_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment8_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment8Enum Array to its string Array representation
*/
func Environment8EnumArrayToValue(environment8Enum []Environment8Enum) []string {
    convArray := make([]string,len( environment8Enum))
    for i:=0; i<len(environment8Enum);i++ {
        convArray[i] = Environment8EnumToValue(environment8Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment8EnumFromValue(value string) Environment8Enum {
    switch value {
        case "kVMware":
            return Environment8_KVMWARE
        case "kHyperV":
            return Environment8_KHYPERV
        case "kSQL":
            return Environment8_KSQL
        case "kView":
            return Environment8_KVIEW
        case "kPuppeteer":
            return Environment8_KPUPPETEER
        case "kPhysical":
            return Environment8_KPHYSICAL
        case "kPure":
            return Environment8_KPURE
        case "kAzure":
            return Environment8_KAZURE
        case "kNetapp":
            return Environment8_KNETAPP
        case "kAgent":
            return Environment8_KAGENT
        case "kGenericNas":
            return Environment8_KGENERICNAS
        case "kAcropolis":
            return Environment8_KACROPOLIS
        case "kPhysicalFiles":
            return Environment8_KPHYSICALFILES
        case "kIsilon":
            return Environment8_KISILON
        case "kKVM":
            return Environment8_KKVM
        case "kAWS":
            return Environment8_KAWS
        case "kExchange":
            return Environment8_KEXCHANGE
        case "kHyperVVSS":
            return Environment8_KHYPERVVSS
        case "kOracle":
            return Environment8_KORACLE
        case "kGCP":
            return Environment8_KGCP
        case "kFlashBlade":
            return Environment8_KFLASHBLADE
        case "kAWSNative":
            return Environment8_KAWSNATIVE
        case "kVCD":
            return Environment8_KVCD
        case "kO365":
            return Environment8_KO365
        case "kO365Outlook":
            return Environment8_KO365OUTLOOK
        case "kHyperFlex":
            return Environment8_KHYPERFLEX
        case "kGCPNative":
            return Environment8_KGCPNATIVE
        case "kAzureNative":
            return Environment8_KAZURENATIVE
        default:
            return Environment8_KVMWARE
    }
}
