// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ProtocolViewProtocolStatsEnum enum
 */
type ProtocolViewProtocolStatsEnum int

/**
 * Value collection for ProtocolViewProtocolStatsEnum enum
 */
const (
    ProtocolViewProtocolStats_KNFS ProtocolViewProtocolStatsEnum = 1 + iota
    ProtocolViewProtocolStats_KSMB
    ProtocolViewProtocolStats_KS3
    ProtocolViewProtocolStats_KISCSI
)

func (r ProtocolViewProtocolStatsEnum) MarshalJSON() ([]byte, error) {
    s := ProtocolViewProtocolStatsEnumToValue(r)
    return json.Marshal(s)
}

func (r *ProtocolViewProtocolStatsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ProtocolViewProtocolStatsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ProtocolViewProtocolStatsEnum to its string representation
 */
func ProtocolViewProtocolStatsEnumToValue(protocolViewProtocolStatsEnum ProtocolViewProtocolStatsEnum) string {
    switch protocolViewProtocolStatsEnum {
        case ProtocolViewProtocolStats_KNFS:
    		return "kNfs"
        case ProtocolViewProtocolStats_KSMB:
    		return "kSmb"
        case ProtocolViewProtocolStats_KS3:
    		return "kS3"
        case ProtocolViewProtocolStats_KISCSI:
    		return "kIscsi"
        default:
        	return "kNfs"
    }
}

/**
 * Converts ProtocolViewProtocolStatsEnum Array to its string Array representation
*/
func ProtocolViewProtocolStatsEnumArrayToValue(protocolViewProtocolStatsEnum []ProtocolViewProtocolStatsEnum) []string {
    convArray := make([]string,len( protocolViewProtocolStatsEnum))
    for i:=0; i<len(protocolViewProtocolStatsEnum);i++ {
        convArray[i] = ProtocolViewProtocolStatsEnumToValue(protocolViewProtocolStatsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ProtocolViewProtocolStatsEnumFromValue(value string) ProtocolViewProtocolStatsEnum {
    switch value {
        case "kNfs":
            return ProtocolViewProtocolStats_KNFS
        case "kSmb":
            return ProtocolViewProtocolStats_KSMB
        case "kS3":
            return ProtocolViewProtocolStats_KS3
        case "kIscsi":
            return ProtocolViewProtocolStats_KISCSI
        default:
            return ProtocolViewProtocolStats_KNFS
    }
}
