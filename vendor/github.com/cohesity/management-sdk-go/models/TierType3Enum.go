package models

import(
    "encoding/json"
)

/**
 * Type definition for TierType3Enum enum
 */
type TierType3Enum int

/**
 * Value collection for TierType3Enum enum
 */
const (
    TierType3_KORACLETIERSTANDARD TierType3Enum = 1 + iota
    TierType3_KORACLETIERARCHIVE
)

func (r TierType3Enum) MarshalJSON() ([]byte, error) { 
    s := TierType3EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *TierType3Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  TierType3EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts TierType3Enum to its string representation
 */
func TierType3EnumToValue(tierType3Enum TierType3Enum) string {
    switch tierType3Enum {
        case TierType3_KORACLETIERSTANDARD:
    		return "kOracleTierStandard"		
        case TierType3_KORACLETIERARCHIVE:
    		return "kOracleTierArchive"		
        default:
        	return "kOracleTierStandard"
    }
}

/**
 * Converts TierType3Enum Array to its string Array representation
*/
func TierType3EnumArrayToValue(tierType3Enum []TierType3Enum) []string {
    convArray := make([]string,len( tierType3Enum))
    for i:=0; i<len(tierType3Enum);i++ {
        convArray[i] = TierType3EnumToValue(tierType3Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TierType3EnumFromValue(value string) TierType3Enum {
    switch value {
        case "kOracleTierStandard":
            return TierType3_KORACLETIERSTANDARD
        case "kOracleTierArchive":
            return TierType3_KORACLETIERARCHIVE
        default:
            return TierType3_KORACLETIERSTANDARD
    }
}
