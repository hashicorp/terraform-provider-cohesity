// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EntityTypeGetFileDistributionStatsEnum enum
 */
type EntityTypeGetFileDistributionStatsEnum int

/**
 * Value collection for EntityTypeGetFileDistributionStatsEnum enum
 */
const (
    EntityTypeGetFileDistributionStats_KCLUSTER EntityTypeGetFileDistributionStatsEnum = 1 + iota
    EntityTypeGetFileDistributionStats_KSTORAGEDOMAIN
)

func (r EntityTypeGetFileDistributionStatsEnum) MarshalJSON() ([]byte, error) {
    s := EntityTypeGetFileDistributionStatsEnumToValue(r)
    return json.Marshal(s)
}

func (r *EntityTypeGetFileDistributionStatsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EntityTypeGetFileDistributionStatsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EntityTypeGetFileDistributionStatsEnum to its string representation
 */
func EntityTypeGetFileDistributionStatsEnumToValue(entityTypeGetFileDistributionStatsEnum EntityTypeGetFileDistributionStatsEnum) string {
    switch entityTypeGetFileDistributionStatsEnum {
        case EntityTypeGetFileDistributionStats_KCLUSTER:
    		return "kCluster"
        case EntityTypeGetFileDistributionStats_KSTORAGEDOMAIN:
    		return "kStorageDomain"
        default:
        	return "kCluster"
    }
}

/**
 * Converts EntityTypeGetFileDistributionStatsEnum Array to its string Array representation
*/
func EntityTypeGetFileDistributionStatsEnumArrayToValue(entityTypeGetFileDistributionStatsEnum []EntityTypeGetFileDistributionStatsEnum) []string {
    convArray := make([]string,len( entityTypeGetFileDistributionStatsEnum))
    for i:=0; i<len(entityTypeGetFileDistributionStatsEnum);i++ {
        convArray[i] = EntityTypeGetFileDistributionStatsEnumToValue(entityTypeGetFileDistributionStatsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EntityTypeGetFileDistributionStatsEnumFromValue(value string) EntityTypeGetFileDistributionStatsEnum {
    switch value {
        case "kCluster":
            return EntityTypeGetFileDistributionStats_KCLUSTER
        case "kStorageDomain":
            return EntityTypeGetFileDistributionStats_KSTORAGEDOMAIN
        default:
            return EntityTypeGetFileDistributionStats_KCLUSTER
    }
}
