// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for NfsAccessEnum enum
 */
type NfsAccessEnum int

/**
 * Value collection for NfsAccessEnum enum
 */
const (
    NfsAccess_KDISABLED NfsAccessEnum = 1 + iota
    NfsAccess_KREADONLY
    NfsAccess_KREADWRITE
)

func (r NfsAccessEnum) MarshalJSON() ([]byte, error) {
    s := NfsAccessEnumToValue(r)
    return json.Marshal(s)
}

func (r *NfsAccessEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  NfsAccessEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts NfsAccessEnum to its string representation
 */
func NfsAccessEnumToValue(nfsAccessEnum NfsAccessEnum) string {
    switch nfsAccessEnum {
        case NfsAccess_KDISABLED:
    		return "kDisabled"
        case NfsAccess_KREADONLY:
    		return "kReadOnly"
        case NfsAccess_KREADWRITE:
    		return "kReadWrite"
        default:
        	return "kDisabled"
    }
}

/**
 * Converts NfsAccessEnum Array to its string Array representation
*/
func NfsAccessEnumArrayToValue(nfsAccessEnum []NfsAccessEnum) []string {
    convArray := make([]string,len( nfsAccessEnum))
    for i:=0; i<len(nfsAccessEnum);i++ {
        convArray[i] = NfsAccessEnumToValue(nfsAccessEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NfsAccessEnumFromValue(value string) NfsAccessEnum {
    switch value {
        case "kDisabled":
            return NfsAccess_KDISABLED
        case "kReadOnly":
            return NfsAccess_KREADONLY
        case "kReadWrite":
            return NfsAccess_KREADWRITE
        default:
            return NfsAccess_KDISABLED
    }
}
