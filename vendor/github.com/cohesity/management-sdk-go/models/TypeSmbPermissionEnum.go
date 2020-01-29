// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeSmbPermissionEnum enum
 */
type TypeSmbPermissionEnum int

/**
 * Value collection for TypeSmbPermissionEnum enum
 */
const (
    TypeSmbPermission_KALLOW TypeSmbPermissionEnum = 1 + iota
    TypeSmbPermission_KDENY
    TypeSmbPermission_KSPECIALTYPE
)

func (r TypeSmbPermissionEnum) MarshalJSON() ([]byte, error) {
    s := TypeSmbPermissionEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeSmbPermissionEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeSmbPermissionEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeSmbPermissionEnum to its string representation
 */
func TypeSmbPermissionEnumToValue(typeSmbPermissionEnum TypeSmbPermissionEnum) string {
    switch typeSmbPermissionEnum {
        case TypeSmbPermission_KALLOW:
    		return "kAllow"
        case TypeSmbPermission_KDENY:
    		return "kDeny"
        case TypeSmbPermission_KSPECIALTYPE:
    		return "kSpecialType"
        default:
        	return "kAllow"
    }
}

/**
 * Converts TypeSmbPermissionEnum Array to its string Array representation
*/
func TypeSmbPermissionEnumArrayToValue(typeSmbPermissionEnum []TypeSmbPermissionEnum) []string {
    convArray := make([]string,len( typeSmbPermissionEnum))
    for i:=0; i<len(typeSmbPermissionEnum);i++ {
        convArray[i] = TypeSmbPermissionEnumToValue(typeSmbPermissionEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeSmbPermissionEnumFromValue(value string) TypeSmbPermissionEnum {
    switch value {
        case "kAllow":
            return TypeSmbPermission_KALLOW
        case "kDeny":
            return TypeSmbPermission_KDENY
        case "kSpecialType":
            return TypeSmbPermission_KSPECIALTYPE
        default:
            return TypeSmbPermission_KALLOW
    }
}
