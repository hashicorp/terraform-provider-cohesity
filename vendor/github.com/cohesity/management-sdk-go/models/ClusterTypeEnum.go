// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ClusterTypeEnum enum
 */
type ClusterTypeEnum int

/**
 * Value collection for ClusterTypeEnum enum
 */
const (
    ClusterType_KPHYSICAL ClusterTypeEnum = 1 + iota
    ClusterType_KVIRTUALROBO
    ClusterType_KMICROSOFTCLOUD
    ClusterType_KAMAZONCLOUD
    ClusterType_KGOOGLECLOUD
)

func (r ClusterTypeEnum) MarshalJSON() ([]byte, error) {
    s := ClusterTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ClusterTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ClusterTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ClusterTypeEnum to its string representation
 */
func ClusterTypeEnumToValue(clusterTypeEnum ClusterTypeEnum) string {
    switch clusterTypeEnum {
        case ClusterType_KPHYSICAL:
    		return "kPhysical"
        case ClusterType_KVIRTUALROBO:
    		return "kVirtualRobo"
        case ClusterType_KMICROSOFTCLOUD:
    		return "kMicrosoftCloud"
        case ClusterType_KAMAZONCLOUD:
    		return "kAmazonCloud"
        case ClusterType_KGOOGLECLOUD:
    		return "kGoogleCloud"
        default:
        	return "kPhysical"
    }
}

/**
 * Converts ClusterTypeEnum Array to its string Array representation
*/
func ClusterTypeEnumArrayToValue(clusterTypeEnum []ClusterTypeEnum) []string {
    convArray := make([]string,len( clusterTypeEnum))
    for i:=0; i<len(clusterTypeEnum);i++ {
        convArray[i] = ClusterTypeEnumToValue(clusterTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ClusterTypeEnumFromValue(value string) ClusterTypeEnum {
    switch value {
        case "kPhysical":
            return ClusterType_KPHYSICAL
        case "kVirtualRobo":
            return ClusterType_KVIRTUALROBO
        case "kMicrosoftCloud":
            return ClusterType_KMICROSOFTCLOUD
        case "kAmazonCloud":
            return ClusterType_KAMAZONCLOUD
        case "kGoogleCloud":
            return ClusterType_KGOOGLECLOUD
        default:
            return ClusterType_KPHYSICAL
    }
}
