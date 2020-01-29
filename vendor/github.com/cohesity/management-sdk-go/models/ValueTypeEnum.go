// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ValueTypeEnum enum
 */
type ValueTypeEnum int

/**
 * Value collection for ValueTypeEnum enum
 */
const (
    ValueType_KINT64 ValueTypeEnum = 1 + iota
    ValueType_KDOUBLE
    ValueType_KSTRING
    ValueType_KBYTES
)

func (r ValueTypeEnum) MarshalJSON() ([]byte, error) {
    s := ValueTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ValueTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ValueTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ValueTypeEnum to its string representation
 */
func ValueTypeEnumToValue(valueTypeEnum ValueTypeEnum) string {
    switch valueTypeEnum {
        case ValueType_KINT64:
    		return "kInt64"
        case ValueType_KDOUBLE:
    		return "kDouble"
        case ValueType_KSTRING:
    		return "kString"
        case ValueType_KBYTES:
    		return "kBytes"
        default:
        	return "kInt64"
    }
}

/**
 * Converts ValueTypeEnum Array to its string Array representation
*/
func ValueTypeEnumArrayToValue(valueTypeEnum []ValueTypeEnum) []string {
    convArray := make([]string,len( valueTypeEnum))
    for i:=0; i<len(valueTypeEnum);i++ {
        convArray[i] = ValueTypeEnumToValue(valueTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ValueTypeEnumFromValue(value string) ValueTypeEnum {
    switch value {
        case "kInt64":
            return ValueType_KINT64
        case "kDouble":
            return ValueType_KDOUBLE
        case "kString":
            return ValueType_KSTRING
        case "kBytes":
            return ValueType_KBYTES
        default:
            return ValueType_KINT64
    }
}
