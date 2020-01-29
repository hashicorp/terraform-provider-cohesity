// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ShareTypeEnum enum
 */
type ShareTypeEnum int

/**
 * Value collection for ShareTypeEnum enum
 */
const (
    ShareType_KNFS ShareTypeEnum = 1 + iota
    ShareType_KCIFS
)

func (r ShareTypeEnum) MarshalJSON() ([]byte, error) {
    s := ShareTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ShareTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ShareTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ShareTypeEnum to its string representation
 */
func ShareTypeEnumToValue(shareTypeEnum ShareTypeEnum) string {
    switch shareTypeEnum {
        case ShareType_KNFS:
    		return "kNFS"
        case ShareType_KCIFS:
    		return "kCIFS"
        default:
        	return "kNFS"
    }
}

/**
 * Converts ShareTypeEnum Array to its string Array representation
*/
func ShareTypeEnumArrayToValue(shareTypeEnum []ShareTypeEnum) []string {
    convArray := make([]string,len( shareTypeEnum))
    for i:=0; i<len(shareTypeEnum);i++ {
        convArray[i] = ShareTypeEnumToValue(shareTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ShareTypeEnumFromValue(value string) ShareTypeEnum {
    switch value {
        case "kNFS":
            return ShareType_KNFS
        case "kCIFS":
            return ShareType_KCIFS
        default:
            return ShareType_KNFS
    }
}
