package models

import(
    "encoding/json"
)

/**
 * Type definition for Day1Enum enum
 */
type Day1Enum int

/**
 * Value collection for Day1Enum enum
 */
const (
    Day1_KSUNDAY Day1Enum = 1 + iota
    Day1_KMONDAY
    Day1_KTUESDAY
    Day1_KWEDNESDAY
    Day1_KTHURSDAY
    Day1_KFRIDAY
    Day1_KSATURDAY
)

func (r Day1Enum) MarshalJSON() ([]byte, error) { 
    s := Day1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Day1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Day1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Day1Enum to its string representation
 */
func Day1EnumToValue(day1Enum Day1Enum) string {
    switch day1Enum {
        case Day1_KSUNDAY:
    		return "kSunday"		
        case Day1_KMONDAY:
    		return "kMonday"		
        case Day1_KTUESDAY:
    		return "kTuesday"		
        case Day1_KWEDNESDAY:
    		return "kWednesday"		
        case Day1_KTHURSDAY:
    		return "kThursday"		
        case Day1_KFRIDAY:
    		return "kFriday"		
        case Day1_KSATURDAY:
    		return "kSaturday"		
        default:
        	return "kSunday"
    }
}

/**
 * Converts Day1Enum Array to its string Array representation
*/
func Day1EnumArrayToValue(day1Enum []Day1Enum) []string {
    convArray := make([]string,len( day1Enum))
    for i:=0; i<len(day1Enum);i++ {
        convArray[i] = Day1EnumToValue(day1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Day1EnumFromValue(value string) Day1Enum {
    switch value {
        case "kSunday":
            return Day1_KSUNDAY
        case "kMonday":
            return Day1_KMONDAY
        case "kTuesday":
            return Day1_KTUESDAY
        case "kWednesday":
            return Day1_KWEDNESDAY
        case "kThursday":
            return Day1_KTHURSDAY
        case "kFriday":
            return Day1_KFRIDAY
        case "kSaturday":
            return Day1_KSATURDAY
        default:
            return Day1_KSUNDAY
    }
}
