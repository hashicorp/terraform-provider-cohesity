// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeNasProtectionSourceEnum enum
 */
type TypeNasProtectionSourceEnum int

/**
 * Value collection for TypeNasProtectionSourceEnum enum
 */
const (
    TypeNasProtectionSource_KGROUP TypeNasProtectionSourceEnum = 1 + iota
    TypeNasProtectionSource_KHOST
    TypeNasProtectionSource_KDFSGROUP
    TypeNasProtectionSource_KDFSTOPDIR
)

func (r TypeNasProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeNasProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeNasProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeNasProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeNasProtectionSourceEnum to its string representation
 */
func TypeNasProtectionSourceEnumToValue(typeNasProtectionSourceEnum TypeNasProtectionSourceEnum) string {
    switch typeNasProtectionSourceEnum {
        case TypeNasProtectionSource_KGROUP:
    		return "kGroup"
        case TypeNasProtectionSource_KHOST:
    		return "kHost"
        case TypeNasProtectionSource_KDFSGROUP:
    		return "kDfsGroup"
        case TypeNasProtectionSource_KDFSTOPDIR:
    		return "kDfsTopDir"
        default:
        	return "kGroup"
    }
}

/**
 * Converts TypeNasProtectionSourceEnum Array to its string Array representation
*/
func TypeNasProtectionSourceEnumArrayToValue(typeNasProtectionSourceEnum []TypeNasProtectionSourceEnum) []string {
    convArray := make([]string,len( typeNasProtectionSourceEnum))
    for i:=0; i<len(typeNasProtectionSourceEnum);i++ {
        convArray[i] = TypeNasProtectionSourceEnumToValue(typeNasProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeNasProtectionSourceEnumFromValue(value string) TypeNasProtectionSourceEnum {
    switch value {
        case "kGroup":
            return TypeNasProtectionSource_KGROUP
        case "kHost":
            return TypeNasProtectionSource_KHOST
        case "kDfsGroup":
            return TypeNasProtectionSource_KDFSGROUP
        case "kDfsTopDir":
            return TypeNasProtectionSource_KDFSTOPDIR
        default:
            return TypeNasProtectionSource_KGROUP
    }
}
