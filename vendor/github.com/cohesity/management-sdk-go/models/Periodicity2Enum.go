package models

import(
    "encoding/json"
)

/**
 * Type definition for Periodicity2Enum enum
 */
type Periodicity2Enum int

/**
 * Value collection for Periodicity2Enum enum
 */
const (
    Periodicity2_KCONTINUOUS Periodicity2Enum = 1 + iota
    Periodicity2_KDAILY
    Periodicity2_KMONTHLY
    Periodicity2_KONEOFF
)

func (r Periodicity2Enum) MarshalJSON() ([]byte, error) { 
    s := Periodicity2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Periodicity2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Periodicity2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Periodicity2Enum to its string representation
 */
func Periodicity2EnumToValue(periodicity2Enum Periodicity2Enum) string {
    switch periodicity2Enum {
        case Periodicity2_KCONTINUOUS:
    		return "kContinuous"		
        case Periodicity2_KDAILY:
    		return "kDaily"		
        case Periodicity2_KMONTHLY:
    		return "kMonthly"		
        case Periodicity2_KONEOFF:
    		return "kOneOff"		
        default:
        	return "kContinuous"
    }
}

/**
 * Converts Periodicity2Enum Array to its string Array representation
*/
func Periodicity2EnumArrayToValue(periodicity2Enum []Periodicity2Enum) []string {
    convArray := make([]string,len( periodicity2Enum))
    for i:=0; i<len(periodicity2Enum);i++ {
        convArray[i] = Periodicity2EnumToValue(periodicity2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Periodicity2EnumFromValue(value string) Periodicity2Enum {
    switch value {
        case "kContinuous":
            return Periodicity2_KCONTINUOUS
        case "kDaily":
            return Periodicity2_KDAILY
        case "kMonthly":
            return Periodicity2_KMONTHLY
        case "kOneOff":
            return Periodicity2_KONEOFF
        default:
            return Periodicity2_KCONTINUOUS
    }
}
