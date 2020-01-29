package models

import(
    "encoding/json"
)

/**
 * Type definition for Severity1Enum enum
 */
type Severity1Enum int

/**
 * Value collection for Severity1Enum enum
 */
const (
    Severity1_KCRITICAL Severity1Enum = 1 + iota
    Severity1_KWARNING
    Severity1_KINFO
)

func (r Severity1Enum) MarshalJSON() ([]byte, error) { 
    s := Severity1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Severity1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Severity1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Severity1Enum to its string representation
 */
func Severity1EnumToValue(severity1Enum Severity1Enum) string {
    switch severity1Enum {
        case Severity1_KCRITICAL:
    		return "kCritical"		
        case Severity1_KWARNING:
    		return "kWarning"		
        case Severity1_KINFO:
    		return "kInfo"		
        default:
        	return "kCritical"
    }
}

/**
 * Converts Severity1Enum Array to its string Array representation
*/
func Severity1EnumArrayToValue(severity1Enum []Severity1Enum) []string {
    convArray := make([]string,len( severity1Enum))
    for i:=0; i<len(severity1Enum);i++ {
        convArray[i] = Severity1EnumToValue(severity1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Severity1EnumFromValue(value string) Severity1Enum {
    switch value {
        case "kCritical":
            return Severity1_KCRITICAL
        case "kWarning":
            return Severity1_KWARNING
        case "kInfo":
            return Severity1_KINFO
        default:
            return Severity1_KCRITICAL
    }
}
