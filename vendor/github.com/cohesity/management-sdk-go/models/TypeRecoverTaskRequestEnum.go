// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeRecoverTaskRequestEnum enum
 */
type TypeRecoverTaskRequestEnum int

/**
 * Value collection for TypeRecoverTaskRequestEnum enum
 */
const (
    TypeRecoverTaskRequest_KRECOVERVMS TypeRecoverTaskRequestEnum = 1 + iota
    TypeRecoverTaskRequest_KMOUNTVOLUMES
)

func (r TypeRecoverTaskRequestEnum) MarshalJSON() ([]byte, error) {
    s := TypeRecoverTaskRequestEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeRecoverTaskRequestEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeRecoverTaskRequestEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeRecoverTaskRequestEnum to its string representation
 */
func TypeRecoverTaskRequestEnumToValue(typeRecoverTaskRequestEnum TypeRecoverTaskRequestEnum) string {
    switch typeRecoverTaskRequestEnum {
        case TypeRecoverTaskRequest_KRECOVERVMS:
    		return "kRecoverVMs"
        case TypeRecoverTaskRequest_KMOUNTVOLUMES:
    		return "kMountVolumes"
        default:
        	return "kRecoverVMs"
    }
}

/**
 * Converts TypeRecoverTaskRequestEnum Array to its string Array representation
*/
func TypeRecoverTaskRequestEnumArrayToValue(typeRecoverTaskRequestEnum []TypeRecoverTaskRequestEnum) []string {
    convArray := make([]string,len( typeRecoverTaskRequestEnum))
    for i:=0; i<len(typeRecoverTaskRequestEnum);i++ {
        convArray[i] = TypeRecoverTaskRequestEnumToValue(typeRecoverTaskRequestEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeRecoverTaskRequestEnumFromValue(value string) TypeRecoverTaskRequestEnum {
    switch value {
        case "kRecoverVMs":
            return TypeRecoverTaskRequest_KRECOVERVMS
        case "kMountVolumes":
            return TypeRecoverTaskRequest_KMOUNTVOLUMES
        default:
            return TypeRecoverTaskRequest_KRECOVERVMS
    }
}
