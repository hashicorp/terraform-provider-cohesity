package models

import(
    "encoding/json"
)

/**
 * Type definition for Type20Enum enum
 */
type Type20Enum int

/**
 * Value collection for Type20Enum enum
 */
const (
    Type20_KRID Type20Enum = 1 + iota
    Type20_KRFC2307
    Type20_KSFU30
    Type20_KCENTRIFY
    Type20_KFIXED
    Type20_KCUSTOMATTRIBUTES
    Type20_KLDAPPROVIDER
)

func (r Type20Enum) MarshalJSON() ([]byte, error) { 
    s := Type20EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type20Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type20EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type20Enum to its string representation
 */
func Type20EnumToValue(type20Enum Type20Enum) string {
    switch type20Enum {
        case Type20_KRID:
    		return "kRid"		
        case Type20_KRFC2307:
    		return "kRfc2307"		
        case Type20_KSFU30:
    		return "kSfu30"		
        case Type20_KCENTRIFY:
    		return "kCentrify"		
        case Type20_KFIXED:
    		return "kFixed"		
        case Type20_KCUSTOMATTRIBUTES:
    		return "kCustomAttributes"		
        case Type20_KLDAPPROVIDER:
    		return "kLdapProvider"		
        default:
        	return "kRid"
    }
}

/**
 * Converts Type20Enum Array to its string Array representation
*/
func Type20EnumArrayToValue(type20Enum []Type20Enum) []string {
    convArray := make([]string,len( type20Enum))
    for i:=0; i<len(type20Enum);i++ {
        convArray[i] = Type20EnumToValue(type20Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type20EnumFromValue(value string) Type20Enum {
    switch value {
        case "kRid":
            return Type20_KRID
        case "kRfc2307":
            return Type20_KRFC2307
        case "kSfu30":
            return Type20_KSFU30
        case "kCentrify":
            return Type20_KCENTRIFY
        case "kFixed":
            return Type20_KFIXED
        case "kCustomAttributes":
            return Type20_KCUSTOMATTRIBUTES
        case "kLdapProvider":
            return Type20_KLDAPPROVIDER
        default:
            return Type20_KRID
    }
}
