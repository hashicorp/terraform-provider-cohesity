// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeGpfsProtectionSourceEnum enum
 */
type TypeGpfsProtectionSourceEnum int

/**
 * Value collection for TypeGpfsProtectionSourceEnum enum
 */
const (
    TypeGpfsProtectionSource_KCLUSTER TypeGpfsProtectionSourceEnum = 1 + iota
    TypeGpfsProtectionSource_KFILESYSTEM
    TypeGpfsProtectionSource_KFILESET
)

func (r TypeGpfsProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeGpfsProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeGpfsProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeGpfsProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeGpfsProtectionSourceEnum to its string representation
 */
func TypeGpfsProtectionSourceEnumToValue(typeGpfsProtectionSourceEnum TypeGpfsProtectionSourceEnum) string {
    switch typeGpfsProtectionSourceEnum {
        case TypeGpfsProtectionSource_KCLUSTER:
    		return "kCluster"
        case TypeGpfsProtectionSource_KFILESYSTEM:
    		return "kFilesystem"
        case TypeGpfsProtectionSource_KFILESET:
    		return "kFileset"
        default:
        	return "kCluster"
    }
}

/**
 * Converts TypeGpfsProtectionSourceEnum Array to its string Array representation
*/
func TypeGpfsProtectionSourceEnumArrayToValue(typeGpfsProtectionSourceEnum []TypeGpfsProtectionSourceEnum) []string {
    convArray := make([]string,len( typeGpfsProtectionSourceEnum))
    for i:=0; i<len(typeGpfsProtectionSourceEnum);i++ {
        convArray[i] = TypeGpfsProtectionSourceEnumToValue(typeGpfsProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeGpfsProtectionSourceEnumFromValue(value string) TypeGpfsProtectionSourceEnum {
    switch value {
        case "kCluster":
            return TypeGpfsProtectionSource_KCLUSTER
        case "kFilesystem":
            return TypeGpfsProtectionSource_KFILESYSTEM
        case "kFileset":
            return TypeGpfsProtectionSource_KFILESET
        default:
            return TypeGpfsProtectionSource_KCLUSTER
    }
}
