// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AlertTypeBucketEnum enum
 */
type AlertTypeBucketEnum int

/**
 * Value collection for AlertTypeBucketEnum enum
 */
const (
    AlertTypeBucket_KSOFTWARE AlertTypeBucketEnum = 1 + iota
    AlertTypeBucket_KHARDWARE
    AlertTypeBucket_KSERVICE
    AlertTypeBucket_KOTHER
)

func (r AlertTypeBucketEnum) MarshalJSON() ([]byte, error) {
    s := AlertTypeBucketEnumToValue(r)
    return json.Marshal(s)
}

func (r *AlertTypeBucketEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AlertTypeBucketEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AlertTypeBucketEnum to its string representation
 */
func AlertTypeBucketEnumToValue(alertTypeBucketEnum AlertTypeBucketEnum) string {
    switch alertTypeBucketEnum {
        case AlertTypeBucket_KSOFTWARE:
    		return "kSoftware"
        case AlertTypeBucket_KHARDWARE:
    		return "kHardware"
        case AlertTypeBucket_KSERVICE:
    		return "kService"
        case AlertTypeBucket_KOTHER:
    		return "kOther"
        default:
        	return "kSoftware"
    }
}

/**
 * Converts AlertTypeBucketEnum Array to its string Array representation
*/
func AlertTypeBucketEnumArrayToValue(alertTypeBucketEnum []AlertTypeBucketEnum) []string {
    convArray := make([]string,len( alertTypeBucketEnum))
    for i:=0; i<len(alertTypeBucketEnum);i++ {
        convArray[i] = AlertTypeBucketEnumToValue(alertTypeBucketEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlertTypeBucketEnumFromValue(value string) AlertTypeBucketEnum {
    switch value {
        case "kSoftware":
            return AlertTypeBucket_KSOFTWARE
        case "kHardware":
            return AlertTypeBucket_KHARDWARE
        case "kService":
            return AlertTypeBucket_KSERVICE
        case "kOther":
            return AlertTypeBucket_KOTHER
        default:
            return AlertTypeBucket_KSOFTWARE
    }
}
