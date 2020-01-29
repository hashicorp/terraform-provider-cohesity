// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeNetappProtectionSourceEnum enum
 */
type TypeNetappProtectionSourceEnum int

/**
 * Value collection for TypeNetappProtectionSourceEnum enum
 */
const (
    TypeNetappProtectionSource_KCLUSTER TypeNetappProtectionSourceEnum = 1 + iota
    TypeNetappProtectionSource_KVSERVER
    TypeNetappProtectionSource_KVOLUME
)

func (r TypeNetappProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypeNetappProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeNetappProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeNetappProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeNetappProtectionSourceEnum to its string representation
 */
func TypeNetappProtectionSourceEnumToValue(typeNetappProtectionSourceEnum TypeNetappProtectionSourceEnum) string {
    switch typeNetappProtectionSourceEnum {
        case TypeNetappProtectionSource_KCLUSTER:
    		return "kCluster"
        case TypeNetappProtectionSource_KVSERVER:
    		return "kVserver"
        case TypeNetappProtectionSource_KVOLUME:
    		return "kVolume"
        default:
        	return "kCluster"
    }
}

/**
 * Converts TypeNetappProtectionSourceEnum Array to its string Array representation
*/
func TypeNetappProtectionSourceEnumArrayToValue(typeNetappProtectionSourceEnum []TypeNetappProtectionSourceEnum) []string {
    convArray := make([]string,len( typeNetappProtectionSourceEnum))
    for i:=0; i<len(typeNetappProtectionSourceEnum);i++ {
        convArray[i] = TypeNetappProtectionSourceEnumToValue(typeNetappProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeNetappProtectionSourceEnumFromValue(value string) TypeNetappProtectionSourceEnum {
    switch value {
        case "kCluster":
            return TypeNetappProtectionSource_KCLUSTER
        case "kVserver":
            return TypeNetappProtectionSource_KVSERVER
        case "kVolume":
            return TypeNetappProtectionSource_KVOLUME
        default:
            return TypeNetappProtectionSource_KCLUSTER
    }
}
