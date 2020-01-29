// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ProtocolSyslogServerEnum enum
 */
type ProtocolSyslogServerEnum int

/**
 * Value collection for ProtocolSyslogServerEnum enum
 */
const (
    ProtocolSyslogServer_KUDP ProtocolSyslogServerEnum = 1 + iota
    ProtocolSyslogServer_KTCP
)

func (r ProtocolSyslogServerEnum) MarshalJSON() ([]byte, error) {
    s := ProtocolSyslogServerEnumToValue(r)
    return json.Marshal(s)
}

func (r *ProtocolSyslogServerEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ProtocolSyslogServerEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ProtocolSyslogServerEnum to its string representation
 */
func ProtocolSyslogServerEnumToValue(protocolSyslogServerEnum ProtocolSyslogServerEnum) string {
    switch protocolSyslogServerEnum {
        case ProtocolSyslogServer_KUDP:
    		return "kUDP"
        case ProtocolSyslogServer_KTCP:
    		return "kTCP"
        default:
        	return "kUDP"
    }
}

/**
 * Converts ProtocolSyslogServerEnum Array to its string Array representation
*/
func ProtocolSyslogServerEnumArrayToValue(protocolSyslogServerEnum []ProtocolSyslogServerEnum) []string {
    convArray := make([]string,len( protocolSyslogServerEnum))
    for i:=0; i<len(protocolSyslogServerEnum);i++ {
        convArray[i] = ProtocolSyslogServerEnumToValue(protocolSyslogServerEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ProtocolSyslogServerEnumFromValue(value string) ProtocolSyslogServerEnum {
    switch value {
        case "kUDP":
            return ProtocolSyslogServer_KUDP
        case "kTCP":
            return ProtocolSyslogServer_KTCP
        default:
            return ProtocolSyslogServer_KUDP
    }
}
