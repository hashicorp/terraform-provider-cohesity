// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeAdRestoreOptionsEnum enum
 */
type TypeAdRestoreOptionsEnum int

/**
 * Value collection for TypeAdRestoreOptionsEnum enum
 */
const (
    TypeAdRestoreOptions_KNONE TypeAdRestoreOptionsEnum = 1 + iota
    TypeAdRestoreOptions_KOBJECTS
    TypeAdRestoreOptions_KOBJECTATTRIBUTES
)

func (r TypeAdRestoreOptionsEnum) MarshalJSON() ([]byte, error) {
    s := TypeAdRestoreOptionsEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeAdRestoreOptionsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeAdRestoreOptionsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeAdRestoreOptionsEnum to its string representation
 */
func TypeAdRestoreOptionsEnumToValue(typeAdRestoreOptionsEnum TypeAdRestoreOptionsEnum) string {
    switch typeAdRestoreOptionsEnum {
        case TypeAdRestoreOptions_KNONE:
    		return "kNone"
        case TypeAdRestoreOptions_KOBJECTS:
    		return "kObjects"
        case TypeAdRestoreOptions_KOBJECTATTRIBUTES:
    		return "kObjectAttributes"
        default:
        	return "kNone"
    }
}

/**
 * Converts TypeAdRestoreOptionsEnum Array to its string Array representation
*/
func TypeAdRestoreOptionsEnumArrayToValue(typeAdRestoreOptionsEnum []TypeAdRestoreOptionsEnum) []string {
    convArray := make([]string,len( typeAdRestoreOptionsEnum))
    for i:=0; i<len(typeAdRestoreOptionsEnum);i++ {
        convArray[i] = TypeAdRestoreOptionsEnumToValue(typeAdRestoreOptionsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeAdRestoreOptionsEnumFromValue(value string) TypeAdRestoreOptionsEnum {
    switch value {
        case "kNone":
            return TypeAdRestoreOptions_KNONE
        case "kObjects":
            return TypeAdRestoreOptions_KOBJECTS
        case "kObjectAttributes":
            return TypeAdRestoreOptions_KOBJECTATTRIBUTES
        default:
            return TypeAdRestoreOptions_KNONE
    }
}
