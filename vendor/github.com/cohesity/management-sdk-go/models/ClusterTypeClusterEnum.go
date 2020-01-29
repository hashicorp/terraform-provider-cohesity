// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ClusterTypeClusterEnum enum
 */
type ClusterTypeClusterEnum int

/**
 * Value collection for ClusterTypeClusterEnum enum
 */
const (
    ClusterTypeCluster_KPHYSICAL ClusterTypeClusterEnum = 1 + iota
    ClusterTypeCluster_KVIRTUALROBO
    ClusterTypeCluster_KMICROSOFTCLOUD
    ClusterTypeCluster_KAMAZONCLOUD
    ClusterTypeCluster_KGOOGLECLOUD
)

func (r ClusterTypeClusterEnum) MarshalJSON() ([]byte, error) {
    s := ClusterTypeClusterEnumToValue(r)
    return json.Marshal(s)
}

func (r *ClusterTypeClusterEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ClusterTypeClusterEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ClusterTypeClusterEnum to its string representation
 */
func ClusterTypeClusterEnumToValue(clusterTypeClusterEnum ClusterTypeClusterEnum) string {
    switch clusterTypeClusterEnum {
        case ClusterTypeCluster_KPHYSICAL:
    		return "kPhysical"
        case ClusterTypeCluster_KVIRTUALROBO:
    		return "kVirtualRobo"
        case ClusterTypeCluster_KMICROSOFTCLOUD:
    		return "kMicrosoftCloud"
        case ClusterTypeCluster_KAMAZONCLOUD:
    		return "kAmazonCloud"
        case ClusterTypeCluster_KGOOGLECLOUD:
    		return "kGoogleCloud"
        default:
        	return "kPhysical"
    }
}

/**
 * Converts ClusterTypeClusterEnum Array to its string Array representation
*/
func ClusterTypeClusterEnumArrayToValue(clusterTypeClusterEnum []ClusterTypeClusterEnum) []string {
    convArray := make([]string,len( clusterTypeClusterEnum))
    for i:=0; i<len(clusterTypeClusterEnum);i++ {
        convArray[i] = ClusterTypeClusterEnumToValue(clusterTypeClusterEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ClusterTypeClusterEnumFromValue(value string) ClusterTypeClusterEnum {
    switch value {
        case "kPhysical":
            return ClusterTypeCluster_KPHYSICAL
        case "kVirtualRobo":
            return ClusterTypeCluster_KVIRTUALROBO
        case "kMicrosoftCloud":
            return ClusterTypeCluster_KMICROSOFTCLOUD
        case "kAmazonCloud":
            return ClusterTypeCluster_KAMAZONCLOUD
        case "kGoogleCloud":
            return ClusterTypeCluster_KGOOGLECLOUD
        default:
            return ClusterTypeCluster_KPHYSICAL
    }
}
