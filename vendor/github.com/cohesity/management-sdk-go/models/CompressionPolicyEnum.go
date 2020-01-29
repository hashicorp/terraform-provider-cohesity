// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for CompressionPolicyEnum enum
 */
type CompressionPolicyEnum int

/**
 * Value collection for CompressionPolicyEnum enum
 */
const (
    CompressionPolicy_KCOMPRESSIONNONE CompressionPolicyEnum = 1 + iota
    CompressionPolicy_KCOMPRESSIONLOW
    CompressionPolicy_KCOMPRESSIONHIGH
)

func (r CompressionPolicyEnum) MarshalJSON() ([]byte, error) {
    s := CompressionPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *CompressionPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  CompressionPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts CompressionPolicyEnum to its string representation
 */
func CompressionPolicyEnumToValue(compressionPolicyEnum CompressionPolicyEnum) string {
    switch compressionPolicyEnum {
        case CompressionPolicy_KCOMPRESSIONNONE:
    		return "kCompressionNone"
        case CompressionPolicy_KCOMPRESSIONLOW:
    		return "kCompressionLow"
        case CompressionPolicy_KCOMPRESSIONHIGH:
    		return "kCompressionHigh"
        default:
        	return "kCompressionNone"
    }
}

/**
 * Converts CompressionPolicyEnum Array to its string Array representation
*/
func CompressionPolicyEnumArrayToValue(compressionPolicyEnum []CompressionPolicyEnum) []string {
    convArray := make([]string,len( compressionPolicyEnum))
    for i:=0; i<len(compressionPolicyEnum);i++ {
        convArray[i] = CompressionPolicyEnumToValue(compressionPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func CompressionPolicyEnumFromValue(value string) CompressionPolicyEnum {
    switch value {
        case "kCompressionNone":
            return CompressionPolicy_KCOMPRESSIONNONE
        case "kCompressionLow":
            return CompressionPolicy_KCOMPRESSIONLOW
        case "kCompressionHigh":
            return CompressionPolicy_KCOMPRESSIONHIGH
        default:
            return CompressionPolicy_KCOMPRESSIONNONE
    }
}
