// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for VaultGroupEnum enum
 */
type VaultGroupEnum int

/**
 * Value collection for VaultGroupEnum enum
 */
const (
    VaultGroup_KAWS VaultGroupEnum = 1 + iota
    VaultGroup_KAZURE
    VaultGroup_KGCP
    VaultGroup_KORACLE
    VaultGroup_KNAS
    VaultGroup_KQSTAR
    VaultGroup_KS3C
    VaultGroup_KOTHER
)

func (r VaultGroupEnum) MarshalJSON() ([]byte, error) {
    s := VaultGroupEnumToValue(r)
    return json.Marshal(s)
}

func (r *VaultGroupEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  VaultGroupEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts VaultGroupEnum to its string representation
 */
func VaultGroupEnumToValue(vaultGroupEnum VaultGroupEnum) string {
    switch vaultGroupEnum {
        case VaultGroup_KAWS:
    		return "kAws"
        case VaultGroup_KAZURE:
    		return "kAzure"
        case VaultGroup_KGCP:
    		return "kGcp"
        case VaultGroup_KORACLE:
    		return "kOracle"
        case VaultGroup_KNAS:
    		return "kNas"
        case VaultGroup_KQSTAR:
    		return "kQStar"
        case VaultGroup_KS3C:
    		return "kS3C"
        case VaultGroup_KOTHER:
    		return "kOther"
        default:
        	return "kAws"
    }
}

/**
 * Converts VaultGroupEnum Array to its string Array representation
*/
func VaultGroupEnumArrayToValue(vaultGroupEnum []VaultGroupEnum) []string {
    convArray := make([]string,len( vaultGroupEnum))
    for i:=0; i<len(vaultGroupEnum);i++ {
        convArray[i] = VaultGroupEnumToValue(vaultGroupEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func VaultGroupEnumFromValue(value string) VaultGroupEnum {
    switch value {
        case "kAws":
            return VaultGroup_KAWS
        case "kAzure":
            return VaultGroup_KAZURE
        case "kGcp":
            return VaultGroup_KGCP
        case "kOracle":
            return VaultGroup_KORACLE
        case "kNas":
            return VaultGroup_KNAS
        case "kQStar":
            return VaultGroup_KQSTAR
        case "kS3C":
            return VaultGroup_KS3C
        case "kOther":
            return VaultGroup_KOTHER
        default:
            return VaultGroup_KAWS
    }
}
