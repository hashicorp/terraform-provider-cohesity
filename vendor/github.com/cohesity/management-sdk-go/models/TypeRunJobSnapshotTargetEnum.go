// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeRunJobSnapshotTargetEnum enum
 */
type TypeRunJobSnapshotTargetEnum int

/**
 * Value collection for TypeRunJobSnapshotTargetEnum enum
 */
const (
    TypeRunJobSnapshotTarget_KLOCAL TypeRunJobSnapshotTargetEnum = 1 + iota
    TypeRunJobSnapshotTarget_KREMOTE
    TypeRunJobSnapshotTarget_KARCHIVAL
    TypeRunJobSnapshotTarget_KCLOUDDEPLOY
)

func (r TypeRunJobSnapshotTargetEnum) MarshalJSON() ([]byte, error) {
    s := TypeRunJobSnapshotTargetEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeRunJobSnapshotTargetEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeRunJobSnapshotTargetEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeRunJobSnapshotTargetEnum to its string representation
 */
func TypeRunJobSnapshotTargetEnumToValue(typeRunJobSnapshotTargetEnum TypeRunJobSnapshotTargetEnum) string {
    switch typeRunJobSnapshotTargetEnum {
        case TypeRunJobSnapshotTarget_KLOCAL:
    		return "kLocal"
        case TypeRunJobSnapshotTarget_KREMOTE:
    		return "kRemote"
        case TypeRunJobSnapshotTarget_KARCHIVAL:
    		return "kArchival"
        case TypeRunJobSnapshotTarget_KCLOUDDEPLOY:
    		return "kCloudDeploy"
        default:
        	return "kLocal"
    }
}

/**
 * Converts TypeRunJobSnapshotTargetEnum Array to its string Array representation
*/
func TypeRunJobSnapshotTargetEnumArrayToValue(typeRunJobSnapshotTargetEnum []TypeRunJobSnapshotTargetEnum) []string {
    convArray := make([]string,len( typeRunJobSnapshotTargetEnum))
    for i:=0; i<len(typeRunJobSnapshotTargetEnum);i++ {
        convArray[i] = TypeRunJobSnapshotTargetEnumToValue(typeRunJobSnapshotTargetEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeRunJobSnapshotTargetEnumFromValue(value string) TypeRunJobSnapshotTargetEnum {
    switch value {
        case "kLocal":
            return TypeRunJobSnapshotTarget_KLOCAL
        case "kRemote":
            return TypeRunJobSnapshotTarget_KREMOTE
        case "kArchival":
            return TypeRunJobSnapshotTarget_KARCHIVAL
        case "kCloudDeploy":
            return TypeRunJobSnapshotTarget_KCLOUDDEPLOY
        default:
            return TypeRunJobSnapshotTarget_KLOCAL
    }
}
