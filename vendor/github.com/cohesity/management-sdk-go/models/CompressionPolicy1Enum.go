package models

import(
    "encoding/json"
)

/**
 * Type definition for CompressionPolicy1Enum enum
 */
type CompressionPolicy1Enum int

/**
 * Value collection for CompressionPolicy1Enum enum
 */
const (
    CompressionPolicy1_KCOMPRESSIONNONE CompressionPolicy1Enum = 1 + iota
    CompressionPolicy1_KCOMPRESSIONLOW
)

func (r CompressionPolicy1Enum) MarshalJSON() ([]byte, error) { 
    s := CompressionPolicy1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *CompressionPolicy1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  CompressionPolicy1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts CompressionPolicy1Enum to its string representation
 */
func CompressionPolicy1EnumToValue(compressionPolicy1Enum CompressionPolicy1Enum) string {
    switch compressionPolicy1Enum {
        case CompressionPolicy1_KCOMPRESSIONNONE:
    		return "kCompressionNone"		
        case CompressionPolicy1_KCOMPRESSIONLOW:
    		return "kCompressionLow"		
        default:
        	return "kCompressionNone"
    }
}

/**
 * Converts CompressionPolicy1Enum Array to its string Array representation
*/
func CompressionPolicy1EnumArrayToValue(compressionPolicy1Enum []CompressionPolicy1Enum) []string {
    convArray := make([]string,len( compressionPolicy1Enum))
    for i:=0; i<len(compressionPolicy1Enum);i++ {
        convArray[i] = CompressionPolicy1EnumToValue(compressionPolicy1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func CompressionPolicy1EnumFromValue(value string) CompressionPolicy1Enum {
    switch value {
        case "kCompressionNone":
            return CompressionPolicy1_KCOMPRESSIONNONE
        case "kCompressionLow":
            return CompressionPolicy1_KCOMPRESSIONLOW
        default:
            return CompressionPolicy1_KCOMPRESSIONNONE
    }
}
