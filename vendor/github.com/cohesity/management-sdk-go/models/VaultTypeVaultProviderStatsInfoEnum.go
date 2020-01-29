// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for VaultTypeVaultProviderStatsInfoEnum enum
 */
type VaultTypeVaultProviderStatsInfoEnum int

/**
 * Value collection for VaultTypeVaultProviderStatsInfoEnum enum
 */
const (
    VaultTypeVaultProviderStatsInfo_KNEARLINE VaultTypeVaultProviderStatsInfoEnum = 1 + iota
    VaultTypeVaultProviderStatsInfo_KGLACIER
    VaultTypeVaultProviderStatsInfo_KS3
    VaultTypeVaultProviderStatsInfo_KAZURESTANDARD
    VaultTypeVaultProviderStatsInfo_KS3COMPATIBLE
    VaultTypeVaultProviderStatsInfo_KQSTARTAPE
    VaultTypeVaultProviderStatsInfo_KGOOGLESTANDARD
    VaultTypeVaultProviderStatsInfo_KGOOGLEDRA
    VaultTypeVaultProviderStatsInfo_KAMAZONS3STANDARDIA
    VaultTypeVaultProviderStatsInfo_KAWSGOVCLOUD
    VaultTypeVaultProviderStatsInfo_KNAS
    VaultTypeVaultProviderStatsInfo_KCOLDLINE
    VaultTypeVaultProviderStatsInfo_KAZUREGOVCLOUD
    VaultTypeVaultProviderStatsInfo_KAZUREARCHIVE
    VaultTypeVaultProviderStatsInfo_KAZURE
    VaultTypeVaultProviderStatsInfo_KGOOGLE
    VaultTypeVaultProviderStatsInfo_KAMAZON
    VaultTypeVaultProviderStatsInfo_KORACLE
    VaultTypeVaultProviderStatsInfo_KORACLETIERSTANDARD
    VaultTypeVaultProviderStatsInfo_KORACLETIERARCHIVE
    VaultTypeVaultProviderStatsInfo_KAMAZONC2S
)

func (r VaultTypeVaultProviderStatsInfoEnum) MarshalJSON() ([]byte, error) {
    s := VaultTypeVaultProviderStatsInfoEnumToValue(r)
    return json.Marshal(s)
}

func (r *VaultTypeVaultProviderStatsInfoEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  VaultTypeVaultProviderStatsInfoEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts VaultTypeVaultProviderStatsInfoEnum to its string representation
 */
func VaultTypeVaultProviderStatsInfoEnumToValue(vaultTypeVaultProviderStatsInfoEnum VaultTypeVaultProviderStatsInfoEnum) string {
    switch vaultTypeVaultProviderStatsInfoEnum {
        case VaultTypeVaultProviderStatsInfo_KNEARLINE:
    		return "kNearline"
        case VaultTypeVaultProviderStatsInfo_KGLACIER:
    		return "kGlacier"
        case VaultTypeVaultProviderStatsInfo_KS3:
    		return "kS3"
        case VaultTypeVaultProviderStatsInfo_KAZURESTANDARD:
    		return "kAzureStandard"
        case VaultTypeVaultProviderStatsInfo_KS3COMPATIBLE:
    		return "kS3Compatible"
        case VaultTypeVaultProviderStatsInfo_KQSTARTAPE:
    		return "kQStarTape"
        case VaultTypeVaultProviderStatsInfo_KGOOGLESTANDARD:
    		return "kGoogleStandard"
        case VaultTypeVaultProviderStatsInfo_KGOOGLEDRA:
    		return "kGoogleDRA"
        case VaultTypeVaultProviderStatsInfo_KAMAZONS3STANDARDIA:
    		return "kAmazonS3StandardIA"
        case VaultTypeVaultProviderStatsInfo_KAWSGOVCLOUD:
    		return "kAWSGovCloud"
        case VaultTypeVaultProviderStatsInfo_KNAS:
    		return "kNAS"
        case VaultTypeVaultProviderStatsInfo_KCOLDLINE:
    		return "kColdline"
        case VaultTypeVaultProviderStatsInfo_KAZUREGOVCLOUD:
    		return "kAzureGovCloud"
        case VaultTypeVaultProviderStatsInfo_KAZUREARCHIVE:
    		return "kAzureArchive"
        case VaultTypeVaultProviderStatsInfo_KAZURE:
    		return "kAzure"
        case VaultTypeVaultProviderStatsInfo_KGOOGLE:
    		return "kGoogle"
        case VaultTypeVaultProviderStatsInfo_KAMAZON:
    		return "kAmazon"
        case VaultTypeVaultProviderStatsInfo_KORACLE:
    		return "kOracle"
        case VaultTypeVaultProviderStatsInfo_KORACLETIERSTANDARD:
    		return "kOracleTierStandard"
        case VaultTypeVaultProviderStatsInfo_KORACLETIERARCHIVE:
    		return "kOracleTierArchive"
        case VaultTypeVaultProviderStatsInfo_KAMAZONC2S:
    		return "kAmazonC2S"
        default:
        	return "kNearline"
    }
}

/**
 * Converts VaultTypeVaultProviderStatsInfoEnum Array to its string Array representation
*/
func VaultTypeVaultProviderStatsInfoEnumArrayToValue(vaultTypeVaultProviderStatsInfoEnum []VaultTypeVaultProviderStatsInfoEnum) []string {
    convArray := make([]string,len( vaultTypeVaultProviderStatsInfoEnum))
    for i:=0; i<len(vaultTypeVaultProviderStatsInfoEnum);i++ {
        convArray[i] = VaultTypeVaultProviderStatsInfoEnumToValue(vaultTypeVaultProviderStatsInfoEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func VaultTypeVaultProviderStatsInfoEnumFromValue(value string) VaultTypeVaultProviderStatsInfoEnum {
    switch value {
        case "kNearline":
            return VaultTypeVaultProviderStatsInfo_KNEARLINE
        case "kGlacier":
            return VaultTypeVaultProviderStatsInfo_KGLACIER
        case "kS3":
            return VaultTypeVaultProviderStatsInfo_KS3
        case "kAzureStandard":
            return VaultTypeVaultProviderStatsInfo_KAZURESTANDARD
        case "kS3Compatible":
            return VaultTypeVaultProviderStatsInfo_KS3COMPATIBLE
        case "kQStarTape":
            return VaultTypeVaultProviderStatsInfo_KQSTARTAPE
        case "kGoogleStandard":
            return VaultTypeVaultProviderStatsInfo_KGOOGLESTANDARD
        case "kGoogleDRA":
            return VaultTypeVaultProviderStatsInfo_KGOOGLEDRA
        case "kAmazonS3StandardIA":
            return VaultTypeVaultProviderStatsInfo_KAMAZONS3STANDARDIA
        case "kAWSGovCloud":
            return VaultTypeVaultProviderStatsInfo_KAWSGOVCLOUD
        case "kNAS":
            return VaultTypeVaultProviderStatsInfo_KNAS
        case "kColdline":
            return VaultTypeVaultProviderStatsInfo_KCOLDLINE
        case "kAzureGovCloud":
            return VaultTypeVaultProviderStatsInfo_KAZUREGOVCLOUD
        case "kAzureArchive":
            return VaultTypeVaultProviderStatsInfo_KAZUREARCHIVE
        case "kAzure":
            return VaultTypeVaultProviderStatsInfo_KAZURE
        case "kGoogle":
            return VaultTypeVaultProviderStatsInfo_KGOOGLE
        case "kAmazon":
            return VaultTypeVaultProviderStatsInfo_KAMAZON
        case "kOracle":
            return VaultTypeVaultProviderStatsInfo_KORACLE
        case "kOracleTierStandard":
            return VaultTypeVaultProviderStatsInfo_KORACLETIERSTANDARD
        case "kOracleTierArchive":
            return VaultTypeVaultProviderStatsInfo_KORACLETIERARCHIVE
        case "kAmazonC2S":
            return VaultTypeVaultProviderStatsInfo_KAMAZONC2S
        default:
            return VaultTypeVaultProviderStatsInfo_KNEARLINE
    }
}
