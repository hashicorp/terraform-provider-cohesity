// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EncryptionPolicyEnum enum
 */
type EncryptionPolicyEnum int

/**
 * Value collection for EncryptionPolicyEnum enum
 */
const (
    EncryptionPolicy_KENCRYPTIONNONE EncryptionPolicyEnum = 1 + iota
    EncryptionPolicy_KENCRYPTIONSTRONG
    EncryptionPolicy_KENCRYPTIONWEAK
)

func (r EncryptionPolicyEnum) MarshalJSON() ([]byte, error) {
    s := EncryptionPolicyEnumToValue(r)
    return json.Marshal(s)
}

func (r *EncryptionPolicyEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EncryptionPolicyEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EncryptionPolicyEnum to its string representation
 */
func EncryptionPolicyEnumToValue(encryptionPolicyEnum EncryptionPolicyEnum) string {
    switch encryptionPolicyEnum {
        case EncryptionPolicy_KENCRYPTIONNONE:
    		return "kEncryptionNone"
        case EncryptionPolicy_KENCRYPTIONSTRONG:
    		return "kEncryptionStrong"
        case EncryptionPolicy_KENCRYPTIONWEAK:
    		return "kEncryptionWeak"
        default:
        	return "kEncryptionNone"
    }
}

/**
 * Converts EncryptionPolicyEnum Array to its string Array representation
*/
func EncryptionPolicyEnumArrayToValue(encryptionPolicyEnum []EncryptionPolicyEnum) []string {
    convArray := make([]string,len( encryptionPolicyEnum))
    for i:=0; i<len(encryptionPolicyEnum);i++ {
        convArray[i] = EncryptionPolicyEnumToValue(encryptionPolicyEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EncryptionPolicyEnumFromValue(value string) EncryptionPolicyEnum {
    switch value {
        case "kEncryptionNone":
            return EncryptionPolicy_KENCRYPTIONNONE
        case "kEncryptionStrong":
            return EncryptionPolicy_KENCRYPTIONSTRONG
        case "kEncryptionWeak":
            return EncryptionPolicy_KENCRYPTIONWEAK
        default:
            return EncryptionPolicy_KENCRYPTIONNONE
    }
}
