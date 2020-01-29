// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RoleNetworkInterfaceEnum enum
 */
type RoleNetworkInterfaceEnum int

/**
 * Value collection for RoleNetworkInterfaceEnum enum
 */
const (
    RoleNetworkInterface_KPRIMARY RoleNetworkInterfaceEnum = 1 + iota
    RoleNetworkInterface_KSECONDARY
)

func (r RoleNetworkInterfaceEnum) MarshalJSON() ([]byte, error) {
    s := RoleNetworkInterfaceEnumToValue(r)
    return json.Marshal(s)
}

func (r *RoleNetworkInterfaceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RoleNetworkInterfaceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RoleNetworkInterfaceEnum to its string representation
 */
func RoleNetworkInterfaceEnumToValue(roleNetworkInterfaceEnum RoleNetworkInterfaceEnum) string {
    switch roleNetworkInterfaceEnum {
        case RoleNetworkInterface_KPRIMARY:
    		return "kPrimary"
        case RoleNetworkInterface_KSECONDARY:
    		return "kSecondary"
        default:
        	return "kPrimary"
    }
}

/**
 * Converts RoleNetworkInterfaceEnum Array to its string Array representation
*/
func RoleNetworkInterfaceEnumArrayToValue(roleNetworkInterfaceEnum []RoleNetworkInterfaceEnum) []string {
    convArray := make([]string,len( roleNetworkInterfaceEnum))
    for i:=0; i<len(roleNetworkInterfaceEnum);i++ {
        convArray[i] = RoleNetworkInterfaceEnumToValue(roleNetworkInterfaceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RoleNetworkInterfaceEnumFromValue(value string) RoleNetworkInterfaceEnum {
    switch value {
        case "kPrimary":
            return RoleNetworkInterface_KPRIMARY
        case "kSecondary":
            return RoleNetworkInterface_KSECONDARY
        default:
            return RoleNetworkInterface_KPRIMARY
    }
}
