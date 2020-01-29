// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RemovalStateEnum enum
 */
type RemovalStateEnum int

/**
 * Value collection for RemovalStateEnum enum
 */
const (
    RemovalState_KDONTREMOVE RemovalStateEnum = 1 + iota
    RemovalState_KMARKEDFORREMOVAL
    RemovalState_KOKTOREMOVE
)

func (r RemovalStateEnum) MarshalJSON() ([]byte, error) {
    s := RemovalStateEnumToValue(r)
    return json.Marshal(s)
}

func (r *RemovalStateEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RemovalStateEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RemovalStateEnum to its string representation
 */
func RemovalStateEnumToValue(removalStateEnum RemovalStateEnum) string {
    switch removalStateEnum {
        case RemovalState_KDONTREMOVE:
    		return "kDontRemove"
        case RemovalState_KMARKEDFORREMOVAL:
    		return "kMarkedForRemoval"
        case RemovalState_KOKTOREMOVE:
    		return "kOkToRemove"
        default:
        	return "kDontRemove"
    }
}

/**
 * Converts RemovalStateEnum Array to its string Array representation
*/
func RemovalStateEnumArrayToValue(removalStateEnum []RemovalStateEnum) []string {
    convArray := make([]string,len( removalStateEnum))
    for i:=0; i<len(removalStateEnum);i++ {
        convArray[i] = RemovalStateEnumToValue(removalStateEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RemovalStateEnumFromValue(value string) RemovalStateEnum {
    switch value {
        case "kDontRemove":
            return RemovalState_KDONTREMOVE
        case "kMarkedForRemoval":
            return RemovalState_KMARKEDFORREMOVAL
        case "kOkToRemove":
            return RemovalState_KOKTOREMOVE
        default:
            return RemovalState_KDONTREMOVE
    }
}
