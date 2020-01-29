// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeViewProtectionSourceEnum enum
 */
type TypeViewProtectionSourceEnum int

/**
 * Value collection for TypeViewProtectionSourceEnum enum
 */
const (
    TypeViewProtectionSource_KVIEWBOX TypeViewProtectionSourceEnum = 1 + iota
    TypeViewProtectionSource_KVIEW
)

func (r TypeViewProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeViewProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeViewProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeViewProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeViewProtectionSourceEnum to its string representation
 */
func TypeViewProtectionSourceEnumToValue(typeViewProtectionSourceEnum TypeViewProtectionSourceEnum) string {
    switch typeViewProtectionSourceEnum {
        case TypeViewProtectionSource_KVIEWBOX:
    		return "kViewBox"
        case TypeViewProtectionSource_KVIEW:
    		return "kView"
        default:
        	return "kViewBox"
    }
}

/**
 * Converts TypeViewProtectionSourceEnum Array to its string Array representation
*/
func TypeViewProtectionSourceEnumArrayToValue(typeViewProtectionSourceEnum []TypeViewProtectionSourceEnum) []string {
    convArray := make([]string,len( typeViewProtectionSourceEnum))
    for i:=0; i<len(typeViewProtectionSourceEnum);i++ {
        convArray[i] = TypeViewProtectionSourceEnumToValue(typeViewProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeViewProtectionSourceEnumFromValue(value string) TypeViewProtectionSourceEnum {
    switch value {
        case "kViewBox":
            return TypeViewProtectionSource_KVIEWBOX
        case "kView":
            return TypeViewProtectionSource_KVIEW
        default:
            return TypeViewProtectionSource_KVIEWBOX
    }
}
