package models

import(
    "encoding/json"
)

/**
 * Type definition for EncryptionPolicy1Enum enum
 */
type EncryptionPolicy1Enum int

/**
 * Value collection for EncryptionPolicy1Enum enum
 */
const (
    EncryptionPolicy1_KENCRYPTIONNONE EncryptionPolicy1Enum = 1 + iota
    EncryptionPolicy1_KENCRYPTIONSTRONG
    EncryptionPolicy1_KENCRYPTIONWEAK
)

func (r EncryptionPolicy1Enum) MarshalJSON() ([]byte, error) { 
    s := EncryptionPolicy1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *EncryptionPolicy1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  EncryptionPolicy1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts EncryptionPolicy1Enum to its string representation
 */
func EncryptionPolicy1EnumToValue(encryptionPolicy1Enum EncryptionPolicy1Enum) string {
    switch encryptionPolicy1Enum {
        case EncryptionPolicy1_KENCRYPTIONNONE:
    		return "kEncryptionNone"		
        case EncryptionPolicy1_KENCRYPTIONSTRONG:
    		return "kEncryptionStrong"		
        case EncryptionPolicy1_KENCRYPTIONWEAK:
    		return "kEncryptionWeak"		
        default:
        	return "kEncryptionNone"
    }
}

/**
 * Converts EncryptionPolicy1Enum Array to its string Array representation
*/
func EncryptionPolicy1EnumArrayToValue(encryptionPolicy1Enum []EncryptionPolicy1Enum) []string {
    convArray := make([]string,len( encryptionPolicy1Enum))
    for i:=0; i<len(encryptionPolicy1Enum);i++ {
        convArray[i] = EncryptionPolicy1EnumToValue(encryptionPolicy1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EncryptionPolicy1EnumFromValue(value string) EncryptionPolicy1Enum {
    switch value {
        case "kEncryptionNone":
            return EncryptionPolicy1_KENCRYPTIONNONE
        case "kEncryptionStrong":
            return EncryptionPolicy1_KENCRYPTIONSTRONG
        case "kEncryptionWeak":
            return EncryptionPolicy1_KENCRYPTIONWEAK
        default:
            return EncryptionPolicy1_KENCRYPTIONNONE
    }
}
