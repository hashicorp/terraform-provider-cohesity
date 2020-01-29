// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for NasProtocolEnum enum
 */
type NasProtocolEnum int

/**
 * Value collection for NasProtocolEnum enum
 */
const (
    NasProtocol_KNFS3 NasProtocolEnum = 1 + iota
    NasProtocol_KCIFS1
)

func (r NasProtocolEnum) MarshalJSON() ([]byte, error) {
    s := NasProtocolEnumToValue(r)
    return json.Marshal(s)
}

func (r *NasProtocolEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  NasProtocolEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts NasProtocolEnum to its string representation
 */
func NasProtocolEnumToValue(nasProtocolEnum NasProtocolEnum) string {
    switch nasProtocolEnum {
        case NasProtocol_KNFS3:
    		return "kNfs3"
        case NasProtocol_KCIFS1:
    		return "kCifs1"
        default:
        	return "kNfs3"
    }
}

/**
 * Converts NasProtocolEnum Array to its string Array representation
*/
func NasProtocolEnumArrayToValue(nasProtocolEnum []NasProtocolEnum) []string {
    convArray := make([]string,len( nasProtocolEnum))
    for i:=0; i<len(nasProtocolEnum);i++ {
        convArray[i] = NasProtocolEnumToValue(nasProtocolEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NasProtocolEnumFromValue(value string) NasProtocolEnum {
    switch value {
        case "kNfs3":
            return NasProtocol_KNFS3
        case "kCifs1":
            return NasProtocol_KCIFS1
        default:
            return NasProtocol_KNFS3
    }
}
