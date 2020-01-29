package models

import(
    "encoding/json"
)

/**
 * Type definition for Environments2Enum enum
 */
type Environments2Enum int

/**
 * Value collection for Environments2Enum enum
 */
const (
    Environments2_KVMWARE Environments2Enum = 1 + iota
    Environments2_KSQL
    Environments2_KVIEW
    Environments2_KPUPPETEER
    Environments2_KPHYSICAL
    Environments2_KPURE
    Environments2_KNETAPP
    Environments2_KGENERICNAS
    Environments2_KHYPERV
    Environments2_KACROPOLIS
    Environments2_KAZURE
)

func (r Environments2Enum) MarshalJSON() ([]byte, error) { 
    s := Environments2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Environments2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Environments2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Environments2Enum to its string representation
 */
func Environments2EnumToValue(environments2Enum Environments2Enum) string {
    switch environments2Enum {
        case Environments2_KVMWARE:
    		return "kVMware"		
        case Environments2_KSQL:
    		return "kSQL"		
        case Environments2_KVIEW:
    		return "kView"		
        case Environments2_KPUPPETEER:
    		return "kPuppeteer"		
        case Environments2_KPHYSICAL:
    		return "kPhysical"		
        case Environments2_KPURE:
    		return "kPure"		
        case Environments2_KNETAPP:
    		return "kNetapp"		
        case Environments2_KGENERICNAS:
    		return "kGenericNas"		
        case Environments2_KHYPERV:
    		return "kHyperV"		
        case Environments2_KACROPOLIS:
    		return "kAcropolis"		
        case Environments2_KAZURE:
    		return "kAzure"		
        default:
        	return "kVMware"
    }
}

/**
 * Converts Environments2Enum Array to its string Array representation
*/
func Environments2EnumArrayToValue(environments2Enum []Environments2Enum) []string {
    convArray := make([]string,len( environments2Enum))
    for i:=0; i<len(environments2Enum);i++ {
        convArray[i] = Environments2EnumToValue(environments2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Environments2EnumFromValue(value string) Environments2Enum {
    switch value {
        case "kVMware":
            return Environments2_KVMWARE
        case "kSQL":
            return Environments2_KSQL
        case "kView":
            return Environments2_KVIEW
        case "kPuppeteer":
            return Environments2_KPUPPETEER
        case "kPhysical":
            return Environments2_KPHYSICAL
        case "kPure":
            return Environments2_KPURE
        case "kNetapp":
            return Environments2_KNETAPP
        case "kGenericNas":
            return Environments2_KGENERICNAS
        case "kHyperV":
            return Environments2_KHYPERV
        case "kAcropolis":
            return Environments2_KACROPOLIS
        case "kAzure":
            return Environments2_KAZURE
        default:
            return Environments2_KVMWARE
    }
}
