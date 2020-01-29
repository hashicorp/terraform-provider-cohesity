// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeIsilonProtectionSourceEnum enum
 */
type TypeIsilonProtectionSourceEnum int

/**
 * Value collection for TypeIsilonProtectionSourceEnum enum
 */
const (
    TypeIsilonProtectionSource_KCLUSTER TypeIsilonProtectionSourceEnum = 1 + iota
    TypeIsilonProtectionSource_KZONE
    TypeIsilonProtectionSource_KMOUNTPOINT
)

func (r TypeIsilonProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeIsilonProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeIsilonProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeIsilonProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeIsilonProtectionSourceEnum to its string representation
 */
func TypeIsilonProtectionSourceEnumToValue(typeIsilonProtectionSourceEnum TypeIsilonProtectionSourceEnum) string {
    switch typeIsilonProtectionSourceEnum {
        case TypeIsilonProtectionSource_KCLUSTER:
    		return "kCluster"
        case TypeIsilonProtectionSource_KZONE:
    		return "kZone"
        case TypeIsilonProtectionSource_KMOUNTPOINT:
    		return "kMountPoint"
        default:
        	return "kCluster"
    }
}

/**
 * Converts TypeIsilonProtectionSourceEnum Array to its string Array representation
*/
func TypeIsilonProtectionSourceEnumArrayToValue(typeIsilonProtectionSourceEnum []TypeIsilonProtectionSourceEnum) []string {
    convArray := make([]string,len( typeIsilonProtectionSourceEnum))
    for i:=0; i<len(typeIsilonProtectionSourceEnum);i++ {
        convArray[i] = TypeIsilonProtectionSourceEnumToValue(typeIsilonProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeIsilonProtectionSourceEnumFromValue(value string) TypeIsilonProtectionSourceEnum {
    switch value {
        case "kCluster":
            return TypeIsilonProtectionSource_KCLUSTER
        case "kZone":
            return TypeIsilonProtectionSource_KZONE
        case "kMountPoint":
            return TypeIsilonProtectionSource_KMOUNTPOINT
        default:
            return TypeIsilonProtectionSource_KCLUSTER
    }
}
