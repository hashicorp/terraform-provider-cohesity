package models

import(
    "encoding/json"
)

/**
 * Type definition for Status7Enum enum
 */
type Status7Enum int

/**
 * Value collection for Status7Enum enum
 */
const (
    Status7_ACTIVE Status7Enum = 1 + iota
    Status7_DEACTIVATED
    Status7_DELETED
)

func (r Status7Enum) MarshalJSON() ([]byte, error) { 
    s := Status7EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Status7Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Status7EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Status7Enum to its string representation
 */
func Status7EnumToValue(status7Enum Status7Enum) string {
    switch status7Enum {
        case Status7_ACTIVE:
    		return "Active"		
        case Status7_DEACTIVATED:
    		return "Deactivated"		
        case Status7_DELETED:
    		return "Deleted"		
        default:
        	return "Active"
    }
}

/**
 * Converts Status7Enum Array to its string Array representation
*/
func Status7EnumArrayToValue(status7Enum []Status7Enum) []string {
    convArray := make([]string,len( status7Enum))
    for i:=0; i<len(status7Enum);i++ {
        convArray[i] = Status7EnumToValue(status7Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Status7EnumFromValue(value string) Status7Enum {
    switch value {
        case "Active":
            return Status7_ACTIVE
        case "Deactivated":
            return Status7_DEACTIVATED
        case "Deleted":
            return Status7_DELETED
        default:
            return Status7_ACTIVE
    }
}
