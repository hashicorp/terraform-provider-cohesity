// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EncryptionPolicyVaultEnum enum
 */
type EncryptionPolicyVaultEnum int

/**
 * Value collection for EncryptionPolicyVaultEnum enum
 */
const (
    EncryptionPolicyVault_KENCRYPTIONNONE EncryptionPolicyVaultEnum = 1 + iota
    EncryptionPolicyVault_KENCRYPTIONSTRONG
    EncryptionPolicyVault_KENCRYPTIONWEAK
)

func (r EncryptionPolicyVaultEnum) MarshalJSON() ([]byte, error) {
    s := EncryptionPolicyVaultEnumToValue(r)
    return json.Marshal(s)
}

func (r *EncryptionPolicyVaultEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EncryptionPolicyVaultEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EncryptionPolicyVaultEnum to its string representation
 */
func EncryptionPolicyVaultEnumToValue(encryptionPolicyVaultEnum EncryptionPolicyVaultEnum) string {
    switch encryptionPolicyVaultEnum {
        case EncryptionPolicyVault_KENCRYPTIONNONE:
    		return "kEncryptionNone"
        case EncryptionPolicyVault_KENCRYPTIONSTRONG:
    		return "kEncryptionStrong"
        case EncryptionPolicyVault_KENCRYPTIONWEAK:
    		return "kEncryptionWeak"
        default:
        	return "kEncryptionNone"
    }
}

/**
 * Converts EncryptionPolicyVaultEnum Array to its string Array representation
*/
func EncryptionPolicyVaultEnumArrayToValue(encryptionPolicyVaultEnum []EncryptionPolicyVaultEnum) []string {
    convArray := make([]string,len( encryptionPolicyVaultEnum))
    for i:=0; i<len(encryptionPolicyVaultEnum);i++ {
        convArray[i] = EncryptionPolicyVaultEnumToValue(encryptionPolicyVaultEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EncryptionPolicyVaultEnumFromValue(value string) EncryptionPolicyVaultEnum {
    switch value {
        case "kEncryptionNone":
            return EncryptionPolicyVault_KENCRYPTIONNONE
        case "kEncryptionStrong":
            return EncryptionPolicyVault_KENCRYPTIONSTRONG
        case "kEncryptionWeak":
            return EncryptionPolicyVault_KENCRYPTIONWEAK
        default:
            return EncryptionPolicyVault_KENCRYPTIONNONE
    }
}
