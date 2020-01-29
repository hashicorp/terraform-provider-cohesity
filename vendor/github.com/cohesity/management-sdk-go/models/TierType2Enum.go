package models

import(
    "encoding/json"
)

/**
 * Type definition for TierType2Enum enum
 */
type TierType2Enum int

/**
 * Value collection for TierType2Enum enum
 */
const (
    TierType2_KGOOGLESTANDARD TierType2Enum = 1 + iota
    TierType2_KGOOGLENEARLINE
    TierType2_KGOOGLECOLDLINE
    TierType2_KGOOGLEREGIONAL
    TierType2_KGOOGLEMULTIREGIONAL
)

func (r TierType2Enum) MarshalJSON() ([]byte, error) { 
    s := TierType2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *TierType2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  TierType2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts TierType2Enum to its string representation
 */
func TierType2EnumToValue(tierType2Enum TierType2Enum) string {
    switch tierType2Enum {
        case TierType2_KGOOGLESTANDARD:
    		return "kGoogleStandard"		
        case TierType2_KGOOGLENEARLINE:
    		return "kGoogleNearline"		
        case TierType2_KGOOGLECOLDLINE:
    		return "kGoogleColdline"		
        case TierType2_KGOOGLEREGIONAL:
    		return "kGoogleRegional"		
        case TierType2_KGOOGLEMULTIREGIONAL:
    		return "kGoogleMultiRegional"		
        default:
        	return "kGoogleStandard"
    }
}

/**
 * Converts TierType2Enum Array to its string Array representation
*/
func TierType2EnumArrayToValue(tierType2Enum []TierType2Enum) []string {
    convArray := make([]string,len( tierType2Enum))
    for i:=0; i<len(tierType2Enum);i++ {
        convArray[i] = TierType2EnumToValue(tierType2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TierType2EnumFromValue(value string) TierType2Enum {
    switch value {
        case "kGoogleStandard":
            return TierType2_KGOOGLESTANDARD
        case "kGoogleNearline":
            return TierType2_KGOOGLENEARLINE
        case "kGoogleColdline":
            return TierType2_KGOOGLECOLDLINE
        case "kGoogleRegional":
            return TierType2_KGOOGLEREGIONAL
        case "kGoogleMultiRegional":
            return TierType2_KGOOGLEMULTIREGIONAL
        default:
            return TierType2_KGOOGLESTANDARD
    }
}
