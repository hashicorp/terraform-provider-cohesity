// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeCloneTaskRequestEnum enum
 */
type TypeCloneTaskRequestEnum int

/**
 * Value collection for TypeCloneTaskRequestEnum enum
 */
const (
    TypeCloneTaskRequest_KCLONEVMS TypeCloneTaskRequestEnum = 1 + iota
    TypeCloneTaskRequest_KCLONEVIEW
)

func (r TypeCloneTaskRequestEnum) MarshalJSON() ([]byte, error) {
    s := TypeCloneTaskRequestEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeCloneTaskRequestEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeCloneTaskRequestEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeCloneTaskRequestEnum to its string representation
 */
func TypeCloneTaskRequestEnumToValue(typeCloneTaskRequestEnum TypeCloneTaskRequestEnum) string {
    switch typeCloneTaskRequestEnum {
        case TypeCloneTaskRequest_KCLONEVMS:
    		return "kCloneVMs"
        case TypeCloneTaskRequest_KCLONEVIEW:
    		return "kCloneView"
        default:
        	return "kCloneVMs"
    }
}

/**
 * Converts TypeCloneTaskRequestEnum Array to its string Array representation
*/
func TypeCloneTaskRequestEnumArrayToValue(typeCloneTaskRequestEnum []TypeCloneTaskRequestEnum) []string {
    convArray := make([]string,len( typeCloneTaskRequestEnum))
    for i:=0; i<len(typeCloneTaskRequestEnum);i++ {
        convArray[i] = TypeCloneTaskRequestEnumToValue(typeCloneTaskRequestEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeCloneTaskRequestEnumFromValue(value string) TypeCloneTaskRequestEnum {
    switch value {
        case "kCloneVMs":
            return TypeCloneTaskRequest_KCLONEVMS
        case "kCloneView":
            return TypeCloneTaskRequest_KCLONEVIEW
        default:
            return TypeCloneTaskRequest_KCLONEVMS
    }
}
