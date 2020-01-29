// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AccessEnum enum
 */
type AccessEnum int

/**
 * Value collection for AccessEnum enum
 */
const (
    Access_KREADONLY AccessEnum = 1 + iota
    Access_KREADWRITE
    Access_KMODIFY
    Access_KFULLCONTROL
    Access_KSPECIALACCESS
)

func (r AccessEnum) MarshalJSON() ([]byte, error) {
    s := AccessEnumToValue(r)
    return json.Marshal(s)
}

func (r *AccessEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AccessEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AccessEnum to its string representation
 */
func AccessEnumToValue(accessEnum AccessEnum) string {
    switch accessEnum {
        case Access_KREADONLY:
    		return "kReadOnly"
        case Access_KREADWRITE:
    		return "kReadWrite"
        case Access_KMODIFY:
    		return "kModify"
        case Access_KFULLCONTROL:
    		return "kFullControl"
        case Access_KSPECIALACCESS:
    		return "kSpecialAccess"
        default:
        	return "kReadOnly"
    }
}

/**
 * Converts AccessEnum Array to its string Array representation
*/
func AccessEnumArrayToValue(accessEnum []AccessEnum) []string {
    convArray := make([]string,len( accessEnum))
    for i:=0; i<len(accessEnum);i++ {
        convArray[i] = AccessEnumToValue(accessEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AccessEnumFromValue(value string) AccessEnum {
    switch value {
        case "kReadOnly":
            return Access_KREADONLY
        case "kReadWrite":
            return Access_KREADWRITE
        case "kModify":
            return Access_KMODIFY
        case "kFullControl":
            return Access_KFULLCONTROL
        case "kSpecialAccess":
            return Access_KSPECIALACCESS
        default:
            return Access_KREADONLY
    }
}
