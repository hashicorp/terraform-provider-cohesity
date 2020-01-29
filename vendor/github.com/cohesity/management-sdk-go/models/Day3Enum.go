package models

import(
    "encoding/json"
)

/**
 * Type definition for Day3Enum enum
 */
type Day3Enum int

/**
 * Value collection for Day3Enum enum
 */
const (
    Day3_KSUNDAY Day3Enum = 1 + iota
    Day3_KMONDAY
    Day3_KTUESDAY
    Day3_KWEDNESDAY
    Day3_KTHURSDAY
    Day3_KFRIDAY
    Day3_KSATURDAY
)

func (r Day3Enum) MarshalJSON() ([]byte, error) { 
    s := Day3EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Day3Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Day3EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Day3Enum to its string representation
 */
func Day3EnumToValue(day3Enum Day3Enum) string {
    switch day3Enum {
        case Day3_KSUNDAY:
    		return "kSunday"		
        case Day3_KMONDAY:
    		return "kMonday"		
        case Day3_KTUESDAY:
    		return "kTuesday"		
        case Day3_KWEDNESDAY:
    		return "kWednesday"		
        case Day3_KTHURSDAY:
    		return "kThursday"		
        case Day3_KFRIDAY:
    		return "kFriday"		
        case Day3_KSATURDAY:
    		return "kSaturday"		
        default:
        	return "kSunday"
    }
}

/**
 * Converts Day3Enum Array to its string Array representation
*/
func Day3EnumArrayToValue(day3Enum []Day3Enum) []string {
    convArray := make([]string,len( day3Enum))
    for i:=0; i<len(day3Enum);i++ {
        convArray[i] = Day3EnumToValue(day3Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Day3EnumFromValue(value string) Day3Enum {
    switch value {
        case "kSunday":
            return Day3_KSUNDAY
        case "kMonday":
            return Day3_KMONDAY
        case "kTuesday":
            return Day3_KTUESDAY
        case "kWednesday":
            return Day3_KWEDNESDAY
        case "kThursday":
            return Day3_KTHURSDAY
        case "kFriday":
            return Day3_KFRIDAY
        case "kSaturday":
            return Day3_KSATURDAY
        default:
            return Day3_KSUNDAY
    }
}
