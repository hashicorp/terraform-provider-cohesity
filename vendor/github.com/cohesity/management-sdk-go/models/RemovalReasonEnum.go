// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for RemovalReasonEnum enum
 */
type RemovalReasonEnum int

/**
 * Value collection for RemovalReasonEnum enum
 */
const (
    RemovalReason_KAUTOHEALTHCHECK RemovalReasonEnum = 1 + iota
    RemovalReason_KUSERGRACEFULREMOVAL
    RemovalReason_KUSERAVOIDACCESS
    RemovalReason_KUSERGRACEFULNODEREMOVAL
    RemovalReason_KUSERREMOVEDOWNNODE
)

func (r RemovalReasonEnum) MarshalJSON() ([]byte, error) {
    s := RemovalReasonEnumToValue(r)
    return json.Marshal(s)
}

func (r *RemovalReasonEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  RemovalReasonEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts RemovalReasonEnum to its string representation
 */
func RemovalReasonEnumToValue(removalReasonEnum RemovalReasonEnum) string {
    switch removalReasonEnum {
        case RemovalReason_KAUTOHEALTHCHECK:
    		return "kAutoHealthCheck"
        case RemovalReason_KUSERGRACEFULREMOVAL:
    		return "kUserGracefulRemoval"
        case RemovalReason_KUSERAVOIDACCESS:
    		return "kUserAvoidAccess"
        case RemovalReason_KUSERGRACEFULNODEREMOVAL:
    		return "kUserGracefulNodeRemoval"
        case RemovalReason_KUSERREMOVEDOWNNODE:
    		return "kUserRemoveDownNode"
        default:
        	return "kAutoHealthCheck"
    }
}

/**
 * Converts RemovalReasonEnum Array to its string Array representation
*/
func RemovalReasonEnumArrayToValue(removalReasonEnum []RemovalReasonEnum) []string {
    convArray := make([]string,len( removalReasonEnum))
    for i:=0; i<len(removalReasonEnum);i++ {
        convArray[i] = RemovalReasonEnumToValue(removalReasonEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func RemovalReasonEnumFromValue(value string) RemovalReasonEnum {
    switch value {
        case "kAutoHealthCheck":
            return RemovalReason_KAUTOHEALTHCHECK
        case "kUserGracefulRemoval":
            return RemovalReason_KUSERGRACEFULREMOVAL
        case "kUserAvoidAccess":
            return RemovalReason_KUSERAVOIDACCESS
        case "kUserGracefulNodeRemoval":
            return RemovalReason_KUSERGRACEFULNODEREMOVAL
        case "kUserRemoveDownNode":
            return RemovalReason_KUSERREMOVEDOWNNODE
        default:
            return RemovalReason_KAUTOHEALTHCHECK
    }
}
