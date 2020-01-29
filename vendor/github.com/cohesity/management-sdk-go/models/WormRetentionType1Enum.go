package models

import(
    "encoding/json"
)

/**
 * Type definition for WormRetentionType1Enum enum
 */
type WormRetentionType1Enum int

/**
 * Value collection for WormRetentionType1Enum enum
 */
const (
    WormRetentionType1_KNONE WormRetentionType1Enum = 1 + iota
    WormRetentionType1_KCOMPLIANCE
    WormRetentionType1_KADMINISTRATIVE
)

func (r WormRetentionType1Enum) MarshalJSON() ([]byte, error) { 
    s := WormRetentionType1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *WormRetentionType1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  WormRetentionType1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts WormRetentionType1Enum to its string representation
 */
func WormRetentionType1EnumToValue(wormRetentionType1Enum WormRetentionType1Enum) string {
    switch wormRetentionType1Enum {
        case WormRetentionType1_KNONE:
    		return "kNone"		
        case WormRetentionType1_KCOMPLIANCE:
    		return "kCompliance"		
        case WormRetentionType1_KADMINISTRATIVE:
    		return "kAdministrative"		
        default:
        	return "kNone"
    }
}

/**
 * Converts WormRetentionType1Enum Array to its string Array representation
*/
func WormRetentionType1EnumArrayToValue(wormRetentionType1Enum []WormRetentionType1Enum) []string {
    convArray := make([]string,len( wormRetentionType1Enum))
    for i:=0; i<len(wormRetentionType1Enum);i++ {
        convArray[i] = WormRetentionType1EnumToValue(wormRetentionType1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func WormRetentionType1EnumFromValue(value string) WormRetentionType1Enum {
    switch value {
        case "kNone":
            return WormRetentionType1_KNONE
        case "kCompliance":
            return WormRetentionType1_KCOMPLIANCE
        case "kAdministrative":
            return WormRetentionType1_KADMINISTRATIVE
        default:
            return WormRetentionType1_KNONE
    }
}
