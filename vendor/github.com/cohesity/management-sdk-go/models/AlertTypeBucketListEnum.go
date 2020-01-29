// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for AlertTypeBucketListEnum enum
 */
type AlertTypeBucketListEnum int

/**
 * Value collection for AlertTypeBucketListEnum enum
 */
const (
    AlertTypeBucketList_KSOFTWARE AlertTypeBucketListEnum = 1 + iota
    AlertTypeBucketList_KHARDWARE
    AlertTypeBucketList_KSERVICE
    AlertTypeBucketList_KOTHER
)

func (r AlertTypeBucketListEnum) MarshalJSON() ([]byte, error) {
    s := AlertTypeBucketListEnumToValue(r)
    return json.Marshal(s)
}

func (r *AlertTypeBucketListEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  AlertTypeBucketListEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts AlertTypeBucketListEnum to its string representation
 */
func AlertTypeBucketListEnumToValue(alertTypeBucketListEnum AlertTypeBucketListEnum) string {
    switch alertTypeBucketListEnum {
        case AlertTypeBucketList_KSOFTWARE:
    		return "kSoftware"
        case AlertTypeBucketList_KHARDWARE:
    		return "kHardware"
        case AlertTypeBucketList_KSERVICE:
    		return "kService"
        case AlertTypeBucketList_KOTHER:
    		return "kOther"
        default:
        	return "kSoftware"
    }
}

/**
 * Converts AlertTypeBucketListEnum Array to its string Array representation
*/
func AlertTypeBucketListEnumArrayToValue(alertTypeBucketListEnum []AlertTypeBucketListEnum) []string {
    convArray := make([]string,len( alertTypeBucketListEnum))
    for i:=0; i<len(alertTypeBucketListEnum);i++ {
        convArray[i] = AlertTypeBucketListEnumToValue(alertTypeBucketListEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func AlertTypeBucketListEnumFromValue(value string) AlertTypeBucketListEnum {
    switch value {
        case "kSoftware":
            return AlertTypeBucketList_KSOFTWARE
        case "kHardware":
            return AlertTypeBucketList_KHARDWARE
        case "kService":
            return AlertTypeBucketList_KSERVICE
        case "kOther":
            return AlertTypeBucketList_KOTHER
        default:
            return AlertTypeBucketList_KSOFTWARE
    }
}
