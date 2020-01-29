package models

import(
    "encoding/json"
)

/**
 * Type definition for RemovalState1Enum enum
 */
type RemovalState1Enum int

/**
 * Value collection for RemovalState1Enum enum
 */
const (
    RemovalState1_KDONTREMOVE RemovalState1Enum = 1 + iota
    RemovalState1_KMARKEDFORREMOVAL
    RemovalState1_KOKTOREMOVE
)

func (r RemovalState1Enum) MarshalJSON() ([]byte, error) { 
    s := RemovalState1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *RemovalState1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  RemovalState1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts RemovalState1Enum to its string representation
 */
func RemovalState1EnumToValue(removalState1Enum RemovalState1Enum) string {
    switch removalState1Enum {
        case RemovalState1_KDONTREMOVE:
    		return "kDontRemove"		
        case RemovalState1_KMARKEDFORREMOVAL:
    		return "kMarkedForRemoval"		
        case RemovalState1_KOKTOREMOVE:
    		return "kOkToRemove"		
        default:
        	return "kDontRemove"
    }
}

/**
 * Converts RemovalState1Enum Array to its string Array representation
*/
func RemovalState1EnumArrayToValue(removalState1Enum []RemovalState1Enum) []string {
    convArray := make([]string,len( removalState1Enum))
    for i:=0; i<len(removalState1Enum);i++ {
        convArray[i] = RemovalState1EnumToValue(removalState1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RemovalState1EnumFromValue(value string) RemovalState1Enum {
    switch value {
        case "kDontRemove":
            return RemovalState1_KDONTREMOVE
        case "kMarkedForRemoval":
            return RemovalState1_KMARKEDFORREMOVAL
        case "kOkToRemove":
            return RemovalState1_KOKTOREMOVE
        default:
            return RemovalState1_KDONTREMOVE
    }
}
