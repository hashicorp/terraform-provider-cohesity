// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeSnapshotTargetSettingsEnum enum
 */
type TypeSnapshotTargetSettingsEnum int

/**
 * Value collection for TypeSnapshotTargetSettingsEnum enum
 */
const (
    TypeSnapshotTargetSettings_KLOCAL TypeSnapshotTargetSettingsEnum = 1 + iota
    TypeSnapshotTargetSettings_KREMOTE
    TypeSnapshotTargetSettings_KARCHIVAL
    TypeSnapshotTargetSettings_KCLOUDDEPLOY
)

func (r TypeSnapshotTargetSettingsEnum) MarshalJSON() ([]byte, error) {
    s := TypeSnapshotTargetSettingsEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeSnapshotTargetSettingsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeSnapshotTargetSettingsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeSnapshotTargetSettingsEnum to its string representation
 */
func TypeSnapshotTargetSettingsEnumToValue(typeSnapshotTargetSettingsEnum TypeSnapshotTargetSettingsEnum) string {
    switch typeSnapshotTargetSettingsEnum {
        case TypeSnapshotTargetSettings_KLOCAL:
    		return "kLocal"
        case TypeSnapshotTargetSettings_KREMOTE:
    		return "kRemote"
        case TypeSnapshotTargetSettings_KARCHIVAL:
    		return "kArchival"
        case TypeSnapshotTargetSettings_KCLOUDDEPLOY:
    		return "kCloudDeploy"
        default:
        	return "kLocal"
    }
}

/**
 * Converts TypeSnapshotTargetSettingsEnum Array to its string Array representation
*/
func TypeSnapshotTargetSettingsEnumArrayToValue(typeSnapshotTargetSettingsEnum []TypeSnapshotTargetSettingsEnum) []string {
    convArray := make([]string,len( typeSnapshotTargetSettingsEnum))
    for i:=0; i<len(typeSnapshotTargetSettingsEnum);i++ {
        convArray[i] = TypeSnapshotTargetSettingsEnumToValue(typeSnapshotTargetSettingsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeSnapshotTargetSettingsEnumFromValue(value string) TypeSnapshotTargetSettingsEnum {
    switch value {
        case "kLocal":
            return TypeSnapshotTargetSettings_KLOCAL
        case "kRemote":
            return TypeSnapshotTargetSettings_KREMOTE
        case "kArchival":
            return TypeSnapshotTargetSettings_KARCHIVAL
        case "kCloudDeploy":
            return TypeSnapshotTargetSettings_KCLOUDDEPLOY
        default:
            return TypeSnapshotTargetSettings_KLOCAL
    }
}
