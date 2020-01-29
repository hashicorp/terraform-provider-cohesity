// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ProtocolNasProtectionSourceEnum enum
 */
type ProtocolNasProtectionSourceEnum int

/**
 * Value collection for ProtocolNasProtectionSourceEnum enum
 */
const (
    ProtocolNasProtectionSource_KNFS3 ProtocolNasProtectionSourceEnum = 1 + iota
    ProtocolNasProtectionSource_KCIFS1
)

func (r ProtocolNasProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := ProtocolNasProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *ProtocolNasProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ProtocolNasProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ProtocolNasProtectionSourceEnum to its string representation
 */
func ProtocolNasProtectionSourceEnumToValue(protocolNasProtectionSourceEnum ProtocolNasProtectionSourceEnum) string {
    switch protocolNasProtectionSourceEnum {
        case ProtocolNasProtectionSource_KNFS3:
    		return "kNfs3"
        case ProtocolNasProtectionSource_KCIFS1:
    		return "kCifs1"
        default:
        	return "kNfs3"
    }
}

/**
 * Converts ProtocolNasProtectionSourceEnum Array to its string Array representation
*/
func ProtocolNasProtectionSourceEnumArrayToValue(protocolNasProtectionSourceEnum []ProtocolNasProtectionSourceEnum) []string {
    convArray := make([]string,len( protocolNasProtectionSourceEnum))
    for i:=0; i<len(protocolNasProtectionSourceEnum);i++ {
        convArray[i] = ProtocolNasProtectionSourceEnumToValue(protocolNasProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ProtocolNasProtectionSourceEnumFromValue(value string) ProtocolNasProtectionSourceEnum {
    switch value {
        case "kNfs3":
            return ProtocolNasProtectionSource_KNFS3
        case "kCifs1":
            return ProtocolNasProtectionSource_KCIFS1
        default:
            return ProtocolNasProtectionSource_KNFS3
    }
}
