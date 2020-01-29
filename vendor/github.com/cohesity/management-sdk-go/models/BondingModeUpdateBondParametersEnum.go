// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for BondingModeUpdateBondParametersEnum enum
 */
type BondingModeUpdateBondParametersEnum int

/**
 * Value collection for BondingModeUpdateBondParametersEnum enum
 */
const (
    BondingModeUpdateBondParameters_KACTIVEBACKUP BondingModeUpdateBondParametersEnum = 1 + iota
    BondingModeUpdateBondParameters_K802_3AD
)

func (r BondingModeUpdateBondParametersEnum) MarshalJSON() ([]byte, error) {
    s := BondingModeUpdateBondParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *BondingModeUpdateBondParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  BondingModeUpdateBondParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts BondingModeUpdateBondParametersEnum to its string representation
 */
func BondingModeUpdateBondParametersEnumToValue(bondingModeUpdateBondParametersEnum BondingModeUpdateBondParametersEnum) string {
    switch bondingModeUpdateBondParametersEnum {
        case BondingModeUpdateBondParameters_KACTIVEBACKUP:
    		return "kActiveBackup"
        case BondingModeUpdateBondParameters_K802_3AD:
    		return "k802_3ad"
        default:
        	return "kActiveBackup"
    }
}

/**
 * Converts BondingModeUpdateBondParametersEnum Array to its string Array representation
*/
func BondingModeUpdateBondParametersEnumArrayToValue(bondingModeUpdateBondParametersEnum []BondingModeUpdateBondParametersEnum) []string {
    convArray := make([]string,len( bondingModeUpdateBondParametersEnum))
    for i:=0; i<len(bondingModeUpdateBondParametersEnum);i++ {
        convArray[i] = BondingModeUpdateBondParametersEnumToValue(bondingModeUpdateBondParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func BondingModeUpdateBondParametersEnumFromValue(value string) BondingModeUpdateBondParametersEnum {
    switch value {
        case "kActiveBackup":
            return BondingModeUpdateBondParameters_KACTIVEBACKUP
        case "k802_3ad":
            return BondingModeUpdateBondParameters_K802_3AD
        default:
            return BondingModeUpdateBondParameters_KACTIVEBACKUP
    }
}
