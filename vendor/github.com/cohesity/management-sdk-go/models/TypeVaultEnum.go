// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for TypeVaultEnum enum
 */
type TypeVaultEnum int

/**
 * Value collection for TypeVaultEnum enum
 */
const (
    TypeVault_KNEARLINE TypeVaultEnum = 1 + iota
    TypeVault_KGLACIER
    TypeVault_KS3
    TypeVault_KAZURESTANDARD
    TypeVault_KS3COMPATIBLE
    TypeVault_KQSTARTAPE
    TypeVault_KGOOGLESTANDARD
    TypeVault_KGOOGLEDRA
    TypeVault_KAMAZONS3STANDARDIA
    TypeVault_KAWSGOVCLOUD
    TypeVault_KNAS
    TypeVault_KCOLDLINE
    TypeVault_KAZUREGOVCLOUD
    TypeVault_KAZUREARCHIVE
    TypeVault_KAZURE
    TypeVault_KGOOGLE
    TypeVault_KAMAZON
    TypeVault_KORACLE
    TypeVault_KORACLETIERSTANDARD
    TypeVault_KORACLETIERARCHIVE
    TypeVault_KAMAZONC2S
)

func (r TypeVaultEnum) MarshalJSON() ([]byte, error) {
    s := TypeVaultEnumToValue(r)
    return json.Marshal(s)
}

func (r *TypeVaultEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  TypeVaultEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts TypeVaultEnum to its string representation
 */
func TypeVaultEnumToValue(typeVaultEnum TypeVaultEnum) string {
    switch typeVaultEnum {
        case TypeVault_KNEARLINE:
    		return "kNearline"
        case TypeVault_KGLACIER:
    		return "kGlacier"
        case TypeVault_KS3:
    		return "kS3"
        case TypeVault_KAZURESTANDARD:
    		return "kAzureStandard"
        case TypeVault_KS3COMPATIBLE:
    		return "kS3Compatible"
        case TypeVault_KQSTARTAPE:
    		return "kQStarTape"
        case TypeVault_KGOOGLESTANDARD:
    		return "kGoogleStandard"
        case TypeVault_KGOOGLEDRA:
    		return "kGoogleDRA"
        case TypeVault_KAMAZONS3STANDARDIA:
    		return "kAmazonS3StandardIA"
        case TypeVault_KAWSGOVCLOUD:
    		return "kAWSGovCloud"
        case TypeVault_KNAS:
    		return "kNAS"
        case TypeVault_KCOLDLINE:
    		return "kColdline"
        case TypeVault_KAZUREGOVCLOUD:
    		return "kAzureGovCloud"
        case TypeVault_KAZUREARCHIVE:
    		return "kAzureArchive"
        case TypeVault_KAZURE:
    		return "kAzure"
        case TypeVault_KGOOGLE:
    		return "kGoogle"
        case TypeVault_KAMAZON:
    		return "kAmazon"
        case TypeVault_KORACLE:
    		return "kOracle"
        case TypeVault_KORACLETIERSTANDARD:
    		return "kOracleTierStandard"
        case TypeVault_KORACLETIERARCHIVE:
    		return "kOracleTierArchive"
        case TypeVault_KAMAZONC2S:
    		return "kAmazonC2S"
        default:
        	return "kNearline"
    }
}

/**
 * Converts TypeVaultEnum Array to its string Array representation
*/
func TypeVaultEnumArrayToValue(typeVaultEnum []TypeVaultEnum) []string {
    convArray := make([]string,len( typeVaultEnum))
    for i:=0; i<len(typeVaultEnum);i++ {
        convArray[i] = TypeVaultEnumToValue(typeVaultEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeVaultEnumFromValue(value string) TypeVaultEnum {
    switch value {
        case "kNearline":
            return TypeVault_KNEARLINE
        case "kGlacier":
            return TypeVault_KGLACIER
        case "kS3":
            return TypeVault_KS3
        case "kAzureStandard":
            return TypeVault_KAZURESTANDARD
        case "kS3Compatible":
            return TypeVault_KS3COMPATIBLE
        case "kQStarTape":
            return TypeVault_KQSTARTAPE
        case "kGoogleStandard":
            return TypeVault_KGOOGLESTANDARD
        case "kGoogleDRA":
            return TypeVault_KGOOGLEDRA
        case "kAmazonS3StandardIA":
            return TypeVault_KAMAZONS3STANDARDIA
        case "kAWSGovCloud":
            return TypeVault_KAWSGOVCLOUD
        case "kNAS":
            return TypeVault_KNAS
        case "kColdline":
            return TypeVault_KCOLDLINE
        case "kAzureGovCloud":
            return TypeVault_KAZUREGOVCLOUD
        case "kAzureArchive":
            return TypeVault_KAZUREARCHIVE
        case "kAzure":
            return TypeVault_KAZURE
        case "kGoogle":
            return TypeVault_KGOOGLE
        case "kAmazon":
            return TypeVault_KAMAZON
        case "kOracle":
            return TypeVault_KORACLE
        case "kOracleTierStandard":
            return TypeVault_KORACLETIERSTANDARD
        case "kOracleTierArchive":
            return TypeVault_KORACLETIERARCHIVE
        case "kAmazonC2S":
            return TypeVault_KAMAZONC2S
        default:
            return TypeVault_KNEARLINE
    }
}
