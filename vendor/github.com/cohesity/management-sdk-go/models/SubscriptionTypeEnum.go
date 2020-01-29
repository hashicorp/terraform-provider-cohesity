// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SubscriptionTypeEnum enum
 */
type SubscriptionTypeEnum int

/**
 * Value collection for SubscriptionTypeEnum enum
 */
const (
    SubscriptionType_KAZURECOMMERCIAL SubscriptionTypeEnum = 1 + iota
    SubscriptionType_KAZUREGOVCLOUD
)

func (r SubscriptionTypeEnum) MarshalJSON() ([]byte, error) {
    s := SubscriptionTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *SubscriptionTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SubscriptionTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SubscriptionTypeEnum to its string representation
 */
func SubscriptionTypeEnumToValue(subscriptionTypeEnum SubscriptionTypeEnum) string {
    switch subscriptionTypeEnum {
        case SubscriptionType_KAZURECOMMERCIAL:
    		return "kAzureCommercial"
        case SubscriptionType_KAZUREGOVCLOUD:
    		return "kAzureGovCloud"
        default:
        	return "kAzureCommercial"
    }
}

/**
 * Converts SubscriptionTypeEnum Array to its string Array representation
*/
func SubscriptionTypeEnumArrayToValue(subscriptionTypeEnum []SubscriptionTypeEnum) []string {
    convArray := make([]string,len( subscriptionTypeEnum))
    for i:=0; i<len(subscriptionTypeEnum);i++ {
        convArray[i] = SubscriptionTypeEnumToValue(subscriptionTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SubscriptionTypeEnumFromValue(value string) SubscriptionTypeEnum {
    switch value {
        case "kAzureCommercial":
            return SubscriptionType_KAZURECOMMERCIAL
        case "kAzureGovCloud":
            return SubscriptionType_KAZUREGOVCLOUD
        default:
            return SubscriptionType_KAZURECOMMERCIAL
    }
}
