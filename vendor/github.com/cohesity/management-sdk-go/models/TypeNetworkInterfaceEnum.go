// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeNetworkInterfaceEnum enum
 */
type TypeNetworkInterfaceEnum int

/**
 * Value collection for TypeNetworkInterfaceEnum enum
 */
const (
    TypeNetworkInterface_KPHYSICALINTERFACE TypeNetworkInterfaceEnum = 1 + iota
    TypeNetworkInterface_KBONDMASTERINTERFACE
    TypeNetworkInterface_KBONDSLAVEINTERFACE
    TypeNetworkInterface_KTAGGEDVLANINTERFACE
)

func (r TypeNetworkInterfaceEnum) MarshalJSON() ([]byte, error) {
    s := TypeNetworkInterfaceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeNetworkInterfaceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeNetworkInterfaceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeNetworkInterfaceEnum to its string representation
 */
func TypeNetworkInterfaceEnumToValue(typeNetworkInterfaceEnum TypeNetworkInterfaceEnum) string {
    switch typeNetworkInterfaceEnum {
        case TypeNetworkInterface_KPHYSICALINTERFACE:
    		return "kPhysicalInterface"
        case TypeNetworkInterface_KBONDMASTERINTERFACE:
    		return "kBondMasterInterface"
        case TypeNetworkInterface_KBONDSLAVEINTERFACE:
    		return "kBondSlaveInterface"
        case TypeNetworkInterface_KTAGGEDVLANINTERFACE:
    		return "kTaggedVlanInterface"
        default:
        	return "kPhysicalInterface"
    }
}

/**
 * Converts TypeNetworkInterfaceEnum Array to its string Array representation
*/
func TypeNetworkInterfaceEnumArrayToValue(typeNetworkInterfaceEnum []TypeNetworkInterfaceEnum) []string {
    convArray := make([]string,len( typeNetworkInterfaceEnum))
    for i:=0; i<len(typeNetworkInterfaceEnum);i++ {
        convArray[i] = TypeNetworkInterfaceEnumToValue(typeNetworkInterfaceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeNetworkInterfaceEnumFromValue(value string) TypeNetworkInterfaceEnum {
    switch value {
        case "kPhysicalInterface":
            return TypeNetworkInterface_KPHYSICALINTERFACE
        case "kBondMasterInterface":
            return TypeNetworkInterface_KBONDMASTERINTERFACE
        case "kBondSlaveInterface":
            return TypeNetworkInterface_KBONDSLAVEINTERFACE
        case "kTaggedVlanInterface":
            return TypeNetworkInterface_KTAGGEDVLANINTERFACE
        default:
            return TypeNetworkInterface_KPHYSICALINTERFACE
    }
}
