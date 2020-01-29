// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for NasTypeEnum enum
 */
type NasTypeEnum int

/**
 * Value collection for NasTypeEnum enum
 */
const (
    NasType_KGROUP NasTypeEnum = 1 + iota
    NasType_KHOST
    NasType_KDFSGROUP
    NasType_KDFSTOPDIR
)

func (r NasTypeEnum) MarshalJSON() ([]byte, error) {
    s := NasTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *NasTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  NasTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts NasTypeEnum to its string representation
 */
func NasTypeEnumToValue(nasTypeEnum NasTypeEnum) string {
    switch nasTypeEnum {
        case NasType_KGROUP:
    		return "kGroup"
        case NasType_KHOST:
    		return "kHost"
        case NasType_KDFSGROUP:
    		return "kDfsGroup"
        case NasType_KDFSTOPDIR:
    		return "kDfsTopDir"
        default:
        	return "kGroup"
    }
}

/**
 * Converts NasTypeEnum Array to its string Array representation
*/
func NasTypeEnumArrayToValue(nasTypeEnum []NasTypeEnum) []string {
    convArray := make([]string,len( nasTypeEnum))
    for i:=0; i<len(nasTypeEnum);i++ {
        convArray[i] = NasTypeEnumToValue(nasTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NasTypeEnumFromValue(value string) NasTypeEnum {
    switch value {
        case "kGroup":
            return NasType_KGROUP
        case "kHost":
            return NasType_KHOST
        case "kDfsGroup":
            return NasType_KDFSGROUP
        case "kDfsTopDir":
            return NasType_KDFSTOPDIR
        default:
            return NasType_KGROUP
    }
}
