// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RemovalStateViewBoxEnum enum
 */
type RemovalStateViewBoxEnum int

/**
 * Value collection for RemovalStateViewBoxEnum enum
 */
const (
    RemovalStateViewBox_KDONTREMOVE RemovalStateViewBoxEnum = 1 + iota
    RemovalStateViewBox_KMARKEDFORREMOVAL
    RemovalStateViewBox_KOKTOREMOVE
)

func (r RemovalStateViewBoxEnum) MarshalJSON() ([]byte, error) {
    s := RemovalStateViewBoxEnumToValue(r)
    return json.Marshal(s)
}

func (r *RemovalStateViewBoxEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RemovalStateViewBoxEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RemovalStateViewBoxEnum to its string representation
 */
func RemovalStateViewBoxEnumToValue(removalStateViewBoxEnum RemovalStateViewBoxEnum) string {
    switch removalStateViewBoxEnum {
        case RemovalStateViewBox_KDONTREMOVE:
    		return "kDontRemove"
        case RemovalStateViewBox_KMARKEDFORREMOVAL:
    		return "kMarkedForRemoval"
        case RemovalStateViewBox_KOKTOREMOVE:
    		return "kOkToRemove"
        default:
        	return "kDontRemove"
    }
}

/**
 * Converts RemovalStateViewBoxEnum Array to its string Array representation
*/
func RemovalStateViewBoxEnumArrayToValue(removalStateViewBoxEnum []RemovalStateViewBoxEnum) []string {
    convArray := make([]string,len( removalStateViewBoxEnum))
    for i:=0; i<len(removalStateViewBoxEnum);i++ {
        convArray[i] = RemovalStateViewBoxEnumToValue(removalStateViewBoxEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RemovalStateViewBoxEnumFromValue(value string) RemovalStateViewBoxEnum {
    switch value {
        case "kDontRemove":
            return RemovalStateViewBox_KDONTREMOVE
        case "kMarkedForRemoval":
            return RemovalStateViewBox_KMARKEDFORREMOVAL
        case "kOkToRemove":
            return RemovalStateViewBox_KOKTOREMOVE
        default:
            return RemovalStateViewBox_KDONTREMOVE
    }
}
