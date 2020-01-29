package models

import(
    "encoding/json"
)

/**
 * Type definition for TierType1Enum enum
 */
type TierType1Enum int

/**
 * Value collection for TierType1Enum enum
 */
const (
    TierType1_KAZURETIERHOT TierType1Enum = 1 + iota
    TierType1_KAZURETIERCOOL
    TierType1_KAZURETIERARCHIVE
)

func (r TierType1Enum) MarshalJSON() ([]byte, error) { 
    s := TierType1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *TierType1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  TierType1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts TierType1Enum to its string representation
 */
func TierType1EnumToValue(tierType1Enum TierType1Enum) string {
    switch tierType1Enum {
        case TierType1_KAZURETIERHOT:
    		return "kAzureTierHot"		
        case TierType1_KAZURETIERCOOL:
    		return "kAzureTierCool"		
        case TierType1_KAZURETIERARCHIVE:
    		return "kAzureTierArchive"		
        default:
        	return "kAzureTierHot"
    }
}

/**
 * Converts TierType1Enum Array to its string Array representation
*/
func TierType1EnumArrayToValue(tierType1Enum []TierType1Enum) []string {
    convArray := make([]string,len( tierType1Enum))
    for i:=0; i<len(tierType1Enum);i++ {
        convArray[i] = TierType1EnumToValue(tierType1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TierType1EnumFromValue(value string) TierType1Enum {
    switch value {
        case "kAzureTierHot":
            return TierType1_KAZURETIERHOT
        case "kAzureTierCool":
            return TierType1_KAZURETIERCOOL
        case "kAzureTierArchive":
            return TierType1_KAZURETIERARCHIVE
        default:
            return TierType1_KAZURETIERHOT
    }
}
