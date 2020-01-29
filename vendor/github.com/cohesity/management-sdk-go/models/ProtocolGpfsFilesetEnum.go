// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ProtocolGpfsFilesetEnum enum
 */
type ProtocolGpfsFilesetEnum int

/**
 * Value collection for ProtocolGpfsFilesetEnum enum
 */
const (
    ProtocolGpfsFileset_KNFS ProtocolGpfsFilesetEnum = 1 + iota
    ProtocolGpfsFileset_KSMB
)

func (r ProtocolGpfsFilesetEnum) MarshalJSON() ([]byte, error) {
    s := ProtocolGpfsFilesetEnumToValue(r)
    return json.Marshal(s)
}

func (r *ProtocolGpfsFilesetEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ProtocolGpfsFilesetEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ProtocolGpfsFilesetEnum to its string representation
 */
func ProtocolGpfsFilesetEnumToValue(protocolGpfsFilesetEnum ProtocolGpfsFilesetEnum) string {
    switch protocolGpfsFilesetEnum {
        case ProtocolGpfsFileset_KNFS:
    		return "kNfs"
        case ProtocolGpfsFileset_KSMB:
    		return "kSmb"
        default:
        	return "kNfs"
    }
}

/**
 * Converts ProtocolGpfsFilesetEnum Array to its string Array representation
*/
func ProtocolGpfsFilesetEnumArrayToValue(protocolGpfsFilesetEnum []ProtocolGpfsFilesetEnum) []string {
    convArray := make([]string,len( protocolGpfsFilesetEnum))
    for i:=0; i<len(protocolGpfsFilesetEnum);i++ {
        convArray[i] = ProtocolGpfsFilesetEnumToValue(protocolGpfsFilesetEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ProtocolGpfsFilesetEnumFromValue(value string) ProtocolGpfsFilesetEnum {
    switch value {
        case "kNfs":
            return ProtocolGpfsFileset_KNFS
        case "kSmb":
            return ProtocolGpfsFileset_KSMB
        default:
            return ProtocolGpfsFileset_KNFS
    }
}
