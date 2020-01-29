// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for BondingModeCreateBondParametersEnum enum
 */
type BondingModeCreateBondParametersEnum int

/**
 * Value collection for BondingModeCreateBondParametersEnum enum
 */
const (
    BondingModeCreateBondParameters_KACTIVEBACKUP BondingModeCreateBondParametersEnum = 1 + iota
    BondingModeCreateBondParameters_K802_3AD
)

func (r BondingModeCreateBondParametersEnum) MarshalJSON() ([]byte, error) {
    s := BondingModeCreateBondParametersEnumToValue(r)
    return json.Marshal(s)
}

func (r *BondingModeCreateBondParametersEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  BondingModeCreateBondParametersEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts BondingModeCreateBondParametersEnum to its string representation
 */
func BondingModeCreateBondParametersEnumToValue(bondingModeCreateBondParametersEnum BondingModeCreateBondParametersEnum) string {
    switch bondingModeCreateBondParametersEnum {
        case BondingModeCreateBondParameters_KACTIVEBACKUP:
    		return "kActiveBackup"
        case BondingModeCreateBondParameters_K802_3AD:
    		return "k802_3ad"
        default:
        	return "kActiveBackup"
    }
}

/**
 * Converts BondingModeCreateBondParametersEnum Array to its string Array representation
*/
func BondingModeCreateBondParametersEnumArrayToValue(bondingModeCreateBondParametersEnum []BondingModeCreateBondParametersEnum) []string {
    convArray := make([]string,len( bondingModeCreateBondParametersEnum))
    for i:=0; i<len(bondingModeCreateBondParametersEnum);i++ {
        convArray[i] = BondingModeCreateBondParametersEnumToValue(bondingModeCreateBondParametersEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func BondingModeCreateBondParametersEnumFromValue(value string) BondingModeCreateBondParametersEnum {
    switch value {
        case "kActiveBackup":
            return BondingModeCreateBondParameters_KACTIVEBACKUP
        case "k802_3ad":
            return BondingModeCreateBondParameters_K802_3AD
        default:
            return BondingModeCreateBondParameters_KACTIVEBACKUP
    }
}
