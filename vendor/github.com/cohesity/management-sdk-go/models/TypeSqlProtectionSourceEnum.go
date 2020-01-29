// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeSqlProtectionSourceEnum enum
 */
type TypeSqlProtectionSourceEnum int

/**
 * Value collection for TypeSqlProtectionSourceEnum enum
 */
const (
    TypeSqlProtectionSource_KINSTANCE TypeSqlProtectionSourceEnum = 1 + iota
    TypeSqlProtectionSource_KDATABASE
    TypeSqlProtectionSource_KAAG
    TypeSqlProtectionSource_KAAGROOTCONTAINER
    TypeSqlProtectionSource_KROOTCONTAINER
)

func (r TypeSqlProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeSqlProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeSqlProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeSqlProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeSqlProtectionSourceEnum to its string representation
 */
func TypeSqlProtectionSourceEnumToValue(typeSqlProtectionSourceEnum TypeSqlProtectionSourceEnum) string {
    switch typeSqlProtectionSourceEnum {
        case TypeSqlProtectionSource_KINSTANCE:
    		return "kInstance"
        case TypeSqlProtectionSource_KDATABASE:
    		return "kDatabase"
        case TypeSqlProtectionSource_KAAG:
    		return "kAAG"
        case TypeSqlProtectionSource_KAAGROOTCONTAINER:
    		return "kAAGRootContainer"
        case TypeSqlProtectionSource_KROOTCONTAINER:
    		return "kRootContainer"
        default:
        	return "kInstance"
    }
}

/**
 * Converts TypeSqlProtectionSourceEnum Array to its string Array representation
*/
func TypeSqlProtectionSourceEnumArrayToValue(typeSqlProtectionSourceEnum []TypeSqlProtectionSourceEnum) []string {
    convArray := make([]string,len( typeSqlProtectionSourceEnum))
    for i:=0; i<len(typeSqlProtectionSourceEnum);i++ {
        convArray[i] = TypeSqlProtectionSourceEnumToValue(typeSqlProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeSqlProtectionSourceEnumFromValue(value string) TypeSqlProtectionSourceEnum {
    switch value {
        case "kInstance":
            return TypeSqlProtectionSource_KINSTANCE
        case "kDatabase":
            return TypeSqlProtectionSource_KDATABASE
        case "kAAG":
            return TypeSqlProtectionSource_KAAG
        case "kAAGRootContainer":
            return TypeSqlProtectionSource_KAAGROOTCONTAINER
        case "kRootContainer":
            return TypeSqlProtectionSource_KROOTCONTAINER
        default:
            return TypeSqlProtectionSource_KINSTANCE
    }
}
