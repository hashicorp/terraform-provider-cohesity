// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypePhysicalProtectionSourceEnum enum
 */
type TypePhysicalProtectionSourceEnum int

/**
 * Value collection for TypePhysicalProtectionSourceEnum enum
 */
const (
    TypePhysicalProtectionSource_KGROUP TypePhysicalProtectionSourceEnum = 1 + iota
    TypePhysicalProtectionSource_KHOST
    TypePhysicalProtectionSource_KWINDOWSCLUSTER
    TypePhysicalProtectionSource_KORACLERACCLUSTER
    TypePhysicalProtectionSource_KORACLEAPCLUSTER
)

func (r TypePhysicalProtectionSourceEnum) MarshalJSON() ([]byte, error) {
    s := TypePhysicalProtectionSourceEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypePhysicalProtectionSourceEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypePhysicalProtectionSourceEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypePhysicalProtectionSourceEnum to its string representation
 */
func TypePhysicalProtectionSourceEnumToValue(typePhysicalProtectionSourceEnum TypePhysicalProtectionSourceEnum) string {
    switch typePhysicalProtectionSourceEnum {
        case TypePhysicalProtectionSource_KGROUP:
    		return "kGroup"
        case TypePhysicalProtectionSource_KHOST:
    		return "kHost"
        case TypePhysicalProtectionSource_KWINDOWSCLUSTER:
    		return "kWindowsCluster"
        case TypePhysicalProtectionSource_KORACLERACCLUSTER:
    		return "kOracleRACCluster"
        case TypePhysicalProtectionSource_KORACLEAPCLUSTER:
    		return "kOracleAPCluster"
        default:
        	return "kGroup"
    }
}

/**
 * Converts TypePhysicalProtectionSourceEnum Array to its string Array representation
*/
func TypePhysicalProtectionSourceEnumArrayToValue(typePhysicalProtectionSourceEnum []TypePhysicalProtectionSourceEnum) []string {
    convArray := make([]string,len( typePhysicalProtectionSourceEnum))
    for i:=0; i<len(typePhysicalProtectionSourceEnum);i++ {
        convArray[i] = TypePhysicalProtectionSourceEnumToValue(typePhysicalProtectionSourceEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypePhysicalProtectionSourceEnumFromValue(value string) TypePhysicalProtectionSourceEnum {
    switch value {
        case "kGroup":
            return TypePhysicalProtectionSource_KGROUP
        case "kHost":
            return TypePhysicalProtectionSource_KHOST
        case "kWindowsCluster":
            return TypePhysicalProtectionSource_KWINDOWSCLUSTER
        case "kOracleRACCluster":
            return TypePhysicalProtectionSource_KORACLERACCLUSTER
        case "kOracleAPCluster":
            return TypePhysicalProtectionSource_KORACLEAPCLUSTER
        default:
            return TypePhysicalProtectionSource_KGROUP
    }
}
