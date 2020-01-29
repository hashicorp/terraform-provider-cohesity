// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for BondingModeNetworkInterfaceEnum enum
 */
type BondingModeNetworkInterfaceEnum int

/**
 * Value collection for BondingModeNetworkInterfaceEnum enum
 */
const (
    BondingModeNetworkInterface_KACTIVEBACKUP BondingModeNetworkInterfaceEnum = 1 + iota
    BondingModeNetworkInterface_K802_3AD
)

func (r BondingModeNetworkInterfaceEnum) MarshalJSON() ([]byte, error) {
    s := BondingModeNetworkInterfaceEnumToValue(r)
    return json.Marshal(s)
}

func (r *BondingModeNetworkInterfaceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  BondingModeNetworkInterfaceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts BondingModeNetworkInterfaceEnum to its string representation
 */
func BondingModeNetworkInterfaceEnumToValue(bondingModeNetworkInterfaceEnum BondingModeNetworkInterfaceEnum) string {
    switch bondingModeNetworkInterfaceEnum {
        case BondingModeNetworkInterface_KACTIVEBACKUP:
    		return "kActiveBackup"
        case BondingModeNetworkInterface_K802_3AD:
    		return "k802_3ad"
        default:
        	return "kActiveBackup"
    }
}

/**
 * Converts BondingModeNetworkInterfaceEnum Array to its string Array representation
*/
func BondingModeNetworkInterfaceEnumArrayToValue(bondingModeNetworkInterfaceEnum []BondingModeNetworkInterfaceEnum) []string {
    convArray := make([]string,len( bondingModeNetworkInterfaceEnum))
    for i:=0; i<len(bondingModeNetworkInterfaceEnum);i++ {
        convArray[i] = BondingModeNetworkInterfaceEnumToValue(bondingModeNetworkInterfaceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func BondingModeNetworkInterfaceEnumFromValue(value string) BondingModeNetworkInterfaceEnum {
    switch value {
        case "kActiveBackup":
            return BondingModeNetworkInterface_KACTIVEBACKUP
        case "k802_3ad":
            return BondingModeNetworkInterface_K802_3AD
        default:
            return BondingModeNetworkInterface_KACTIVEBACKUP
    }
}
