// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ProtocolIsilonMountPointEnum enum
 */
type ProtocolIsilonMountPointEnum int

/**
 * Value collection for ProtocolIsilonMountPointEnum enum
 */
const (
    ProtocolIsilonMountPoint_KNFS ProtocolIsilonMountPointEnum = 1 + iota
    ProtocolIsilonMountPoint_KSMB
)

func (r ProtocolIsilonMountPointEnum) MarshalJSON() ([]byte, error) {
    s := ProtocolIsilonMountPointEnumToValue(r)
    return json.Marshal(s)
}

func (r *ProtocolIsilonMountPointEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ProtocolIsilonMountPointEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ProtocolIsilonMountPointEnum to its string representation
 */
func ProtocolIsilonMountPointEnumToValue(protocolIsilonMountPointEnum ProtocolIsilonMountPointEnum) string {
    switch protocolIsilonMountPointEnum {
        case ProtocolIsilonMountPoint_KNFS:
    		return "kNfs"
        case ProtocolIsilonMountPoint_KSMB:
    		return "kSmb"
        default:
        	return "kNfs"
    }
}

/**
 * Converts ProtocolIsilonMountPointEnum Array to its string Array representation
*/
func ProtocolIsilonMountPointEnumArrayToValue(protocolIsilonMountPointEnum []ProtocolIsilonMountPointEnum) []string {
    convArray := make([]string,len( protocolIsilonMountPointEnum))
    for i:=0; i<len(protocolIsilonMountPointEnum);i++ {
        convArray[i] = ProtocolIsilonMountPointEnumToValue(protocolIsilonMountPointEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ProtocolIsilonMountPointEnumFromValue(value string) ProtocolIsilonMountPointEnum {
    switch value {
        case "kNfs":
            return ProtocolIsilonMountPoint_KNFS
        case "kSmb":
            return ProtocolIsilonMountPoint_KSMB
        default:
            return ProtocolIsilonMountPoint_KNFS
    }
}
