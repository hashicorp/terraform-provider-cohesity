package models

import(
    "encoding/json"
)

/**
 * Type definition for NasProtocol1Enum enum
 */
type NasProtocol1Enum int

/**
 * Value collection for NasProtocol1Enum enum
 */
const (
    NasProtocol1_KNFS3 NasProtocol1Enum = 1 + iota
    NasProtocol1_KCIFS1
)

func (r NasProtocol1Enum) MarshalJSON() ([]byte, error) { 
    s := NasProtocol1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *NasProtocol1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  NasProtocol1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts NasProtocol1Enum to its string representation
 */
func NasProtocol1EnumToValue(nasProtocol1Enum NasProtocol1Enum) string {
    switch nasProtocol1Enum {
        case NasProtocol1_KNFS3:
    		return "kNfs3"		
        case NasProtocol1_KCIFS1:
    		return "kCifs1"		
        default:
        	return "kNfs3"
    }
}

/**
 * Converts NasProtocol1Enum Array to its string Array representation
*/
func NasProtocol1EnumArrayToValue(nasProtocol1Enum []NasProtocol1Enum) []string {
    convArray := make([]string,len( nasProtocol1Enum))
    for i:=0; i<len(nasProtocol1Enum);i++ {
        convArray[i] = NasProtocol1EnumToValue(nasProtocol1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NasProtocol1EnumFromValue(value string) NasProtocol1Enum {
    switch value {
        case "kNfs3":
            return NasProtocol1_KNFS3
        case "kCifs1":
            return NasProtocol1_KCIFS1
        default:
            return NasProtocol1_KNFS3
    }
}
