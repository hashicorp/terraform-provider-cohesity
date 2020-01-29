package models

import(
    "encoding/json"
)

/**
 * Type definition for Type12Enum enum
 */
type Type12Enum int

/**
 * Value collection for Type12Enum enum
 */
const (
    Type12_KREADWRITE Type12Enum = 1 + iota
    Type12_KLOADSHARING
    Type12_KDATAPROTECTION
    Type12_KDATACACHE
    Type12_KTMP
    Type12_KUNKNOWNTYPE
)

func (r Type12Enum) MarshalJSON() ([]byte, error) { 
    s := Type12EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type12Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type12EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type12Enum to its string representation
 */
func Type12EnumToValue(type12Enum Type12Enum) string {
    switch type12Enum {
        case Type12_KREADWRITE:
    		return "kReadWrite"		
        case Type12_KLOADSHARING:
    		return "kLoadSharing"		
        case Type12_KDATAPROTECTION:
    		return "kDataProtection"		
        case Type12_KDATACACHE:
    		return "kDataCache"		
        case Type12_KTMP:
    		return "kTmp"		
        case Type12_KUNKNOWNTYPE:
    		return "kUnknownType"		
        default:
        	return "kReadWrite"
    }
}

/**
 * Converts Type12Enum Array to its string Array representation
*/
func Type12EnumArrayToValue(type12Enum []Type12Enum) []string {
    convArray := make([]string,len( type12Enum))
    for i:=0; i<len(type12Enum);i++ {
        convArray[i] = Type12EnumToValue(type12Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type12EnumFromValue(value string) Type12Enum {
    switch value {
        case "kReadWrite":
            return Type12_KREADWRITE
        case "kLoadSharing":
            return Type12_KLOADSHARING
        case "kDataProtection":
            return Type12_KDATAPROTECTION
        case "kDataCache":
            return Type12_KDATACACHE
        case "kTmp":
            return Type12_KTMP
        case "kUnknownType":
            return Type12_KUNKNOWNTYPE
        default:
            return Type12_KREADWRITE
    }
}
