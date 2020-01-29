// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SqlOptionsEnum enum
 */
type SqlOptionsEnum int

/**
 * Value collection for SqlOptionsEnum enum
 */
const (
    SqlOptions_KCREATE SqlOptionsEnum = 1 + iota
    SqlOptions_KUPDATE
    SqlOptions_KFINALIZE
)

func (r SqlOptionsEnum) MarshalJSON() ([]byte, error) {
    s := SqlOptionsEnumToValue(r)
    return json.Marshal(s)
}

func (r *SqlOptionsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SqlOptionsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SqlOptionsEnum to its string representation
 */
func SqlOptionsEnumToValue(sqlOptionsEnum SqlOptionsEnum) string {
    switch sqlOptionsEnum {
        case SqlOptions_KCREATE:
    		return "kCreate"
        case SqlOptions_KUPDATE:
    		return "kUpdate"
        case SqlOptions_KFINALIZE:
    		return "kFinalize"
        default:
        	return "kCreate"
    }
}

/**
 * Converts SqlOptionsEnum Array to its string Array representation
*/
func SqlOptionsEnumArrayToValue(sqlOptionsEnum []SqlOptionsEnum) []string {
    convArray := make([]string,len( sqlOptionsEnum))
    for i:=0; i<len(sqlOptionsEnum);i++ {
        convArray[i] = SqlOptionsEnumToValue(sqlOptionsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SqlOptionsEnumFromValue(value string) SqlOptionsEnum {
    switch value {
        case "kCreate":
            return SqlOptions_KCREATE
        case "kUpdate":
            return SqlOptions_KUPDATE
        case "kFinalize":
            return SqlOptions_KFINALIZE
        default:
            return SqlOptions_KCREATE
    }
}
