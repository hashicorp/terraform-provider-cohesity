// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for GlacierRetrievalTypeEnum enum
 */
type GlacierRetrievalTypeEnum int

/**
 * Value collection for GlacierRetrievalTypeEnum enum
 */
const (
    GlacierRetrievalType_KSTANDARD GlacierRetrievalTypeEnum = 1 + iota
    GlacierRetrievalType_KBULK
    GlacierRetrievalType_KEXPEDITED
)

func (r GlacierRetrievalTypeEnum) MarshalJSON() ([]byte, error) {
    s := GlacierRetrievalTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *GlacierRetrievalTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  GlacierRetrievalTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts GlacierRetrievalTypeEnum to its string representation
 */
func GlacierRetrievalTypeEnumToValue(glacierRetrievalTypeEnum GlacierRetrievalTypeEnum) string {
    switch glacierRetrievalTypeEnum {
        case GlacierRetrievalType_KSTANDARD:
    		return "kStandard"
        case GlacierRetrievalType_KBULK:
    		return "kBulk"
        case GlacierRetrievalType_KEXPEDITED:
    		return "kExpedited"
        default:
        	return "kStandard"
    }
}

/**
 * Converts GlacierRetrievalTypeEnum Array to its string Array representation
*/
func GlacierRetrievalTypeEnumArrayToValue(glacierRetrievalTypeEnum []GlacierRetrievalTypeEnum) []string {
    convArray := make([]string,len( glacierRetrievalTypeEnum))
    for i:=0; i<len(glacierRetrievalTypeEnum);i++ {
        convArray[i] = GlacierRetrievalTypeEnumToValue(glacierRetrievalTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func GlacierRetrievalTypeEnumFromValue(value string) GlacierRetrievalTypeEnum {
    switch value {
        case "kStandard":
            return GlacierRetrievalType_KSTANDARD
        case "kBulk":
            return GlacierRetrievalType_KBULK
        case "kExpedited":
            return GlacierRetrievalType_KEXPEDITED
        default:
            return GlacierRetrievalType_KSTANDARD
    }
}
