package models

import(
    "encoding/json"
)

/**
 * Type definition for Upgradability1Enum enum
 */
type Upgradability1Enum int

/**
 * Value collection for Upgradability1Enum enum
 */
const (
    Upgradability1_KUPGRADABLE Upgradability1Enum = 1 + iota
    Upgradability1_KCURRENT
    Upgradability1_KUNKNOWN
    Upgradability1_KNONUPGRADABLEINVALIDVERSION
    Upgradability1_KNONUPGRADABLEAGENTISNEWER
    Upgradability1_KNONUPGRADABLEAGENTISOLD
)

func (r Upgradability1Enum) MarshalJSON() ([]byte, error) { 
    s := Upgradability1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Upgradability1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Upgradability1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Upgradability1Enum to its string representation
 */
func Upgradability1EnumToValue(upgradability1Enum Upgradability1Enum) string {
    switch upgradability1Enum {
        case Upgradability1_KUPGRADABLE:
    		return "kUpgradable"		
        case Upgradability1_KCURRENT:
    		return "kCurrent"		
        case Upgradability1_KUNKNOWN:
    		return "kUnknown"		
        case Upgradability1_KNONUPGRADABLEINVALIDVERSION:
    		return "kNonUpgradableInvalidVersion"		
        case Upgradability1_KNONUPGRADABLEAGENTISNEWER:
    		return "kNonUpgradableAgentIsNewer"		
        case Upgradability1_KNONUPGRADABLEAGENTISOLD:
    		return "kNonUpgradableAgentIsOld"		
        default:
        	return "kUpgradable"
    }
}

/**
 * Converts Upgradability1Enum Array to its string Array representation
*/
func Upgradability1EnumArrayToValue(upgradability1Enum []Upgradability1Enum) []string {
    convArray := make([]string,len( upgradability1Enum))
    for i:=0; i<len(upgradability1Enum);i++ {
        convArray[i] = Upgradability1EnumToValue(upgradability1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Upgradability1EnumFromValue(value string) Upgradability1Enum {
    switch value {
        case "kUpgradable":
            return Upgradability1_KUPGRADABLE
        case "kCurrent":
            return Upgradability1_KCURRENT
        case "kUnknown":
            return Upgradability1_KUNKNOWN
        case "kNonUpgradableInvalidVersion":
            return Upgradability1_KNONUPGRADABLEINVALIDVERSION
        case "kNonUpgradableAgentIsNewer":
            return Upgradability1_KNONUPGRADABLEAGENTISNEWER
        case "kNonUpgradableAgentIsOld":
            return Upgradability1_KNONUPGRADABLEAGENTISOLD
        default:
            return Upgradability1_KUPGRADABLE
    }
}
