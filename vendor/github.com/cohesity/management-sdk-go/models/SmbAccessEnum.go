// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SmbAccessEnum enum
 */
type SmbAccessEnum int

/**
 * Value collection for SmbAccessEnum enum
 */
const (
    SmbAccess_KDISABLED SmbAccessEnum = 1 + iota
    SmbAccess_KREADONLY
    SmbAccess_KREADWRITE
)

func (r SmbAccessEnum) MarshalJSON() ([]byte, error) {
    s := SmbAccessEnumToValue(r)
    return json.Marshal(s)
}

func (r *SmbAccessEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SmbAccessEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SmbAccessEnum to its string representation
 */
func SmbAccessEnumToValue(smbAccessEnum SmbAccessEnum) string {
    switch smbAccessEnum {
        case SmbAccess_KDISABLED:
    		return "kDisabled"
        case SmbAccess_KREADONLY:
    		return "kReadOnly"
        case SmbAccess_KREADWRITE:
    		return "kReadWrite"
        default:
        	return "kDisabled"
    }
}

/**
 * Converts SmbAccessEnum Array to its string Array representation
*/
func SmbAccessEnumArrayToValue(smbAccessEnum []SmbAccessEnum) []string {
    convArray := make([]string,len( smbAccessEnum))
    for i:=0; i<len(smbAccessEnum);i++ {
        convArray[i] = SmbAccessEnumToValue(smbAccessEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SmbAccessEnumFromValue(value string) SmbAccessEnum {
    switch value {
        case "kDisabled":
            return SmbAccess_KDISABLED
        case "kReadOnly":
            return SmbAccess_KREADONLY
        case "kReadWrite":
            return SmbAccess_KREADWRITE
        default:
            return SmbAccess_KDISABLED
    }
}
