// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for DbTypeEnum enum
 */
type DbTypeEnum int

/**
 * Value collection for DbTypeEnum enum
 */
const (
    DbType_KSINGLEINSTANCE DbTypeEnum = 1 + iota
    DbType_KRACDATABASE
)

func (r DbTypeEnum) MarshalJSON() ([]byte, error) {
    s := DbTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *DbTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  DbTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts DbTypeEnum to its string representation
 */
func DbTypeEnumToValue(dbTypeEnum DbTypeEnum) string {
    switch dbTypeEnum {
        case DbType_KSINGLEINSTANCE:
    		return "kSingleInstance"
        case DbType_KRACDATABASE:
    		return "kRACDatabase"
        default:
        	return "kSingleInstance"
    }
}

/**
 * Converts DbTypeEnum Array to its string Array representation
*/
func DbTypeEnumArrayToValue(dbTypeEnum []DbTypeEnum) []string {
    convArray := make([]string,len( dbTypeEnum))
    for i:=0; i<len(dbTypeEnum);i++ {
        convArray[i] = DbTypeEnumToValue(dbTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DbTypeEnumFromValue(value string) DbTypeEnum {
    switch value {
        case "kSingleInstance":
            return DbType_KSINGLEINSTANCE
        case "kRACDatabase":
            return DbType_KRACDATABASE
        default:
            return DbType_KSINGLEINSTANCE
    }
}
