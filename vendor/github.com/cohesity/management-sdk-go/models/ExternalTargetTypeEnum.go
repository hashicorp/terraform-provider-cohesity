// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for ExternalTargetTypeEnum enum
 */
type ExternalTargetTypeEnum int

/**
 * Value collection for ExternalTargetTypeEnum enum
 */
const (
    ExternalTargetType_KNEARLINE ExternalTargetTypeEnum = 1 + iota
    ExternalTargetType_KGLACIER
    ExternalTargetType_KS3
    ExternalTargetType_KAZURESTANDARD
    ExternalTargetType_KS3COMPATIBLE
    ExternalTargetType_KQSTARTAPE
    ExternalTargetType_KGOOGLESTANDARD
    ExternalTargetType_KGOOGLEDRA
    ExternalTargetType_KAMAZONS3STANDARDIA
    ExternalTargetType_KAWSGOVCLOUD
    ExternalTargetType_KNAS
    ExternalTargetType_KCOLDLINE
    ExternalTargetType_KAZUREGOVCLOUD
    ExternalTargetType_KAZUREARCHIVE
    ExternalTargetType_KAZURE
    ExternalTargetType_KGOOGLE
    ExternalTargetType_KAMAZON
    ExternalTargetType_KORACLE
    ExternalTargetType_KORACLETIERSTANDARD
    ExternalTargetType_KORACLETIERARCHIVE
    ExternalTargetType_KAMAZONC2S
)

func (r ExternalTargetTypeEnum) MarshalJSON() ([]byte, error) {
    s := ExternalTargetTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *ExternalTargetTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  ExternalTargetTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts ExternalTargetTypeEnum to its string representation
 */
func ExternalTargetTypeEnumToValue(externalTargetTypeEnum ExternalTargetTypeEnum) string {
    switch externalTargetTypeEnum {
        case ExternalTargetType_KNEARLINE:
    		return "kNearline"
        case ExternalTargetType_KGLACIER:
    		return "kGlacier"
        case ExternalTargetType_KS3:
    		return "kS3"
        case ExternalTargetType_KAZURESTANDARD:
    		return "kAzureStandard"
        case ExternalTargetType_KS3COMPATIBLE:
    		return "kS3Compatible"
        case ExternalTargetType_KQSTARTAPE:
    		return "kQStarTape"
        case ExternalTargetType_KGOOGLESTANDARD:
    		return "kGoogleStandard"
        case ExternalTargetType_KGOOGLEDRA:
    		return "kGoogleDRA"
        case ExternalTargetType_KAMAZONS3STANDARDIA:
    		return "kAmazonS3StandardIA"
        case ExternalTargetType_KAWSGOVCLOUD:
    		return "kAWSGovCloud"
        case ExternalTargetType_KNAS:
    		return "kNAS"
        case ExternalTargetType_KCOLDLINE:
    		return "kColdline"
        case ExternalTargetType_KAZUREGOVCLOUD:
    		return "kAzureGovCloud"
        case ExternalTargetType_KAZUREARCHIVE:
    		return "kAzureArchive"
        case ExternalTargetType_KAZURE:
    		return "kAzure"
        case ExternalTargetType_KGOOGLE:
    		return "kGoogle"
        case ExternalTargetType_KAMAZON:
    		return "kAmazon"
        case ExternalTargetType_KORACLE:
    		return "kOracle"
        case ExternalTargetType_KORACLETIERSTANDARD:
    		return "kOracleTierStandard"
        case ExternalTargetType_KORACLETIERARCHIVE:
    		return "kOracleTierArchive"
        case ExternalTargetType_KAMAZONC2S:
    		return "kAmazonC2S"
        default:
        	return "kNearline"
    }
}

/**
 * Converts ExternalTargetTypeEnum Array to its string Array representation
*/
func ExternalTargetTypeEnumArrayToValue(externalTargetTypeEnum []ExternalTargetTypeEnum) []string {
    convArray := make([]string,len( externalTargetTypeEnum))
    for i:=0; i<len(externalTargetTypeEnum);i++ {
        convArray[i] = ExternalTargetTypeEnumToValue(externalTargetTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ExternalTargetTypeEnumFromValue(value string) ExternalTargetTypeEnum {
    switch value {
        case "kNearline":
            return ExternalTargetType_KNEARLINE
        case "kGlacier":
            return ExternalTargetType_KGLACIER
        case "kS3":
            return ExternalTargetType_KS3
        case "kAzureStandard":
            return ExternalTargetType_KAZURESTANDARD
        case "kS3Compatible":
            return ExternalTargetType_KS3COMPATIBLE
        case "kQStarTape":
            return ExternalTargetType_KQSTARTAPE
        case "kGoogleStandard":
            return ExternalTargetType_KGOOGLESTANDARD
        case "kGoogleDRA":
            return ExternalTargetType_KGOOGLEDRA
        case "kAmazonS3StandardIA":
            return ExternalTargetType_KAMAZONS3STANDARDIA
        case "kAWSGovCloud":
            return ExternalTargetType_KAWSGOVCLOUD
        case "kNAS":
            return ExternalTargetType_KNAS
        case "kColdline":
            return ExternalTargetType_KCOLDLINE
        case "kAzureGovCloud":
            return ExternalTargetType_KAZUREGOVCLOUD
        case "kAzureArchive":
            return ExternalTargetType_KAZUREARCHIVE
        case "kAzure":
            return ExternalTargetType_KAZURE
        case "kGoogle":
            return ExternalTargetType_KGOOGLE
        case "kAmazon":
            return ExternalTargetType_KAMAZON
        case "kOracle":
            return ExternalTargetType_KORACLE
        case "kOracleTierStandard":
            return ExternalTargetType_KORACLETIERSTANDARD
        case "kOracleTierArchive":
            return ExternalTargetType_KORACLETIERARCHIVE
        case "kAmazonC2S":
            return ExternalTargetType_KAMAZONC2S
        default:
            return ExternalTargetType_KNEARLINE
    }
}
