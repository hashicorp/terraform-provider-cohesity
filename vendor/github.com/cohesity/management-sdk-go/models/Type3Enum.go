package models

import(
    "encoding/json"
)

/**
 * Type definition for Type3Enum enum
 */
type Type3Enum int

/**
 * Value collection for Type3Enum enum
 */
const (
    Type3_KSTORAGEARRAY Type3Enum = 1 + iota
    Type3_KFILESYSTEM
)

func (r Type3Enum) MarshalJSON() ([]byte, error) { 
    s := Type3EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Type3Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Type3EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Type3Enum to its string representation
 */
func Type3EnumToValue(type3Enum Type3Enum) string {
    switch type3Enum {
        case Type3_KSTORAGEARRAY:
    		return "kStorageArray"		
        case Type3_KFILESYSTEM:
    		return "kFileSystem"		
        default:
        	return "kStorageArray"
    }
}

/**
 * Converts Type3Enum Array to its string Array representation
*/
func Type3EnumArrayToValue(type3Enum []Type3Enum) []string {
    convArray := make([]string,len( type3Enum))
    for i:=0; i<len(type3Enum);i++ {
        convArray[i] = Type3EnumToValue(type3Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Type3EnumFromValue(value string) Type3Enum {
    switch value {
        case "kStorageArray":
            return Type3_KSTORAGEARRAY
        case "kFileSystem":
            return Type3_KFILESYSTEM
        default:
            return Type3_KSTORAGEARRAY
    }
}
