// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypePureProtectionSourceEnum enum
 */
type TypePureProtectionSourceEnum int

/**
 * Value collection for TypePureProtectionSourceEnum enum
 */
const (
    TypePureProtectionSource_KSTORAGEARRAY TypePureProtectionSourceEnum = 1 + iota
    TypePureProtectionSource_KVOLUME
)

func (r TypePureProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypePureProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypePureProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypePureProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypePureProtectionSourceEnum to its string representation
 */
func TypePureProtectionSourceEnumToValue(typePureProtectionSourceEnum TypePureProtectionSourceEnum) string {
    switch typePureProtectionSourceEnum {
        case TypePureProtectionSource_KSTORAGEARRAY:
    		return "kStorageArray"
        case TypePureProtectionSource_KVOLUME:
    		return "kVolume"
        default:
        	return "kStorageArray"
    }
}

/**
 * Converts TypePureProtectionSourceEnum Array to its string Array representation
*/
func TypePureProtectionSourceEnumArrayToValue(typePureProtectionSourceEnum []TypePureProtectionSourceEnum) []string {
    convArray := make([]string,len( typePureProtectionSourceEnum))
    for i:=0; i<len(typePureProtectionSourceEnum);i++ {
        convArray[i] = TypePureProtectionSourceEnumToValue(typePureProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypePureProtectionSourceEnumFromValue(value string) TypePureProtectionSourceEnum {
    switch value {
        case "kStorageArray":
            return TypePureProtectionSource_KSTORAGEARRAY
        case "kVolume":
            return TypePureProtectionSource_KVOLUME
        default:
            return TypePureProtectionSource_KSTORAGEARRAY
    }
}
