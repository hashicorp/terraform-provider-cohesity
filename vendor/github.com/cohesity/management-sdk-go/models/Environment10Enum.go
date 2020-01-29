package models

import(
    "encoding/json"
)

/**
 * Type definition for Environment10Enum enum
 */
type Environment10Enum int

/**
 * Value collection for Environment10Enum enum
 */
const (
    Environment10_KVMWARE Environment10Enum = 1 + iota
    Environment10_KHYPERV
    Environment10_KSQL
    Environment10_KVIEW
    Environment10_KPUPPETEER
    Environment10_KPHYSICAL
    Environment10_KPURE
    Environment10_KAZURE
    Environment10_KNETAPP
    Environment10_KAGENT
    Environment10_KGENERICNAS
    Environment10_KACROPOLIS
    Environment10_KPHYSICALFILES
    Environment10_KISILON
    Environment10_KKVM
    Environment10_KAWS
    Environment10_KEXCHANGE
    Environment10_KHYPERVVSS
    Environment10_KORACLE
    Environment10_KGCP
    Environment10_KFLASHBLADE
    Environment10_KAWSNATIVE
    Environment10_KVCD
    Environment10_KO365
    Environment10_KO365OUTLOOK
    Environment10_KHYPERFLEX
    Environment10_KGCPNATIVE
    Environment10_KAZURENATIVE
)

func (r Environment10Enum) MarshalJSON() ([]byte, error) { 
    s := Environment10EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environment10Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environment10EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environment10Enum to its string representation
 */
func Environment10EnumToValue(environment10Enum Environment10Enum) string {
    switch environment10Enum {
        case Environment10_KVMWARE:
    		return "kVMware"		
        case Environment10_KHYPERV:
    		return "kHyperV"		
        case Environment10_KSQL:
    		return "kSQL"		
        case Environment10_KVIEW:
    		return "kView"		
        case Environment10_KPUPPETEER:
    		return "kPuppeteer"		
        case Environment10_KPHYSICAL:
    		return "kPhysical"		
        case Environment10_KPURE:
    		return "kPure"		
        case Environment10_KAZURE:
    		return "kAzure"		
        case Environment10_KNETAPP:
    		return "kNetapp"		
        case Environment10_KAGENT:
    		return "kAgent"		
        case Environment10_KGENERICNAS:
    		return "kGenericNas"		
        case Environment10_KACROPOLIS:
    		return "kAcropolis"		
        case Environment10_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Environment10_KISILON:
    		return "kIsilon"		
        case Environment10_KKVM:
    		return "kKVM"		
        case Environment10_KAWS:
    		return "kAWS"		
        case Environment10_KEXCHANGE:
    		return "kExchange"		
        case Environment10_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Environment10_KORACLE:
    		return "kOracle"		
        case Environment10_KGCP:
    		return "kGCP"		
        case Environment10_KFLASHBLADE:
    		return "kFlashBlade"		
        case Environment10_KAWSNATIVE:
    		return "kAWSNative"		
        case Environment10_KVCD:
    		return "kVCD"		
        case Environment10_KO365:
    		return "kO365"		
        case Environment10_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Environment10_KHYPERFLEX:
    		return "kHyperFlex"		
        case Environment10_KGCPNATIVE:
    		return "kGCPNative"		
        case Environment10_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environment10Enum Array to its string Array representation
*/
func Environment10EnumArrayToValue(environment10Enum []Environment10Enum) []string {
    convArray := make([]string,len( environment10Enum))
    for i:=0; i<len(environment10Enum);i++ {
        convArray[i] = Environment10EnumToValue(environment10Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environment10EnumFromValue(value string) Environment10Enum {
    switch value {
        case "kVMware":
            return Environment10_KVMWARE
        case "kHyperV":
            return Environment10_KHYPERV
        case "kSQL":
            return Environment10_KSQL
        case "kView":
            return Environment10_KVIEW
        case "kPuppeteer":
            return Environment10_KPUPPETEER
        case "kPhysical":
            return Environment10_KPHYSICAL
        case "kPure":
            return Environment10_KPURE
        case "kAzure":
            return Environment10_KAZURE
        case "kNetapp":
            return Environment10_KNETAPP
        case "kAgent":
            return Environment10_KAGENT
        case "kGenericNas":
            return Environment10_KGENERICNAS
        case "kAcropolis":
            return Environment10_KACROPOLIS
        case "kPhysicalFiles":
            return Environment10_KPHYSICALFILES
        case "kIsilon":
            return Environment10_KISILON
        case "kKVM":
            return Environment10_KKVM
        case "kAWS":
            return Environment10_KAWS
        case "kExchange":
            return Environment10_KEXCHANGE
        case "kHyperVVSS":
            return Environment10_KHYPERVVSS
        case "kOracle":
            return Environment10_KORACLE
        case "kGCP":
            return Environment10_KGCP
        case "kFlashBlade":
            return Environment10_KFLASHBLADE
        case "kAWSNative":
            return Environment10_KAWSNATIVE
        case "kVCD":
            return Environment10_KVCD
        case "kO365":
            return Environment10_KO365
        case "kO365Outlook":
            return Environment10_KO365OUTLOOK
        case "kHyperFlex":
            return Environment10_KHYPERFLEX
        case "kGCPNative":
            return Environment10_KGCPNATIVE
        case "kAzureNative":
            return Environment10_KAZURENATIVE
        default:
            return Environment10_KVMWARE
    }
}
