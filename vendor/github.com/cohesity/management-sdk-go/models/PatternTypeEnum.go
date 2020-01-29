package models

import(
    "encoding/json"
)

/**
 * Type definition for PatternTypeEnum enum
 */
type PatternTypeEnum int

/**
 * Value collection for PatternTypeEnum enum
 */
const (
    PatternType_REGULAR PatternTypeEnum = 1 + iota
    PatternType_TEMPLATE
)

func (r PatternTypeEnum) MarshalJSON() ([]byte, error) { 
    s := PatternTypeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *PatternTypeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  PatternTypeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts PatternTypeEnum to its string representation
 */
func PatternTypeEnumToValue(patternTypeEnum PatternTypeEnum) string {
    switch patternTypeEnum {
        case PatternType_REGULAR:
    		return "REGULAR"		
        case PatternType_TEMPLATE:
    		return "TEMPLATE"		
        default:
        	return "REGULAR"
    }
}

/**
 * Converts PatternTypeEnum Array to its string Array representation
*/
func PatternTypeEnumArrayToValue(patternTypeEnum []PatternTypeEnum) []string {
    convArray := make([]string,len( patternTypeEnum))
    for i:=0; i<len(patternTypeEnum);i++ {
        convArray[i] = PatternTypeEnumToValue(patternTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PatternTypeEnumFromValue(value string) PatternTypeEnum {
    switch value {
        case "REGULAR":
            return PatternType_REGULAR
        case "TEMPLATE":
            return PatternType_TEMPLATE
        default:
            return PatternType_REGULAR
    }
}
