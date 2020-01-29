// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ConsumerTypeGetConsumerStatsEnum enum
 */
type ConsumerTypeGetConsumerStatsEnum int

/**
 * Value collection for ConsumerTypeGetConsumerStatsEnum enum
 */
const (
    ConsumerTypeGetConsumerStats_KVIEWS ConsumerTypeGetConsumerStatsEnum = 1 + iota
    ConsumerTypeGetConsumerStats_KPROTECTIONRUNS
    ConsumerTypeGetConsumerStats_KREPLICATIONRUNS
)

func (r ConsumerTypeGetConsumerStatsEnum) MarshalJSON() ([]byte, error) {
    s := ConsumerTypeGetConsumerStatsEnumToValue(r)
    return json.Marshal(s)
}

func (r *ConsumerTypeGetConsumerStatsEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ConsumerTypeGetConsumerStatsEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ConsumerTypeGetConsumerStatsEnum to its string representation
 */
func ConsumerTypeGetConsumerStatsEnumToValue(consumerTypeGetConsumerStatsEnum ConsumerTypeGetConsumerStatsEnum) string {
    switch consumerTypeGetConsumerStatsEnum {
        case ConsumerTypeGetConsumerStats_KVIEWS:
    		return "kViews"
        case ConsumerTypeGetConsumerStats_KPROTECTIONRUNS:
    		return "kProtectionRuns"
        case ConsumerTypeGetConsumerStats_KREPLICATIONRUNS:
    		return "kReplicationRuns"
        default:
        	return "kViews"
    }
}

/**
 * Converts ConsumerTypeGetConsumerStatsEnum Array to its string Array representation
*/
func ConsumerTypeGetConsumerStatsEnumArrayToValue(consumerTypeGetConsumerStatsEnum []ConsumerTypeGetConsumerStatsEnum) []string {
    convArray := make([]string,len( consumerTypeGetConsumerStatsEnum))
    for i:=0; i<len(consumerTypeGetConsumerStatsEnum);i++ {
        convArray[i] = ConsumerTypeGetConsumerStatsEnumToValue(consumerTypeGetConsumerStatsEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ConsumerTypeGetConsumerStatsEnumFromValue(value string) ConsumerTypeGetConsumerStatsEnum {
    switch value {
        case "kViews":
            return ConsumerTypeGetConsumerStats_KVIEWS
        case "kProtectionRuns":
            return ConsumerTypeGetConsumerStats_KPROTECTIONRUNS
        case "kReplicationRuns":
            return ConsumerTypeGetConsumerStats_KREPLICATIONRUNS
        default:
            return ConsumerTypeGetConsumerStats_KVIEWS
    }
}
