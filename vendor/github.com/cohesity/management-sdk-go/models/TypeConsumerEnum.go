// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeConsumerEnum enum
 */
type TypeConsumerEnum int

/**
 * Value collection for TypeConsumerEnum enum
 */
const (
    TypeConsumer_KVIEWS TypeConsumerEnum = 1 + iota
    TypeConsumer_KPROTECTIONRUNS
    TypeConsumer_KREPLICATIONRUNS
)

func (r TypeConsumerEnum) MarshalJSON() ([]byte, error) {
    s := TypeConsumerEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeConsumerEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeConsumerEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeConsumerEnum to its string representation
 */
func TypeConsumerEnumToValue(typeConsumerEnum TypeConsumerEnum) string {
    switch typeConsumerEnum {
        case TypeConsumer_KVIEWS:
    		return "kViews"
        case TypeConsumer_KPROTECTIONRUNS:
    		return "kProtectionRuns"
        case TypeConsumer_KREPLICATIONRUNS:
    		return "kReplicationRuns"
        default:
        	return "kViews"
    }
}

/**
 * Converts TypeConsumerEnum Array to its string Array representation
*/
func TypeConsumerEnumArrayToValue(typeConsumerEnum []TypeConsumerEnum) []string {
    convArray := make([]string,len( typeConsumerEnum))
    for i:=0; i<len(typeConsumerEnum);i++ {
        convArray[i] = TypeConsumerEnumToValue(typeConsumerEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeConsumerEnumFromValue(value string) TypeConsumerEnum {
    switch value {
        case "kViews":
            return TypeConsumer_KVIEWS
        case "kProtectionRuns":
            return TypeConsumer_KPROTECTIONRUNS
        case "kReplicationRuns":
            return TypeConsumer_KREPLICATIONRUNS
        default:
            return TypeConsumer_KVIEWS
    }
}
