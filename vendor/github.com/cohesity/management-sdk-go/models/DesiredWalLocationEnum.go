// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for DesiredWalLocationEnum enum
 */
type DesiredWalLocationEnum int

/**
 * Value collection for DesiredWalLocationEnum enum
 */
const (
    DesiredWalLocation_KHOMEPARTITION DesiredWalLocationEnum = 1 + iota
    DesiredWalLocation_KDISK
    DesiredWalLocation_KSCRIBE
    DesiredWalLocation_KSCRIBETABLE
)

func (r DesiredWalLocationEnum) MarshalJSON() ([]byte, error) {
    s := DesiredWalLocationEnumToValue(r)
    return json.Marshal(s)
}

func (r *DesiredWalLocationEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  DesiredWalLocationEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts DesiredWalLocationEnum to its string representation
 */
func DesiredWalLocationEnumToValue(desiredWalLocationEnum DesiredWalLocationEnum) string {
    switch desiredWalLocationEnum {
        case DesiredWalLocation_KHOMEPARTITION:
    		return "kHomePartition"
        case DesiredWalLocation_KDISK:
    		return "kDisk"
        case DesiredWalLocation_KSCRIBE:
    		return "kScribe"
        case DesiredWalLocation_KSCRIBETABLE:
    		return "kScribeTable"
        default:
        	return "kHomePartition"
    }
}

/**
 * Converts DesiredWalLocationEnum Array to its string Array representation
*/
func DesiredWalLocationEnumArrayToValue(desiredWalLocationEnum []DesiredWalLocationEnum) []string {
    convArray := make([]string,len( desiredWalLocationEnum))
    for i:=0; i<len(desiredWalLocationEnum);i++ {
        convArray[i] = DesiredWalLocationEnumToValue(desiredWalLocationEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DesiredWalLocationEnumFromValue(value string) DesiredWalLocationEnum {
    switch value {
        case "kHomePartition":
            return DesiredWalLocation_KHOMEPARTITION
        case "kDisk":
            return DesiredWalLocation_KDISK
        case "kScribe":
            return DesiredWalLocation_KSCRIBE
        case "kScribeTable":
            return DesiredWalLocation_KSCRIBETABLE
        default:
            return DesiredWalLocation_KHOMEPARTITION
    }
}
