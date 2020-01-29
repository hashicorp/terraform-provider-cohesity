package models

import(
    "encoding/json"
)

/**
 * Type definition for Type23Enum enum
 */
type Type23Enum int

/**
 * Value collection for Type23Enum enum
 */
const (
    Type23_KVMWARE Type23Enum = 1 + iota
    Type23_KHYPERV
    Type23_KSQL
    Type23_KVIEW
    Type23_KPUPPETEER
    Type23_KPHYSICAL
    Type23_KPURE
    Type23_KAZURE
    Type23_KNETAPP
    Type23_KAGENT
    Type23_KGENERICNAS
    Type23_KACROPOLIS
    Type23_KPHYSICALFILES
    Type23_KISILON
    Type23_KKVM
    Type23_KAWS
    Type23_KEXCHANGE
    Type23_KHYPERVVSS
    Type23_KORACLE
    Type23_KGCP
    Type23_KFLASHBLADE
    Type23_KAWSNATIVE
    Type23_KVCD
    Type23_KO365
    Type23_KO365OUTLOOK
    Type23_KHYPERFLEX
    Type23_KGCPNATIVE
    Type23_KAZURENATIVE
)

func (r Type23Enum) MarshalJSON() ([]byte, error) { 
    s := Type23EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type23Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type23EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type23Enum to its string representation
 */
func Type23EnumToValue(type23Enum Type23Enum) string {
    switch type23Enum {
        case Type23_KVMWARE:
    		return "kVMware"		
        case Type23_KHYPERV:
    		return "kHyperV"		
        case Type23_KSQL:
    		return "kSQL"		
        case Type23_KVIEW:
    		return "kView"		
        case Type23_KPUPPETEER:
    		return "kPuppeteer"		
        case Type23_KPHYSICAL:
    		return "kPhysical"		
        case Type23_KPURE:
    		return "kPure"		
        case Type23_KAZURE:
    		return "kAzure"		
        case Type23_KNETAPP:
    		return "kNetapp"		
        case Type23_KAGENT:
    		return "kAgent"		
        case Type23_KGENERICNAS:
    		return "kGenericNas"		
        case Type23_KACROPOLIS:
    		return "kAcropolis"		
        case Type23_KPHYSICALFILES:
    		return "kPhysicalFiles"		
        case Type23_KISILON:
    		return "kIsilon"		
        case Type23_KKVM:
    		return "kKVM"		
        case Type23_KAWS:
    		return "kAWS"		
        case Type23_KEXCHANGE:
    		return "kExchange"		
        case Type23_KHYPERVVSS:
    		return "kHyperVVSS"		
        case Type23_KORACLE:
    		return "kOracle"		
        case Type23_KGCP:
    		return "kGCP"		
        case Type23_KFLASHBLADE:
    		return "kFlashBlade"		
        case Type23_KAWSNATIVE:
    		return "kAWSNative"		
        case Type23_KVCD:
    		return "kVCD"		
        case Type23_KO365:
    		return "kO365"		
        case Type23_KO365OUTLOOK:
    		return "kO365Outlook"		
        case Type23_KHYPERFLEX:
    		return "kHyperFlex"		
        case Type23_KGCPNATIVE:
    		return "kGCPNative"		
        case Type23_KAZURENATIVE:
    		return "kAzureNative"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Type23Enum Array to its string Array representation
*/
func Type23EnumArrayToValue(type23Enum []Type23Enum) []string {
    convArray := make([]string,len( type23Enum))
    for i:=0; i<len(type23Enum);i++ {
        convArray[i] = Type23EnumToValue(type23Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type23EnumFromValue(value string) Type23Enum {
    switch value {
        case "kVMware":
            return Type23_KVMWARE
        case "kHyperV":
            return Type23_KHYPERV
        case "kSQL":
            return Type23_KSQL
        case "kView":
            return Type23_KVIEW
        case "kPuppeteer":
            return Type23_KPUPPETEER
        case "kPhysical":
            return Type23_KPHYSICAL
        case "kPure":
            return Type23_KPURE
        case "kAzure":
            return Type23_KAZURE
        case "kNetapp":
            return Type23_KNETAPP
        case "kAgent":
            return Type23_KAGENT
        case "kGenericNas":
            return Type23_KGENERICNAS
        case "kAcropolis":
            return Type23_KACROPOLIS
        case "kPhysicalFiles":
            return Type23_KPHYSICALFILES
        case "kIsilon":
            return Type23_KISILON
        case "kKVM":
            return Type23_KKVM
        case "kAWS":
            return Type23_KAWS
        case "kExchange":
            return Type23_KEXCHANGE
        case "kHyperVVSS":
            return Type23_KHYPERVVSS
        case "kOracle":
            return Type23_KORACLE
        case "kGCP":
            return Type23_KGCP
        case "kFlashBlade":
            return Type23_KFLASHBLADE
        case "kAWSNative":
            return Type23_KAWSNATIVE
        case "kVCD":
            return Type23_KVCD
        case "kO365":
            return Type23_KO365
        case "kO365Outlook":
            return Type23_KO365OUTLOOK
        case "kHyperFlex":
            return Type23_KHYPERFLEX
        case "kGCPNative":
            return Type23_KGCPNATIVE
        case "kAzureNative":
            return Type23_KAZURENATIVE
        default:
            return Type23_KVMWARE
    }
}
