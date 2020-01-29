// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ProtocolViewStatInfoEnum enum
 */
type ProtocolViewStatInfoEnum int

/**
 * Value collection for ProtocolViewStatInfoEnum enum
 */
const (
    ProtocolViewStatInfo_KNFS ProtocolViewStatInfoEnum = 1 + iota
    ProtocolViewStatInfo_KSMB
    ProtocolViewStatInfo_KS3
    ProtocolViewStatInfo_KISCSI
)

func (r ProtocolViewStatInfoEnum) MarshalJSON() ([]byte, error) {
    s := ProtocolViewStatInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *ProtocolViewStatInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ProtocolViewStatInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ProtocolViewStatInfoEnum to its string representation
 */
func ProtocolViewStatInfoEnumToValue(protocolViewStatInfoEnum ProtocolViewStatInfoEnum) string {
    switch protocolViewStatInfoEnum {
        case ProtocolViewStatInfo_KNFS:
    		return "kNfs"
        case ProtocolViewStatInfo_KSMB:
    		return "kSmb"
        case ProtocolViewStatInfo_KS3:
    		return "kS3"
        case ProtocolViewStatInfo_KISCSI:
    		return "kIscsi"
        default:
        	return "kNfs"
    }
}

/**
 * Converts ProtocolViewStatInfoEnum Array to its string Array representation
*/
func ProtocolViewStatInfoEnumArrayToValue(protocolViewStatInfoEnum []ProtocolViewStatInfoEnum) []string {
    convArray := make([]string,len( protocolViewStatInfoEnum))
    for i:=0; i<len(protocolViewStatInfoEnum);i++ {
        convArray[i] = ProtocolViewStatInfoEnumToValue(protocolViewStatInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ProtocolViewStatInfoEnumFromValue(value string) ProtocolViewStatInfoEnum {
    switch value {
        case "kNfs":
            return ProtocolViewStatInfo_KNFS
        case "kSmb":
            return ProtocolViewStatInfo_KSMB
        case "kS3":
            return ProtocolViewStatInfo_KS3
        case "kIscsi":
            return ProtocolViewStatInfo_KISCSI
        default:
            return ProtocolViewStatInfo_KNFS
    }
}
