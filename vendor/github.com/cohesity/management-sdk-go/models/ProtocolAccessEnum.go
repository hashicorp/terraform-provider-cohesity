// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ProtocolAccessEnum enum
 */
type ProtocolAccessEnum int

/**
 * Value collection for ProtocolAccessEnum enum
 */
const (
    ProtocolAccess_KALL ProtocolAccessEnum = 1 + iota
    ProtocolAccess_KNFSONLY
    ProtocolAccess_KSMBONLY
    ProtocolAccess_KS3ONLY
)

func (r ProtocolAccessEnum) MarshalJSON() ([]byte, error) {
    s := ProtocolAccessEnumToValue(r)
    return json.Marshal(s)
}

func (r *ProtocolAccessEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ProtocolAccessEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ProtocolAccessEnum to its string representation
 */
func ProtocolAccessEnumToValue(protocolAccessEnum ProtocolAccessEnum) string {
    switch protocolAccessEnum {
        case ProtocolAccess_KALL:
    		return "kAll"
        case ProtocolAccess_KNFSONLY:
    		return "kNFSOnly"
        case ProtocolAccess_KSMBONLY:
    		return "kSMBOnly"
        case ProtocolAccess_KS3ONLY:
    		return "kS3Only"
        default:
        	return "kAll"
    }
}

/**
 * Converts ProtocolAccessEnum Array to its string Array representation
*/
func ProtocolAccessEnumArrayToValue(protocolAccessEnum []ProtocolAccessEnum) []string {
    convArray := make([]string,len( protocolAccessEnum))
    for i:=0; i<len(protocolAccessEnum);i++ {
        convArray[i] = ProtocolAccessEnumToValue(protocolAccessEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ProtocolAccessEnumFromValue(value string) ProtocolAccessEnum {
    switch value {
        case "kAll":
            return ProtocolAccess_KALL
        case "kNFSOnly":
            return ProtocolAccess_KNFSONLY
        case "kSMBOnly":
            return ProtocolAccess_KSMBONLY
        case "kS3Only":
            return ProtocolAccess_KS3ONLY
        default:
            return ProtocolAccess_KALL
    }
}
