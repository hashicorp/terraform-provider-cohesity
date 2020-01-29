package models

import(
    "encoding/json"
)

/**
 * Type definition for PrivilegesTypeEnum enum
 */
type PrivilegesTypeEnum int

/**
 * Value collection for PrivilegesTypeEnum enum
 */
const (
    PrivilegesType_KNONE PrivilegesTypeEnum = 1 + iota
    PrivilegesType_KALL
    PrivilegesType_KSPECIFIC
)

func (r PrivilegesTypeEnum) MarshalJSON() ([]byte, error) { 
    s := PrivilegesTypeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *PrivilegesTypeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  PrivilegesTypeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts PrivilegesTypeEnum to its string representation
 */
func PrivilegesTypeEnumToValue(privilegesTypeEnum PrivilegesTypeEnum) string {
    switch privilegesTypeEnum {
        case PrivilegesType_KNONE:
    		return "kNone"		
        case PrivilegesType_KALL:
    		return "kAll"		
        case PrivilegesType_KSPECIFIC:
    		return "kSpecific"		
        default:
        	return "kNone"
    }
}

/**
 * Converts PrivilegesTypeEnum Array to its string Array representation
*/
func PrivilegesTypeEnumArrayToValue(privilegesTypeEnum []PrivilegesTypeEnum) []string {
    convArray := make([]string,len( privilegesTypeEnum))
    for i:=0; i<len(privilegesTypeEnum);i++ {
        convArray[i] = PrivilegesTypeEnumToValue(privilegesTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func PrivilegesTypeEnumFromValue(value string) PrivilegesTypeEnum {
    switch value {
        case "kNone":
            return PrivilegesType_KNONE
        case "kAll":
            return PrivilegesType_KALL
        case "kSpecific":
            return PrivilegesType_KSPECIFIC
        default:
            return PrivilegesType_KNONE
    }
}
