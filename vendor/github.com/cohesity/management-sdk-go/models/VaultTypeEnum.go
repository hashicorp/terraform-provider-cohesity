// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for VaultTypeEnum enum
 */
type VaultTypeEnum int

/**
 * Value collection for VaultTypeEnum enum
 */
const (
    VaultType_KCLOUD VaultTypeEnum = 1 + iota
    VaultType_KTAPE
    VaultType_KNAS
)

func (r VaultTypeEnum) MarshalJSON() ([]byte, error) {
    s := VaultTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *VaultTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  VaultTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts VaultTypeEnum to its string representation
 */
func VaultTypeEnumToValue(vaultTypeEnum VaultTypeEnum) string {
    switch vaultTypeEnum {
        case VaultType_KCLOUD:
    		return "kCloud"
        case VaultType_KTAPE:
    		return "kTape"
        case VaultType_KNAS:
    		return "kNas"
        default:
        	return "kCloud"
    }
}

/**
 * Converts VaultTypeEnum Array to its string Array representation
*/
func VaultTypeEnumArrayToValue(vaultTypeEnum []VaultTypeEnum) []string {
    convArray := make([]string,len( vaultTypeEnum))
    for i:=0; i<len(vaultTypeEnum);i++ {
        convArray[i] = VaultTypeEnumToValue(vaultTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func VaultTypeEnumFromValue(value string) VaultTypeEnum {
    switch value {
        case "kCloud":
            return VaultType_KCLOUD
        case "kTape":
            return VaultType_KTAPE
        case "kNas":
            return VaultType_KNAS
        default:
            return VaultType_KCLOUD
    }
}
