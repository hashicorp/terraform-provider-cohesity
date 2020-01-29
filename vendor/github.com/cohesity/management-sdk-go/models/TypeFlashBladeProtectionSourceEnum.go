// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeFlashBladeProtectionSourceEnum enum
 */
type TypeFlashBladeProtectionSourceEnum int

/**
 * Value collection for TypeFlashBladeProtectionSourceEnum enum
 */
const (
    TypeFlashBladeProtectionSource_KSTORAGEARRAY TypeFlashBladeProtectionSourceEnum = 1 + iota
    TypeFlashBladeProtectionSource_KFILESYSTEM
)

func (r TypeFlashBladeProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeFlashBladeProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeFlashBladeProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeFlashBladeProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeFlashBladeProtectionSourceEnum to its string representation
 */
func TypeFlashBladeProtectionSourceEnumToValue(typeFlashBladeProtectionSourceEnum TypeFlashBladeProtectionSourceEnum) string {
    switch typeFlashBladeProtectionSourceEnum {
        case TypeFlashBladeProtectionSource_KSTORAGEARRAY:
    		return "kStorageArray"
        case TypeFlashBladeProtectionSource_KFILESYSTEM:
    		return "kFileSystem"
        default:
        	return "kStorageArray"
    }
}

/**
 * Converts TypeFlashBladeProtectionSourceEnum Array to its string Array representation
*/
func TypeFlashBladeProtectionSourceEnumArrayToValue(typeFlashBladeProtectionSourceEnum []TypeFlashBladeProtectionSourceEnum) []string {
    convArray := make([]string,len( typeFlashBladeProtectionSourceEnum))
    for i:=0; i<len(typeFlashBladeProtectionSourceEnum);i++ {
        convArray[i] = TypeFlashBladeProtectionSourceEnumToValue(typeFlashBladeProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeFlashBladeProtectionSourceEnumFromValue(value string) TypeFlashBladeProtectionSourceEnum {
    switch value {
        case "kStorageArray":
            return TypeFlashBladeProtectionSource_KSTORAGEARRAY
        case "kFileSystem":
            return TypeFlashBladeProtectionSource_KFILESYSTEM
        default:
            return TypeFlashBladeProtectionSource_KSTORAGEARRAY
    }
}
