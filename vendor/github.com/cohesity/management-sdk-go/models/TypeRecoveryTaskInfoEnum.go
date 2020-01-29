// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeRecoveryTaskInfoEnum enum
 */
type TypeRecoveryTaskInfoEnum int

/**
 * Value collection for TypeRecoveryTaskInfoEnum enum
 */
const (
    TypeRecoveryTaskInfo_LOCAL TypeRecoveryTaskInfoEnum = 1 + iota
    TypeRecoveryTaskInfo_ARCHIVE
)

func (r TypeRecoveryTaskInfoEnum) MarshalJSON() ([]byte, error) {
    s := TypeRecoveryTaskInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeRecoveryTaskInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeRecoveryTaskInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeRecoveryTaskInfoEnum to its string representation
 */
func TypeRecoveryTaskInfoEnumToValue(typeRecoveryTaskInfoEnum TypeRecoveryTaskInfoEnum) string {
    switch typeRecoveryTaskInfoEnum {
        case TypeRecoveryTaskInfo_LOCAL:
    		return "local"
        case TypeRecoveryTaskInfo_ARCHIVE:
    		return "archive"
        default:
        	return "local"
    }
}

/**
 * Converts TypeRecoveryTaskInfoEnum Array to its string Array representation
*/
func TypeRecoveryTaskInfoEnumArrayToValue(typeRecoveryTaskInfoEnum []TypeRecoveryTaskInfoEnum) []string {
    convArray := make([]string,len( typeRecoveryTaskInfoEnum))
    for i:=0; i<len(typeRecoveryTaskInfoEnum);i++ {
        convArray[i] = TypeRecoveryTaskInfoEnumToValue(typeRecoveryTaskInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeRecoveryTaskInfoEnumFromValue(value string) TypeRecoveryTaskInfoEnum {
    switch value {
        case "local":
            return TypeRecoveryTaskInfo_LOCAL
        case "archive":
            return TypeRecoveryTaskInfo_ARCHIVE
        default:
            return TypeRecoveryTaskInfo_LOCAL
    }
}
