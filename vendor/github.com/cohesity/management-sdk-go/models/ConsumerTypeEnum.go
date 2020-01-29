// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ConsumerTypeEnum enum
 */
type ConsumerTypeEnum int

/**
 * Value collection for ConsumerTypeEnum enum
 */
const (
    ConsumerType_KVIEWS ConsumerTypeEnum = 1 + iota
    ConsumerType_KPROTECTIONRUNS
    ConsumerType_KREPLICATIONRUNS
)

func (r ConsumerTypeEnum) MarshalJSON() ([]byte, error) {
    s := ConsumerTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ConsumerTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ConsumerTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ConsumerTypeEnum to its string representation
 */
func ConsumerTypeEnumToValue(consumerTypeEnum ConsumerTypeEnum) string {
    switch consumerTypeEnum {
        case ConsumerType_KVIEWS:
    		return "kViews"
        case ConsumerType_KPROTECTIONRUNS:
    		return "kProtectionRuns"
        case ConsumerType_KREPLICATIONRUNS:
    		return "kReplicationRuns"
        default:
        	return "kViews"
    }
}

/**
 * Converts ConsumerTypeEnum Array to its string Array representation
*/
func ConsumerTypeEnumArrayToValue(consumerTypeEnum []ConsumerTypeEnum) []string {
    convArray := make([]string,len( consumerTypeEnum))
    for i:=0; i<len(consumerTypeEnum);i++ {
        convArray[i] = ConsumerTypeEnumToValue(consumerTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ConsumerTypeEnumFromValue(value string) ConsumerTypeEnum {
    switch value {
        case "kViews":
            return ConsumerType_KVIEWS
        case "kProtectionRuns":
            return ConsumerType_KPROTECTIONRUNS
        case "kReplicationRuns":
            return ConsumerType_KREPLICATIONRUNS
        default:
            return ConsumerType_KVIEWS
    }
}
