// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeOracleProtectionSourceEnum enum
 */
type TypeOracleProtectionSourceEnum int

/**
 * Value collection for TypeOracleProtectionSourceEnum enum
 */
const (
    TypeOracleProtectionSource_KRACROOTCONTAINER TypeOracleProtectionSourceEnum = 1 + iota
    TypeOracleProtectionSource_KROOTCONTAINER
    TypeOracleProtectionSource_KHOST
    TypeOracleProtectionSource_KDATABASE
    TypeOracleProtectionSource_KTABLESPACE
    TypeOracleProtectionSource_KTABLE
)

func (r TypeOracleProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeOracleProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeOracleProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeOracleProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeOracleProtectionSourceEnum to its string representation
 */
func TypeOracleProtectionSourceEnumToValue(typeOracleProtectionSourceEnum TypeOracleProtectionSourceEnum) string {
    switch typeOracleProtectionSourceEnum {
        case TypeOracleProtectionSource_KRACROOTCONTAINER:
    		return "kRACRootContainer"
        case TypeOracleProtectionSource_KROOTCONTAINER:
    		return "kRootContainer"
        case TypeOracleProtectionSource_KHOST:
    		return "kHost"
        case TypeOracleProtectionSource_KDATABASE:
    		return "kDatabase"
        case TypeOracleProtectionSource_KTABLESPACE:
    		return "kTableSpace"
        case TypeOracleProtectionSource_KTABLE:
    		return "kTable"
        default:
        	return "kRACRootContainer"
    }
}

/**
 * Converts TypeOracleProtectionSourceEnum Array to its string Array representation
*/
func TypeOracleProtectionSourceEnumArrayToValue(typeOracleProtectionSourceEnum []TypeOracleProtectionSourceEnum) []string {
    convArray := make([]string,len( typeOracleProtectionSourceEnum))
    for i:=0; i<len(typeOracleProtectionSourceEnum);i++ {
        convArray[i] = TypeOracleProtectionSourceEnumToValue(typeOracleProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeOracleProtectionSourceEnumFromValue(value string) TypeOracleProtectionSourceEnum {
    switch value {
        case "kRACRootContainer":
            return TypeOracleProtectionSource_KRACROOTCONTAINER
        case "kRootContainer":
            return TypeOracleProtectionSource_KROOTCONTAINER
        case "kHost":
            return TypeOracleProtectionSource_KHOST
        case "kDatabase":
            return TypeOracleProtectionSource_KDATABASE
        case "kTableSpace":
            return TypeOracleProtectionSource_KTABLESPACE
        case "kTable":
            return TypeOracleProtectionSource_KTABLE
        default:
            return TypeOracleProtectionSource_KRACROOTCONTAINER
    }
}
