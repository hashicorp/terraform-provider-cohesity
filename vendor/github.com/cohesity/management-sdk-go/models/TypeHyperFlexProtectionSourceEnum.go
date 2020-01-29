// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeHyperFlexProtectionSourceEnum enum
 */
type TypeHyperFlexProtectionSourceEnum int

/**
 * Value collection for TypeHyperFlexProtectionSourceEnum enum
 */
const (
    TypeHyperFlexProtectionSource_KSERVER TypeHyperFlexProtectionSourceEnum = 1 + iota
)

func (r TypeHyperFlexProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeHyperFlexProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeHyperFlexProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeHyperFlexProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeHyperFlexProtectionSourceEnum to its string representation
 */
func TypeHyperFlexProtectionSourceEnumToValue(typeHyperFlexProtectionSourceEnum TypeHyperFlexProtectionSourceEnum) string {
    switch typeHyperFlexProtectionSourceEnum {
        case TypeHyperFlexProtectionSource_KSERVER:
    		return "kServer"
        default:
        	return "kServer"
    }
}

/**
 * Converts TypeHyperFlexProtectionSourceEnum Array to its string Array representation
*/
func TypeHyperFlexProtectionSourceEnumArrayToValue(typeHyperFlexProtectionSourceEnum []TypeHyperFlexProtectionSourceEnum) []string {
    convArray := make([]string,len( typeHyperFlexProtectionSourceEnum))
    for i:=0; i<len(typeHyperFlexProtectionSourceEnum);i++ {
        convArray[i] = TypeHyperFlexProtectionSourceEnumToValue(typeHyperFlexProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeHyperFlexProtectionSourceEnumFromValue(value string) TypeHyperFlexProtectionSourceEnum {
    switch value {
        case "kServer":
            return TypeHyperFlexProtectionSource_KSERVER
        default:
            return TypeHyperFlexProtectionSource_KSERVER
    }
}
