// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for BondingModeEnum enum
 */
type BondingModeEnum int

/**
 * Value collection for BondingModeEnum enum
 */
const (
    BondingMode_ACTIVEBACKUP BondingModeEnum = 1 + iota
    BondingMode_ENUM_802_3AD
    BondingMode_BALANCEALB
)

func (r BondingModeEnum) MarshalJSON() ([]byte, error) {
    s := BondingModeEnumToValue(r)
    return json.Marshal(s)
}

func (r *BondingModeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  BondingModeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts BondingModeEnum to its string representation
 */
func BondingModeEnumToValue(bondingModeEnum BondingModeEnum) string {
    switch bondingModeEnum {
        case BondingMode_ACTIVEBACKUP:
    		return "ActiveBackup"
        case BondingMode_ENUM_802_3AD:
    		return "Enum_802_3ad"
        case BondingMode_BALANCEALB:
    		return "BalanceAlb"
        default:
        	return "ActiveBackup"
    }
}

/**
 * Converts BondingModeEnum Array to its string Array representation
*/
func BondingModeEnumArrayToValue(bondingModeEnum []BondingModeEnum) []string {
    convArray := make([]string,len( bondingModeEnum))
    for i:=0; i<len(bondingModeEnum);i++ {
        convArray[i] = BondingModeEnumToValue(bondingModeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func BondingModeEnumFromValue(value string) BondingModeEnum {
    switch value {
        case "ActiveBackup":
            return BondingMode_ACTIVEBACKUP
        case "Enum_802_3ad":
            return BondingMode_ENUM_802_3AD
        case "BalanceAlb":
            return BondingMode_BALANCEALB
        default:
            return BondingMode_ACTIVEBACKUP
    }
}
