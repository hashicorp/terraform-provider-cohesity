// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TierTypeEnum enum
 */
type TierTypeEnum int

/**
 * Value collection for TierTypeEnum enum
 */
const (
    TierType_KAMAZONS3STANDARD TierTypeEnum = 1 + iota
    TierType_KAMAZONS3STANDARDIA
    TierType_KAMAZONGLACIER
    TierType_KAMAZONS3ONEZONEIA
    TierType_KAMAZONS3INTELLIGENTTIERING
    TierType_KAMAZONS3GLACIERDEEPARCHIVE
)

func (r TierTypeEnum) MarshalJSON() ([]byte, error) {
    s := TierTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *TierTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TierTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TierTypeEnum to its string representation
 */
func TierTypeEnumToValue(tierTypeEnum TierTypeEnum) string {
    switch tierTypeEnum {
        case TierType_KAMAZONS3STANDARD:
    		return "kAmazonS3Standard"
        case TierType_KAMAZONS3STANDARDIA:
    		return "kAmazonS3StandardIA"
        case TierType_KAMAZONGLACIER:
    		return "kAmazonGlacier"
        case TierType_KAMAZONS3ONEZONEIA:
    		return "kAmazonS3OneZoneIA"
        case TierType_KAMAZONS3INTELLIGENTTIERING:
    		return "kAmazonS3IntelligentTiering"
        case TierType_KAMAZONS3GLACIERDEEPARCHIVE:
    		return "kAmazonS3GlacierDeepArchive"
        default:
        	return "kAmazonS3Standard"
    }
}

/**
 * Converts TierTypeEnum Array to its string Array representation
*/
func TierTypeEnumArrayToValue(tierTypeEnum []TierTypeEnum) []string {
    convArray := make([]string,len( tierTypeEnum))
    for i:=0; i<len(tierTypeEnum);i++ {
        convArray[i] = TierTypeEnumToValue(tierTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TierTypeEnumFromValue(value string) TierTypeEnum {
    switch value {
        case "kAmazonS3Standard":
            return TierType_KAMAZONS3STANDARD
        case "kAmazonS3StandardIA":
            return TierType_KAMAZONS3STANDARDIA
        case "kAmazonGlacier":
            return TierType_KAMAZONGLACIER
        case "kAmazonS3OneZoneIA":
            return TierType_KAMAZONS3ONEZONEIA
        case "kAmazonS3IntelligentTiering":
            return TierType_KAMAZONS3INTELLIGENTTIERING
        case "kAmazonS3GlacierDeepArchive":
            return TierType_KAMAZONS3GLACIERDEEPARCHIVE
        default:
            return TierType_KAMAZONS3STANDARD
    }
}
