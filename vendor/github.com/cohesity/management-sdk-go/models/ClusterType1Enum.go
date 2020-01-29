package models

import(
    "encoding/json"
)

/**
 * Type definition for ClusterType1Enum enum
 */
type ClusterType1Enum int

/**
 * Value collection for ClusterType1Enum enum
 */
const (
    ClusterType1_KPHYSICAL ClusterType1Enum = 1 + iota
    ClusterType1_KVIRTUALROBO
    ClusterType1_KMICROSOFTCLOUD
    ClusterType1_KAMAZONCLOUD
    ClusterType1_KGOOGLECLOUD
)

func (r ClusterType1Enum) MarshalJSON() ([]byte, error) { 
    s := ClusterType1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *ClusterType1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  ClusterType1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ClusterType1Enum to its string representation
 */
func ClusterType1EnumToValue(clusterType1Enum ClusterType1Enum) string {
    switch clusterType1Enum {
        case ClusterType1_KPHYSICAL:
    		return "kPhysical"		
        case ClusterType1_KVIRTUALROBO:
    		return "kVirtualRobo"		
        case ClusterType1_KMICROSOFTCLOUD:
    		return "kMicrosoftCloud"		
        case ClusterType1_KAMAZONCLOUD:
    		return "kAmazonCloud"		
        case ClusterType1_KGOOGLECLOUD:
    		return "kGoogleCloud"		
        default:
        	return "kPhysical"
    }
}

/**
 * Converts ClusterType1Enum Array to its string Array representation
*/
func ClusterType1EnumArrayToValue(clusterType1Enum []ClusterType1Enum) []string {
    convArray := make([]string,len( clusterType1Enum))
    for i:=0; i<len(clusterType1Enum);i++ {
        convArray[i] = ClusterType1EnumToValue(clusterType1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ClusterType1EnumFromValue(value string) ClusterType1Enum {
    switch value {
        case "kPhysical":
            return ClusterType1_KPHYSICAL
        case "kVirtualRobo":
            return ClusterType1_KVIRTUALROBO
        case "kMicrosoftCloud":
            return ClusterType1_KMICROSOFTCLOUD
        case "kAmazonCloud":
            return ClusterType1_KAMAZONCLOUD
        case "kGoogleCloud":
            return ClusterType1_KGOOGLECLOUD
        default:
            return ClusterType1_KPHYSICAL
    }
}
