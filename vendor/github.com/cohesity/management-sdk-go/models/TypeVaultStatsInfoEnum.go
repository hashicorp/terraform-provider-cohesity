// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeVaultStatsInfoEnum enum
 */
type TypeVaultStatsInfoEnum int

/**
 * Value collection for TypeVaultStatsInfoEnum enum
 */
const (
    TypeVaultStatsInfo_KNEARLINE TypeVaultStatsInfoEnum = 1 + iota
    TypeVaultStatsInfo_KGLACIER
    TypeVaultStatsInfo_KS3
    TypeVaultStatsInfo_KAZURESTANDARD
    TypeVaultStatsInfo_KS3COMPATIBLE
    TypeVaultStatsInfo_KQSTARTAPE
    TypeVaultStatsInfo_KGOOGLESTANDARD
    TypeVaultStatsInfo_KGOOGLEDRA
    TypeVaultStatsInfo_KAMAZONS3STANDARDIA
    TypeVaultStatsInfo_KAWSGOVCLOUD
    TypeVaultStatsInfo_KNAS
    TypeVaultStatsInfo_KCOLDLINE
    TypeVaultStatsInfo_KAZUREGOVCLOUD
    TypeVaultStatsInfo_KAZUREARCHIVE
    TypeVaultStatsInfo_KAZURE
    TypeVaultStatsInfo_KGOOGLE
    TypeVaultStatsInfo_KAMAZON
    TypeVaultStatsInfo_KORACLE
    TypeVaultStatsInfo_KORACLETIERSTANDARD
    TypeVaultStatsInfo_KORACLETIERARCHIVE
    TypeVaultStatsInfo_KAMAZONC2S
)

func (r TypeVaultStatsInfoEnum) MarshalJSON() ([]byte, error) {
    s := TypeVaultStatsInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeVaultStatsInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeVaultStatsInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeVaultStatsInfoEnum to its string representation
 */
func TypeVaultStatsInfoEnumToValue(typeVaultStatsInfoEnum TypeVaultStatsInfoEnum) string {
    switch typeVaultStatsInfoEnum {
        case TypeVaultStatsInfo_KNEARLINE:
    		return "kNearline"
        case TypeVaultStatsInfo_KGLACIER:
    		return "kGlacier"
        case TypeVaultStatsInfo_KS3:
    		return "kS3"
        case TypeVaultStatsInfo_KAZURESTANDARD:
    		return "kAzureStandard"
        case TypeVaultStatsInfo_KS3COMPATIBLE:
    		return "kS3Compatible"
        case TypeVaultStatsInfo_KQSTARTAPE:
    		return "kQStarTape"
        case TypeVaultStatsInfo_KGOOGLESTANDARD:
    		return "kGoogleStandard"
        case TypeVaultStatsInfo_KGOOGLEDRA:
    		return "kGoogleDRA"
        case TypeVaultStatsInfo_KAMAZONS3STANDARDIA:
    		return "kAmazonS3StandardIA"
        case TypeVaultStatsInfo_KAWSGOVCLOUD:
    		return "kAWSGovCloud"
        case TypeVaultStatsInfo_KNAS:
    		return "kNAS"
        case TypeVaultStatsInfo_KCOLDLINE:
    		return "kColdline"
        case TypeVaultStatsInfo_KAZUREGOVCLOUD:
    		return "kAzureGovCloud"
        case TypeVaultStatsInfo_KAZUREARCHIVE:
    		return "kAzureArchive"
        case TypeVaultStatsInfo_KAZURE:
    		return "kAzure"
        case TypeVaultStatsInfo_KGOOGLE:
    		return "kGoogle"
        case TypeVaultStatsInfo_KAMAZON:
    		return "kAmazon"
        case TypeVaultStatsInfo_KORACLE:
    		return "kOracle"
        case TypeVaultStatsInfo_KORACLETIERSTANDARD:
    		return "kOracleTierStandard"
        case TypeVaultStatsInfo_KORACLETIERARCHIVE:
    		return "kOracleTierArchive"
        case TypeVaultStatsInfo_KAMAZONC2S:
    		return "kAmazonC2S"
        default:
        	return "kNearline"
    }
}

/**
 * Converts TypeVaultStatsInfoEnum Array to its string Array representation
*/
func TypeVaultStatsInfoEnumArrayToValue(typeVaultStatsInfoEnum []TypeVaultStatsInfoEnum) []string {
    convArray := make([]string,len( typeVaultStatsInfoEnum))
    for i:=0; i<len(typeVaultStatsInfoEnum);i++ {
        convArray[i] = TypeVaultStatsInfoEnumToValue(typeVaultStatsInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeVaultStatsInfoEnumFromValue(value string) TypeVaultStatsInfoEnum {
    switch value {
        case "kNearline":
            return TypeVaultStatsInfo_KNEARLINE
        case "kGlacier":
            return TypeVaultStatsInfo_KGLACIER
        case "kS3":
            return TypeVaultStatsInfo_KS3
        case "kAzureStandard":
            return TypeVaultStatsInfo_KAZURESTANDARD
        case "kS3Compatible":
            return TypeVaultStatsInfo_KS3COMPATIBLE
        case "kQStarTape":
            return TypeVaultStatsInfo_KQSTARTAPE
        case "kGoogleStandard":
            return TypeVaultStatsInfo_KGOOGLESTANDARD
        case "kGoogleDRA":
            return TypeVaultStatsInfo_KGOOGLEDRA
        case "kAmazonS3StandardIA":
            return TypeVaultStatsInfo_KAMAZONS3STANDARDIA
        case "kAWSGovCloud":
            return TypeVaultStatsInfo_KAWSGOVCLOUD
        case "kNAS":
            return TypeVaultStatsInfo_KNAS
        case "kColdline":
            return TypeVaultStatsInfo_KCOLDLINE
        case "kAzureGovCloud":
            return TypeVaultStatsInfo_KAZUREGOVCLOUD
        case "kAzureArchive":
            return TypeVaultStatsInfo_KAZUREARCHIVE
        case "kAzure":
            return TypeVaultStatsInfo_KAZURE
        case "kGoogle":
            return TypeVaultStatsInfo_KGOOGLE
        case "kAmazon":
            return TypeVaultStatsInfo_KAMAZON
        case "kOracle":
            return TypeVaultStatsInfo_KORACLE
        case "kOracleTierStandard":
            return TypeVaultStatsInfo_KORACLETIERSTANDARD
        case "kOracleTierArchive":
            return TypeVaultStatsInfo_KORACLETIERARCHIVE
        case "kAmazonC2S":
            return TypeVaultStatsInfo_KAMAZONC2S
        default:
            return TypeVaultStatsInfo_KNEARLINE
    }
}
