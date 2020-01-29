// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for CompressionPolicyVaultEnum enum
 */
type CompressionPolicyVaultEnum int

/**
 * Value collection for CompressionPolicyVaultEnum enum
 */
const (
    CompressionPolicyVault_KCOMPRESSIONNONE CompressionPolicyVaultEnum = 1 + iota
    CompressionPolicyVault_KCOMPRESSIONLOW
    CompressionPolicyVault_KCOMPRESSIONHIGH
)

func (r CompressionPolicyVaultEnum) MarshalJSON() ([]byte, error) {
    s := CompressionPolicyVaultEnumToValue(r)
    return json.Marshal(s)
}

func (r *CompressionPolicyVaultEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  CompressionPolicyVaultEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts CompressionPolicyVaultEnum to its string representation
 */
func CompressionPolicyVaultEnumToValue(compressionPolicyVaultEnum CompressionPolicyVaultEnum) string {
    switch compressionPolicyVaultEnum {
        case CompressionPolicyVault_KCOMPRESSIONNONE:
    		return "kCompressionNone"
        case CompressionPolicyVault_KCOMPRESSIONLOW:
    		return "kCompressionLow"
        case CompressionPolicyVault_KCOMPRESSIONHIGH:
    		return "kCompressionHigh"
        default:
        	return "kCompressionNone"
    }
}

/**
 * Converts CompressionPolicyVaultEnum Array to its string Array representation
*/
func CompressionPolicyVaultEnumArrayToValue(compressionPolicyVaultEnum []CompressionPolicyVaultEnum) []string {
    convArray := make([]string,len( compressionPolicyVaultEnum))
    for i:=0; i<len(compressionPolicyVaultEnum);i++ {
        convArray[i] = CompressionPolicyVaultEnumToValue(compressionPolicyVaultEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func CompressionPolicyVaultEnumFromValue(value string) CompressionPolicyVaultEnum {
    switch value {
        case "kCompressionNone":
            return CompressionPolicyVault_KCOMPRESSIONNONE
        case "kCompressionLow":
            return CompressionPolicyVault_KCOMPRESSIONLOW
        case "kCompressionHigh":
            return CompressionPolicyVault_KCOMPRESSIONHIGH
        default:
            return CompressionPolicyVault_KCOMPRESSIONNONE
    }
}
