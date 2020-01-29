// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ConsumerTypeGetTenantStatsEnum enum
 */
type ConsumerTypeGetTenantStatsEnum int

/**
 * Value collection for ConsumerTypeGetTenantStatsEnum enum
 */
const (
    ConsumerTypeGetTenantStats_KVIEWS ConsumerTypeGetTenantStatsEnum = 1 + iota
    ConsumerTypeGetTenantStats_KPROTECTIONRUNS
    ConsumerTypeGetTenantStats_KREPLICATIONRUNS
)

func (r ConsumerTypeGetTenantStatsEnum) MarshalJSON() ([]byte, error) {
    s := ConsumerTypeGetTenantStatsEnumToValue(r)
    return json.Marshal(s)
}

func (r *ConsumerTypeGetTenantStatsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ConsumerTypeGetTenantStatsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ConsumerTypeGetTenantStatsEnum to its string representation
 */
func ConsumerTypeGetTenantStatsEnumToValue(consumerTypeGetTenantStatsEnum ConsumerTypeGetTenantStatsEnum) string {
    switch consumerTypeGetTenantStatsEnum {
        case ConsumerTypeGetTenantStats_KVIEWS:
    		return "kViews"
        case ConsumerTypeGetTenantStats_KPROTECTIONRUNS:
    		return "kProtectionRuns"
        case ConsumerTypeGetTenantStats_KREPLICATIONRUNS:
    		return "kReplicationRuns"
        default:
        	return "kViews"
    }
}

/**
 * Converts ConsumerTypeGetTenantStatsEnum Array to its string Array representation
*/
func ConsumerTypeGetTenantStatsEnumArrayToValue(consumerTypeGetTenantStatsEnum []ConsumerTypeGetTenantStatsEnum) []string {
    convArray := make([]string,len( consumerTypeGetTenantStatsEnum))
    for i:=0; i<len(consumerTypeGetTenantStatsEnum);i++ {
        convArray[i] = ConsumerTypeGetTenantStatsEnumToValue(consumerTypeGetTenantStatsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ConsumerTypeGetTenantStatsEnumFromValue(value string) ConsumerTypeGetTenantStatsEnum {
    switch value {
        case "kViews":
            return ConsumerTypeGetTenantStats_KVIEWS
        case "kProtectionRuns":
            return ConsumerTypeGetTenantStats_KPROTECTIONRUNS
        case "kReplicationRuns":
            return ConsumerTypeGetTenantStats_KREPLICATIONRUNS
        default:
            return ConsumerTypeGetTenantStats_KVIEWS
    }
}
