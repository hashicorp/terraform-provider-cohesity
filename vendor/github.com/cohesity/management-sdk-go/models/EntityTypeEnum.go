// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for EntityTypeEnum enum
 */
type EntityTypeEnum int

/**
 * Value collection for EntityTypeEnum enum
 */
const (
    EntityType_CLUSTER EntityTypeEnum = 1 + iota
    EntityType_STORAGEDOMAIN
)

func (r EntityTypeEnum) MarshalJSON() ([]byte, error) {
    s := EntityTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *EntityTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  EntityTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts EntityTypeEnum to its string representation
 */
func EntityTypeEnumToValue(entityTypeEnum EntityTypeEnum) string {
    switch entityTypeEnum {
        case EntityType_CLUSTER:
    		return "Cluster"
        case EntityType_STORAGEDOMAIN:
    		return "StorageDomain"
        default:
        	return "Cluster"
    }
}

/**
 * Converts EntityTypeEnum Array to its string Array representation
*/
func EntityTypeEnumArrayToValue(entityTypeEnum []EntityTypeEnum) []string {
    convArray := make([]string,len( entityTypeEnum))
    for i:=0; i<len(entityTypeEnum);i++ {
        convArray[i] = EntityTypeEnumToValue(entityTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func EntityTypeEnumFromValue(value string) EntityTypeEnum {
    switch value {
        case "Cluster":
            return EntityType_CLUSTER
        case "StorageDomain":
            return EntityType_STORAGEDOMAIN
        default:
            return EntityType_CLUSTER
    }
}
