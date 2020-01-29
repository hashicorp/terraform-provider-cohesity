// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ProtocolEnum enum
 */
type ProtocolEnum int

/**
 * Value collection for ProtocolEnum enum
 */
const (
    Protocol_KNFS ProtocolEnum = 1 + iota
    Protocol_KCIFS2
    Protocol_KHTTP
)

func (r ProtocolEnum) MarshalJSON() ([]byte, error) {
    s := ProtocolEnumToValue(r)
    return json.Marshal(s)
}

func (r *ProtocolEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ProtocolEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ProtocolEnum to its string representation
 */
func ProtocolEnumToValue(protocolEnum ProtocolEnum) string {
    switch protocolEnum {
        case Protocol_KNFS:
    		return "kNfs"
        case Protocol_KCIFS2:
    		return "kCifs2"
        case Protocol_KHTTP:
    		return "kHttp"
        default:
        	return "kNfs"
    }
}

/**
 * Converts ProtocolEnum Array to its string Array representation
*/
func ProtocolEnumArrayToValue(protocolEnum []ProtocolEnum) []string {
    convArray := make([]string,len( protocolEnum))
    for i:=0; i<len(protocolEnum);i++ {
        convArray[i] = ProtocolEnumToValue(protocolEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ProtocolEnumFromValue(value string) ProtocolEnum {
    switch value {
        case "kNfs":
            return Protocol_KNFS
        case "kCifs2":
            return Protocol_KCIFS2
        case "kHttp":
            return Protocol_KHTTP
        default:
            return Protocol_KNFS
    }
}
